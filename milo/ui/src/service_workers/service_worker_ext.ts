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

import './force_update';
import { getAuthStateCache, getAuthStateCacheSync, setAuthStateCache } from '../auth_state_cache';
import { PrpcClientExt } from '../libs/prpc_client_ext';
import { queryAuthState } from '../services/milo_internal';
import { ResultDb } from '../services/resultdb';
import { Prefetcher } from './prefetch';

importScripts('/configs.js');

// TSC isn't able to determine the scope properly.
// Perform manual casting to fix typing.
const _self = self as unknown as ServiceWorkerGlobalScope;

const prefetcher = new Prefetcher(CONFIGS, _self.fetch.bind(_self));

_self.addEventListener('fetch', async (e) => {
  if (prefetcher.respondWithPrefetched(e)) {
    return;
  }

  const url = new URL(e.request.url);

  // Ensure all clients served by this service worker use the same config.
  if (url.pathname === '/configs.js') {
    const res = new Response(`const CONFIGS=${JSON.stringify(CONFIGS)};`);
    res.headers.set('content-type', 'application/javascript');
    e.respondWith(res);
    return;
  }

  const rawArtifactMatch = /^\/ui\/artifact\/raw\/(.*)$/i.exec(url.pathname);
  if (rawArtifactMatch) {
    const artifactName = rawArtifactMatch[1];
    e.respondWith(handleRawArtifact(artifactName));
    return;
  }

  prefetcher.prefetchResources(url);
});

const resultDb = new ResultDb(
  new PrpcClientExt(
    { host: CONFIGS.RESULT_DB.HOST, fetchImpl: _self.fetch.bind(_self) },
    () => getAuthStateCacheSync()?.accessToken || ''
  )
);

/**
 * Handles raw artifact requests.
 */
async function handleRawArtifact(artifactName: string): Promise<Response> {
  const authState = await getAuthStateCache();
  if (!authState) {
    await setAuthStateCache(await queryAuthState());
  }

  // TODO(weiweilin): if needed, we can cache Artifact until
  // Artifact.urlExpiration.
  const artifact = await resultDb.getArtifact({ name: artifactName }, { acceptCache: false, skipUpdate: true });
  return fetch(artifact.fetchUrl);
}
