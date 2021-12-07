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

import { css, customElement, html, svg } from 'lit-element';
import { DateTime, Duration } from 'luxon';
import { observable } from 'mobx';

import { MiloBaseElement } from '../../components/milo_base';
import { consumeTestHistoryPageState, TestHistoryPageState } from '../../context/test_history_page_state';
import { consumer } from '../../libs/context';
import { displayDuration, parseProtoDuration } from '../../libs/time_utils';
import commonStyle from '../../styles/common_style.css';
import { CELL_PADDING, CELL_SIZE, INNER_CELL_SIZE } from './constants';

@customElement('milo-th-duration-graph')
@consumer
export class TestHistoryDurationGraphElement extends MiloBaseElement {
  @observable.ref @consumeTestHistoryPageState() pageState!: TestHistoryPageState;

  protected render() {
    const variants = this.pageState.filteredVariants;
    return html`
      <svg id="graph" height=${CELL_SIZE * variants.length}>
        ${variants.map(
          ([vHash], i) => svg`
            <g transform="translate(1, ${i * CELL_SIZE})">
              <rect
                x="-1"
                height=${CELL_SIZE}
                width=${CELL_SIZE * this.pageState.days + 2}
                fill=${i % 2 === 0 ? 'var(--block-background-color)' : 'transparent'}
              />
              ${this.pageState.dates.map((d, j) => this.renderEntries(vHash, d, j))}
            </g>
          `
        )}
      </svg>
    `;
  }

  private renderEntries(vHash: string, date: DateTime, index: number) {
    const entries = this.pageState.testHistoryLoader!.getEntries(vHash, date)?.filter(this.pageState.tvhEntryFilter);
    const durations = this.pageState.passOnlyDuration
      ? entries?.filter((e) => e.avgDurationPass).map((e) => e.avgDurationPass!)
      : entries?.filter((e) => e.avgDuration).map((e) => e.avgDuration!);
    if (!durations || durations.length === 0) {
      return null;
    }

    const averageDurationMs =
      durations.map((duration) => parseProtoDuration(duration)).reduce((duration, total) => total + duration, 0) /
      durations.length;

    if (!this.pageState.durationInitialized) {
      this.pageState.maxDurationMs = averageDurationMs;
      this.pageState.minDurationMs = averageDurationMs;
      this.pageState.durationInitialized = true;
    } else {
      if (averageDurationMs > this.pageState.maxDurationMs) {
        this.pageState.maxDurationMs = averageDurationMs;
      }
      if (averageDurationMs < this.pageState.minDurationMs) {
        this.pageState.minDurationMs = averageDurationMs;
      }
    }

    return svg`
      <rect
        x=${index * CELL_SIZE + CELL_PADDING}
        y=${CELL_PADDING}
        width=${INNER_CELL_SIZE}
        height=${INNER_CELL_SIZE}
        fill=${this.pageState.scaleDurationColor(averageDurationMs)}
        style="cursor: pointer;"
        @click=${() => {
          this.pageState.selectedTvhEntries = entries!;
        }}
      >
        <title>Average Duration: ${displayDuration(
          Duration.fromMillis(averageDurationMs)
        )}\nClick to view test run details.</title>
      </rect>
    `;
  }

  static styles = [
    commonStyle,
    css`
      #graph {
        width: 100%;
      }

      .count-label {
        fill: white;
        text-anchor: middle;
        alignment-baseline: central;
      }
    `,
  ];
}
