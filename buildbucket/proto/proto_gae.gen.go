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

// Code generated by gae/tools/proto-gae/proto_gae.go. DO NOT EDIT.

// +build !copybara

package buildbucketpb

import (
	"google.golang.org/protobuf/proto"

	"go.chromium.org/luci/gae/service/datastore"
)

var _ datastore.PropertyConverter = (*Bucket)(nil)

// ToProperty implements datastore.PropertyConverter. It causes an embedded
// 'Bucket' to serialize to an unindexed '[]byte' when used with the
// "go.chromium.org/luci/gae" library.
func (p *Bucket) ToProperty() (prop datastore.Property, err error) {
	data, err := proto.Marshal(p)
	if err == nil {
		prop.SetValue(data, datastore.NoIndex)
	}
	return
}

// FromProperty implements datastore.PropertyConverter. It parses a '[]byte'
// into an embedded 'Bucket' when used with the "go.chromium.org/luci/gae" library.
func (p *Bucket) FromProperty(prop datastore.Property) error {
	data, err := prop.Project(datastore.PTBytes)
	if err != nil {
		return err
	}
	return proto.Unmarshal(data.([]byte), p)
}

var _ datastore.PropertyConverter = (*Build)(nil)

// ToProperty implements datastore.PropertyConverter. It causes an embedded
// 'Build' to serialize to an unindexed '[]byte' when used with the
// "go.chromium.org/luci/gae" library.
func (p *Build) ToProperty() (prop datastore.Property, err error) {
	data, err := proto.Marshal(p)
	if err == nil {
		prop.SetValue(data, datastore.NoIndex)
	}
	return
}

// FromProperty implements datastore.PropertyConverter. It parses a '[]byte'
// into an embedded 'Build' when used with the "go.chromium.org/luci/gae" library.
func (p *Build) FromProperty(prop datastore.Property) error {
	data, err := prop.Project(datastore.PTBytes)
	if err != nil {
		return err
	}
	return proto.Unmarshal(data.([]byte), p)
}

var _ datastore.PropertyConverter = (*BuildInfra)(nil)

// ToProperty implements datastore.PropertyConverter. It causes an embedded
// 'BuildInfra' to serialize to an unindexed '[]byte' when used with the
// "go.chromium.org/luci/gae" library.
func (p *BuildInfra) ToProperty() (prop datastore.Property, err error) {
	data, err := proto.Marshal(p)
	if err == nil {
		prop.SetValue(data, datastore.NoIndex)
	}
	return
}

// FromProperty implements datastore.PropertyConverter. It parses a '[]byte'
// into an embedded 'BuildInfra' when used with the "go.chromium.org/luci/gae" library.
func (p *BuildInfra) FromProperty(prop datastore.Property) error {
	data, err := prop.Project(datastore.PTBytes)
	if err != nil {
		return err
	}
	return proto.Unmarshal(data.([]byte), p)
}

var _ datastore.PropertyConverter = (*Builder)(nil)

// ToProperty implements datastore.PropertyConverter. It causes an embedded
// 'Builder' to serialize to an unindexed '[]byte' when used with the
// "go.chromium.org/luci/gae" library.
func (p *Builder) ToProperty() (prop datastore.Property, err error) {
	data, err := proto.Marshal(p)
	if err == nil {
		prop.SetValue(data, datastore.NoIndex)
	}
	return
}

// FromProperty implements datastore.PropertyConverter. It parses a '[]byte'
// into an embedded 'Builder' when used with the "go.chromium.org/luci/gae" library.
func (p *Builder) FromProperty(prop datastore.Property) error {
	data, err := prop.Project(datastore.PTBytes)
	if err != nil {
		return err
	}
	return proto.Unmarshal(data.([]byte), p)
}
