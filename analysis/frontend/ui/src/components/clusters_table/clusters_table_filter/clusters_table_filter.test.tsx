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

import '@testing-library/jest-dom';

import {
  render,
  screen,
} from '@testing-library/react';

import { identityFunction } from '@/testing_tools/functions';

import ClustersTableFilter from './clusters_table_filter';

describe('Test ClustersTableFilter component', () => {
  it('should display the failures filter', async () => {
    render(
        <ClustersTableFilter
          failureFilter=""
          setFailureFilter={identityFunction}/>,
    );

    await screen.findByTestId('clusters_table_filter');

    expect(screen.getByTestId('failure_filter')).toBeInTheDocument();
  });

  it('given an existing filter, the filter should be pre-populated', async () => {
    render(
        <ClustersTableFilter
          failureFilter="some restriction"
          setFailureFilter={identityFunction}/>,
    );

    await screen.findByTestId('clusters_table_filter');

    expect(screen.getByTestId('failure_filter_input')).toHaveValue('some restriction');
  });
});
