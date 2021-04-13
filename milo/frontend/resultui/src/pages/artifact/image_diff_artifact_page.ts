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

import { MobxLitElement } from '@adobe/lit-mobx';
import { PreventAndRedirectCommands, RouterLocation } from '@vaadin/router';
import { css, customElement, html } from 'lit-element';
import { computed, observable } from 'mobx';
import { fromPromise, FULFILLED, PENDING } from 'mobx-utils';

import '../../components/image_diff_viewer';
import '../../components/status_bar';
import './artifact_page_layout';
import { AppState, consumeAppState } from '../../context/app_state';
import { NOT_FOUND_URL } from '../../routes';
import { constructArtifactName, parseArtifactName } from '../../services/resultdb';
import commonStyle from '../../styles/common_style.css';

/**
 * Renders an image diff artifact set, including expected image, actual image
 * and image diff.
 */
// TODO(weiweilin): improve error handling.
@customElement('milo-image-diff-artifact-page')
@consumeAppState
export class ImageDiffArtifactPage extends MobxLitElement {
  @observable.ref appState!: AppState;
  @observable.ref diffArtifactName!: string;

  @computed private get expectedArtifactName() {
    return constructArtifactName({ ...this.diffArtifactIdent, artifactId: this.expectedArtifactId });
  }
  @computed private get actualArtifactName() {
    return constructArtifactName({ ...this.diffArtifactIdent, artifactId: this.actualArtifactId });
  }

  @computed private get diffArtifactId() {
    return this.diffArtifactIdent.artifactId;
  }
  @observable.ref private expectedArtifactId!: string;
  @observable.ref private actualArtifactId!: string;

  @computed private get diffArtifactIdent() {
    return parseArtifactName(this.diffArtifactName);
  }

  @computed
  private get diffArtifact$() {
    if (!this.appState.resultDb) {
      return fromPromise(Promise.race([]));
    }
    return fromPromise(this.appState.resultDb.getArtifact({ name: this.diffArtifactName }));
  }
  @computed private get diffArtifact() {
    return this.diffArtifact$.state === FULFILLED ? this.diffArtifact$.value : null;
  }

  @computed
  private get expectedArtifact$() {
    if (!this.appState.resultDb) {
      return fromPromise(Promise.race([]));
    }
    return fromPromise(this.appState.resultDb.getArtifact({ name: this.expectedArtifactName }));
  }
  @computed private get expectedArtifact() {
    return this.expectedArtifact$.state === FULFILLED ? this.expectedArtifact$.value : null;
  }

  @computed
  private get actualArtifact$() {
    if (!this.appState.resultDb) {
      return fromPromise(Promise.race([]));
    }
    return fromPromise(this.appState.resultDb.getArtifact({ name: this.actualArtifactName }));
  }
  @computed private get actualArtifact() {
    return this.actualArtifact$.state === FULFILLED ? this.actualArtifact$.value : null;
  }

  @computed get isLoading() {
    return [this.expectedArtifact$.state, this.actualArtifact$.state, this.diffArtifact$.state].includes(PENDING);
  }

  onBeforeEnter(location: RouterLocation, cmd: PreventAndRedirectCommands) {
    const diffArtifactName = location.params['artifact_name'];
    const search = new URLSearchParams(location.search);
    const expectedArtifactId = search.get('expected_artifact_id');
    const actualArtifactId = search.get('actual_artifact_id');
    if (typeof diffArtifactName !== 'string' || !expectedArtifactId || !actualArtifactId) {
      return cmd.redirect(NOT_FOUND_URL);
    }

    this.expectedArtifactId = expectedArtifactId;
    this.actualArtifactId = actualArtifactId;
    this.diffArtifactName = diffArtifactName;
    return;
  }

  protected render() {
    return html`
      <milo-artifact-page-layout
        .ident=${this.diffArtifactIdent}
        .artifactIds=${[this.expectedArtifactId, this.actualArtifactId, this.diffArtifactId]}
        .isLoading=${this.isLoading}
      >
        ${this.isLoading
          ? ''
          : html`
              <milo-image-diff-viewer
                .expected=${this.expectedArtifact}
                .actual=${this.actualArtifact}
                .diff=${this.diffArtifact}
              >
              </milo-image-diff-viewer>
            `}
      </milo-artifact-page-layout>
    `;
  }

  static styles = [
    commonStyle,
    css`
      :host {
        display: block;
      }
    `,
  ];
}
