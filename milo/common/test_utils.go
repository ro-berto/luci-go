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

package common

import (
	"context"

	"github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/common/data/caching/lru"
	"go.chromium.org/luci/server/caching"
	"go.chromium.org/luci/server/caching/cachingtest"
)

// shouldResembleMatcher is a gomock.Matcher that performs
// convey.ShouldResemble(actual, expected).
type shouldResembleMatcher struct {
	expected interface{}
}

// NewShouldResemberMatcher constrcuts a gomock.Matcher that performs
// convey.ShouldResemble(actual, expected).
func NewShouldResemberMatcher(expected interface{}) shouldResembleMatcher {
	return shouldResembleMatcher{
		expected: expected,
	}
}

// Matches implements gomock.Matcher
func (e shouldResembleMatcher) Matches(actual interface{}) bool {
	convey.ShouldResemble(actual, e.expected)
	return true
}

// String implements gomock.Matcher
func (e shouldResembleMatcher) String() string {
	return "ShouldResemble"
}

// SetUpTestGlobalCache sets up GlobalCache in the context.
func SetUpTestGlobalCache(ctx context.Context) context.Context {
	caches := make(map[string]caching.BlobCache)
	return caching.WithGlobalCache(ctx, func(namespace string) caching.BlobCache {
		cache, ok := caches[namespace]
		if !ok {
			cache = &cachingtest.BlobCache{LRU: lru.New(0)}
			caches[namespace] = cache
		}
		return cache
	})
}
