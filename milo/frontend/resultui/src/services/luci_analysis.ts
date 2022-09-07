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

import stableStringify from 'fast-json-stable-stringify';

import { batched, BatchOption } from '../libs/batched_fn';
import { cached, CacheOption } from '../libs/cached_fn';
import { PrpcClientExt } from '../libs/prpc_client_ext';

export interface Variant {
  readonly def: { [key: string]: string };
}

export type VariantPredicate =
  | { readonly equals: Variant }
  | { readonly contains: Variant }
  | { readonly hashEquals: string };

export const enum SubmittedFilter {
  SUBMITTED_FILTER_UNSPECIFIED = 'SUBMITTED_FILTER_UNSPECIFIED',
  ONLY_SUBMITTED = 'ONLY_SUBMITTED',
  ONLY_UNSUBMITTED = 'ONLY_UNSUBMITTED',
}

export enum TestVerdictStatus {
  TEST_VERDICT_STATUS_UNSPECIFIED = 'TEST_VERDICT_STATUS_UNSPECIFIED',
  UNEXPECTED = 'UNEXPECTED',
  UNEXPECTEDLY_SKIPPED = 'UNEXPECTEDLY_SKIPPED',
  FLAKY = 'FLAKY',
  EXONERATED = 'EXONERATED',
  EXPECTED = 'EXPECTED',
}

export interface TimeRange {
  readonly earliest?: string;
  readonly latest?: string;
}

export interface TestVerdictPredicate {
  readonly subRealm?: string;
  readonly variantPredicate?: VariantPredicate;
  readonly submittedFilter?: SubmittedFilter;
  readonly partitionTimeRange?: TimeRange;
}

export interface QueryTestHistoryRequest {
  readonly project: string;
  readonly testId: string;
  readonly predicate: TestVerdictPredicate;
  readonly pageSize?: number;
  readonly pageToken?: string;
}

export interface TestVerdict {
  readonly testId: string;
  readonly variantHash: string;
  readonly invocationId: string;
  readonly status: TestVerdictStatus;
  readonly partitionTime: string;
  readonly passedAvgDuration?: string;
}

export interface QueryTestHistoryResponse {
  readonly verdicts?: readonly TestVerdict[];
  readonly nextPageToken?: string;
}

export interface QueryTestHistoryStatsRequest {
  readonly project: string;
  readonly testId: string;
  readonly predicate: TestVerdictPredicate;
  readonly pageSize?: number;
  readonly pageToken?: string;
}

export interface QueryTestHistoryStatsResponseGroup {
  readonly partitionTime: string;
  readonly variantHash: string;
  readonly unexpectedCount?: number;
  readonly unexpectedlySkippedCount?: number;
  readonly flakyCount?: number;
  readonly exoneratedCount?: number;
  readonly expectedCount?: number;
  readonly passedAvgDuration?: string;
}

export interface QueryTestHistoryStatsResponse {
  readonly groups?: readonly QueryTestHistoryStatsResponseGroup[];
  readonly nextPageToken?: string;
}

export interface QueryVariantsRequest {
  readonly project: string;
  readonly testId: string;
  readonly subRealm?: string;
  readonly variantPredicate?: VariantPredicate;
  readonly pageSize?: number;
  readonly pageToken?: string;
}

export interface QueryVariantsResponseVariantInfo {
  readonly variantHash: string;
  readonly variant?: Variant;
}

export interface QueryVariantsResponse {
  readonly variants?: readonly QueryVariantsResponseVariantInfo[];
  readonly nextPageToken?: string;
}

export interface TestVerdictBundle {
  readonly verdict: TestVerdict;
  readonly variant: Variant;
}

export interface FailureReason {
  readonly primaryErrorMessage: string;
}

export interface QueryTestsRequest {
  readonly project: string;
  readonly testIdSubstring: string;
  readonly subRealm?: string;
  readonly pageSize?: number;
  readonly pageToken?: string;
}

export interface QueryTestsResponse {
  readonly testIds?: string[];
  readonly nextPageToken?: string;
}

export interface ClusterRequest {
  readonly project: string;
  readonly testResults: ReadonlyArray<{
    readonly requestTag?: string;
    readonly testId: string;
    readonly failureReason?: FailureReason;
  }>;
}

export interface Cluster {
  readonly clusterId: ClusterId;
  readonly bug?: AssociatedBug;
}

export interface ClusterResponse {
  readonly clusteredTestResults: ReadonlyArray<{
    readonly requestTag?: string;
    readonly clusters: readonly Cluster[];
  }>;
  readonly clusteringVersion: ClusteringVersion;
}

