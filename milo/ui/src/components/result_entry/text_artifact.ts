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

import '@material/mwc-icon';
import { MobxLitElement } from '@adobe/lit-mobx';
import { css, html } from 'lit';
import { customElement } from 'lit/decorators.js';
import { computed, makeObservable, observable } from 'mobx';
import { fromPromise, IPromiseBasedObservable } from 'mobx-utils';

import '../dot_spinner';
import { consumeArtifacts, consumeArtifactsFinalized } from '../../context/artifact_provider';
import { ARTIFACT_LENGTH_LIMIT } from '../../libs/constants';
import { consumer } from '../../libs/context';
import { reportRenderError } from '../../libs/error_handler';
import { unwrapObservable } from '../../libs/milo_mobx_utils';
import { urlSetSearchQueryParam } from '../../libs/utils';
import { Artifact } from '../../services/resultdb';
import commonStyle from '../../styles/common_style.css';

/**
 * Renders a text artifact.
 */
@customElement('text-artifact')
@consumer
export class TextArtifactElement extends MobxLitElement {
  static get properties() {
    return {
      artifactId: {
        attribute: 'artifact-id',
        type: String,
      },
      invLevel: {
        attribute: 'inv-level',
        type: Boolean,
      },
    };
  }

  @observable.ref
  @consumeArtifacts()
  artifacts!: Map<string, Artifact>;

  @observable.ref
  @consumeArtifactsFinalized()
  finalized = false;

  set artifactId(newVal: string) {
    this._artifactId = newVal;
  }

  set invLevel(newVal: boolean) {
    this._invLevel = newVal;
  }

  @observable.ref _artifactId!: string;
  @observable.ref _invLevel = false;

  @computed
  private get fetchUrl(): string {
    const artifact = this.artifacts.get((this._invLevel ? 'inv-level/' : '') + this._artifactId);
    // TODO(crbug/1206109): use permanent raw artifact URL.
    return artifact ? artifact.fetchUrl : '';
  }

  @computed
  private get content$(): IPromiseBasedObservable<string> {
    if (!this.fetchUrl) {
      return fromPromise(Promise.race([]));
    }
    return fromPromise(
      fetch(urlSetSearchQueryParam(this.fetchUrl, 'n', ARTIFACT_LENGTH_LIMIT)).then((res) => res.text())
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
    const label = this._invLevel ? 'Inv-level artifact' : 'Artifact';

    if (this.finalized && this.fetchUrl === '') {
      return html`<div>${label}: <i>${this._artifactId}</i> not found.</div>`;
    }

    if (this.content === null) {
      return html`<div id="load">Loading <milo-dot-spinner></milo-dot-spinner></div>`;
    }

    if (this.content === '') {
      return html`<div>${label}: <i>${this._artifactId}</i> is empty.</div>`;
    }

    return html`<pre>${this.content}</pre>`;
  });

  static styles = [
    commonStyle,
    css`
      #load {
        color: var(--active-text-color);
      }
      pre {
        margin: 0;
        font-size: 12px;
        white-space: pre-wrap;
        overflow-wrap: anywhere;
      }
    `,
  ];
}
