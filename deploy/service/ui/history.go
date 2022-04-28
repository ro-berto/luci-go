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

package ui

import (
	"fmt"
	"strconv"
	"strings"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/server/router"
	"go.chromium.org/luci/server/templates"

	"go.chromium.org/luci/deploy/api/modelpb"
)

// actuationOutcome is an outcome of an actuation of a single asset.
type actuationOutcome string

const (
	outcomeUnknown   actuationOutcome = "UNKNOWN"   // should be unreachable
	outcomeUnchanged actuationOutcome = "UNCHANGED" // matches intent
	outcomeDisabled  actuationOutcome = "DISABLED"  // disabled in the config
	outcomeLocked    actuationOutcome = "LOCKED"    // has outstanding locks
	outcomeBroken    actuationOutcome = "BROKEN"    // broken configuration
	outcomeUpdating  actuationOutcome = "UPDATING"  // running right now
	outcomeUpdated   actuationOutcome = "UPDATED"   // applied changes
	outcomeFailed    actuationOutcome = "FAILED"    // failed to apply changes
)

func deriveOutcome(r *modelpb.AssetHistory) actuationOutcome {
	switch r.Decision.GetDecision() {
	case modelpb.ActuationDecision_SKIP_UPTODATE:
		return outcomeUnchanged

	case modelpb.ActuationDecision_SKIP_DISABLED:
		return outcomeDisabled

	case modelpb.ActuationDecision_SKIP_LOCKED:
		return outcomeLocked

	case modelpb.ActuationDecision_SKIP_BROKEN:
		return outcomeBroken

	case modelpb.ActuationDecision_ACTUATE_FORCE, modelpb.ActuationDecision_ACTUATE_STALE:
		switch r.Actuation.GetState() {
		case modelpb.Actuation_EXECUTING:
			return outcomeUpdating
		case modelpb.Actuation_SUCCEEDED:
			return outcomeUpdated
		case modelpb.Actuation_FAILED, modelpb.Actuation_EXPIRED:
			return outcomeFailed
		default:
			return outcomeUnknown // this should not be possible
		}

	default:
		return outcomeUnknown // this should not be possible
	}
}

// tableClass is the corresponding Bootstrap CSS table class.
func (o actuationOutcome) tableClass() string {
	switch o {
	case outcomeUnchanged, outcomeUpdated:
		return "" // default transparent background
	case outcomeDisabled:
		return "table-secondary"
	case outcomeLocked:
		return "table-warning"
	case outcomeUnknown, outcomeBroken, outcomeFailed:
		return "table-danger"
	case outcomeUpdating:
		return "table-info"
	default:
		panic("impossible")
	}
}

// badgeClass is the corresponding Bootstrap CSS badge class.
func (o actuationOutcome) badgeClass() string {
	switch o {
	case outcomeUnchanged, outcomeDisabled:
		return "bg-secondary"
	case outcomeLocked:
		return "bg-warning text-dark"
	case outcomeUnknown, outcomeBroken, outcomeFailed:
		return "bg-danger"
	case outcomeUpdating:
		return "bg-info text-dark"
	case outcomeUpdated:
		return "bg-success"
	default:
		panic("impossible")
	}
}

type commitDetails struct {
	Subject       linkHref // commit subject linking to gitiles
	Rev           string   // full commit revision
	AuthorEmail   string   // author email address
	CommitMessage string   // full commit message
}

func getCommitDetails(dep *modelpb.Deployment) commitDetails {
	message := strings.TrimSpace(dep.GetLatestCommit().GetCommitMessage())
	lines := strings.SplitN(message, "\n", 2)

	subject := commitHref(dep)
	if len(lines) == 0 {
		subject.Text = "-"
	} else {
		subject.Text = lines[0]
	}

	return commitDetails{
		Subject:       subject,
		Rev:           dep.GetConfigRev(),
		AuthorEmail:   dep.GetLatestCommit().GetAuthorEmail(),
		CommitMessage: message,
	}
}

type historyOverview struct {
	Href       string           // a link to the dedicate history entry page
	Age        linkHref         // when it started
	Commit     commitDetails    // commit subject, commit message, etc.
	Outcome    actuationOutcome // summary of what happened
	TableClass string           // CSS class for the table row
	BadgeClass string           // CSS class for the state cell
}

func deriveHistoryOverview(asset *modelpb.Asset, rec *modelpb.AssetHistory) *historyOverview {
	out := &historyOverview{
		Href:    fmt.Sprintf("/a/%s/history/%d", asset.Id, rec.HistoryId),
		Age:     timestampHref(rec.Actuation.Created, "", ""),
		Commit:  getCommitDetails(rec.Actuation.Deployment),
		Outcome: deriveOutcome(rec),
	}
	out.TableClass = out.Outcome.tableClass()
	out.BadgeClass = out.Outcome.badgeClass()
	return out
}

// historyListingPage renders the history listing page.
func (ui *UI) historyListingPage(ctx *router.Context, assetID string) error {
	ref := assetRefFromID(assetID)

	// TODO: Implement.

	templates.MustRender(ctx.Context, ctx.Writer, "pages/history-listing.html", map[string]interface{}{
		"Breadcrumbs": historyListingBreadcrumbs(ref),
		"Ref":         ref,
	})
	return nil
}

// historyEntryPage renders a page with a single actuation history entry.
func (ui *UI) historyEntryPage(ctx *router.Context, assetID, historyID string) error {
	entryID, err := strconv.ParseInt(historyID, 10, 64)
	if err != nil {
		return errors.Annotate(err, "not a valid history entry ID").Err()
	}

	ref := assetRefFromID(assetID)

	// TODO: Implement.

	templates.MustRender(ctx.Context, ctx.Writer, "pages/history-entry.html", map[string]interface{}{
		"Breadcrumbs": historyEntryBreadcrumbs(ref, entryID),
		"Ref":         ref,
	})
	return nil
}
