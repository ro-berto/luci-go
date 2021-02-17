// Copyright 2020 The LUCI Authors.
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

/**
 * @fileoverview This file contains functions/classes that helps loading test
 * results and exonerations from resultDb to a TestNode.
 */

import { action, computed, observable } from 'mobx';

import { QueryTestVariantsRequest, TestVariant, TestVariantStatus, UISpecificService } from '../services/resultdb';

/**
 * The stage of the next test variant. The stage can be
 * 1. LoadingXXX: the status of the next test variant will be no worse than XXX.
 * 2. Done: all test variants have been loaded.
 */
export const enum LoadingStage {
  LoadingUnexpected = 0,
  LoadingUnexpectedlySkipped = 1,
  LoadingFlaky = 2,
  LoadingExonerated = 3,
  LoadingExpected = 4,
  Done = 5,
}

const VARIANT_STATUS_LOADING_STAGE_MAP = Object.freeze({
  [TestVariantStatus.TEST_VARIANT_STATUS_UNSPECIFIED]: LoadingStage.LoadingUnexpected,
  [TestVariantStatus.UNEXPECTED]: LoadingStage.LoadingUnexpected,
  [TestVariantStatus.UNEXPECTEDLY_SKIPPED]: LoadingStage.LoadingUnexpectedlySkipped,
  [TestVariantStatus.FLAKY]: LoadingStage.LoadingFlaky,
  [TestVariantStatus.EXONERATED]: LoadingStage.LoadingExonerated,
  [TestVariantStatus.EXPECTED]: LoadingStage.LoadingExpected,
});

/**
 * Keeps the progress of the iterator and loads tests into the test node on
 * request.
 */
export class TestLoader {
  @computed get isLoading() { return !this.loadedAllVariants && this.loadingReqCount !== 0; }
  @observable.ref private loadingReqCount = 0;

  @computed get firstRequestSent() { return this._firstRequestSent; }
  @observable.ref private _firstRequestSent = false;

  /**
   * The queryTestVariants RPC sorted the variant by status. We can use this to
   * tell the possible status of the next test variants and therefore avoid
   * unnecessary loading.
   */
  @computed get stage() { return this._stage; }
  @observable.ref private _stage = LoadingStage.LoadingUnexpected;

  @observable.shallow readonly unexpectedTestVariants: TestVariant[] = [];
  @observable.shallow readonly unexpectedlySkippedTestVariants: TestVariant[] = [];
  @observable.shallow readonly flakyTestVariants: TestVariant[] = [];
  @observable.shallow readonly exoneratedTestVariants: TestVariant[] = [];
  @observable.shallow readonly expectedTestVariants: TestVariant[] = [];

  @computed get loadedAllVariants() { return this.stage === LoadingStage.Done; }
  @computed get loadedAllUnexpectedVariants() { return this.stage > LoadingStage.LoadingUnexpected; }
  @computed get loadedAllUnexpectedlySkippedVariants() { return this.stage > LoadingStage.LoadingUnexpectedlySkipped; }
  @computed get loadedAllFlakyVariants() { return this.stage > LoadingStage.LoadingFlaky; }
  @computed get loadedAllExoneratedVariants() { return this.stage > LoadingStage.LoadingExonerated; }
  @computed get loadedAllExpectedVariants() { return this.stage > LoadingStage.LoadingExpected; }
  @computed get firstPageLoaded() {
    return this.unexpectedTestVariants.length > 0 || this.loadedAllUnexpectedVariants;
  }
  @computed get firstPageIsEmpty() {
    return this.loadedAllUnexpectedVariants &&
      this.unexpectedTestVariants.length === 0 &&
      this.unexpectedlySkippedTestVariants.length === 0 &&
      this.flakyTestVariants.length === 0 &&
      this.exoneratedTestVariants.length === 0;
  }

  // undefined means the end has been reached.
  // empty string is the token for the first page.
  private nextPageToken: string | undefined = '';

  constructor(
    private readonly req: QueryTestVariantsRequest,
    private readonly uiSpecificService: UISpecificService,
  ) {}

  private loadPromise = Promise.resolve();

  /**
   * Loads the next batch of tests from the iterator to the node.
   */
  loadNextPage() {
    this._firstRequestSent = true;
    if (this.stage === LoadingStage.Done) {
      return this.loadPromise;
    }
    this.loadingReqCount++;
    this.loadPromise = this.loadPromise.then(() => this.loadNextPageInternal());
    return this.loadPromise.then(() => this.loadingReqCount--);
  }

  /**
   * Loads pages repeatedly until we receive some variants with the given variant status.
   *
   * Will always load at least one page.
   */
  async loadPagesUntilStatus(status: TestVariantStatus) {
    if (this.stage === VARIANT_STATUS_LOADING_STAGE_MAP[status]) {
      // If we expect the next batch to be at least the status we want, load one page only.
      await this.loadNextPage();
      return;
    }

    // Load pages until the next expected status is at least the one we're after.
    do {
      await this.loadNextPage();
    } while (this.stage < VARIANT_STATUS_LOADING_STAGE_MAP[status]);

    // If we wanted to load up to Expected, and none have arrived yet
    // (i.e. we're at the point where ResultDB has just returned the final
    // non-expected variants to us), load one more page.
    if (status === TestVariantStatus.EXPECTED && this.expectedTestVariants.length === 0) {
      await this.loadNextPage();
    }
  }

  /**
   * Loads the next batch of tests from the iterator to the node.
   *
   * @precondition there should not exist a running instance of
   * this.loadMoreInternal
   */
  private async loadNextPageInternal() {
    if (this.nextPageToken === undefined) {
      return;
    }

    const res = await this.uiSpecificService
      .queryTestVariants({...this.req, pageToken: this.nextPageToken});
    this.nextPageToken = res.nextPageToken;

    const testVariants = res.testVariants || [];
    this.processTestVariants(testVariants);
    if (this.nextPageToken === undefined) {
      this._stage = LoadingStage.Done;
      return;
    }
    if (testVariants.length < (this.req.pageSize || 1000)) {
      // When the service returns an incomplete page and nextPageToken is not
      // undefined, the following pages must be expected test variants.
      // Without this special case, the UI may incorrectly indicate that not all
      // variants have been loaded for statuses worse than Expected.
      this._stage = LoadingStage.LoadingExpected;
      return;
    }
  }

  @action
  private processTestVariants(testVariants: readonly TestVariant[]) {
    for (const testVariant of testVariants) {
      switch (testVariant.status) {
        case TestVariantStatus.UNEXPECTED:
          this._stage = LoadingStage.LoadingUnexpected;
          this.unexpectedTestVariants.push(testVariant);
          break;
        case TestVariantStatus.UNEXPECTEDLY_SKIPPED:
          this._stage = LoadingStage.LoadingUnexpectedlySkipped;
          this.unexpectedlySkippedTestVariants.push(testVariant);
          break;
        case TestVariantStatus.FLAKY:
          this._stage = LoadingStage.LoadingFlaky;
          this.flakyTestVariants.push(testVariant);
          break;
        case TestVariantStatus.EXONERATED:
          this._stage = LoadingStage.LoadingExonerated;
          this.exoneratedTestVariants.push(testVariant);
          break;
        case TestVariantStatus.EXPECTED:
          this._stage = LoadingStage.LoadingExpected;
          this.expectedTestVariants.push(testVariant);
          break;
        default:
          break;
      }
    }
  }
}
