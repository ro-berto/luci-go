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


import './culprit_verification_table.css';

import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';

import { CulpritVerificationTableRow } from './culprit_verification_table_row/culprit_verification_table_row';

import {
  Analysis,
  HeuristicSuspect,
} from '../../services/luci_bisection';

interface Props {
  result?: Analysis;
}

function getRows(suspects: HeuristicSuspect[]) {
  return suspects.map((suspect) => (
    <CulpritVerificationTableRow
      key={suspect.gitilesCommit.id}
      suspect={suspect}
    />
  ));
}

export const CulpritVerificationTable = ({ result }: Props) => {
  // TODO: Support nth-section suspects
  const suspects = result?.heuristicResult?.suspects ?? []
  if (!suspects) {
    return <>Could not find any culprit verification result</>
  }
  return (
    <TableContainer component={Paper} className='culprit-verification-table-container'>
      <Table className='culprit-verification-table' size='small'>
        <TableHead>
          <TableRow>
            <TableCell>Suspect CL</TableCell>
            <TableCell>Type</TableCell>
            <TableCell>Verification Status</TableCell>
            <TableCell>Reruns</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {getRows(suspects)}
        </TableBody>
      </Table>
    </TableContainer>
  );
};
