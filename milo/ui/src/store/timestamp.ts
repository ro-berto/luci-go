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

import { DateTime } from 'luxon';
import { Instance, SnapshotIn, SnapshotOut, types } from 'mobx-state-tree';

export const Timestamp = types
  .model('Timestamp', {
    id: types.optional(types.identifierNumber, () => Math.random()),
    value: types.optional(types.number, () => Date.now()),
  })
  .views((self) => ({
    get dateTime() {
      return DateTime.fromMillis(self.value);
    },
  }))
  .actions((self) => ({
    refresh() {
      self.value = Date.now();
    },
  }));

export type TimestampInstance = Instance<typeof Timestamp>;
export type TimestampSnapshotIn = SnapshotIn<typeof Timestamp>;
export type TimestampSnapshotOut = SnapshotOut<typeof Timestamp>;
