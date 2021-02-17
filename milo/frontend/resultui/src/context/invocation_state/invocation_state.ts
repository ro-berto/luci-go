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


import { autorun, computed, observable } from 'mobx';
import { fromPromise, FULFILLED, IPromiseBasedObservable } from 'mobx-utils';

import { consumeContext, provideContext } from '../../libs/context';
import { parseSearchQuery } from '../../libs/search_query';
import { TestLoader } from '../../models/test_loader';
import { TestNode } from '../../models/test_node';
import { Invocation, TestVariant } from '../../services/resultdb';
import { AppState } from '../app_state/app_state';

/**
 * Records state of an invocation.
 */
export class InvocationState {
  @observable.ref invocationId = '';
  @observable.ref initialized = false;
  @observable.ref searchText = '';

  @observable.ref searchFilter = (_v: TestVariant) => true;

  private filterVariant(variant: TestVariant): boolean {
    return variant.testId.startsWith(this.selectedNode.path) && this.searchFilter(variant);
  }

  private disposer = () => {};
  constructor(private appState: AppState) {
    this.disposer = autorun(() => {
      try {
        this.searchFilter = parseSearchQuery(this.searchText);
      } catch (e) {
        //TODO(weiweilin): display the error to the user.
        console.error(e);
      }
    });
  }

  @observable.ref private isDisposed = false;

  /**
   * Perform cleanup.
   * Must be called before the object is GCed.
   */
  dispose() {
    this.isDisposed = true;
    this.disposer();

    // Evaluates @computed({keepAlive: true}) properties after this.isDisposed
    // is set to true so they no longer subscribes to any external observable.
    // tslint:disable: no-unused-expression
    this.testLoader;
    // tslint:enable: no-unused-expression
  }


  @computed
  get invocationName(): string | null {
    if (!this.invocationId) {
      return null;
    }
    return 'invocations/' + this.invocationId;
  }

  @computed
  get invocationRes(): IPromiseBasedObservable<Invocation> {
    if (!this.appState.resultDb || !this.invocationName) {
      // Returns a promise that never resolves when resultDb isn't ready.
      return fromPromise(Promise.race([]));
    }
    return fromPromise(this.appState.resultDb.getInvocation({name: this.invocationName}));
  }

  @computed
  get invocation(): Invocation | null {
    if (this.invocationRes.state !== FULFILLED) {
      return null;
    }
    return this.invocationRes.value;
  }

  @observable.ref selectedNode!: TestNode;

  @computed({keepAlive: true})
  get testLoader(): TestLoader | null {
    if (this.isDisposed || !this.invocationName || !this.appState.uiSpecificService) {
      return null;
    }
    return new TestLoader(TestNode.newRoot(), {invocations: [this.invocationName]}, this.appState.uiSpecificService);
  }

  @computed get filteredUnexpectedVariants() {
    return (this.testLoader?.unexpectedTestVariants || [])
      .filter(v => this.filterVariant(v));
  }

  @computed get filteredUnexpectedlySkippedVariants() {
    return (this.testLoader?.unexpectedlySkippedTestVariants || [])
      .filter(v => this.filterVariant(v));
  }

  @computed get filteredFlakyVariants() {
    return (this.testLoader?.flakyTestVariants || [])
      .filter(v => this.filterVariant(v));
  }
  @computed get filteredExoneratedVariants() {
    return (this.testLoader?.exoneratedTestVariants || [])
      .filter(v => this.filterVariant(v));
  }
  @computed get filteredExpectedVariants() {
    return (this.testLoader?.expectedTestVariants || [])
      .filter(v => this.filterVariant(v));
  }
}

export const consumeInvocationState = consumeContext<'invocationState', InvocationState>('invocationState');
export const provideInvocationState = provideContext<'invocationState', InvocationState>('invocationState');
