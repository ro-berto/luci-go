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

package bq

import (
	"context"
	"sync"

	"google.golang.org/protobuf/proto"
)

// Fake is a fake BQ client for tests.
type Fake struct {
	// mu protects access/mutation to this Fake.
	mu sync.RWMutex
	// sent is a map of "project.dataset.table" keys to slices of sent rows.
	sent map[string][]proto.Message
}

// SendRow provides a mock SendRow implementation for tests.
func (f *Fake) SendRow(ctx context.Context, row Row) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	key := row.CloudProject + "." + row.Dataset + "." + row.Table
	if f.sent == nil {
		f.sent = make(map[string][]proto.Message)
	}
	f.sent[key] = append(f.sent[key], row.Payload)
	return nil
}

// Rows returns the stored rows for a given dataset and table.
//
// cloudProject can be empty, implying the same cloud project as the running
// code.
func (f *Fake) Rows(cloudProject, dataset, table string) []proto.Message {
	f.mu.RLock()
	defer f.mu.RUnlock()
	rows := f.sent[cloudProject+"."+dataset+"."+table]
	ret := make([]proto.Message, len(rows))
	copy(ret, rows)
	return ret
}

// RowsCount returns the number of stored rows for a given dataset and table.
//
// cloudProject can be empty, implying the same cloud project as the running
// code.
func (f *Fake) RowsCount(cloudProject, dataset, table string) int {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return len(f.sent[cloudProject+"."+dataset+"."+table])
}

// TotalSent returns total number of all rows sent across all destinations.
func (f *Fake) TotalSent() int {
	f.mu.RLock()
	defer f.mu.RUnlock()
	cnt := 0
	for _, rows := range f.sent {
		cnt += len(rows)
	}
	return cnt
}

// Ensure that Fake implement the Client interface.
var _ Client = (*Fake)(nil)
