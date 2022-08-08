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

// Package cipderr contains an enumeration with possible CIPD error categories.
package cipderr

import (
	"context"

	"go.chromium.org/luci/common/errors"
)

// Code is returned as part of JSON output by CIPD CLI.
//
// It as an enumeration with broad categories of possible errors.
type Code string

const (
	// Authentication or authorization error when contacting the backend.
	Auth Code = "auth_error"
	// An error doing local I/O (i.e. writing or reading files).
	IO Code = "io_error"
	// An IO error reading or writing from CIPD CAS.
	CAS Code = "cas_error"
	// An incorrectly formatted version name, instance ID, etc.
	BadArgument Code = "bad_argument_error"
	// A requested package is missing or its version can't be resolved.
	InvalidVersion Code = "invalid_version_error"
	// An error getting a response from the backend.
	BackendUnavailable Code = "backend_unavailable_error"
	// A generic fatal RPC error, e.g. violation of some precodition.
	RPC Code = "rpc_error"
	// Something (e.g. a resolved pins file) needs to be regenerated.
	Stale Code = "stale_error"
	// A hash of downloaded file doesn't match the expected value.
	HashMismatch Code = "hash_mismatch_error"
	// The admission plugin forbid installation of a package.
	NotAdmitted Code = "not_admitted_error"
	// A timeout of some sort.
	Timeout Code = "timeout_error"
	// Unrecognized (possibly transient) error.
	Unknown Code = "unknown_error"
)

// Details can be optionally attached to an error.
type Details struct {
	Package string `json:"package,omitempty"`
	Version string `json:"version,omitempty"`
	Subdir  string `json:"subdir,omitempty"`
}

var tagKey = errors.NewTagKey("cipderr.Code")

type tagValue struct {
	code    Code
	details *Details
}

// GenerateErrorTagValue is part of errors.TagValueGenerator, allowing this
// pair to be used as en error tag.
func (v tagValue) GenerateErrorTagValue() errors.TagValue {
	return errors.TagValue{Key: tagKey, Value: v}
}

// GenerateErrorTagValue is part of errors.TagValueGenerator, allowing this
// code to be used as en error tag.
func (c Code) GenerateErrorTagValue() errors.TagValue {
	return errors.TagValue{Key: tagKey, Value: tagValue{code: c}}
}

// WithDetails returns a error tag that attaches this code together with some
// details.
func (c Code) WithDetails(d Details) errors.TagValueGenerator {
	return tagValue{code: c, details: &d}
}

// ToCode examines a CIPD error to get a representative error code.
func ToCode(err error) Code {
	if val, ok := errors.TagValueIn(tagKey, err); ok {
		return val.(tagValue).code
	}
	deadline := errors.Any(err, func(err error) bool {
		return err == context.DeadlineExceeded || err == context.Canceled
	})
	if deadline {
		return Timeout
	}
	return Unknown
}

// ToDetails extracts error details, if available.
func ToDetails(err error) *Details {
	if val, ok := errors.TagValueIn(tagKey, err); ok {
		return val.(tagValue).details
	}
	return nil
}

// AttachDetails attaches details to an error, preserving its code.
//
// Overrides any previous details. Does nothing if `err` is nil.
func AttachDetails(err *error, d Details) {
	if *err != nil {
		*err = errors.Annotate(*err, "").Tag(ToCode(*err).WithDetails(d)).Err()
	}
}
