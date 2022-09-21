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

import React, { useState } from 'react';

import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Search from '@mui/icons-material/Search';
import TextField from '@mui/material/TextField';

import { ListAnalysesTable } from '../components/analyses_tables/list_analyses_table';
import { SearchAnalysisTable } from '../components/analyses_tables/search_analysis_table';

export const FailureAnalysesPage = () => {
  const [bbid, setBbid] = useState<string>('');
  const [isValidBbid, setIsValidBbid] = useState<boolean>(true);
  const [bbidToSearch, setBbidToSearch] = useState<string>('');

  const digitsRegex = new RegExp('^([0-9,\\s]*)$');
  const handleChangeBbid = (event: React.ChangeEvent<HTMLInputElement>) => {
    // Replace whitespace in Buildbucket ID
    const cleansedBbid = event.target.value.replace(/\s/g, '');
    setBbid(cleansedBbid);
    setIsValidBbid(digitsRegex.test(cleansedBbid));
  };

  const searchForAnalysis = () => {
    if (isValidBbid) {
      setBbidToSearch(bbid);
    }
  };

  return (
    <main>
      <Box
        justifyContent='center'
        sx={{
          paddingLeft: '1rem',
          paddingRight: '1rem',
          paddingBottom: '2rem',
        }}
      >
        <TextField
          fullWidth
          size='small'
          variant='outlined'
          label='Buildbucket ID'
          onChange={handleChangeBbid}
          onKeyDown={(e) => {
            if (['Tab', 'Enter'].includes(e.key)) {
              searchForAnalysis();
            }
          }}
          error={!isValidBbid}
          helperText={
            !isValidBbid ? 'Invalid Buildbucket ID - enter digits only' : ''
          }
          InputProps={{
            endAdornment: (
              <Button
                color='primary'
                disabled={!isValidBbid || bbid === bbidToSearch}
                onClick={(_) => {
                  searchForAnalysis();
                }}
                startIcon={<Search />}
                sx={{
                  paddingLeft: '2rem',
                  paddingRight: '2rem',
                  borderRadius: 'none',
                }}
              >
                Search
              </Button>
            ),
          }}
          sx={{
            '& .MuiOutlinedInput-root': {
              paddingRight: '0',
            },
          }}
        />
      </Box>
      {bbidToSearch !== '' ? (
        <SearchAnalysisTable bbid={bbidToSearch} />
      ) : (
        <ListAnalysesTable />
      )}
    </main>
  );
};
