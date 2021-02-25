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

package dsset

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"go.chromium.org/luci/gae/filter/txndefer"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/data/rand/mathrand"
	"go.chromium.org/luci/common/data/stringset"

	. "github.com/smartystreets/goconvey/convey"
)

func testingContext() context.Context {
	c := txndefer.FilterRDS(memory.Use(context.Background()))
	datastore.GetTestable(c).AutoIndex(true)
	datastore.GetTestable(c).Consistent(true)
	c = clock.Set(c, testclock.New(time.Unix(1442270520, 0).UTC()))
	c = mathrand.Set(c, rand.New(rand.NewSource(1000)))
	return c
}

// pop pops a bunch of items from the set and returns items that were popped.
func pop(c context.Context, s *Set, listing *Listing, ids []string) (popped []string, err error) {
	op, err := s.BeginPop(c, listing)
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		if op.Pop(id) {
			popped = append(popped, id)
		}
	}
	if err = FinishPop(c, op); err != nil {
		return nil, err
	}
	return popped, nil
}

func TestSet(t *testing.T) {
	t.Parallel()

	Convey("item one lifecycle", t, func() {
		c := testingContext()

		set := Set{
			Parent:          datastore.NewKey(c, "Parent", "parent", 0, nil),
			TombstonesDelay: time.Minute,
		}

		// Add one item.
		So(set.Add(c, []Item{{ID: "abc"}}), ShouldBeNil)

		// The item is returned by the listing.
		listing, err := set.List(c)
		So(err, ShouldBeNil)
		So(listing.Items, ShouldResemble, []Item{{ID: "abc"}})
		So(listing.Garbage, ShouldBeNil)

		// Pop it!
		err = datastore.RunInTransaction(c, func(c context.Context) error {
			popped, err := pop(c, &set, listing, []string{"abc"})
			So(err, ShouldBeNil)
			So(popped, ShouldResemble, []string{"abc"})
			return nil
		}, nil)
		So(err, ShouldBeNil)

		// The listing no longer returns it.
		listing, err = set.List(c)
		So(err, ShouldBeNil)
		So(listing.Items, ShouldBeNil)

		// The listing no longer returns the item, and there's no tombstones to
		// cleanup.
		listing, err = set.List(c)
		So(err, ShouldBeNil)
		So(listing.Items, ShouldBeNil)
		So(listing.Garbage, ShouldBeNil)

		// Attempt to add it back (should be ignored).
		So(set.Add(c, []Item{{ID: "abc"}}), ShouldBeNil)

		// The listing still doesn't return it, but we now have a tombstone to
		// cleanup (again).
		listing, err = set.List(c)
		So(err, ShouldBeNil)
		So(listing.Items, ShouldBeNil)
		So(len(listing.Garbage), ShouldEqual, 1)
		So(listing.Garbage[0].old, ShouldBeFalse)
		So(listing.Garbage[0].storage, ShouldNotBeNil)

		// Popping it again doesn't work either.
		err = datastore.RunInTransaction(c, func(c context.Context) error {
			popped, err := pop(c, &set, listing, []string{"abc"})
			So(err, ShouldBeNil)
			So(popped, ShouldBeNil)
			return nil
		}, nil)
		So(err, ShouldBeNil)

		// Cleaning up the storage, again. This should make List stop returning
		// the tombstone (since it has no storage items associated with it and it's
		// not ready to be evicted yet).
		So(CleanupGarbage(c, listing.Garbage), ShouldBeNil)
		listing, err = set.List(c)
		So(err, ShouldBeNil)
		So(listing.Items, ShouldBeNil)
		So(listing.Garbage, ShouldBeNil)

		// Time passes, tombstone expires.
		clock.Get(c).(testclock.TestClock).Add(2 * time.Minute)

		// Listing now returns expired tombstone.
		listing, err = set.List(c)
		So(err, ShouldBeNil)
		So(listing.Items, ShouldBeNil)
		So(len(listing.Garbage), ShouldEqual, 1)
		So(listing.Garbage[0].storage, ShouldBeNil) // cleaned already

		// Cleanup storage keys.
		So(CleanupGarbage(c, listing.Garbage), ShouldBeNil)

		// Cleanup the tombstones themselves.
		err = datastore.RunInTransaction(c, func(c context.Context) error {
			popped, err := pop(c, &set, listing, nil)
			So(err, ShouldBeNil)
			So(popped, ShouldBeNil)
			return nil
		}, nil)
		So(err, ShouldBeNil)

		// No tombstones returned any longer.
		listing, err = set.List(c)
		So(err, ShouldBeNil)
		So(listing.Items, ShouldBeNil)
		So(listing.Garbage, ShouldBeNil)

		// And the item can be added back now, since no trace of it is left.
		So(set.Add(c, []Item{{ID: "abc"}}), ShouldBeNil)

		// Yep, it is there.
		listing, err = set.List(c)
		So(err, ShouldBeNil)
		So(listing.Items, ShouldResemble, []Item{{ID: "abc"}})
		So(listing.Garbage, ShouldBeNil)
	})

	Convey("delete items non-transactionally", t, func() {
		c := testingContext()

		set := Set{
			Parent:          datastore.MakeKey(c, "Parent", "parent"),
			TombstonesDelay: time.Minute,
		}

		// Add 3 items.
		So(set.Add(c, []Item{{ID: "abc"}}), ShouldBeNil)
		So(set.Add(c, []Item{{ID: "def"}}), ShouldBeNil)
		So(set.Add(c, []Item{{ID: "ghi"}}), ShouldBeNil)

		l, err := set.List(c)
		So(err, ShouldBeNil)
		So(l.Items, ShouldHaveLength, 3)

		// Delete 2 items before transacting.
		i := 0
		err = set.Delete(c, func() string {
			switch i = i + 1; i {
			case 1:
				return "def"
			case 2:
				return "abc"
			default:
				return ""
			}
		})
		So(err, ShouldBeNil)

		l2, err := set.List(c)
		So(err, ShouldBeNil)
		So(l2.Items, ShouldResemble, []Item{{ID: "ghi"}})
	})
}

