// Copyright 2021 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/sync/errgroup"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/config"
	"go.chromium.org/luci/cv/internal/gerrit/cfgmatcher"
	"go.chromium.org/luci/cv/internal/gerrit/trigger"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/impl/state"
)

// UpdateConfig implements Handler interface.
func (impl *Impl) UpdateConfig(ctx context.Context, rs *state.RunState, hash string) (*Result, error) {
	// First, check if config update is feasible given Run Status.
	switch status := rs.Run.Status; {
	case status == run.Status_STATUS_UNSPECIFIED:
		err := errors.Reason("CRITICAL: Received UpdateConfig event but Run is in unspecified status").Err()
		common.LogError(ctx, err)
		panic(err)
	case status == run.Status_SUBMITTING:
		// Can't update config while submitting.
		return &Result{State: rs, PreserveEvents: true}, nil
	case run.IsEnded(status):
		logging.Debugf(ctx, "Skipping updating config because Run is %s", status)
		return &Result{State: rs}, nil
	}

	// Second, load runCLs and new config groups.
	eg, egCtx := errgroup.WithContext(ctx)
	var runCLs []*run.RunCL
	eg.Go(func() error {
		var err error
		runCLs, err = run.LoadRunCLs(egCtx, rs.Run.ID, rs.Run.CLs)
		return err
	})
	var newMeta config.Meta
	var newConfigGroups []*config.ConfigGroup
	upToDate := false
	eg.Go(func() error {
		switch metas, err := config.GetHashMetas(egCtx, rs.Run.ID.LUCIProject(), rs.Run.ConfigGroupID.Hash(), hash); {
		case err != nil:
			return err
		case metas[0].EVersion >= metas[1].EVersion:
			// Current config at least as recent as the "new" one.
			upToDate = true
			return nil
		default:
			newMeta = metas[1]
			newConfigGroups, err = newMeta.GetConfigGroups(egCtx)
			return err
		}
	})
	switch err := eg.Wait(); {
	case err != nil:
		return nil, err
	case upToDate:
		return &Result{State: rs}, nil
	}

	matcher := cfgmatcher.LoadMatcherFromConfigGroups(ctx, newConfigGroups, &newMeta)
	cgsMap := make(map[string]*config.ConfigGroup, len(newConfigGroups))
	for _, cg := range newConfigGroups {
		cgsMap[cg.ID.Name()] = cg
	}

	// Third, try upgrading config.
	// A Run remains feasible iff all of:
	//  * each CL is still watched by one ConfigGroup;
	//  * all CLs are watched by the same ConfigGroup,
	//    although its name may have changed from the current Run.ConfigGroupID;
	//  * each CL's .Trigger is still the same.
	m := make(map[config.ConfigGroupID][]*run.RunCL, 1)
	var unwatched, multiple, diffTrigger []*run.RunCL
	for _, cl := range runCLs {
		switch cgids := matchingConfigGroups(matcher, cl); len(cgids) {
		case 0:
			unwatched = append(unwatched, cl)
		case 1:
			cgid := cgids[0]
			tr := trigger.Find(cl.Detail.GetGerrit().GetInfo(), cgsMap[cgid.Name()].Content)
			if hasTriggerChanged(cl, tr) {
				diffTrigger = append(diffTrigger, cl)
			} else {
				m[cgids[0]] = append(m[cgids[0]], cl)
			}
		default:
			multiple = append(multiple, cl)
		}
	}
	if len(unwatched) == 0 && len(multiple) == 0 && len(diffTrigger) == 0 && len(m) == 1 {
		// Run is still feasible.
		rs = rs.ShallowCopy()
		for cgid := range m { // extra first and only key from the map.
			rs.Run.ConfigGroupID = cgid
		}
		logging.Infof(ctx, "Upgrading to new ConfigGroupID %q", rs.Run.ConfigGroupID)
		return &Result{State: rs}, nil
	}
	// Run is no longer feasible and should be cancelled.
	reason := formatNoLongerFeasibleRunReason(unwatched, multiple, diffTrigger, m, &newMeta)
	logging.Warningf(ctx, reason)
	return impl.Cancel(ctx, rs)
}

func matchingConfigGroups(matcher *cfgmatcher.Matcher, cl *run.RunCL) []config.ConfigGroupID {
	ci := cl.Detail.GetGerrit().GetInfo()
	if ci == nil {
		panic(fmt.Errorf("only gerrit CLs are supported, not %s", cl.Detail))
	}
	return matcher.Match(cl.Detail.GetGerrit().GetHost(), ci.GetProject(), ci.GetRef())
}

func formatNoLongerFeasibleRunReason(
	unwatched, multiple, diffTrigger []*run.RunCL,
	m map[config.ConfigGroupID][]*run.RunCL,
	newMeta *config.Meta,
) string {
	// Computes detailed reason why to assist in debugging.
	buf := strings.Builder{}
	// TODO(tandrii): it'd very useful to list GitRevision here for faster
	// debugging, and definitely has to be done if this message is sent to the
	// users.
	fmt.Fprintf(&buf, "Run is no longer feasible due to project config change to %q\n", newMeta.Hash())
	if len(unwatched) > 0 {
		fmt.Fprintf(&buf, "%d CLs no longer watched:\n", len(unwatched))
		for _, cl := range unwatched {
			fmt.Fprintf(&buf, "  * %s\n", cl.ExternalID.MustURL())
		}
	}
	if len(multiple) > 0 {
		fmt.Fprintf(&buf, "%d CLs now match 2+ config groups:\n", len(multiple))
		for _, cl := range multiple {
			fmt.Fprintf(&buf, "  * %s\n", cl.ExternalID.MustURL())
		}
	}
	if len(diffTrigger) > 0 {
		fmt.Fprintf(&buf, "%d CLs have new trigger:\n", len(diffTrigger))
		for _, cl := range diffTrigger {
			fmt.Fprintf(&buf, "  * %s\n", cl.ExternalID.MustURL())
		}
	}
	if len(m) > 1 {
		fmt.Fprintf(&buf, "Run CLs now belong to %d different CQ config groups", len(m))
		for cgid, cls := range m {
			fmt.Fprintf(&buf, "  * %s:\n", cgid.Name())
			for _, cl := range cls {
				fmt.Fprintf(&buf, "      * %s\n", cl.ExternalID.MustURL())
			}
		}
	}
	return buf.String()
}