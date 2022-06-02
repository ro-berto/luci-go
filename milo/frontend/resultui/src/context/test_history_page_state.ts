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

import { interpolateOranges, scaleLinear, scaleSequential } from 'd3';
import { DateTime } from 'luxon';
import { autorun, comparer, computed, observable, reaction } from 'mobx';

import { createContextLink } from '../libs/context';
import { parseVariantFilter, parseVariantPredicate } from '../libs/queries/th_filter_query';
import { TestHistoryEntriesLoader } from '../models/test_history_entries_loader';
import { TestHistoryStatsLoader } from '../models/test_history_stats_loader';
import { VariantLoader } from '../models/variant_loader';
import { getCriticalVariantKeys, ResultDb } from '../services/resultdb';
import {
  QueryTestHistoryStatsResponseGroup,
  TestHistoryService,
  TestVerdictBundle,
  TestVerdictStatus,
  Variant,
  VariantPredicate,
} from '../services/weetbix';

export const enum GraphType {
  STATUS = 'STATUS',
  DURATION = 'DURATION',
}

export const enum XAxisType {
  DATE = 'DATE',
  COMMIT = 'COMMIT',
}

// Use SCALE_COLOR to discard colors avoid using white color when the input is
// close to 0.
const SCALE_COLOR = scaleLinear().range([0.1, 1]).domain([0, 1]);

/**
 * Records the test history page state.
 */
export class TestHistoryPageState {
  readonly project: string;
  readonly subRealm: string;

  readonly latestDate = DateTime.now().toUTC().startOf('day');
  @observable.ref days = 14;

  @observable.ref filterText = '';
  @observable.ref private variantFilter = (_v: Variant, _hash: string) => true;
  @observable.struct private variantPredicate: VariantPredicate = { contains: { def: {} } };

  @computed({ keepAlive: true })
  get statsLoader() {
    if (this.isDisposed) {
      return null;
    }

    return new TestHistoryStatsLoader(
      this.project,
      this.subRealm,
      this.testId,
      this.latestDate,
      this.variantPredicate,
      this.testHistoryService
    );
  }

  @computed({ keepAlive: true })
  get variantLoader() {
    if (this.isDisposed) {
      return null;
    }
    return new VariantLoader(this.project, this.subRealm, this.testId, this.variantPredicate, this.testHistoryService);
  }

  @computed get filteredVariants() {
    return this.variantLoader!.variants.filter(([hash, v]) => this.variantFilter(v, hash));
  }

  @observable.ref private discoverVariantReqCount = 0;
  @computed get isDiscoveringVariants() {
    return this.discoverVariantReqCount > 0;
  }
  @observable.ref private _loadedAllVariants = false;
  @computed get loadedAllVariants() {
    return this._loadedAllVariants;
  }

  @observable.ref graphType = GraphType.STATUS;
  @observable.ref xAxisType = XAxisType.DATE;

  @observable.ref countUnexpected = true;
  @observable.ref countUnexpectedlySkipped = true;
  @observable.ref countFlaky = true;

  // Keep track of the max and min duration to render the duration graph.
  @observable.ref durationInitialized = false;
  @observable.ref maxDurationMs = 100;
  @observable.ref minDurationMs = 0;

  resetDurations() {
    this.durationInitialized = false;
    this.maxDurationMs = 100;
    this.minDurationMs = 0;
  }

  @computed get scaleDurationColor() {
    return scaleSequential((x) => interpolateOranges(SCALE_COLOR(x))).domain([this.minDurationMs, this.maxDurationMs]);
  }

  @computed({ equals: comparer.shallow })
  get criticalVariantKeys(): readonly string[] {
    return getCriticalVariantKeys(this.variantLoader!.variants.map(([_, v]) => v));
  }

