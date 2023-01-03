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

import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TableSortLabel from '@mui/material/TableSortLabel';

import {
  MetricId,
} from '@/services/shared_models';

import {
  Metric,
} from '@/services/metrics';

interface Props {
    orderBy?: OrderBy,
    metrics: Metric[],
    handleOrderByChanged: (orderBy: OrderBy) => void,
}

export interface OrderBy {
  metric: MetricId,
  isAscending: boolean,
}

const ClustersTableHead = ({
  orderBy,
  metrics,
  handleOrderByChanged,
}: Props) => {
  const toggleSort = (metric: MetricId) => {
    if (orderBy && orderBy.metric === metric) {
      handleOrderByChanged({
        metric: metric,
        isAscending: !orderBy.isAscending,
      });
    } else {
      handleOrderByChanged({
        metric: metric,
        isAscending: false,
      });
    }
  };

  return (
    <TableHead data-testid="clusters_table_head">
      <TableRow>
        <TableCell>Cluster</TableCell>
        <TableCell sx={{ width: '150px' }}>Bug</TableCell>
        {
          metrics.map((metric: Metric) => (
            <TableCell
              key={metric.metricId}
              sortDirection={(orderBy && (orderBy.metric === metric.metricId)) ? (orderBy.isAscending ? 'asc' : 'desc') : false}
              sx={{ cursor: 'pointer', width: '100px' }}>
              <TableSortLabel
                aria-label={`Sort by ${metric.humanReadableName}`}
                active={orderBy && (orderBy.metric === metric.metricId)}
                direction={(orderBy && orderBy.isAscending) ? 'asc' : 'desc'}
                onClick={() => toggleSort(metric.metricId)}>
                {metric.humanReadableName}
              </TableSortLabel>
            </TableCell>
          ))
        }
      </TableRow>
    </TableHead>
  );
};

export default ClustersTableHead;
