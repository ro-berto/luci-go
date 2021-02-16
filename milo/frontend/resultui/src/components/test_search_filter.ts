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

import { MobxLitElement } from '@adobe/lit-mobx';
import { customElement, html } from 'lit-element';
import { computed, observable } from 'mobx';

import { consumeInvocationState, InvocationState } from '../context/invocation_state/invocation_state';
import { suggestSearchQuery } from '../libs/search_query';
import './auto_complete';
import { Suggestion } from './auto_complete';
import './hotkey';

export interface TestFilter {
  showExpected: boolean;
  showExonerated: boolean;
  showFlaky: boolean;
}

/**
 * An element that let the user search tests with DSL.
 */
@customElement('milo-test-search-filter')
@consumeInvocationState
export class TestSearchFilterElement extends MobxLitElement {
  @observable.ref invocationState!: InvocationState;
  @observable.ref searchText!: string;

  @computed private get lastSubQuery() {
    return this.invocationState.searchText.split(' ').pop() || '';
  }
  @computed private get queryPrefix() {
    const searchTextPrefixLen = this.invocationState.searchText.length - this.lastSubQuery.length;
    return this.invocationState.searchText.slice(0, searchTextPrefixLen);
  }
  @computed private get suggestions() {
    return suggestSearchQuery(this.lastSubQuery);
  }

  protected render() {
    return html`
      <milo-hotkey
        key="/"
        .handler=${() => {
          // Set a tiny timeout to ensure '/' isn't recorded by the input box.
          setTimeout(() => this.shadowRoot?.getElementById('search-box')?.focus());
        }}
      >
        <milo-auto-complete
          id="search-box"
          .value=${this.invocationState.searchText}
          .placeHolder=${'Press / to search test results...'}
          .suggestions=${this.suggestions}
          .onValueUpdate=${(newVal: string) => this.invocationState.searchText = newVal}
          .onSuggestionSelected=${(suggestion: Suggestion) => {
            this.invocationState.searchText = this.queryPrefix + suggestion.value + ' ';
          }}
        >
        </milo-auto-complete>
      </milo-hotkey>
    `;
  }
}
