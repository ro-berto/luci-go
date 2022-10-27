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
  waitForElementToBeRemoved,
} from '@testing-library/react';

import { ExoneratedTestVariantBuilder } from '../model/mocks';

import ExonerationsTableRow from './exonerations_table_row';

describe('Test ExonerationsTableRows component', () => {
  it('displays test variant statistics', async () => {
    const testVariant = new ExoneratedTestVariantBuilder().almostMeetsFailureThreshold().build();

    render(
        <table>
          <tbody>
            <ExonerationsTableRow
              project='testproject'
              testVariant={testVariant}/>
          </tbody>
        </table>,
    );

    expect(await screen.findByText('someTestId')).toBeInTheDocument();
    expect(screen.getByText('keya: valuea, keyb: valueb')).toBeInTheDocument();
    expect(screen.getByText('No, but close to')).toBeInTheDocument();
    expect(screen.getByText('100001')).toBeInTheDocument();
  });

  it('details popup opens and closes', async () => {
    const testVariant = new ExoneratedTestVariantBuilder().almostMeetsFlakyThreshold().build();

    render(
        <table>
          <tbody>
            <ExonerationsTableRow
              project='testproject'
              testVariant={testVariant}/>
          </tbody>
        </table>,
    );

    expect(await screen.findByText('someTestId')).toBeInTheDocument();

    // Open dialog.
    screen.getByText('Why?').click();
    expect(await screen.findByText('Why is this test variant close to being exonerated?')).toBeInTheDocument();

    // Close dialog.
    screen.getByText('Close').click();
    await waitForElementToBeRemoved(() => screen.queryByText('Why is this test variant close to being exonerated?'));
    expect(screen.queryByText('Why is this test variant close to being exonerated?')).toBeNull();
  });
});