export interface ClusteringVersion {
  readonly algorithmsVersion: string;
  readonly rulesVersion: string;
  readonly configVersion: string;
}

export interface ClusterId {
  readonly algorithm: string;
  readonly id: string;
}

export interface AssociatedBug {
  readonly system: string;
  readonly id: string;
  readonly linkText: string;
  readonly url: string;
}

export class TestHistoryService {
  private static SERVICE = 'luci.analysis.v1.TestHistory';

  private readonly cachedCallFn: (opt: CacheOption, method: string, message: object) => Promise<unknown>;

  constructor(client: PrpcClientExt) {
    this.cachedCallFn = cached(
      (method: string, message: object) => client.call(TestHistoryService.SERVICE, method, message),
      {
        key: (method, message) => `${method}-${stableStringify(message)}`,
      }
    );
  }

  async query(req: QueryTestHistoryRequest, cacheOpt: CacheOption = {}): Promise<QueryTestHistoryResponse> {
    return (await this.cachedCallFn(cacheOpt, 'Query', req)) as QueryTestHistoryResponse;
  }

  async queryStats(
    req: QueryTestHistoryStatsRequest,
    cacheOpt: CacheOption = {}
  ): Promise<QueryTestHistoryStatsResponse> {
    return (await this.cachedCallFn(cacheOpt, 'QueryStats', req)) as QueryTestHistoryStatsResponse;
  }

  async queryVariants(req: QueryVariantsRequest, cacheOpt: CacheOption = {}): Promise<QueryVariantsResponse> {
    return (await this.cachedCallFn(cacheOpt, 'QueryVariants', req)) as QueryVariantsResponse;
  }

  async queryTests(req: QueryTestsRequest, cacheOpt: CacheOption = {}): Promise<QueryTestsResponse> {
    return (await this.cachedCallFn(cacheOpt, 'QueryTests', req)) as QueryTestsResponse;
  }
}

export class ClustersService {
  private static SERVICE = 'luci.analysis.v1.Clusters';

  private readonly cachedBatchedCluster: (
    cacheOpt: CacheOption,
    batchOpt: BatchOption,
    req: ClusterRequest
  ) => Promise<ClusterResponse>;

  constructor(client: PrpcClientExt) {
    const CLUSTER_BATCH_LIMIT = 1000;

    const batchedCluster = batched<[ClusterRequest], ClusterResponse>({
      fn: (req: ClusterRequest) => client.call(ClustersService.SERVICE, 'Cluster', req),
      combineParamSets: ([req1], [req2]) => {
        const canCombine =
          req1.testResults.length + req2.testResults.length <= CLUSTER_BATCH_LIMIT && req1.project === req2.project;
        if (!canCombine) {
          return { ok: false } as ResultErr<void>;
        }
        return {
          ok: true,
          value: [
            {
              project: req1.project,
              testResults: [...req1.testResults, ...req2.testResults],
            },
          ] as [ClusterRequest],
        };
      },
      splitReturn: (paramSets, ret) => {
        let pivot = 0;
        const splitRets: ClusterResponse[] = [];
        for (const [req] of paramSets) {
          splitRets.push({
            clusteringVersion: ret.clusteringVersion,
            clusteredTestResults: ret.clusteredTestResults.slice(pivot, pivot + req.testResults.length),
          });
          pivot += req.testResults.length;
        }

        return splitRets;
      },
    });

    this.cachedBatchedCluster = cached((batchOpt: BatchOption, req: ClusterRequest) => batchedCluster(batchOpt, req), {
      key: (_batchOpt, req) => stableStringify(req),
    });
  }

  async cluster(req: ClusterRequest, batchOpt: BatchOption = {}, cacheOpt: CacheOption = {}): Promise<ClusterResponse> {
    return (await this.cachedBatchedCluster(cacheOpt, batchOpt, req)) as ClusterResponse;
  }
}

/**
 * Construct a link to a luci-analysis rule page.
 */
export function makeRuleLink(project: string, ruleId: string) {
  return `https://${CONFIGS.LUCI_ANALYSIS.HOST}/p/${project}/rules/${ruleId}`;
}

/**
 * Construct a link to a luci-analysis cluster.
 */
export function makeClusterLink(project: string, clusterId: ClusterId) {
  return `https://${CONFIGS.LUCI_ANALYSIS.HOST}/p/${project}/clusters/${clusterId.algorithm}/${clusterId.id}`;
}
