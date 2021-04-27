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

package prpc

import (
	"encoding/base64"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"

	"go.chromium.org/luci/common/errors"
)

// isReservedMetadataKey returns true for disallowed metadata keys.
//
// Keys are given in HTTP header canonical format.
//
// Setting metadata with such keys may break the protocol.
// See also Client.prepareRequest.
func isReservedMetadataKey(k string) bool {
	switch {
	case strings.HasPrefix(k, "X-Prpc-"):
		return true

	case k == "Accept",
		k == "Accept-Encoding",
		k == "Content-Encoding",
		k == "Content-Length",
		k == "Content-Type",
		k == "X-Content-Type-Options":
		return true

	default:
		return false
	}
}

// metaIntoHeaders merges outgoing metadata into the given set of headers.
//
// Encodes metadata entries with keys that end with "-bin".
func metaIntoHeaders(md metadata.MD, h http.Header) error {
	for k, vs := range md {
		canon := http.CanonicalHeaderKey(k)
		if isReservedMetadataKey(canon) {
			return errors.Reason("using reserved metadata key %q", k).Err()
		}
		if !strings.HasSuffix(canon, "-Bin") {
			h[canon] = append(h[canon], vs...)
		} else {
			for _, v := range vs {
				h[canon] = append(h[canon], base64.StdEncoding.EncodeToString([]byte(v)))
			}
		}
	}
	return nil
}

// headerIntoMeta merges values of the given header key into the metadata.
//
// Decodes entries with keys that end with "-bin". See also readStatusDetails
// which also does this decoding specifically for "X-Prpc-Status-Details-Bin".
func headerIntoMeta(key string, values []string, md metadata.MD) error {
	key = strings.ToLower(key)
	if !strings.HasSuffix(key, "-bin") {
		md[key] = append(md[key], values...)
		return nil
	}
	for _, v := range values {
		decoded, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			return err
		}
		md[key] = append(md[key], string(decoded))
	}
	return nil
}

// headersIntoMetadata returns a new metadata.MD constructed from given headers.
//
// All reserved headers are silently skipped.
func headersIntoMetadata(h http.Header) (metadata.MD, error) {
	if len(h) == 0 {
		return nil, nil
	}
	md := make(metadata.MD, len(h))
	for k, v := range h {
		if isReservedMetadataKey(http.CanonicalHeaderKey(k)) {
			continue
		}
		if err := headerIntoMeta(k, v, md); err != nil {
			return nil, errors.Annotate(err, "can't decode header %q", k).Err()
		}
	}
	return md, nil
}
