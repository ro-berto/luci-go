// Copyright 2015 The LUCI Authors.
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

package deps

import (
	"context"
	"fmt"

	dm "go.chromium.org/luci/dm/api/service/v1"
	"go.chromium.org/luci/dm/appengine/distributor/fake"
	"go.chromium.org/luci/dm/appengine/model"
	"go.chromium.org/luci/dm/appengine/mutate"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"
	"go.chromium.org/luci/tumble"
)

func testSetup() (ttest *tumble.Testing, c context.Context, dist *fake.Distributor, s testDepsServer) {
	ttest, c, dist = fake.Setup(mutate.FinishExecutionFn)
	s = testDepsServer{newDecoratedDeps()}
	return
}

func reader(c context.Context) context.Context {
	return auth.WithState(c, &authtest.FakeState{
		Identity:       "test@example.com",
		IdentityGroups: []string{"reader_group"},
	})
}

func writer(c context.Context) context.Context {
	return auth.WithState(c, &authtest.FakeState{
		Identity:       "test@example.com",
		IdentityGroups: []string{"reader_group", "writer_group"},
	})
}

type testDepsServer struct {
	dm.DepsServer
}

func (s testDepsServer) ensureQuest(c context.Context, name string, aids ...uint32) string {
	desc := fake.QuestDesc(name)
	q := model.NewQuest(c, desc)
	qsts, err := s.EnsureGraphData(writer(c), &dm.EnsureGraphDataReq{
		Quest:        []*dm.Quest_Desc{desc},
		QuestAttempt: []*dm.AttemptList_Nums{{Nums: aids}},
	})
	if err != nil {
		panic(err)
	}
	for qid := range qsts.Result.Quests {
		if qid != q.ID {
			panic(fmt.Errorf("non matching quest ID!? got %q, expected %q", qid, q.ID))
		}
		return qid
	}
	panic("impossible")
}
