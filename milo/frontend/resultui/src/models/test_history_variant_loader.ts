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

import { DateTime } from 'luxon';
import { comparer, computed, observable, untracked } from 'mobx';

import { Variant } from '../services/resultdb';
import { TestHistoryService, TestVariantHistoryEntry } from '../services/test_history_service';

/**
 * Test history loader for a specific variant.
 */
export class TestHistoryVariantLoader {
  /**
   * datetime str -> test variant history entries
   */
  private readonly cache = new Map<string, TestVariantHistoryEntry[]>();

  private worker: AsyncIterableIterator<null>;

  /**
   * Test histories created after `loadedTime` are all loaded.
   * Test histories created before `loadedTime` are yet to be loaded.
   */
  // Initialize to a future date. All the entries created after 1 year in the
  // future can be considered loaded because we know they don't exist.
  @observable.ref private loadedTime = DateTime.now().plus({ years: 1 });
  @computed private get loadedTimeGroupId() {
    return this.resolve(this.loadedTime);
  }

  /**
   * When `this.worker.next()` is called, it will keep loading until
   * 1. `loadedTime` < `targetTime`, and
   * 2. `loadedTime` and `targetTime` resolves to different strings.
   */
  @observable.ref private targetTime = this.loadedTime;
  @computed private get targetTimeGroupId() {
    return this.resolve(this.targetTime);
  }

  constructor(
    readonly realm: string,
    readonly testId: string,
    readonly variant: Variant,
    /**
     * Resolve controls the size of the time step when grouping and querying
     * test history entries. For example, if all timestamps between
     * [2021-11-05T00:00:00Z, 2021-11-06T00:00:00Z) resolves to '2021-11-05',
     * All test history entries in that time range will be grouped together.
     * They will all be returned when `getEntries` is called with a timestamp
     * between [2021-11-05T00:00:00Z, 2021-11-06T00:00:00Z).
     *
     * Note: if time1 and time2 both resolve to the same string, any time
     * between time1 and time2 must resolves to the same string.
     */
    readonly resolve: (time: DateTime) => string,
    readonly testHistoryService: TestHistoryService
  ) {
    this.worker = this.workerGen();
  }

  /**
   * Generates a worker that loads the entries between
   * [`now`, `this.targetTime`] then yields back.
   *
   * `this.targetTime` can be updated so the worker can load the entries between
   * [`last target time`, `this.targetTime`] when `.next()` is called.
   */
  private async *workerGen() {
    let pageToken = '';
    for (;;) {
      // We've loaded all required entries. Yield back.
      while (this.targetTime > this.loadedTime && this.targetTimeGroupId !== this.loadedTimeGroupId) {
        yield null;
      }

      const res = await this.testHistoryService.queryTestHistory({
        realm: this.realm,
        testId: this.testId,
        variantPredicate: { equals: this.variant },
        timeRange: {},
        pageToken: pageToken,
        // TODO(weiweilin): the RPC is currently implemented in the frontend.
        // Use a small page size so we don't need to wait for several
        // GetTestResultHistory calls to gather 1 page. Adjust the value once
        // we have the actual RPC implemented on the server side.
        pageSize: 5,
      });

      for (const entry of res.entries) {
        this.addEntry(entry);
      }

      if (!res.nextPageToken) {
        // We've loaded all the entries. Set the loaded time to the earliest
        // possible time.
        this.loadedTime = DateTime.fromMillis(0);
        return;
      }

      pageToken = res.nextPageToken;
    }
  }

  private earliestEntryIdentifiers = new Set<string>();

  /**
   * Adds the entry to the cache. If the entry was already added, ignore it.
   */
  private addEntry(entry: TestVariantHistoryEntry) {
    const entryTime = DateTime.fromISO(entry.invocationTimestamp);

    // Compare the timestamp and see if the entry was already loaded.
    if (entryTime > this.loadedTime) {
      return;
    }

    // Join invocations IDs as entry IDs. Once the improve test history RPC is
    // implemented, we could use the root invocation ID instead.
    const entryId = entry.invocationIds.join(' ');

    // Compare the entry identifier if we can't tell whether the entry was
    // already loaded from the timestamp alone.
    if (entryTime.toMillis() === this.loadedTime.toMillis()) {
      if (this.earliestEntryIdentifiers.has(entryId)) {
        return;
      }
      this.earliestEntryIdentifiers.add(entryId);
    } else {
      this.earliestEntryIdentifiers = new Set([entryId]);
      this.loadedTime = entryTime;
    }

    let dateCache = this.cache.get(this.loadedTimeGroupId);
    if (!dateCache) {
      dateCache = [];
      this.cache.set(this.loadedTimeGroupId, dateCache);
    }

    dateCache.push(entry);
  }

  /**
   * Load all entries that were created after `time`.
   */
  async loadUntil(time: DateTime) {
    if (time >= this.targetTime) {
      return;
    }
    this.targetTime = time;
    await this.worker.next();
    return;
  }

  /**
   * Get all entries associated with the time slot.
   *
   * If there are no entries associated with time slot, return an empty array.
   * If the entries associated with the time slot hasn't been loaded yet, return
   * null.
   */
  getEntries(time: DateTime, noLoading = false): readonly TestVariantHistoryEntry[] | null {
    const timeStr = this.resolve(time);
    const loaded = computed(() => time > this.loadedTime && this.loadedTimeGroupId !== timeStr).get();
    if (!loaded) {
      if (!noLoading) {
        untracked(() => this.loadUntil(time));
      }
      return null;
    }
    const dateStr = this.resolve(time);
    const ret = computed(() => this.cache.get(dateStr) || [], { equals: comparer.shallow }).get();
    return ret;
  }
}