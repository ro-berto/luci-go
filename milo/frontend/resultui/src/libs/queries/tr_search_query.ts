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

import { html } from 'lit-html';

import { Suggestion } from '../../components/auto_complete';
import { TestVariant } from '../../services/resultdb';
import { highlight } from '../lit_utils';
import { parseProtoDuration } from '../time_utils';
import { KV_SYNTAX_EXPLANATION, parseKeyValue } from './utils';

const SPECIAL_QUERY_RE = /^(-?)([a-zA-Z]+):(.+)$/;

export type TestVariantFilter = (v: TestVariant) => boolean;

export function parseTestResultSearchQuery(searchQuery: string): TestVariantFilter {
  const filters = searchQuery.split(' ').map((query) => {
    const match = query.match(SPECIAL_QUERY_RE);

    const [, neg, type, value] = match || ['', '', '', query];
    const valueUpper = value.toUpperCase();
    const negate = neg === '-';
    switch (type.toUpperCase()) {
      // Whether the test ID or test name contains the query as a substring
      // (case insensitive).
      case '': {
        return (v: TestVariant) => {
          const matched =
            v.testId.toUpperCase().includes(valueUpper) || v.testMetadata?.name?.toUpperCase().includes(valueUpper);
          return negate !== Boolean(matched);
        };
      }
      // Whether the test variant has the specified status.
      case 'STATUS': {
        const statuses = valueUpper.split(',');
        return (v: TestVariant) => negate !== statuses.includes(v.status);
      }
      // Whether there's at least one a test result of the specified status.
      case 'RSTATUS': {
        const statuses = valueUpper.split(',');
        return (v: TestVariant) => negate !== (v.results || []).some((r) => statuses.includes(r.result.status));
      }
      // Whether the test ID contains the query as a substring (case
      // insensitive).
      case 'ID': {
        return (v: TestVariant) => negate !== v.testId.toUpperCase().includes(valueUpper);
      }
      // Whether the test ID matches the specified ID (case sensitive).
      case 'EXACTID': {
        return (v: TestVariant) => negate !== (v.testId === value);
      }
      // Whether the test variant has a matching variant key-value pair.
      case 'V': {
        const [vKey, vValue] = parseKeyValue(value);

        // Otherwise, the value must match the specified value (case sensitive).
        return vValue === null
          ? (v: TestVariant) => negate !== (v.variant?.def?.[vKey] !== undefined)
          : (v: TestVariant) => negate !== (v.variant?.def?.[vKey] === vValue);
      }
      // Whether the test variant has the specified variant hash.
      case 'VHASH': {
        return (v: TestVariant) => negate !== (v.variantHash.toUpperCase() === valueUpper);
      }
      // Whether the test name contains the query as a substring (case
      // insensitive).
      case 'NAME': {
        return (v: TestVariant) => negate !== (v.testMetadata?.name || '').toUpperCase().includes(valueUpper);
      }
      // Whether the test name matches the specified name (case sensitive).
      case 'EXACTNAME': {
        return (v: TestVariant) => negate !== (v.testMetadata?.name === value);
      }
      // Whether the test has a run with a matching tag (case sensitive).
      case 'TAG': {
        const [tKey, tValue] = parseKeyValue(value);

        if (tValue) {
          return (v: TestVariant) =>
            negate === !v.results?.some((r) => r.result.tags?.some((t) => t.key === tKey && t.value === tValue));
        } else {
          return (v: TestVariant) => negate === !v.results?.some((r) => r.result.tags?.some((t) => t.key === tKey));
        }
      }
      // Whether the test has at least one run with a duration in the specified
      // range.
      case 'DURATION': {
        const match = value.match(/^(\d+(?:\.\d+)?)-(\d+(?:\.\d+)?)?$/);
        if (!match) {
          throw new Error(`invalid duration range: ${value}`);
        }
        const [, minDurationStr, maxDurationStr] = match;
        const minDuration = Number(minDurationStr) * 1000;
        const maxDuration = maxDurationStr ? Number(maxDurationStr || '0') * 1000 : Infinity;
        return (v: TestVariant) =>
          negate ===
          !v.results?.some((r) => {
            if (!r.result.duration) {
              return false;
            }
            const durationMs = parseProtoDuration(r.result.duration);
            return durationMs >= minDuration && durationMs <= maxDuration;
          });
      }
      default: {
        throw new Error(`invalid query type: ${type}`);
      }
    }
  });
  return (v) => filters.every((f) => f(v));
}

