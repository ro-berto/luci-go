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

import Workbox from 'workbox-webpack-plugin';
import { DefinePlugin } from 'webpack';
import merge from 'webpack-merge';

import common from './webpack.common';

export default merge(common, {
  devtool: 'source-map',
  mode: 'production',
  plugins: [
    new DefinePlugin({ PRODUCTION: JSON.stringify(true) }),
    new Workbox.GenerateSW({
      clientsClaim: true,
      navigateFallback: '/ui/index.html',
      // Workbox source map changes every build.
      // This causes noise in the auto-roller.
      // https://github.com/GoogleChrome/workbox/issues/2784
      sourcemap: true,
    }),
  ],
});
