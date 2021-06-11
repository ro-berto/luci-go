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

declare global {
  // eslint-disable-next-line @typescript-eslint/no-namespace
  namespace Cypress {
    interface Chainable {
      /**
       * Stubs all pRPC requests to buildbucket, resultdb, and Milo with
       * stubRequests command. Cache names are set to 'buildbucket', 'resultdb',
       * and 'milo'.
       */
      stubPrpcServices: typeof stubPrpcServices;
    }
  }
}

export const STUB_REQUEST_OPTIONS = {
  matchHeaders: ['host', 'accept', 'content-type', 'origin', 'authorization'],
};

/**
 * Stubs all pRPC requests to buildbucket, resultdb, and Milo.
 */
function stubPrpcServices() {
  // TODO(weiweilin): read host names from configs.
  cy.stubRequests('https://cr-buildbucket-dev.appspot.com/prpc/**', 'buildbucket', STUB_REQUEST_OPTIONS);
  cy.stubRequests('https://staging.results.api.cr.dev/prpc/**', 'resultdb', STUB_REQUEST_OPTIONS);
  cy.stubRequests('https://localhost:8080/prpc/**', 'milo', STUB_REQUEST_OPTIONS);
}

/**
 * Adds stubPrpcServices command to Cypress.
 */
export function addStubPrpcServicesCommand() {
  Cypress.Commands.add('stubPrpcServices', stubPrpcServices);
}
