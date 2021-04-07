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

import { Workbox } from 'workbox-window';

import './routes';

window.SW_PROMISE = new Promise((resolve) => {
  // Don't cache resources in development mode. Otherwise we will need to
  // refresh the page manually for changes to take effect.
  if (ENABLE_UI_SW && 'serviceWorker' in navigator) {
    const wb = new Workbox('/ui/service-worker.js');
    wb.register({ immediate: true }).then((registration) => {
      // eslint-disable-next-line no-console
      console.log('UI SW registered: ', registration);
      resolve(wb);
    });
  }
});

if ('serviceWorker' in navigator) {
  if (!document.cookie.includes('showNewBuildPage=false')) {
    window.addEventListener(
      'load',
      async () => {
        const registration = await navigator.serviceWorker.register('/redirect-sw.js');
        // eslint-disable-next-line no-console
        console.log('Redirect SW registered: ', registration);
      },
      { once: true }
    );
  }
}
