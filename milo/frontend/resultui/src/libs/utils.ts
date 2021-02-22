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

import MarkdownIt from 'markdown-it';

import { defaultTarget } from './markdown_it_plugins/default_target';
import { sanitizeHTML } from './sanitize_html';

const md = MarkdownIt({html: true, linkify: true})
  .use(defaultTarget, '_blank');

export function renderMarkdown(markdown: string) {
  return sanitizeHTML(renderMarkdownUnsanitized(markdown));
}

export function renderMarkdownUnsanitized(markdown: string): string {
  return md.render(markdown);
}

/**
 * Extend URL with methods that can be chained.
 */
export class ChainableURL extends URL {
  withSearchParam(key: string, value: string, override = false) {
    if (override) {
      this.searchParams.set(key, value);
    } else {
      this.searchParams.append(key, value);
    }
    return this;
  }
}

// Generates URL for collecting feedback.
export function genFeedbackUrl() {
  const feedbackComment = encodeURIComponent(
`From Link: ${document.location.href}
Please enter a description of the problem, with repro steps if applicable.
`);
  return `https://bugs.chromium.org/p/chromium/issues/entry?template=Build%20Infrastructure&components=Infra%3EPlatform%3EMilo%3EResultUI&labels=Pri-2,Type-Bug&comment=${feedbackComment}`;
}
