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

import createDomPurify from 'dompurify';
import { unsafeHTML } from 'lit-html/directives/unsafe-html';

const domPurify = createDomPurify(window);

// Mitigate target="_blank" vulnerability.
domPurify.addHook('afterSanitizeAttributes', (node) => {
  if (!['A', 'FORM', 'AREA'].includes(node.tagName)) {
    return;
  }

  // Note: rel="noopener" is added when the target is not set because <base> can
  // set the default target to _blank.
  if (['_self', '_top', '_parent'].includes(node.getAttribute('target') || '')) {
    return;
  }

  const existingRef = node.getAttribute('rel') || '';
  if (!/\bnoopener\b/i.test(existingRef)) {
    node.setAttribute('rel', (existingRef + ' noopener').trim());
  }
});

/**
 * Sanitizes the input HTML string and renders it.
 */
export function sanitizeHTML(html: string) {
  return unsafeHTML(domPurify.sanitize(html, {ADD_ATTR: ['target', 'artifact-id'], ADD_TAGS: ['text-artifact']}));
}