  @observable.ref private customColumnKeys?: readonly string[];
  @computed get defaultColumnKeys(): readonly string[] {
    return this.criticalVariantKeys.map((k) => 'v.' + k);
  }
  setColumnKeys(v: readonly string[]): void {
    this.customColumnKeys = v;
  }
  @computed({ equals: comparer.shallow }) get columnKeys(): readonly string[] {
    return this.customColumnKeys || this.defaultColumnKeys;
  }

  @observable.ref private customColumnWidths: { readonly [key: string]: number } = {};
  @computed get columnWidths() {
    return this.columnKeys.map((col) => this.customColumnWidths[col] ?? 100);
  }
  setColumnWidths(v: { readonly [key: string]: number }): void {
    this.customColumnWidths = v;
  }

  @observable.ref private customSortingKeys?: readonly string[];
  readonly defaultSortingKeys: readonly string[] = [];
  @computed get sortingKeys(): readonly string[] {
    return this.customSortingKeys || this.defaultSortingKeys;
  }
  setSortingKeys(v: readonly string[]): void {
    this.customSortingKeys = v;
  }

  @observable.ref isDisposed = false;
  private disposers: Array<() => void> = [];
  constructor(
    readonly realm: string,
    readonly testId: string,
    readonly testHistoryService: TestHistoryService,
    readonly resultDb: ResultDb
  ) {
    [this.project, this.subRealm] = realm.split(':', 2);

    // Ensure the first page of test history entry details are loaded / being
    // loaded.
    this.disposers.push(
      reaction(
        () => this.entriesLoader,
        () => this.entriesLoader?.loadFirstPage(),
        { fireImmediately: true }
      )
    );

    // Ensure the first page of variants are loaded / being loaded.
    this.disposers.push(
      reaction(
        () => this.variantLoader,
        () => this.discoverVariants(),
        { fireImmediately: true }
      )
    );

    // Keep the filters in sync.
    this.disposers.push(
      autorun(() => {
        try {
          const newVariantFilter = parseVariantFilter(this.filterText);
          const newVariantPredicate = parseVariantPredicate(this.filterText);

          // Only update the filters after the query is successfully parsed.
          this.variantFilter = newVariantFilter;
          this.variantPredicate = newVariantPredicate;
        } catch (e) {
          //TODO(weiweilin): display the error to the user.
          console.error(e);
        }
      })
    );
  }

  async discoverVariants() {
    this.discoverVariantReqCount++;
    const req = this.variantLoader!.discoverVariants();
    req.finally(() => this.discoverVariantReqCount--);
    this._loadedAllVariants = await req;
    return this._loadedAllVariants;
  }

  @computed({ keepAlive: true })
  private get entriesLoader() {
    if (this.isDisposed || !this.selectedGroup) {
      return null;
    }
    const [project, subRealm] = this.realm.split(':', 2);
    return new TestHistoryEntriesLoader(
      project,
      subRealm,
      this.testId,
      DateTime.fromISO(this.selectedGroup.partitionTime),
      this.variantLoader!.getVariant(this.selectedGroup.variantHash)!,
      this.testHistoryService
    );
  }

  @observable.ref selectedGroup: QueryTestHistoryStatsResponseGroup | null = null;

  @computed get verdictBundles(): ReadonlyArray<TestVerdictBundle> {
    if (!this.entriesLoader?.verdictBundles.length) {
      return [];
    }

    const cmpFn = createTVCmpFn(this.sortingKeys);
    return [...this.entriesLoader?.verdictBundles].sort(cmpFn);
  }

  @computed
  get loadedTestVerdictCount() {
    return this.entriesLoader?.verdictBundles.length || 0;
  }
  @computed
  get selectedTestVerdictCount() {
    return (
      (this.selectedGroup?.unexpectedCount || 0) +
      (this.selectedGroup?.unexpectedlySkippedCount || 0) +
      (this.selectedGroup?.flakyCount || 0) +
      (this.selectedGroup?.exoneratedCount || 0) +
      (this.selectedGroup?.expectedCount || 0)
    );
  }

