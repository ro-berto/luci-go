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

import { fixture, fixtureCleanup, html } from '@open-wc/testing/index-no-side-effects';
import { expect } from 'chai';

import './step_cluster';
import { BuildStatus } from '../../../services/buildbucket';
import { Store } from '../../../store';
import { StepExt } from '../../../store/build_state';
import { BuildPageStepClusterElement } from './step_cluster';

function createStep(index: number, startTime: string, endTime: string) {
  return new StepExt({
    step: {
      name: 'step' + index,
      status: BuildStatus.Success,
      startTime,
      endTime,
    },
    listNumber: index + 1 + '.',
    selfName: 'step' + index,
    depth: 0,
    index,
  });
}

const step1 = createStep(0, '2022-01-01T00:00:00Z', '2022-01-01T00:02:00Z');
const step2 = createStep(0, '2022-01-01T00:01:00Z', '2022-01-01T00:03:00Z');
const step3 = createStep(0, '2022-01-01T00:02:00Z', '2022-01-01T00:04:00Z');

describe('bp_step_cluster', () => {
  it('should calculate duration based on start & end time', async () => {
    const ele = await fixture<BuildPageStepClusterElement>(html`
      <milo-bp-step-cluster .store=${Store.create()} .steps=${[step1, step2, step3]}></milo-bp-step-cluster>
    `);
    after(fixtureCleanup);

    // The duration should equals endTime - startTime. Not a sum of all
    // durations. Because steps can run in parallel.
    expect(ele.duration?.toMillis()).to.eq(240000);
  });
});
