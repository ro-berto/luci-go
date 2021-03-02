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

import { assert } from 'chai';
import chai from 'chai';
import sinon from 'sinon';

import { chaiRecursiveDeepInclude } from '../libs/test_utils/chai_recursive_deep_include';
import { TestVariantStatus, UISpecificService } from '../services/resultdb';
import { LoadingStage, TestLoader } from './test_loader';

chai.use(chaiRecursiveDeepInclude);

const variant1 = {
  testId: 'a',
  variant: {'def': {'key1': 'val1'}},
  variantHash: 'key1:val1',
  status: TestVariantStatus.UNEXPECTED,
};

const variant2 = {
  testId: 'a',
  variant: {def: {'key1': 'val2'}},
  variantHash: 'key1:val2',
  status: TestVariantStatus.UNEXPECTED,
};

const variant3 = {
  testId: 'a',
  variant: {def: {'key1': 'val3'}},
  variantHash: 'key1:val3',
  status: TestVariantStatus.UNEXPECTED,
};

const variant4 = {
  testId: 'b',
  variant: {'def': {'key1': 'val2'}},
  variantHash: 'key1:val2',
  status: TestVariantStatus.FLAKY,
};

const variant5 = {
  testId: 'matched-id',
  variant: {def: {'key1': 'val2', 'key2': 'val1'}},
  variantHash: 'key1:val2|key2:val1',
  status: TestVariantStatus.FLAKY,
};

const variant6 = {
  testId: 'c',
  variant: {def: {'key1': 'val2', 'key2': 'val2'}},
  variantHash: 'key1:val2|key2:val2',
  status: TestVariantStatus.EXONERATED,
};

const variant7 = {
  testId: 'd',
  variant: {def: {'key1': 'val1'}},
  variantHash: 'key1:val1',
  status: TestVariantStatus.EXONERATED,
};

const variant8 = {
  testId: 'd',
  variant: {def: {'key1': 'val2'}},
  variantHash: 'key1:val2',
  status: TestVariantStatus.EXPECTED,
};

const variant9 = {
  testId: 'e',
  variant: {def: {'key1': 'val2'}},
  variantHash: 'key1:val2',
  status: TestVariantStatus.EXPECTED,
};

const variant10 = {
  testId: 'f',
  variant: {def: {'key1': 'val2'}},
  variantHash: 'key1:val2',
  status: TestVariantStatus.EXPECTED,
};

const variant11 = {
  testId: 'g',
  variant: {def: {'key1': 'val2'}},
  variantHash: 'key1:val2',
  status: TestVariantStatus.EXPECTED,
};

const variant12 = {
  testId: 'matched-id',
  variant: {def: {'key1': 'val2'}},
  variantHash: 'key1:val2',
  status: TestVariantStatus.EXPECTED,
};

