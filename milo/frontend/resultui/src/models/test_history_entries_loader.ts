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
import { computed, observable } from 'mobx';

import { ResultDb, TestVariant } from '../services/resultdb';
import { TestHistoryService, Variant } from '../services/weetbix';

/**
 * A utility class that helps loading test history entry details.
 */
export class TestHistoryEntriesLoader {
  constructor(
    readonly project: string,
    readonly subRealm: string,
    readonly testId: string,
    readonly date: DateTime,
    readonly variant: Variant,
    readonly testHistoryService: TestHistoryService,
    readonly resultDb: ResultDb,
    readonly pageSize = 10
  ) {}

  @observable.shallow private _testVariants: TestVariant[] = [];
  @computed get testVariants(): readonly TestVariant[] {
    return this._testVariants;
  }

  private loadPromise = Promise.resolve();
  private firstLoadPromise?: Promise<void>;

  @observable.ref private loadingReqCount = 0;
  get isLoading() {
    return this.loadingReqCount !== 0;
  }
  @computed get loadedAllTestVariants() {
    return this.pageToken === null;
  }
  @computed get loadedFirstPage() {
    return this._testVariants.length > 0;
  }

  private pageToken: string | null = '';
  private readonly historyReq = {
    project: this.project,
    testId: this.testId,
    predicate: {
      subRealm: this.subRealm,
      variantPredicate: {
        equals: this.variant,
      },
      partitionTimeRange: {
        earliest: this.date.toISO(),
        latest: this.date.minus({ days: -1 }).toISO(),
      },
    },
    pageSize: this.pageSize,
  };

  /**
   * Loads the next batch of tests.
   *
   * @precondition there should not exist a running instance of
   * this.loadNextPage
   */
  private async loadNextPageImpl() {
    if (this.pageToken === null) {
      return;
    }

    const historyRes = await this.testHistoryService.query({
      ...this.historyReq,
      pageToken: this.pageToken,
    });

    const verdicts = historyRes.verdicts || [];

    const variants = verdicts.map(async (v) => {
      const variantRes = await this.resultDb.batchGetTestVariants({
        invocation: 'invocations/' + v.invocationId,
        testVariants: [
          {
            testId: v.testId,
            variantHash: v.variantHash,
          },
        ],
        resultLimit: 100,
      });
      return {
        ...variantRes.testVariants![0],
        partitionTime: v.partitionTime,
      };
    });

    this._testVariants.push(...(await Promise.all(variants)));

    this.pageToken = historyRes.nextPageToken || null;
  }
  // Don't mark as async so loadingReqCount and firstLoadPromise can be updated
  // immediately.
  loadNextPage(): Promise<void> {
    if (this.loadedAllTestVariants) {
      return this.loadPromise;
    }

    this.loadingReqCount++;
    this.loadPromise = this.loadPromise
      .then(() => this.loadNextPageImpl())
      .then(() => {
        this.loadingReqCount--;
      });
    if (!this.firstLoadPromise) {
      this.firstLoadPromise = this.loadPromise;
    }

    return this.loadPromise;
  }

  loadFirstPage(): Promise<void> {
    return this.firstLoadPromise || this.loadNextPage();
  }
}
