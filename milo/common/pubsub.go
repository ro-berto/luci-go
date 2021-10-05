// Copyright 2019 The LUCI Authors.
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

package common

import (
	"encoding/base64"
)

type PubSubMessage struct {
	Attributes map[string]interface{} `json:"attributes"`
	Data       string                 `json:"data"`
	MessageID  string                 `json:"message_id"`
}

type PubSubSubscription struct {
	Message PubSubMessage `json:"message"`
}

// GetData returns the expanded form of Data (decoded from base64).
func (m *PubSubSubscription) GetData() ([]byte, error) {
	return base64.StdEncoding.DecodeString(m.Message.Data)
}
