// Copyright 2017 The LUCI Authors.
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

package serviceaccounts

import (
	"context"
	"net"
	"time"

	"cloud.google.com/go/bigquery"

	"go.chromium.org/luci/common/proto/google"

	tokenserver "go.chromium.org/luci/tokenserver/api"
	"go.chromium.org/luci/tokenserver/api/admin/v1"
	bqpb "go.chromium.org/luci/tokenserver/api/bq"
	"go.chromium.org/luci/tokenserver/api/minter/v1"

	"go.chromium.org/luci/tokenserver/appengine/impl/utils"
	"go.chromium.org/luci/tokenserver/appengine/impl/utils/bq"
)

// MintedOAuthTokenInfo is passed to LogOAuthToken.
//
// It carries all information about the returned token.
type MintedOAuthTokenInfo struct {
	RequestedAt time.Time                              // when the RPC happened
	Request     *minter.MintOAuthTokenViaGrantRequest  // RPC input, as is
	Response    *minter.MintOAuthTokenViaGrantResponse // RPC output, as is
	GrantBody   *tokenserver.OAuthTokenGrantBody       // deserialized grant
	ConfigRev   string                                 // revision of the service_accounts.cfg used
	Rule        *admin.ServiceAccountRule              // the particular rule used to authorize the request
	PeerIP      net.IP                                 // caller IP address
	RequestID   string                                 // GAE request ID that handled the RPC
	AuthDBRev   int64                                  // revision of groups database (or 0 if unknown)
}

// toBigQueryMessage returns a message to upload to BigQuery.
func (i *MintedOAuthTokenInfo) toBigQueryMessage() *bqpb.OAuthToken {
	return &bqpb.OAuthToken{
		Fingerprint:      utils.TokenFingerprint(i.Response.AccessToken),
		GrantFingerprint: utils.TokenFingerprint(i.Request.GrantToken),
		ServiceAccount:   i.GrantBody.ServiceAccount,
		OauthScopes:      i.Request.OauthScope,
		ProxyIdentity:    i.GrantBody.Proxy,
		EndUserIdentity:  i.GrantBody.EndUser,

		// Note: we are not using 'issued_at' because the returned token is often
		// fetched from cache (and thus it was issued some time ago, not now). This
		// timestamp is not preserved in the cache, since it can be calculated from
		// 'expiration' if necessary.
		RequestedAt: google.NewTimestamp(i.RequestedAt),
		Expiration:  i.Response.Expiry,

		// Information supplied by the caller.
		AuditTags: i.Request.AuditTags,

		// Information about the service account rule used.
		ConfigRev:  i.ConfigRev,
		ConfigRule: i.Rule.Name,

		// Information about the request handler environment.
		PeerIp:         i.PeerIP.String(),
		ServiceVersion: i.Response.ServiceVersion,
		GaeRequestId:   i.RequestID,
		AuthDbRev:      i.AuthDBRev,
	}
}

// GrantLogger records info about the OAuth token grant to BigQuery.
type OAuthTokenLogger func(context.Context, *MintedOAuthTokenInfo) error

// NewOAuthTokenLogger returns a callback that records info about OAuth tokens
// to BigQuery.
//
// Tokens themselves are not logged. Only first 16 bytes of their SHA256 hashes
// (aka 'fingerprint') are. They are used only to identify tokens in logs.
//
// When dryRun is true, logs to the local text log only, not to BigQuery
// (to avoid accidentally pushing fake data to real BigQuery dataset).
func NewOAuthTokenLogger(client *bigquery.Client, dryRun bool) OAuthTokenLogger {
	inserter := bq.Inserter{
		Table:  client.Dataset("tokens").Table("oauth_tokens"),
		DryRun: dryRun,
	}
	return func(ctx context.Context, i *MintedOAuthTokenInfo) error {
		return inserter.Insert(ctx, i.toBigQueryMessage())
	}
}
