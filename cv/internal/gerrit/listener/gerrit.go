// Copyright 2022 The LUCI Authors.
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

package listener

import (
	"context"
	"strconv"
	"strings"

	"cloud.google.com/go/pubsub"
	"google.golang.org/protobuf/encoding/protojson"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/common/retry/transient"

	"go.chromium.org/luci/cv/internal/changelist"
	listenerpb "go.chromium.org/luci/cv/settings/listener"
)

// gerritProcessor implements processor interface for Gerrit subscription.
type gerritProcessor struct {
	sch       scheduler
	host      string
	prjFinder *projectFinder
}

// process processes a given Gerrit pubsub message and schedules UpdateCLTask(s)
// for all the LUCI projects watching the Gerrit repo.
func (p *gerritProcessor) process(ctx context.Context, m *pubsub.Message) error {
	if len(m.Data) == 0 {
		return nil
	}
	msg := &gerritpb.SourceRepoEvent{}
	if err := protojson.Unmarshal(m.Data, msg); err != nil {
		return errors.Annotate(err, "json.Unmarshal").Err()
	}

	var repo string
	switch chunks := strings.SplitN(msg.Name, "/", 4); {
	case len(chunks) != 4, chunks[0] != "projects", chunks[2] != "repos":
		// This is the format of Gerrit pubsub payload. If the format unmatches,
		// it's likely a bug in CV or Gerrit.
		return errors.Reason("invalid SourceRepoEvent name: %q", msg.Name).Err()
	default:
		repo = chunks[3]
	}

	// If no project is watching the repo, don't bother parsing the payload.
	prjs, err := p.prjFinder.lookup(ctx, p.host, repo)
	switch {
	case err != nil:
		return errors.Annotate(err, "projectFinder.lookup").Err()
	case len(prjs) == 0:
		return nil
	}

	// MetaRevIDs by ExternalIDs.
	var eidToMetaRevID map[string]string
	if e := msg.GetRefUpdateEvent(); e != nil {
		eidToMetaRevID = make(map[string]string, len(e.GetRefUpdates()))
		for ref, ev := range e.GetRefUpdates() {
			// CV is only interested in CL update events, of which ref name
			// ends with "/meta" in the following format.
			// : "refs/changes/<val>/<change_num>/meta"
			chunks := strings.SplitN(ref, "/", 5)
			switch {
			case len(chunks) != 5,
				chunks[0] != "refs",
				chunks[1] != "changes",
				chunks[4] != "meta":
				continue
			}
			change, err := strconv.ParseInt(chunks[3], 10, 63)
			if err != nil {
				// Must be a bug either in Gerrit or CV.
				return errors.Annotate(err, "invalid change num (%s): %s", chunks[3], msg).Err()
			}
			eid, err := changelist.GobID(p.host, change)
			if err != nil {
				return errors.Annotate(err, "changelist.GobID").Err()
			}

			switch prev, exist := eidToMetaRevID[string(eid)]; {
			case exist && prev != ev.NewId:
				// RefUpdateEvent is a map type. Therefore, a single pubsub
				// message can have at most one update event for each of the CLs
				// listed.
				//
				// If a duplicate ExternalID with different RevID is found,
				// there is a bug in CV or Gerrit.
				return errors.Reason("found multiple meta-rev-ids (%q, %q) for %q: %s",
					prev, ev.NewId, eid, msg).Err()
			case exist && prev == ev.NewId:
				// Still strange, but ok.
				logging.Warningf(ctx, "duplicate update events found for %q: %s", eid, msg)
			case !exist:
				eidToMetaRevID[string(eid)] = ev.NewId
			}
		}
	}

	for eid, meta := range eidToMetaRevID {
		for _, prj := range prjs {
			task := &changelist.UpdateCLTask{
				LuciProject: prj,
				ExternalId:  eid,
				Requester:   changelist.UpdateCLTask_PUBSUB_POLL,
				Hint:        &changelist.UpdateCLTask_Hint{MetaRevId: meta},
			}
			if err := p.sch.Schedule(ctx, task); err != nil {
				return errors.Annotate(err, "Schedule").Tag(transient.Tag).Err()
			}
		}
	}
	return nil
}

func newGerritSubscriber(c *pubsub.Client, sch scheduler, prjFinder *projectFinder, settings *listenerpb.Settings_GerritSubscription) *subscriber {
	subID := settings.GetSubscriptionId()
	if subID == "" {
		subID = settings.GetHost()
	}
	sber := &subscriber{
		sub: c.Subscription(subID),
		proc: &gerritProcessor{
			sch:       sch,
			host:      settings.GetHost(),
			prjFinder: prjFinder,
		},
	}
	sber.sub.ReceiveSettings.NumGoroutines = defaultNumGoroutines
	sber.sub.ReceiveSettings.MaxOutstandingMessages = defaultMaxOutstandingMessages
	if val := settings.GetReceiveSettings().GetNumGoroutines(); val != 0 {
		sber.sub.ReceiveSettings.NumGoroutines = int(val)
	}
	if val := settings.GetReceiveSettings().GetMaxOutstandingMessages(); val != 0 {
		sber.sub.ReceiveSettings.MaxOutstandingMessages = int(val)
	}
	return sber
}

// sameGerritSubscriberSettings returns true if a given GerritSubscriber is
// configured with given settings.
func sameGerritSubscriberSettings(sber *subscriber, settings *listenerpb.Settings_GerritSubscription) bool {
	subID := settings.GetSubscriptionId()
	if subID == "" {
		subID = settings.GetHost()
	}
	return (sber.proc.(*gerritProcessor).host == settings.GetHost() &&
		sber.sub.ID() == subID &&
		sber.sameReceiveSettings(settings.GetReceiveSettings()))
}
