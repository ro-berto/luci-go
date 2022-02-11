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

package graph

import (
	"testing"

	_ "go.chromium.org/luci/gae/service/datastore/crbug1242998safeget"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubgraphOperations(t *testing.T) {
	t.Parallel()

	Convey("Testing addNode", t, func() {
		subgraph := &Subgraph{
			Nodes:     []*SubgraphNode{},
			nodesToID: map[NodeKey]int{},
		}

		Convey("Testing adding group.", func() {
			testGroup := "test-group"
			nodeID, placed := subgraph.addNode(Group, testGroup)
			So(nodeID, ShouldEqual, 0)
			So(placed, ShouldBeTrue)
		})

		Convey("Testing adding user.", func() {
			testUser := "user:m1@example.com"
			nodeID, placed := subgraph.addNode(Identity, testUser)
			So(nodeID, ShouldEqual, 0)
			So(placed, ShouldBeTrue)
		})

		Convey("Testing adding glob.", func() {
			testGlob := "user:*@example.com"
			nodeID, placed := subgraph.addNode(Glob, testGlob)
			So(nodeID, ShouldEqual, 0)
			So(placed, ShouldBeTrue)
		})

		Convey("Testing adding same node.", func() {
			testGroup := "test-group"
			nodeID, placed := subgraph.addNode(Group, testGroup)
			So(nodeID, ShouldEqual, 0)
			So(placed, ShouldBeTrue)
			nodeID, placed = subgraph.addNode(Group, testGroup)
			So(nodeID, ShouldEqual, 0)
			So(placed, ShouldBeFalse)
		})

		Convey("Testing key for nodesToID.", func() {
			testGroup := "test-group"
			testGlob := "user:*@example.com"
			testUser := "user:m1@example.com"
			subgraph.addNode(Group, testGroup)
			subgraph.addNode(Glob, testGlob)
			subgraph.addNode(Identity, testUser)

			expectedNodeMap := map[NodeKey]int{
				{Group, testGroup}:   0,
				{Glob, testGlob}:     1,
				{Identity, testUser}: 2,
			}

			So(subgraph.nodesToID, ShouldResemble, expectedNodeMap)
		})

	})

	Convey("Testing addEdge", t, func() {
		testGlob := "user:*@example.com"
		testGroup0 := "test-group-0"
		testUser := "user:m1@example.com"
		testGroup1 := "test-group-1"
		testGroup2 := "test-group-2"

		subgraph := &Subgraph{
			Nodes: []*SubgraphNode{
				{
					NodeKey: NodeKey{
						Kind:  Glob,
						Value: testGlob,
					},
				},
				{
					NodeKey: NodeKey{
						Kind:  Group,
						Value: testGroup0,
					},
				},
				{
					NodeKey: NodeKey{
						Kind:  Identity,
						Value: testUser,
					},
				},
				{
					NodeKey: NodeKey{
						Kind:  Group,
						Value: testGroup1,
					},
				},
				{
					NodeKey: NodeKey{
						Kind:  Group,
						Value: testGroup2,
					},
				},
			},
		}

		Convey("Testing basic edge adding.", func() {
			subgraph.addEdge(0, In, 1)
			subgraph.addEdge(2, In, 1)

			expectedSubgraph := &Subgraph{
				Nodes: []*SubgraphNode{
					{ // 0
						NodeKey: NodeKey{
							Kind:  Glob,
							Value: testGlob,
						},
						Edges: map[EdgeTag][]int{
							In: {1},
						},
					},
					{ // 1
						NodeKey: NodeKey{
							Kind:  Group,
							Value: testGroup0,
						},
					},
					{ // 2
						NodeKey: NodeKey{
							Kind:  Identity,
							Value: testUser,
						},
						Edges: map[EdgeTag][]int{
							In: {1},
						},
					},
					{ // 3
						NodeKey: NodeKey{
							Kind:  Group,
							Value: testGroup1,
						},
					},
					{ // 4
						NodeKey: NodeKey{
							Kind:  Group,
							Value: testGroup2,
						},
					},
				},
			}
			So(subgraph.Nodes, ShouldResemble, expectedSubgraph.Nodes)
		})

		// Make sure that the order that of the edges stays consistent and is predictable.
		Convey("Testing stability.", func() {
			subgraph.addEdge(0, In, 4)
			subgraph.addEdge(0, In, 2)
			subgraph.addEdge(0, In, 3)
			subgraph.addEdge(2, In, 4)
			subgraph.addEdge(2, In, 3)
			subgraph.addEdge(2, In, 0)
			subgraph.addEdge(2, In, 1)
			expectedSubgraph := &Subgraph{
				Nodes: []*SubgraphNode{
					{ // 0
						NodeKey: NodeKey{
							Kind:  Glob,
							Value: testGlob,
						},
						Edges: map[EdgeTag][]int{
							In: {2, 3, 4},
						},
					},
					{ // 1
						NodeKey: NodeKey{
							Kind:  Group,
							Value: testGroup0,
						},
					},
					{ // 2
						NodeKey: NodeKey{
							Kind:  Identity,
							Value: testUser,
						},
						Edges: map[EdgeTag][]int{
							In: {0, 1, 3, 4},
						},
					},
					{ // 3
						NodeKey: NodeKey{
							Kind:  Group,
							Value: testGroup1,
						},
					},
					{ // 4
						NodeKey: NodeKey{
							Kind:  Group,
							Value: testGroup2,
						},
					},
				},
			}
			So(subgraph.Nodes, ShouldResemble, expectedSubgraph.Nodes)
		})
	})
}