// Queries with predefined value.
const QUERY_SUGGESTIONS = [
  { value: 'Status:UNEXPECTED', explanation: 'Include only tests with unexpected status' },
  { value: '-Status:UNEXPECTED', explanation: 'Exclude tests with unexpected status' },
  { value: 'Status:UNEXPECTEDLY_SKIPPED', explanation: 'Include only tests with unexpectedly skipped status' },
  { value: '-Status:UNEXPECTEDLY_SKIPPED', explanation: 'Exclude tests with unexpectedly skipped status' },
  { value: 'Status:FLAKY', explanation: 'Include only tests with flaky status' },
  { value: '-Status:FLAKY', explanation: 'Exclude tests with flaky status' },
  { value: 'Status:EXONERATED', explanation: 'Include only tests with exonerated status' },
  { value: '-Status:EXONERATED', explanation: 'Exclude tests with exonerated status' },
  { value: 'Status:EXPECTED', explanation: 'Include only tests with expected status' },
  { value: '-Status:EXPECTED', explanation: 'Exclude tests with expected status' },

  { value: 'RStatus:Pass', explanation: 'Include only tests with at least one passed run' },
  { value: '-RStatus:Pass', explanation: 'Exclude tests with at least one passed run' },
  { value: 'RStatus:Fail', explanation: 'Include only tests with at least one failed run' },
  { value: '-RStatus:Fail', explanation: 'Exclude tests with at least one failed run' },
  { value: 'RStatus:Crash', explanation: 'Include only tests with at least one crashed run' },
  { value: '-RStatus:Crash', explanation: 'Exclude tests with at least one crashed run' },
  { value: 'RStatus:Abort', explanation: 'Include only tests with at least one aborted run' },
  { value: '-RStatus:Abort', explanation: 'Exclude tests with at least one aborted run' },
  { value: 'RStatus:Skip', explanation: 'Include only tests with at least one skipped run' },
  { value: '-RStatus:Skip', explanation: 'Exclude tests with at least one skipped run' },
];

// Queries with arbitrary value.
const QUERY_TYPE_SUGGESTIONS = [
  {
    type: 'V:',
    explanation: `Include only tests with a matching variant key-value pair (${KV_SYNTAX_EXPLANATION})`,
  },
  {
    type: '-V:',
    explanation: `Exclude tests with a matching variant key-value pair (${KV_SYNTAX_EXPLANATION})`,
  },

  {
    type: 'Tag:',
    explanation: `Include only tests with a run that has a matching tag key-value pair (${KV_SYNTAX_EXPLANATION})`,
  },
  {
    type: '-Tag:',
    explanation: `Exclude tests with a run that has a matching tag key-value pair (${KV_SYNTAX_EXPLANATION})`,
  },

  { type: 'ID:', explanation: 'Include only tests with the specified substring in their ID (case insensitive)' },
  { type: '-ID:', explanation: 'Exclude tests with the specified substring in their ID (case insensitive)' },

  { type: 'Name:', explanation: 'Include only tests with the specified substring in their Name (case insensitive)' },
  { type: '-Name:', explanation: 'Exclude tests with the specified substring in their Name (case insensitive)' },

  { type: 'ExactID:', explanation: 'Include only tests with the specified ID (case sensitive)' },
  { type: '-ExactID:', explanation: 'Exclude tests with the specified ID (case sensitive)' },

  { type: 'Duration:', explanation: 'Include only tests with a run that has a duration in the specified range' },
  { type: '-Duration:', explanation: 'Exclude tests with a run that has a duration in the specified range' },

  { type: 'ExactName:', explanation: 'Include only tests with the specified name (case sensitive)' },
  { type: '-ExactName:', explanation: 'Exclude tests with the specified name (case sensitive)' },

  { type: 'VHash:', explanation: 'Include only tests with the specified variant hash' },
  { type: '-VHash:', explanation: 'Exclude tests with the specified variant hash' },
];

