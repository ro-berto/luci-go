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

// TODO(weiweilin): add integration tests to ensure the SW works properly.

import { getAuthState, setAuthState } from './auth_state';
import { cached } from './libs/cached_fn';
import { genCacheKeyForPrpcRequest } from './libs/prpc_utils';
import { timeout } from './libs/utils';
import { CACHED_PRPC_URLS, prefetchResources } from './prefetch';

importScripts('/configs.js');

// TSC isn't able to determine the scope properly.
// Perform manual casting to fix typing.
const _self = (self as unknown) as ServiceWorkerGlobalScope;

const PRPC_CACHE_KEY_PREFIX = 'prpc-cache-key';

export interface SetAuthStateEventData {
  type: 'SET_AUTH_STATE';
  authState: AuthState | null;
}

const cachedFetch = cached(
  // _cacheKey and _expiresIn are not used here but is used in the expire
  // and key functions below.
  // they are listed here to help TSC generates the correct type definition.
  (info: RequestInfo, init: RequestInit | undefined, _cacheKey: unknown, _expiresIn: number) => fetch(info, init),
  {
    key: (_info, _init, cacheKey) => cacheKey,
    expire: ([, , , expiresIn]) => timeout(expiresIn),
  }
);

_self.addEventListener('message', async (e) => {
  switch (e.data.type) {
    case 'SET_AUTH_STATE': {
      const data = e.data as SetAuthStateEventData;
      setAuthState(data.authState);
      break;
    }
    default:
      console.warn('unexpected message type', e.data.type, e.data, e);
  }
});

_self.addEventListener('fetch', async (e) => {
  const url = new URL(e.request.url);
  // Serve cached auth data.
  if (url.pathname === '/ui/cached-auth-state.js') {
    e.respondWith(
      (async () => {
        const authState = await getAuthState();
        return new Response(`const CACHED_AUTH_STATE=${JSON.stringify(authState)};`, {
          headers: { 'content-type': 'application/javascript' },
        });
      })()
    );
  }

  // Ensure all clients served by this service worker use the same config.
  if (url.pathname === '/configs.js') {
    const res = new Response(`const CONFIGS=${JSON.stringify(CONFIGS)};`);
    res.headers.set('content-type', 'application/javascript');
    e.respondWith(res);
  }

  prefetchResources(url);
  if (CACHED_PRPC_URLS.includes(e.request.url)) {
    e.respondWith(
      (async () => {
        const res = await cachedFetch(
          // The response can't be reused, don't keep it in cache.
          { skipUpdate: true, invalidateCache: true },
          e.request,
          undefined,
          await genCacheKeyForPrpcRequest(PRPC_CACHE_KEY_PREFIX, e.request.clone()),
          0
        );
        return res;
      })()
    );
  }
});
