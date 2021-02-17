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

package base

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"go.starlark.net/starlark"
)

// ExpandDirectories recursively traverses directories in `paths` discovering
// *.star files in them.
//
// If `paths` is empty, expands `.`.
//
// Returns the overall list of discovered files.
func ExpandDirectories(paths []string) ([]string, error) {
	if len(paths) == 0 {
		paths = []string{"."}
	}
	var files []string
	for _, p := range paths {
		switch info, err := os.Stat(p); {
		case err != nil:
			return nil, err
		case !info.IsDir():
			files = append(files, p)
		default:
			err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
				if err == nil && !info.IsDir() && filepath.Ext(info.Name()) == ".star" {
					files = append(files, path)
				}
				return err
			})
			if err != nil {
				return nil, err
			}
		}
	}
	return files, nil
}

// PathLoader is an interpreter.Loader that loads files using file system paths.
func PathLoader(path string) (starlark.StringDict, string, error) {
	body, err := ioutil.ReadFile(path)
	return nil, string(body), err
}
