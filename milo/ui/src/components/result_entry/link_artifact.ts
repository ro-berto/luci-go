// Copyright 2023 The LUCI Authors.
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
import { css, html } from 'lit';
import { customElement } from 'lit/decorators.js';
import { computed, makeObservable, observable } from 'mobx';
import { fromPromise, IPromiseBasedObservable } from 'mobx-utils';

import '../expandable_entry';
import { ARTIFACT_LENGTH_LIMIT } from '../../libs/constants';
import { reportRenderError } from '../../libs/error_handler';
import { unwrapObservable } from '../../libs/milo_mobx_utils';
import { urlSetSearchQueryParam } from '../../libs/utils';
import { Artifact } from '../../services/resultdb';
import commonStyle from '../../styles/common_style.css';

/**
 * Renders a link artifact.
 */
@customElement('milo-link-artifact')
export class LinkArtifactElement extends MobxLitElement {
  @observable.ref artifact!: Artifact;

  @observable.ref private loadError = false;

  @computed
  private get content$(): IPromiseBasedObservable<string> {
    return fromPromise(
      // TODO(crbug/1206109): use permanent raw artifact URL.
      fetch(
        urlSetSearchQueryParam(
          this.artifact.fetchUrl, 'n', ARTIFACT_LENGTH_LIMIT
        )
      ).then((res) => {
        if (!res.ok) {
          this.loadError = true;
          return '';
        }
        return res.text();
      })
    );
  }

  @computed
  private get content() {
    return unwrapObservable(this.content$, null);
  }

  constructor() {
    super();
    makeObservable(this);
  }

  protected render = reportRenderError(this, () => {
    if (this.loadError) {
      return html`
        <span class="load-error">
          Error loading ${this.artifact.artifactId} link
        </span>`;
    }

    if (this.content) {
      return html`
        <a href=${this.content} target="_blank">
          ${this.artifact.artifactId}
        </a>`;
    }

    return html`<span class="greyed-out">Loading...</span>`;
  });

  static styles = [
    commonStyle,
    css`
      .greyed-out {
        color: var(--greyed-out-text-color);
      }

      .load-error {
        color: var(--failure-color);
      }
    `,
  ];
}
