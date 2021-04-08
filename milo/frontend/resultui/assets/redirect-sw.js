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

/**
 * @fileoverview
 * This service worker redirects milo links to ResultUI links.
 * This is 200-400ms faster than redirecting on the server side.
 */

// TODO(crbug/1108198): we don't need this after removing the /ui prefix.

// Eslint isn't able to type-check webworker scripts.
/* eslint-disable */

self.addEventListener('fetch', (event) => {
  const url = new URL(event.request.url);

  const isResultUI =
    // Short build link.
    url.pathname.match(/^\/b\//) ||
    // Long build link.
    url.pathname.match(/^\/p\/[^/]+\/builders\/[^/]+\/[^/]+\//) ||
    // Invocation link.
    url.pathname.match(/^\/inv\//) ||
    // Artifact link.
    url.pathname.match(/^\/artifact\//);

  if (isResultUI) {
    url.pathname = '/ui' + url.pathname;
    event.respondWith(Response.redirect(url.toString()));
    return;
  }
  event.respondWith(fetch(event.request));
});
