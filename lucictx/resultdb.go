// Copyright 2020 The LUCI Authors.
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

package lucictx

import (
	"context"
)

// GetResultDB returns the current ResultDB from LUCI_CONTEXT if it was present.
// nil, otherwise.
func GetResultDB(ctx context.Context) *ResultDB {
	ret := ResultDB{}
	ok, err := Lookup(ctx, "resultdb", &ret)
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	return &ret
}

// SetResultDB sets the ResultDB in the LUCI_CONTEXT.
func SetResultDB(ctx context.Context, db *ResultDB) context.Context {
	return Set(ctx, "resultdb", db)
}