describe('TestLoader', () => {
  let testLoader: TestLoader;
  let stub = sinon.stub();
  const req = {invocations: ['invocation'], pageSize: 4};
  describe('when first page contains variants', () => {
    beforeEach(() => {
      stub = sinon.stub();
      stub.onCall(0).resolves({testVariants: [variant1, variant2, variant3, variant4], nextPageToken: 'page2'});
      stub.onCall(1).resolves({testVariants: [variant5, variant6, variant7], nextPageToken: 'page3'});
      stub.onCall(2).resolves({testVariants: [variant8, variant9, variant10, variant11], nextPageToken: 'page4'});
      stub.onCall(3).resolves({testVariants: [variant12], nextPageToken: undefined});
      testLoader = new TestLoader(
        req,
        {queryTestVariants: stub} as Partial<UISpecificService> as UISpecificService,
      );
    });

    it('should preserve loading progress', async () => {
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingUnexpected);
      assert.strictEqual(stub.callCount, 0);

      await testLoader.loadNextTestVariants();
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4]);
      assert.deepEqual(testLoader.exoneratedTestVariants, []);
      assert.deepEqual(testLoader.expectedTestVariants, []);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingFlaky);
      assert.strictEqual(stub.callCount, 1);
      assert.deepEqual(stub.getCall(0).args[0], {...req, pageToken: ''});

      await testLoader.loadNextTestVariants();
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, []);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingExpected);
      assert.strictEqual(stub.callCount, 2);
      assert.deepEqual(stub.getCall(1).args[0], {...req, pageToken: 'page2'});

      await testLoader.loadNextTestVariants();
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, [variant8, variant9, variant10, variant11]);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingExpected);
      assert.strictEqual(stub.callCount, 3);
      assert.deepEqual(stub.getCall(2).args[0], {...req, pageToken: 'page3'});

      await testLoader.loadNextTestVariants();
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, [variant8, variant9, variant10, variant11, variant12]);
      assert.strictEqual(testLoader.stage, LoadingStage.Done);
      assert.strictEqual(stub.callCount, 4);
      assert.deepEqual(stub.getCall(3).args[0], {...req, pageToken: 'page4'});

      // Should not load when the iterator is exhausted.
      await testLoader.loadNextTestVariants();
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, [variant8, variant9, variant10, variant11, variant12]);
      assert.strictEqual(testLoader.stage, LoadingStage.Done);
      assert.strictEqual(stub.callCount, 4);
    });

    it('should handle concurrent loadNextPage calls correctly', async () => {
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingUnexpected);

      const loadReq1 = testLoader.loadNextTestVariants();
      const loadReq2 = testLoader.loadNextTestVariants();
      const loadReq3 = testLoader.loadNextTestVariants();
      const loadReq4 = testLoader.loadNextTestVariants();
      const loadReq5 = testLoader.loadNextTestVariants();
      assert.isTrue(testLoader.isLoading);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingUnexpected);

      await loadReq1;
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4]);
      assert.deepEqual(testLoader.exoneratedTestVariants, []);
      assert.deepEqual(testLoader.expectedTestVariants, []);
      // loadReq2 has not finished loading yet.
      assert.isTrue(testLoader.isLoading);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingFlaky);

      await loadReq2;
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, []);
      // loadReq3 has not finished loading yet.
      assert.isTrue(testLoader.isLoading);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingExpected);

      await loadReq3;
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, [variant8, variant9, variant10, variant11]);
      // loadReq4 has not finished loading yet.
      assert.isTrue(testLoader.isLoading);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingExpected);

      await loadReq4;
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, [variant8, variant9, variant10, variant11, variant12]);
      // The list is exhausted, loadReq5 should not change the loading state.
      assert.isFalse(testLoader.isLoading);
      assert.strictEqual(testLoader.stage, LoadingStage.Done);

      await loadReq5;
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, [variant8, variant9, variant10, variant11, variant12]);
      assert.isFalse(testLoader.isLoading);
      assert.strictEqual(testLoader.stage, LoadingStage.Done);

      assert.strictEqual(stub.callCount, 4);
      assert.deepEqual(stub.getCall(0).args[0], {...req, pageToken: ''});
      assert.deepEqual(stub.getCall(1).args[0], {...req, pageToken: 'page2'});
      assert.deepEqual(stub.getCall(2).args[0], {...req, pageToken: 'page3'});
      assert.deepEqual(stub.getCall(3).args[0], {...req, pageToken: 'page4'});
    });

    it('should correctly set firstRequestSent', async () => {
      assert.isFalse(testLoader.firstRequestSent);
      testLoader.loadNextTestVariants();
      assert.isTrue(testLoader.firstRequestSent);
    });

    it('should load until UNEXPECTED correctly', async () => {
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingUnexpected);
      assert.strictEqual(stub.callCount, 0);

      await testLoader.loadNextTestVariants(TestVariantStatus.UNEXPECTED);
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4]);
      assert.deepEqual(testLoader.exoneratedTestVariants, []);
      assert.deepEqual(testLoader.expectedTestVariants, []);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingFlaky);
    });

    it('should load until FLAKY correctly', async () => {
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingUnexpected);
      assert.strictEqual(stub.callCount, 0);

      await testLoader.loadNextTestVariants(TestVariantStatus.FLAKY);
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4]);
      assert.deepEqual(testLoader.exoneratedTestVariants, []);
      assert.deepEqual(testLoader.expectedTestVariants, []);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingFlaky);
    });

    it('should load until EXONERATED correctly', async () => {
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingUnexpected);
      assert.strictEqual(stub.callCount, 0);

      await testLoader.loadNextTestVariants(TestVariantStatus.EXONERATED);
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, []);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingExpected);
    });

    it('should load until EXPECTED correctly', async () => {
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingUnexpected);
      assert.strictEqual(stub.callCount, 0);

      await testLoader.loadNextTestVariants(TestVariantStatus.EXPECTED);
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, [variant8, variant9, variant10, variant11]);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingExpected);
    });

    it('should load one page of EXPECTED variants when EXPECTED is next', async () => {
      await testLoader.loadNextTestVariants(TestVariantStatus.EXONERATED);
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, []);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingExpected);

      await testLoader.loadNextTestVariants(TestVariantStatus.EXPECTED);
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, [variant8, variant9, variant10, variant11]);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingExpected);
    });

    it('should load one page of EXPECTED variants when some EXPECTED variants have already been loaded', async () => {
      await testLoader.loadNextTestVariants(TestVariantStatus.EXPECTED);
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, [variant8, variant9, variant10, variant11]);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingExpected);

      await testLoader.loadNextTestVariants(TestVariantStatus.EXPECTED);
      assert.deepEqual(testLoader.unexpectedTestVariants, [variant1, variant2, variant3]);
      assert.deepEqual(testLoader.flakyTestVariants, [variant4, variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, [variant6, variant7]);
      assert.deepEqual(testLoader.expectedTestVariants, [variant8, variant9, variant10, variant11, variant12]);
      assert.strictEqual(testLoader.stage, LoadingStage.Done);
    });

    it('should load at least one test variant that matches the filter', async () => {
      testLoader.filter = (v) => v.testId === 'matched-id';
      await testLoader.loadNextTestVariants();

      assert.deepEqual(testLoader.unexpectedTestVariants, []);
      assert.deepEqual(testLoader.flakyTestVariants, [variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, []);
      assert.deepEqual(testLoader.expectedTestVariants, []);

      await testLoader.loadNextTestVariants();

      assert.deepEqual(testLoader.unexpectedTestVariants, []);
      assert.deepEqual(testLoader.flakyTestVariants, [variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, []);
      assert.deepEqual(testLoader.expectedTestVariants, [variant12]);
    });

    it('should load at least one test variant that matches the filter', async () => {
      testLoader.filter = (v) => v.testId === 'matched-id';
      await testLoader.loadNextTestVariants();

      assert.deepEqual(testLoader.unexpectedTestVariants, []);
      assert.deepEqual(testLoader.flakyTestVariants, [variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, []);
      assert.deepEqual(testLoader.expectedTestVariants, []);

      await testLoader.loadNextTestVariants();

      assert.deepEqual(testLoader.unexpectedTestVariants, []);
      assert.deepEqual(testLoader.flakyTestVariants, [variant5]);
      assert.deepEqual(testLoader.exoneratedTestVariants, []);
      assert.deepEqual(testLoader.expectedTestVariants, [variant12]);
    });

    it('should stop loading at the final page when no test variants matches the filter', async () => {
      testLoader.filter = () => false;

      // Detect infinite loop and abort.
      const oldLoadNextTestVariants = testLoader.loadNextTestVariants.bind(testLoader);
      let callCount = 0;
      testLoader.loadNextTestVariants = (...params) => {
        callCount++;
        if (callCount > 10) {
          throw new Error('too many load next page calls');
        }
        return oldLoadNextTestVariants(...params);
      };

      await testLoader.loadNextTestVariants();

      assert.strictEqual(testLoader.testVariantCount, 0);
      assert.strictEqual(testLoader.unfilteredTestVariantCount, 12);
      assert.strictEqual(testLoader.isLoading, false);
      assert.strictEqual(testLoader.stage, LoadingStage.Done);
    });
  });

  describe('when first page contains no variants', () => {
    beforeEach(() => {
      stub = sinon.stub();
      stub.onCall(0).resolves({nextPageToken: 'page2'});
      stub.onCall(1).resolves({testVariants: [variant8], nextPageToken: 'page3'});
      stub.onCall(2).resolves({testVariants: [variant9], nextPageToken: undefined});
      testLoader = new TestLoader(
        req,
        {queryTestVariants: stub} as Partial<UISpecificService> as UISpecificService,
      );
    });

    it('should correctly handle a response with 0 variants', async () => {
      assert.strictEqual(stub.callCount, 0);

      await testLoader.loadNextTestVariants();
      assert.deepEqual(testLoader.unexpectedTestVariants, []);
      assert.deepEqual(testLoader.flakyTestVariants, []);
      assert.deepEqual(testLoader.exoneratedTestVariants, []);
      assert.deepEqual(testLoader.expectedTestVariants, [variant8]);
      assert.strictEqual(testLoader.stage, LoadingStage.LoadingExpected);
    });
  });
});
