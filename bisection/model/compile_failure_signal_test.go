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

package model

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddLine(t *testing.T) {
	Convey("Add line or file path", t, func() {
		signal := &CompileFailureSignal{}
		signal.AddLine("a/b", 12)
		So(signal.Files, ShouldResemble, map[string][]int{"a/b": {12}})
		signal.AddLine("a/b", 14)
		So(signal.Files, ShouldResemble, map[string][]int{"a/b": {12, 14}})
		signal.AddLine("c/d", 8)
		So(signal.Files, ShouldResemble, map[string][]int{"a/b": {12, 14}, "c/d": {8}})
		signal.AddLine("a/b", 14)
		So(signal.Files, ShouldResemble, map[string][]int{"a/b": {12, 14}, "c/d": {8}})
		signal.AddFilePath("x/y")
		So(signal.Files, ShouldResemble, map[string][]int{"a/b": {12, 14}, "c/d": {8}, "x/y": {}})
		signal.AddFilePath("x/y")
		So(signal.Files, ShouldResemble, map[string][]int{"a/b": {12, 14}, "c/d": {8}, "x/y": {}})
	})
}

func TestCalculateDependencyMap(t *testing.T) {
	Convey("Calculate dependency map", t, func() {
		signal := &CompileFailureSignal{
			Edges: []*CompileFailureEdge{
				{
					Dependencies: []string{
						"x/y/a.h",
						"xx/yy/b.h",
					},
				},
				{
					Dependencies: []string{
						"y/z/a.cc",
						"zz/y/c.yy",
						"x/y/a.h",
					},
				},
			},
		}
		signal.CalculateDependencyMap(context.Background())
		So(signal.DependencyMap, ShouldResemble, map[string][]string{
			"a": {"x/y/a.h", "y/z/a.cc"},
			"b": {"xx/yy/b.h"},
			"c": {"zz/y/c.yy"},
		})
	})
}