func TestStress(t *testing.T) {
	t.Parallel()

	Convey("stress", t, func() {
		// Add 1000 items in parallel from N goroutines, and (also in parallel),
		// run N instances of "List and pop all", collecting the result in single
		// list. There should be no duplicates in the final list!
		c := testingContext()

		set := Set{
			Parent:          datastore.MakeKey(c, "Parent", "parent"),
			TombstonesDelay: time.Minute,
		}

		producers := 3
		consumers := 5
		items := 100

		wakeups := make(chan string)

		lock := sync.Mutex{}
		var consumed []string

		for i := 0; i < producers; i++ {
			go func() {
				for j := 0; j < items; j++ {
					set.Add(c, []Item{{ID: fmt.Sprintf("%d", j)}})
					// Wake up 3 consumers, so they "fight".
					wakeups <- "wake"
					wakeups <- "wake"
					wakeups <- "wake"
				}
				for i := 0; i < consumers; i++ {
					wakeups <- "done"
				}
			}()
		}

		consume := func() {
			listing, err := set.List(c)
			if err != nil || len(listing.Items) == 0 {
				return
			}

			keys := make([]string, len(listing.Items))
			for i, itm := range listing.Items {
				keys[i] = itm.ID
			}

			// Try to pop all.
			var popped []string
			err = datastore.RunInTransaction(c, func(c context.Context) error {
				var err error
				popped, err = pop(c, &set, listing, keys)
				return err
			}, nil)

			// Consider items consumed only if transaction has landed.
			if err == nil && len(popped) != 0 {
				lock.Lock()
				consumed = append(consumed, popped...)
				lock.Unlock()
			}
		}

		wg := sync.WaitGroup{}
		wg.Add(consumers)
		for i := 0; i < consumers; i++ {
			go func() {
				defer wg.Done()
				done := false
				for !done {
					done = (<-wakeups) == "done"
					consume()
				}
			}()
		}

		wg.Wait() // this waits for completion of the entire pipeline

		// Make sure 'consumed' is the initially produced set.
		dedup := stringset.New(len(consumed))
		for _, itm := range consumed {
			dedup.Add(itm)
		}
		So(dedup.Len(), ShouldEqual, len(consumed)) // no dups
		So(len(consumed), ShouldEqual, items)       // all are accounted for
	})
}
