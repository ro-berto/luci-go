// Copyright 2022 The LUCI Authors.
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

import { sassPlugin } from 'esbuild-sass-plugin';
import esbuild from 'esbuild';

esbuild.build({
  entryPoints: ['index.tsx'],
  bundle: true,
  inject: ['src/tools/react_shim.ts'],
  outfile: 'dist/main.js',
  minify: true,
  sourcemap: true,
  plugins: [sassPlugin()],
  loader: {
    '.png': 'dataurl',
    '.woff': 'dataurl',
    '.woff2': 'dataurl',
    '.eot': 'dataurl',
    '.ttf': 'dataurl',
    '.svg': 'dataurl',
  },
// eslint-disable-next-line @typescript-eslint/no-unused-vars
}).catch((_) => {
  // eslint-disable-next-line no-undef
  process.exit(1);
});
