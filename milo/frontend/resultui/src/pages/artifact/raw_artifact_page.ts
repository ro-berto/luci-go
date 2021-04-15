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

import { GrpcError, RpcCode } from '@chopsui/prpc-client';
import { Router } from '@vaadin/router';
import { css, customElement, html } from 'lit-element';
import { autorun, computed, observable } from 'mobx';
import { fromPromise, PENDING, REJECTED } from 'mobx-utils';

import '../../components/dot_spinner';
import '../../components/status_bar';
import { MiloBaseElement } from '../../components/milo_base';
import { AppState, consumeAppState } from '../../context/app_state';
import { consumeContext } from '../../libs/context';
import { router } from '../../routes';
import { ArtifactIdentifier, constructArtifactName } from '../../services/resultdb';
import commonStyle from '../../styles/common_style.css';

/**
 * Renders a raw artifact.
 */
// TODO(weiweilin): improve error handling.
@customElement('milo-raw-artifact-page')
@consumeAppState
@consumeContext('artifactIdent')
export class RawArtifactPageElement extends MiloBaseElement {
  @observable.ref appState!: AppState;
  @observable.ref artifactIdent!: ArtifactIdentifier;

  @computed
  private get artifact$() {
    if (!this.appState.resultDb) {
      return fromPromise(Promise.race([]));
    }
    return fromPromise(this.appState.resultDb.getArtifact({ name: constructArtifactName(this.artifactIdent) }));
  }

  connectedCallback() {
    super.connectedCallback();

    // TODO(weiweilin): add integration tests to ensure redirection works properly.
    this.addDisposer(
      autorun(() => {
        if (this.artifact$.state === PENDING) {
          return;
        }

        if (this.artifact$.state === REJECTED) {
          const err = this.artifact$.value as GrpcError;
          const mayRequireSignin = [RpcCode.NOT_FOUND, RpcCode.PERMISSION_DENIED, RpcCode.UNAUTHENTICATED].includes(
            err.code
          );
          if (mayRequireSignin && this.appState.userId === '') {
            Router.go(`${router.urlForName('login')}?${new URLSearchParams([['redirect', window.location.href]])}`);
            return;
          }
          this.dispatchEvent(
            new ErrorEvent('error', {
              message: err.message,
              composed: true,
              bubbles: true,
            })
          );
          return;
        }

        window.open(this.artifact$.value.fetchUrl, '_self');
      })
    );
  }

  protected render() {
    return html`<div id="content">Loading artifact <milo-dot-spinner></milo-dot-spinner></div>`;
  }

  static styles = [
    commonStyle,
    css`
      #content {
        margin: 20px;
        color: var(--active-color);
      }
    `,
  ];
}
