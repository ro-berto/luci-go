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

package rpcs

import (
	"go.chromium.org/luci/common/errors"
)

// GetOAuthTokenRequest is parameters for GetOAuthToken RPC call.
type GetOAuthTokenRequest struct {
	BaseRequest

	Scopes []string `json:"scopes"`
}

// Validate checks that the request is structurally valid.
func (r *GetOAuthTokenRequest) Validate() error {
	if len(r.Scopes) == 0 {
		return errors.New(`field "scopes" is required`)
	}
	return r.BaseRequest.Validate()
}

// GetOAuthTokenResponse is returned by GetOAuthToken RPC call.
type GetOAuthTokenResponse struct {
	BaseResponse

	AccessToken string `json:"access_token"`
	Expiry      int64  `json:"expiry"`
}
