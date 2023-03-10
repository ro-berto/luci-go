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

package pagination

import (
	"fmt"

	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/grpc/appstatus"
)

// InvalidPageSize annotates the error with InvalidArgument appstatus.
func InvalidPageSize(err error) error {
	return appstatus.Attachf(err, codes.InvalidArgument, "invalid page size, must be >= 0")
}

type requestWithPageSize interface {
	GetPageSize() int32
}

// ValidatePageSize validates and caps page size from the given request.
func ValidatePageSize(req requestWithPageSize, defaultValue, maxValue int32) (int32, error) {
	if defaultValue > maxValue {
		panic(fmt.Errorf("invalid use: defaultValue %d must be <= maxValue %d", defaultValue, maxValue))
	}
	switch v := req.GetPageSize(); {
	case v < 0:
		return 0, InvalidPageSize(fmt.Errorf("page_size %d", v))
	case v == 0:
		return defaultValue, nil
	case v < maxValue:
		return v, nil
	default:
		return maxValue, nil
	}
}