  @computed get isLoading() {
    return this.entriesLoader?.isLoading ?? false;
  }
  get loadedAllTestVerdicts() {
    return this.entriesLoader?.loadedAllTestVerdicts ?? true;
  }
  get loadedFirstPage() {
    return this.entriesLoader?.loadedFirstPage ?? true;
  }

  async loadFirstPage() {
    await this.entriesLoader?.loadFirstPage();
  }
  async loadNextPage() {
    await this.entriesLoader?.loadNextPage();
  }

  // Don't display history URL when the user is already on the history page.
  getHistoryUrl() {
    return '';
  }

  /**
   * Perform cleanup.
   * Must be called before the object is GCed.
   */
  dispose() {
    this.isDisposed = true;
    for (const disposer of this.disposers) {
      disposer();
    }

    // Evaluates @computed({keepAlive: true}) properties after this.isDisposed
    // is set to true so they no longer subscribes to any external observable.
    this.entriesLoader;
  }
}

export const [provideTestHistoryPageState, consumeTestHistoryPageState] = createContextLink<TestHistoryPageState>();

// Note: once we have more than 9 statuses, we need to add '0' prefix so '10'
// won't appear before '2' after sorting.
export const TEST_VERDICT_STATUS_CMP_STRING = {
  [TestVerdictStatus.TEST_VERDICT_STATUS_UNSPECIFIED]: '0',
  [TestVerdictStatus.UNEXPECTED]: '1',
  [TestVerdictStatus.UNEXPECTEDLY_SKIPPED]: '2',
  [TestVerdictStatus.FLAKY]: '3',
  [TestVerdictStatus.EXONERATED]: '4',
  [TestVerdictStatus.EXPECTED]: '5',
};
/**
 * Create a test variant compare function for the given sorting key list.
 *
 * A sorting key must be one of the following:
 * 1. '{property_key}': sort by property_key in ascending order.
 * 2. '-{property_key}': sort by property_key in descending order.
 */
export function createTVCmpFn(
  sortingKeys: readonly string[]
): (v1: TestVerdictBundle, v2: TestVerdictBundle) => number {
  const sorters: Array<[number, (v: TestVerdictBundle) => { toString(): string }]> = sortingKeys.map((key) => {
    const [mul, propKey] = key.startsWith('-') ? [-1, key.slice(1)] : [1, key];
    const propGetter = createTVPropGetter(propKey);

    // Status should be be sorted by their significance not by their string
    // representation.
    if (propKey.toLowerCase() === 'status') {
      return [mul, (v) => TEST_VERDICT_STATUS_CMP_STRING[propGetter(v) as TestVerdictStatus]];
    }
    return [mul, propGetter];
  });
  return (v1, v2) => {
    for (const [mul, propGetter] of sorters) {
      const cmp = propGetter(v1).toString().localeCompare(propGetter(v2).toString()) * mul;
      if (cmp !== 0) {
        return cmp;
      }
    }
    return 0;
  };
}

/**
 * Create a test verdict property getter for the given property key.
 *
 * A property key must be one of the following:
 * 1. 'status': status of the test verdict.
 * 2. 'partitionTime': partition time of the test verdict.
 * 3. 'v.{variant_key}': def[variant_key] of associated variant of the test
 * verdict (e.g. v.gpu).
 */
export function createTVPropGetter(propKey: string): (v: TestVerdictBundle) => ToString {
  if (propKey.match(/^v[.]/i)) {
    const variantKey = propKey.slice(2);
    return ({ variant }) => variant.def[variantKey] || '';
  }
  propKey = propKey.toLowerCase();
  switch (propKey) {
    case 'status':
      return ({ verdict }) => verdict.status;
    case 'patitiontime':
      return ({ verdict }) => verdict.partitionTime;
    default:
      console.warn('invalid property key', propKey);
      return () => '';
  }
}
