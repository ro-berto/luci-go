// Copyright 2023 The LUCI Authors.
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
 * This adds testing-library matchers to jest.
 */
import '@testing-library/jest-dom';

import 'node-fetch';

/**
 * This import must be declared before any use of `fetch`
 * in your code if you want to mock the `fetch` calls.
 *
 * If any call is made to `fetch` before this gets to run (for example,
 * if you run it in a constructor of a class), your mocks will fail.
 */
import fetchMock from 'fetch-mock-jest';

import {
  render,
  screen,
} from '@testing-library/react';

import { QueryClient, QueryClientProvider } from 'react-query';
import ExampleReactQuery from './example_react_query';

describe('<Example />', () => {
  afterEach(() => {
    fetchMock.mockClear();
    fetchMock.reset();
  });

  it('Wait for component to load', async () => {
    fetchMock.mock('/path/to/api', {});
    const client = new QueryClient({
      defaultOptions: {
        queries: {
          retry: false,
        },
      },
    });
    render(
        <QueryClientProvider client={client}>
          <ExampleReactQuery />
        </QueryClientProvider>,
    );

    await screen.findByRole('article');

    expect(screen.getByTestId('test-component')).toBeInTheDocument();
  });
});
