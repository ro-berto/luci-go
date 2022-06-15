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

import {
  axisRight,
  NumberValue,
  scaleTime,
  select as d3Select,
  timeHour,
  timeMillisecond,
  timeMinute,
  timeSecond,
} from 'd3';
import { css, customElement, html } from 'lit-element';
import { computed, observable } from 'mobx';

import { MiloBaseElement } from '../../components/milo_base';
import { consumeTestHistoryPageState, TestHistoryPageState } from '../../context/test_history_page_state';
import { HOUR_MS, MINUTE_MS, PREDEFINED_TIME_INTERVALS, SECOND_MS } from '../../libs/constants';
import { consumer } from '../../libs/context';
import { roundUp } from '../../libs/utils';
import { CELL_SIZE, X_AXIS_HEIGHT } from './constants';

const DURATION_LEGEND_PADDING = 10;
const DURATION_RECT_WIDTH = 20;
const Y_AXIS_WIDTH = 50;
const MIN_TICK_SIZE = 20;

@customElement('milo-th-duration-legend')
@consumer
export class TestHistoryDurationLegendElement extends MiloBaseElement {
  @observable.ref @consumeTestHistoryPageState() pageState!: TestHistoryPageState;

  @computed private get scaleDurationY() {
    return scaleTime()
      .range([0, this.scaleHeight])
      .domain([this.pageState.maxDurationMs, this.pageState.minDurationMs]);
  }

  @computed private get scaleHeight() {
    return this.pageState.filteredVariants.length * CELL_SIZE + X_AXIS_HEIGHT - 2 * DURATION_LEGEND_PADDING;
  }

  @computed private get tickStepSizeMs() {
    const durationDiff = this.pageState.maxDurationMs - this.pageState.minDurationMs;
    const durationLegendHeight = this.scaleDurationY.range()[1];
    const minInterval = (durationDiff / durationLegendHeight) * MIN_TICK_SIZE;
    return roundUp(minInterval, PREDEFINED_TIME_INTERVALS);
  }

  @computed private get tickInterval() {
    if (this.tickStepSizeMs >= HOUR_MS) {
      return timeHour.every(this.tickStepSizeMs / HOUR_MS);
    }
    if (this.tickStepSizeMs >= MINUTE_MS) {
      return timeMinute.every(this.tickStepSizeMs / MINUTE_MS);
    }
    if (this.tickStepSizeMs >= SECOND_MS) {
      return timeSecond.every(this.tickStepSizeMs / SECOND_MS);
    }
    return timeMillisecond.every(this.tickStepSizeMs);
  }

  @computed private get tickFormat() {
    if (this.tickStepSizeMs >= HOUR_MS) {
      return (v: NumberValue) => v.valueOf() / HOUR_MS + 'hr';
    }
    if (this.tickStepSizeMs >= MINUTE_MS) {
      return (v: NumberValue) => v.valueOf() / MINUTE_MS + 'min';
    }
    if (this.tickStepSizeMs >= SECOND_MS) {
      return (v: NumberValue) => v.valueOf() / SECOND_MS + 's';
    }
    return (v: NumberValue) => v.valueOf() + 'ms';
  }

  @computed private get durationAxisEle() {
    return d3Select(document.createElementNS('http://www.w3.org/2000/svg', 'g')).call(
      axisRight(this.scaleDurationY).ticks(this.tickInterval).tickFormat(this.tickFormat)
    );
  }

  protected render() {
    const variants = this.pageState.filteredVariants;
    return html`
      <svg height=${X_AXIS_HEIGHT + CELL_SIZE * variants.length} width=${Y_AXIS_WIDTH + DURATION_RECT_WIDTH}>
        <defs>
          <linearGradient id="duration-gradient" gradientTransform="rotate(90)">
            <stop offset="0%" stop-color=${this.pageState.scaleDurationColor(this.pageState.maxDurationMs)}></stop>
            <stop
              offset="50%"
              stop-color=${this.pageState.scaleDurationColor(
                (this.pageState.maxDurationMs + this.pageState.minDurationMs) / 2
              )}
            ></stop>
            <stop offset="100%" stop-color=${this.pageState.scaleDurationColor(this.pageState.minDurationMs)}></stop>
          </linearGradient>
        </defs>
        <g transform="translate(${DURATION_RECT_WIDTH}, ${DURATION_LEGEND_PADDING})">
          <rect
            x=${-DURATION_RECT_WIDTH}
            y="-0.5"
            width=${DURATION_RECT_WIDTH}
            height=${this.scaleDurationY.range()[1] + 1}
            fill="url(#duration-gradient)"
          ></rect>
          ${this.durationAxisEle}
        </g>
      </svg>
    `;
  }

  static styles = css`
    :host {
      width: fit-content;
      height: fit-content;
    }
  `;
}
