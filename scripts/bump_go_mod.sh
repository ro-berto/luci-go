#!/bin/bash
# Copyright 2022 The LUCI Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script updates selected dependencies in go.mod to latest versions.
#
# Dependencies that can be added to this list:
#  * Base dependencies shared by majority of packages in luci-go.
#  * Dependencies critical for security.
#  * Dependencies with a track record of not causing troubles during updates.
#
# Ideally running this script should always be a stress-free experience.

set -e

deps=(
  cloud.google.com/go/bigquery@latest
  cloud.google.com/go/bigtable@latest
  cloud.google.com/go/cloudtasks@latest
  cloud.google.com/go/compute@latest
  cloud.google.com/go/compute/metadata@latest
  cloud.google.com/go/datastore@latest
  cloud.google.com/go/errorreporting@latest
  cloud.google.com/go/iam@latest
  cloud.google.com/go/kms@latest
  cloud.google.com/go/logging@latest
  cloud.google.com/go/profiler@latest
  cloud.google.com/go/pubsub@latest
  cloud.google.com/go/secretmanager@latest
  cloud.google.com/go/spanner@latest
  cloud.google.com/go/storage@latest
  github.com/alicebob/miniredis/v2@latest
  github.com/danjacques/gofslock@latest
  github.com/dustin/go-humanize@latest
  github.com/envoyproxy/protoc-gen-validate@latest
  github.com/golang/protobuf@latest
  github.com/gomodule/redigo@latest
  github.com/google/go-cmp@latest
  github.com/google/tink/go@latest
  github.com/google/uuid@latest
  github.com/gorhill/cronexpr@latest
  github.com/jordan-wright/email@latest
  github.com/julienschmidt/httprouter@latest
  github.com/klauspost/compress@latest
  github.com/luci/gtreap@latest
  github.com/maruel/subcommands@latest
  github.com/mattn/go-tty@latest
  github.com/mgutz/ansi@latest
  github.com/Microsoft/go-winio@latest
  github.com/mitchellh/go-homedir@latest
  github.com/op/go-logging@latest
  github.com/pmezard/go-difflib@latest
  github.com/protocolbuffers/txtpbfmt@latest
  github.com/russross/blackfriday/v2@latest
  github.com/sergi/go-diff@latest
  github.com/smartystreets/assertions@latest
  github.com/smartystreets/goconvey@latest
  github.com/yosuke-furukawa/json5@latest
  go.starlark.net@latest
  golang.org/x/crypto@latest
  golang.org/x/net@latest
  golang.org/x/oauth2@latest
  golang.org/x/sync@latest
  golang.org/x/sys@latest
  golang.org/x/term@latest
  golang.org/x/time@latest
  golang.org/x/tools@latest
  google.golang.org/api@latest
  google.golang.org/appengine@504804fb50  # "@latest" picks up quite old v1.6.7
  google.golang.org/genproto@latest
  google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  google.golang.org/grpc@latest
  google.golang.org/protobuf@latest
  gopkg.in/yaml.v2@latest
)

for mod in ${deps[@]}; do
  echo go get -d ${mod}
  go get -d ${mod}
done

echo go mod tidy
go mod tidy