export function suggestTestResultSearchQuery(query: string): readonly Suggestion[] {
  if (query === '') {
    // Return some example queries when the query is empty.
    return [
      {
        isHeader: true,
        display: html`<strong>Advanced Syntax</strong>`,
      },
      {
        value: '-Status:EXPECTED',
        explanation: "Use '-' prefix to negate the filter",
      },
      {
        value: 'Status:UNEXPECTED -RStatus:Skipped',
        explanation: 'Use space to separate filters. Filters are logically joined with AND',
      },

      // Put this section behind `Advanced Syntax` so `Advanced Syntax` won't
      // be hidden after the size of supported filter types grows.
      {
        isHeader: true,
        display: html`<strong>Supported Filter Types</strong>`,
      },
      {
        value: 'test-id-substr',
        explanation: 'Include only tests with the specified substring in their ID or name (case insensitive)',
      },
      {
        value: 'V:query-encoded-variant-key=query-encoded-variant-value',
        explanation: 'Include only tests with a matching test variant key-value pair (case sensitive)',
      },
      {
        value: 'V:query-encoded-variant-key',
        explanation: 'Include only tests with the specified variant key (case sensitive)',
      },
      {
        value: 'Tag:query-encoded-tag-key=query-encoded-tag-value',
        explanation: 'Include only tests with a run that has a matching tag key-value pair (case sensitive)',
      },
      {
        value: 'Tag:query-encoded-tag-key',
        explanation: 'Include only tests with a run that has the specified tag key (case sensitive)',
      },
      {
        value: 'ID:test-id-substr',
        explanation: 'Include only tests with the specified substring in their ID (case insensitive)',
      },
      {
        value: 'Status:UNEXPECTED,UNEXPECTEDLY_SKIPPED,FLAKY,EXONERATED,EXPECTED',
        explanation: 'Include only tests with the specified status',
      },
      {
        value: 'RStatus:Pass,Fail,Crash,Abort,Skip',
        explanation: 'Include only tests with at least one run of the specified status',
      },
      {
        value: 'Name:test-name-substr',
        explanation: 'Include only tests with the specified substring in their name (case insensitive)',
      },
      {
        value: 'Duration:0.05-15',
        explanation: 'Include only tests with a run that has a duration in the specified range (in seconds)',
      },
      {
        value: 'Duration:0.05-',
        explanation: 'Max duration can be omitted',
      },
      {
        value: 'ExactID:test-id',
        explanation: 'Include only tests with the specified test ID (case sensitive)',
      },
      {
        value: 'ExactName:test-name',
        explanation: 'Include only tests with the specified name (case sensitive)',
      },
      {
        value: 'VHash:2660cde9da304c42',
        explanation: 'Include only tests with the specified variant hash',
      },
    ];
  }

  const subQuery = query.split(' ').pop()!;
  if (subQuery === '') {
    return [];
  }

  const suggestions: Suggestion[] = [];

  // Suggest queries with predefined value.
  const subQueryUpper = subQuery.toUpperCase();
  suggestions.push(...QUERY_SUGGESTIONS.filter(({ value }) => value.toUpperCase().includes(subQueryUpper)));

  // Suggest queries with arbitrary value.
  const match = subQuery.match(/^([^:]*:?)(.*)$/);
  if (match) {
    const [, subQueryType, subQueryValue] = match as [string, string, string];
    const typeUpper = subQueryType.toUpperCase();
    suggestions.push(
      ...QUERY_TYPE_SUGGESTIONS.flatMap(({ type, explanation }) => {
        if (type.toUpperCase().includes(typeUpper)) {
          return [{ value: type + subQueryValue, explanation }];
        }

        if (subQueryValue === '') {
          return [{ value: type + subQueryType, explanation }];
        }

        return [];
      })
    );
  }

  return suggestions.map((s) => ({ ...s, display: s.display || highlight(s.value!, subQuery) }));
}
