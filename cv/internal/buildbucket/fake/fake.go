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

package bbfake

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"

	bbpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/errors"
	"google.golang.org/protobuf/proto"

	"go.chromium.org/luci/cv/internal/buildbucket"
)

const requestDeduplicationWindow = 1 * time.Minute

type fakeApp struct {
	hostname      string
	nextBuildID   int64 // for generating monotonically decreasing build ID
	requestCache  timedMap
	buildStoreMu  sync.RWMutex
	buildStore    map[int64]*bbpb.Build // build ID -> build
	configStoreMu sync.RWMutex
	configStore   map[string]*bbpb.BuildbucketCfg // project name -> config
}

type Fake struct {
	hostsMu sync.Mutex
	hosts   map[string]*fakeApp // hostname -> fakeApp
}

// NewClientFactory returns a factory that creates a client for this buildbucket
// fake.
func (f *Fake) NewClientFactory() buildbucket.ClientFactory {
	return clientFactory{
		fake: f,
	}
}

// MustNewClient is a shorthand of `fake.NewClientFactory().MakeClient(...)`.
//
// Panics if fails to create new client.
func (f *Fake) MustNewClient(ctx context.Context, host, luciProject string) *Client {
	factory := clientFactory{
		fake: f,
	}
	client, err := factory.MakeClient(ctx, host, luciProject)
	if err != nil {
		panic(errors.Annotate(err, "failed to create new buildbucket client").Err())
	}
	return client.(*Client)
}

// AddBuilder adds a new builder configuration to fake Buildbucket host.
//
// Overwrites the existing builder if the same builder already exists.
// `properties` should be marshallable by `encoding/json`.
func (f *Fake) AddBuilder(host string, builder *bbpb.BuilderID, properties interface{}) *Fake {
	fa := f.ensureApp(host)
	fa.configStoreMu.Lock()
	defer fa.configStoreMu.Unlock()
	if _, ok := fa.configStore[builder.GetProject()]; !ok {
		fa.configStore[builder.GetProject()] = &bbpb.BuildbucketCfg{}
	}
	cfg := fa.configStore[builder.GetProject()]
	var bucket *bbpb.Bucket
	for _, b := range cfg.GetBuckets() {
		if b.Name == builder.GetBucket() {
			bucket = b
			break
		}
	}
	if bucket == nil {
		bucket = &bbpb.Bucket{
			Name:     builder.GetBucket(),
			Swarming: &bbpb.Swarming{},
		}
		cfg.Buckets = append(cfg.GetBuckets(), bucket)
	}

	builderCfg := &bbpb.BuilderConfig{
		Name: builder.GetBuilder(),
	}
	if properties != nil {
		bProperties, err := json.Marshal(properties)
		if err != nil {
			panic(err)
		}
		builderCfg.Properties = string(bProperties)
	}
	for i, b := range bucket.GetSwarming().GetBuilders() {
		if b.Name == builder.GetBuilder() {
			bucket.GetSwarming().GetBuilders()[i] = builderCfg
			return f
		}
	}
	bucket.GetSwarming().Builders = append(bucket.GetSwarming().GetBuilders(), builderCfg)
	return f
}

// AddBuild adds a build to fake Buildbucket host.
//
// Reads Buildbucket hostname from `infra.buildbucket.hostname`.
// Overwrites the existing build if the build with same ID already exists.
//
// TODO(yiwzhang): make it private so that external package should
// always use schedule build to create new build s.t. the build ID
// will be monotically decreasing.
func (f *Fake) AddBuild(build *bbpb.Build) {
	host := build.GetInfra().GetBuildbucket().GetHostname()
	if host == "" {
		panic(fmt.Errorf("missing host for build %d", build.Id))
	}
	fa := f.ensureApp(host)
	fa.buildStoreMu.Lock()
	fa.buildStore[build.GetId()] = build
	fa.buildStoreMu.Unlock()
}

func (f *Fake) ensureApp(host string) *fakeApp {
	f.hostsMu.Lock()
	defer f.hostsMu.Unlock()
	if _, ok := f.hosts[host]; !ok {
		if f.hosts == nil {
			f.hosts = make(map[string]*fakeApp)
		}
		f.hosts[host] = &fakeApp{
			hostname:    host,
			nextBuildID: math.MaxInt64 - 1,
			buildStore:  make(map[int64]*bbpb.Build),
			configStore: make(map[string]*bbpb.BuildbucketCfg),
		}
	}
	return f.hosts[host]
}

func (fa *fakeApp) getBuild(id int64) *bbpb.Build {
	fa.buildStoreMu.RLock()
	defer fa.buildStoreMu.RUnlock()
	if build, ok := fa.buildStore[id]; ok {
		return proto.Clone(build).(*bbpb.Build)
	}
	return nil
}

func (fa *fakeApp) iterBuildStore(cb func(*bbpb.Build)) {
	fa.buildStoreMu.RLock()
	defer fa.buildStoreMu.RUnlock()
	for _, build := range fa.buildStore {
		cb(proto.Clone(build).(*bbpb.Build))
	}
}

func (fa *fakeApp) updateBuild(id int64, cb func(*bbpb.Build)) *bbpb.Build {
	fa.buildStoreMu.Lock()
	defer fa.buildStoreMu.Unlock()
	if build, ok := fa.buildStore[id]; ok {
		cb(build)
		// store a copy to avoid cb keeps the reference to the build and mutate it
		// later.
		fa.buildStore[id] = proto.Clone(build).(*bbpb.Build)
		return build
	}
	return nil
}

// insertBuild also generates a monotically decreasing build ID.
//
// Caches the build for `requestDeduplicationWindow` to deduplicate request
// with same request ID later.
func (fa *fakeApp) insertBuild(ctx context.Context, build *bbpb.Build, requestID string) *bbpb.Build {
	fa.buildStoreMu.Lock()
	defer fa.buildStoreMu.Unlock()
	build.Id = fa.nextBuildID
	fa.nextBuildID--
	if _, ok := fa.buildStore[build.Id]; ok {
		panic(fmt.Sprintf("build %d already exists", build.Id))
	}
	cloned := proto.Clone(build).(*bbpb.Build)
	fa.buildStore[build.Id] = cloned
	if requestID != "" {
		fa.requestCache.set(ctx, requestID, cloned, requestDeduplicationWindow)
	}
	return build
}

func (fa *fakeApp) findDupRequest(ctx context.Context, requestID string) *bbpb.Build {
	if requestID == "" {
		return nil
	}
	if b, ok := fa.requestCache.get(ctx, requestID); ok {
		return proto.Clone(b.(*bbpb.Build)).(*bbpb.Build)
	}
	return nil
}

func (fa *fakeApp) loadBuilderCfg(builderID *bbpb.BuilderID) *bbpb.BuilderConfig {
	fa.configStoreMu.RLock()
	defer fa.configStoreMu.RUnlock()
	cfg, ok := fa.configStore[builderID.GetProject()]
	if !ok {
		return nil
	}
	for _, bucket := range cfg.GetBuckets() {
		if bucket.GetName() != builderID.GetBucket() {
			continue
		}
		for _, builder := range bucket.GetSwarming().GetBuilders() {
			if builder.GetName() == builderID.GetBuilder() {
				return proto.Clone(builder).(*bbpb.BuilderConfig)
			}
		}
	}
	return nil
}
