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

package app

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	// Needed to ensure task class is registered.
	_ "go.chromium.org/luci/analysis/internal/services/resultingester"
)

func makeReq(blob []byte) io.ReadCloser {
	msg := struct {
		Message struct {
			Data []byte
		}
		Attributes map[string]interface{}
	}{struct{ Data []byte }{Data: blob}, nil}
	jmsg, _ := json.Marshal(msg)
	return ioutil.NopCloser(bytes.NewReader(jmsg))
}
