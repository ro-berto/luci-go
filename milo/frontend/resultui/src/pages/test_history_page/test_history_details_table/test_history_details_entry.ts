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

import '@material/mwc-icon';
import { MobxLitElement } from '@adobe/lit-mobx';
import { css, customElement, html } from 'lit-element';
import { repeat } from 'lit-html/directives/repeat';
import { unsafeHTML } from 'lit-html/directives/unsafe-html';
import { DateTime } from 'luxon';
import { computed, makeObservable, observable } from 'mobx';
import { fromPromise } from 'mobx-utils';

import '../../../components/expandable_entry';
import '../../../components/result_entry';
import { VARIANT_STATUS_CLASS_MAP, VARIANT_STATUS_ICON_MAP, VERDICT_VARIANT_STATUS_MAP } from '../../../libs/constants';
import { unwrapObservable } from '../../../libs/milo_mobx_utils';
import { lazyRendering, RenderPlaceHolder } from '../../../libs/observer_element';
import { LONG_TIME_FORMAT, SHORT_TIME_FORMAT } from '../../../libs/time_utils';
import { router } from '../../../routes';
import { TestVerdictBundle } from '../../../services/luci_analysis';
import { RESULT_LIMIT } from '../../../services/resultdb';
import { consumeStore, StoreInstance } from '../../../store';
import colorClasses from '../../../styles/color_classes.css';
import commonStyle from '../../../styles/common_style.css';

// This list defines the order in which variant def keys should be displayed.
// Any unrecognized keys will be listed after the ones defined below.
const ORDERED_VARIANT_DEF_KEYS = Object.freeze(['bucket', 'builder', 'test_suite']);

/**
 * Renders an expandable entry of the given test variant.
 */
@customElement('milo-test-history-details-entry')
@lazyRendering
export class TestHistoryDetailsEntryElement extends MobxLitElement implements RenderPlaceHolder {
  @observable.ref @consumeStore() store!: StoreInstance;

  @observable.ref verdictBundle!: TestVerdictBundle;
  @observable.ref columnGetters: Array<(v: TestVerdictBundle) => unknown> = [];

  @observable.ref private _expanded = false;
  @computed get expanded() {
    return this._expanded;
  }
  set expanded(newVal: boolean) {
    this._expanded = newVal;
    // Always render the content once it was expanded so the descendants' states
    // don't get reset after the node is collapsed.
    this.shouldRenderContent = this.shouldRenderContent || newVal;
  }

  @observable.ref private shouldRenderContent = false;

  @computed
  private get testVariant$() {
    if (!this.store.services.resultDb) {
      return fromPromise(Promise.race([]));
    }
    const verdict = this.verdictBundle.verdict;
    const testVariant = this.store.services.resultDb
      .batchGetTestVariants({
        invocation: 'invocations/' + verdict.invocationId,
        testVariants: [
          {
            testId: verdict.testId,
            variantHash: verdict.variantHash,
          },
        ],
        resultLimit: RESULT_LIMIT,
      })
      .then((res) => res.testVariants![0]);
    return fromPromise(testVariant);
  }

  @computed
  private get testVariant() {
    return unwrapObservable(this.testVariant$, null);
  }

  @computed
  private get sourceUrl() {
    const testLocation = this.testVariant?.testMetadata?.location;
    if (!testLocation) {
      return null;
    }
    return (
      testLocation.repo +
      '/+/HEAD' +
      testLocation.fileName.slice(1) +
      (testLocation.line ? '#' + testLocation.line : '')
    );
  }

  @computed
  private get hasSingleChild() {
    return (this.testVariant?.results?.length ?? 0) + (this.testVariant?.exonerations?.length ?? 0) === 1;
  }

  @computed
  private get variantDef() {
    const def = this.verdictBundle.variant.def || {};
    const res: Array<[string, string]> = [];
    const seen = new Set();
    for (const key of ORDERED_VARIANT_DEF_KEYS) {
      if (Object.prototype.hasOwnProperty.call(def, key)) {
        res.push([key, def[key]]);
        seen.add(key);
      }
    }
    for (const [key, value] of Object.entries(def)) {
      if (!seen.has(key)) {
        res.push([key, value]);
      }
    }
    return res;
  }

  @computed
  private get expandedResultIndex() {
    // If there's only a single result, just expand it (even if it passed).
    if (this.hasSingleChild) {
      return 0;
    }
    // Otherwise expand the first failed result, or -1 if there aren't any.
    return this.testVariant?.results?.findIndex((e) => !e.result.expected) ?? -1;
  }

