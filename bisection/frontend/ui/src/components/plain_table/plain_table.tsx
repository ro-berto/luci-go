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


import { styled } from '@mui/material/styles';

import Table from '@mui/material/Table';
import { tableCellClasses } from '@mui/material/TableCell';

export const PlainTable = styled(Table)({
  [`& .${tableCellClasses.head}`]: {
    fontSize: '1rem',
    fontWeight: 'normal',
    color: 'dimgray',
    opacity: '80%',
    border: 'none',
    padding: 0,
  },
  [`& .${tableCellClasses.body}`]: {
    fontSize: '1rem',
    border: 'none',
    padding: 0,
  },
});
