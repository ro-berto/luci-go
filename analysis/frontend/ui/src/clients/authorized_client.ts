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

import { PrpcClient } from '@chopsui/prpc-client';

import { obtainAuthState } from '../api/auth_state';

export class AuthorizedPrpcClient {
  client: PrpcClient;
  // Should the ID token be used to authorize the request, or the access token?
  useIDToken: boolean;

  // Initialises a new AuthorizedPrpcClient that connects to host.
  // To connect to LUCI Analysis, leave host unspecified.
  constructor(host?: string, useIDToken?: boolean) {
    // Only allow insecure connections in LUCI Analysis in local development,
    // where risk of man-in-the-middle attack to server is negligible.
    const insecure = document.location.protocol === 'http:' && !host;
    const hostname = document.location.hostname;
    if (insecure && hostname !== 'localhost' && hostname !== '127.0.0.1') {
      // Server misconfiguration.
      throw new Error('LUCI Analysis should never be served over http: outside local development.');
    }
    this.client = new PrpcClient({
      host: host,
      insecure: insecure,
    });
    this.useIDToken = useIDToken === true;
  }

  async call(service: string, method: string, message: object, additionalHeaders?: {
        [key: string]: string;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } | undefined): Promise<any> {
    if (!window.isAnonymous) {
      // Although PrpcClient allows us to pass a token to the constructor,
      // we prefer to inject it at request time to ensure the most recent
      // token is used.
      const authState = await obtainAuthState();
      let token: string;
      if (this.useIDToken) {
        token = authState.idToken;
      } else {
        token = authState.accessToken;
      }
      additionalHeaders = {
        Authorization: 'Bearer ' + token,
        ...additionalHeaders,
      };
    }
    return this.client.call(service, method, message, additionalHeaders);
  }
}
