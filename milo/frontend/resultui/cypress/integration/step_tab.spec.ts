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

describe('Steps & Logs Tab', () => {
  it('should accept arbitrary path suffix', () => {
    cy.visit(
      // eslint-disable-next-line max-len
      '/p/chromium/builders/ci/android-marshmallow-arm64-rel-swarming/12479/steps/an/arbitrary/path/suffix'
    );
    cy.location('pathname').should(
      'equal',
      '/ui/p/chromium/builders/ci/android-marshmallow-arm64-rel-swarming/12479/steps'
    );
  });
});
