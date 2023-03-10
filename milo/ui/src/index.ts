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

import { configure } from 'mobx';

import './routes';
import './stackdriver_errors';
import { initDefaultTrustedTypesPolicy } from './libs/sanitize_html';

initDefaultTrustedTypesPolicy();

// TODO(crbug/1347294): encloses all state modifying actions in mobx actions
// then delete this.
configure({ enforceActions: 'never' });

// Reload the page after a new version is activated to avoid different versions
// of the code talking to each other.
navigator.serviceWorker?.addEventListener('controllerchange', () => window.location.reload());
