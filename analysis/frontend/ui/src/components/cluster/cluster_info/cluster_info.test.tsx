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

import fetchMock from 'fetch-mock-jest';

import { screen } from '@testing-library/react';

import { renderWithRouterAndClient } from '@/testing_tools/libs/mock_router';
import { mockFetchAuthState } from '@/testing_tools/mocks/authstate_mock';
import {
  getMockCluster,
  mockBatchGetCluster,
} from '@/testing_tools/mocks/cluster_mock';

import ClusterInfo from './cluster_info';

describe('test ClusterInfo component', () => {
  beforeEach(() => {
    mockFetchAuthState();
  });
  afterEach(() => {
    fetchMock.mockClear();
    fetchMock.reset();
  });

  it('Given reason based cluster then should display the data', async () => {
    const project = 'chromium';
    const algorithm = 'reason-v3';
    const id = '14ee3dde813f66adc0595e4a21aa1743';
    const mockCluster = getMockCluster(id, project, algorithm, 'ninja://chrome/android');

    mockBatchGetCluster(project, algorithm, id, mockCluster);

    renderWithRouterAndClient(
        <ClusterInfo />,
        `/p/${project}/clusters/${algorithm}/${id}`,
        '/p/:project/clusters/:algorithm/:id',
    );

    await screen.findByText('Failure reason cluster');

    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    expect(screen.getByText(mockCluster.title!)).toBeInTheDocument();
  });

  it('Given test based cluster then should display the data', async () => {
    const project = 'chromium';
    const algorithm = 'testname-v3';
    const id = '14ee3dde813f66adc0595e4a21aa1743';
    const mockCluster = getMockCluster(id, project, algorithm, 'ninja://chrome/android');

    mockBatchGetCluster(project, algorithm, id, mockCluster);

    renderWithRouterAndClient(
        <ClusterInfo />,
        `/p/${project}/clusters/${algorithm}/${id}`,
        '/p/:project/clusters/:algorithm/:id',
    );

    await screen.findByText('Test name cluster');

    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    expect(screen.getByText(mockCluster.title!)).toBeInTheDocument();
  });
});
