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

package changelist

import (
	"fmt"
	"strconv"
	"strings"

	"go.chromium.org/luci/common/errors"
)

// ExternalID is a unique CL ID deterministically constructed based on CL data.
//
// Currently, only Gerrit is supported.
type ExternalID string

// GobID makes an ExternalID for a Gerrit CL.
//
// Host is typically "something-review.googlesource.com".
// Change is a number, e.g. 2515619 for
// https://chromium-review.googlesource.com/c/infra/luci/luci-go/+/2515619
func GobID(host string, change int64) (ExternalID, error) {
	if strings.ContainsRune(host, '/') {
		return "", errors.Reason("invalid host %q: must not contain /", host).Err()
	}
	return ExternalID(fmt.Sprintf("gerrit/%s/%d", host, change)), nil
}

// MustGobID is like GobID but panics on error.
func MustGobID(host string, change int64) ExternalID {
	ret, err := GobID(host, change)
	if err != nil {
		panic(err)
	}
	return ret
}

// ParseGobID returns Gerrit host and change if this is a GobID.
func (eid ExternalID) ParseGobID() (host string, change int64, err error) {
	parts := strings.Split(string(eid), "/")
	if len(parts) != 3 || parts[0] != "gerrit" {
		err = errors.Reason("%q is not a valid GobID", eid).Err()
		return
	}
	host = parts[1]
	change, err = strconv.ParseInt(parts[2], 10, 63)
	if err != nil {
		err = errors.Annotate(err, "%q is not a valid GobID", eid).Err()
	}
	return
}

// URL returns URL of the CL.
func (eid ExternalID) URL() (string, error) {
	parts := strings.Split(string(eid), "/")
	if len(parts) < 2 {
		return "", errors.Reason("invalid ExternalID: %q", eid).Err()
	}
	switch kind := parts[0]; kind {
	case "gerrit":
		return fmt.Sprintf("https://%s/c/%s", parts[1], parts[2]), nil
	default:
		return "", errors.Reason("unrecognized ExternalID: %q", eid).Err()
	}
}

// MustURL is like `URL()` but panic on err.
func (eid ExternalID) MustURL() string {
	ret, err := eid.URL()
	if err != nil {
		panic(err)
	}
	return ret
}