  @computed private get columnValues() {
    return this.columnGetters.map((fn) => fn(this.verdictBundle));
  }

  @computed private get dateTime() {
    if (!this.verdictBundle.verdict.partitionTime) {
      return null;
    }
    return DateTime.fromISO(this.verdictBundle.verdict.partitionTime).toUTC();
  }

  @computed private get invocationUrl() {
    const invId = this.verdictBundle.verdict.invocationId;
    const match = invId.match(/^build-(?<id>\d+)/);
    if (!match) {
      return router.urlForName('invocation-test-results', { invocation_id: invId });
    }
    return router.urlForName('build-short-link', { build_id: match.groups!['id'], path: ['test-results'] });
  }

  constructor() {
    super();
    makeObservable(this);
  }

  private renderBody() {
    if (!this.shouldRenderContent) {
      return html``;
    }

    if (!this.testVariant) {
      return html`Loading <milo-dot-spinner></milo-dot-spinner>`;
    }

    return html`
      <div id="basic-info">
        ${this.sourceUrl ? html`<a href=${this.sourceUrl} target="_blank">source</a>` : ''}
        ${this.variantDef.length !== 0 ? '|' : ''}
        <span class="greyed-out">
          ${this.variantDef.map(
            ([k, v]) => html`
              <span class="kv">
                <span class="kv-key">${k}</span>
                <span class="kv-value">${v}</span>
              </span>
            `
          )}
        </span>
      </div>
      ${repeat(
        this.testVariant.exonerations || [],
        (e) => e.exonerationId,
        (e) => html`
          <div class="explanation-html">
            ${unsafeHTML(
              e.explanationHtml || 'This test variant had unexpected results, but was exonerated (reason not provided).'
            )}
          </div>
        `
      )}
      ${this.testVariant.results?.length === RESULT_LIMIT
        ? html`<div id="result-limit-warning">Only the first ${RESULT_LIMIT} results are displayed.</div>`
        : ''}
      ${repeat(
        this.testVariant.results || [],
        (r) => r.result.resultId,
        (r, i) => html`
          <milo-result-entry
            .id=${i + 1}
            .testResult=${r.result}
            .project=${this.store.testHistoryPage.project || ''}
            .expanded=${i === this.expandedResultIndex}
          ></milo-result-entry>
        `
      )}
    `;
  }

  renderPlaceHolder() {
    return '';
  }

  protected render() {
    const status = VERDICT_VARIANT_STATUS_MAP[this.verdictBundle.verdict.status];
    return html`
      <milo-expandable-entry .expanded=${this.expanded} .onToggle=${(expanded: boolean) => (this.expanded = expanded)}>
        <div id="header" slot="header">
          <mwc-icon class=${VARIANT_STATUS_CLASS_MAP[status]}>${VARIANT_STATUS_ICON_MAP[status]}</mwc-icon>
          <div title=${this.dateTime?.toFormat(LONG_TIME_FORMAT) || ''}>
            ${this.dateTime?.toFormat(SHORT_TIME_FORMAT) || ''}
          </div>
          <a href=${this.invocationUrl} @click=${(e: Event) => e.stopImmediatePropagation()} target="_blank">
            ${this.verdictBundle.verdict.invocationId}
          </a>
          ${this.columnValues.map((v) => html`<div title=${v}>${v}</div>`)}
        </div>
        <div id="body" slot="content">${this.renderBody()}</div>
      </milo-expandable-entry>
    `;
  }

  static styles = [
    commonStyle,
    colorClasses,
    css`
      :host {
        display: block;
        min-height: 24px;
      }

      #header {
        display: grid;
        grid-template-columns: var(--thdt-columns);
        grid-gap: 5px;
        font-size: 16px;
        line-height: 24px;
      }
      #header > * {
        overflow: hidden;
        text-overflow: ellipsis;
      }

      #body {
        overflow: hidden;
      }

      #basic-info {
        font-weight: 500;
        line-height: 24px;
        margin-left: 5px;
      }

      .kv-key::after {
        content: ':';
      }
      .kv-value::after {
        content: ',';
      }
      .kv:last-child > .kv-value::after {
        content: '';
      }
      #def-table {
        margin-left: 29px;
      }

      .greyed-out {
        color: var(--greyed-out-text-color);
      }

      .explanation-html {
        background-color: var(--block-background-color);
        padding: 5px;
      }

      #result-limit-warning {
        padding: 5px;
        background-color: var(--warning-color);
        font-weight: 500;
      }
    `,
  ];
}
