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

import './analysis_details.css';

import { useState } from 'react';
import { useParams } from 'react-router-dom';
import { useQuery } from 'react-query';

import Alert from '@mui/material/Alert';
import AlertTitle from '@mui/material/AlertTitle';
import Box from '@mui/material/Box';
import CircularProgress from '@mui/material/CircularProgress';
import Tab from '@mui/material/Tab';
import Tabs from '@mui/material/Tabs';
import Typography from '@mui/material/Typography';

import { AnalysisOverview } from '../../components/analysis_overview/analysis_overview';
import { HeuristicAnalysisTable } from '../../components/heuristic_analysis_table/heuristic_analysis_table';
// TODO: uncomment below once revert CL and prime suspect information is added to the analysis response
// import { RevertCLOverview } from '../../components/revert_cl_overview/revert_cl_overview';
// import { SuspectsOverview } from '../../components/suspects_overview/suspects_overview';

import {
  getLUCIBisectionService,
  QueryAnalysisRequest,
} from '../../services/luci_bisection';

interface TabPanelProps {
  children?: React.ReactNode;
  name: string;
  value: string;
}

function TabPanel(props: TabPanelProps) {
  const { children, value, name } = props;

  return (
    <div hidden={value !== name} className='tab-panel'>
      {value === name && <div className='tab-panel-contents'>{children}</div>}
    </div>
  );
}

export const AnalysisDetailsPage = () => {
  enum AnalysisComponentTabs {
    HEURISTIC = 'Heuristic analysis',
    NTH_SECTION = 'Nth section analysis',
    CULPRIT_VERIFICATION = 'Culprit verification',
  }

  const [currentTab, setCurrentTab] = useState(AnalysisComponentTabs.HEURISTIC);

  const handleTabChange = (
    _: React.SyntheticEvent,
    newTab: AnalysisComponentTabs
  ) => {
    setCurrentTab(newTab);
  };

  const { bbid } = useParams();

  const bisectionService = getLUCIBisectionService();

  const {
    isLoading,
    isError,
    isSuccess,
    data: response,
    error,
  } = useQuery(['analysis', bbid], async () => {
    const request: QueryAnalysisRequest = {
      buildFailure: {
        bbid: bbid!,
        failedStepName: 'compile',
      },
    };

    return await bisectionService.queryAnalysis(request);
  });

  let analysis = null;
  if (isSuccess && response) {
    analysis = response!.analyses[0];
  }

  // TODO: display alert if the build ID queried is not the first failed build
  //       linked to the failure analysis

  return (
    <main>
      {isError && (
        <div className='section'>
          <Alert severity='error'>
            <AlertTitle>Failed to load analysis details</AlertTitle>
            {/* TODO: display more error detail for input issues e.g.
                  Build not found, No analysis for that build, etc */}
            An error occurred when querying for the analysis details using build
            ID "{bbid}":
            <Box sx={{ padding: '1rem' }}>{`${error}`}</Box>
          </Alert>
        </div>
      )}
      {isLoading && (
        <Box
          display='flex'
          justifyContent='center'
          alignItems='center'
          height='80vh'
        >
          <CircularProgress />
        </Box>
      )}
      {isSuccess && (
        <>
          <div className='section'>
            <Typography variant='h4' gutterBottom>
              Analysis Details
            </Typography>
            <AnalysisOverview analysis={analysis!} />
          </div>
          {/* TODO: add revert CL and prime suspect information to the analysis
                    response, then display it below
          {analysis.revertCL! && (
            <div className='section'>
              <Typography variant='h4' gutterBottom>
                Revert CL
              </Typography>
              <RevertCLOverview revertCL={analysis.revertCL} />
            </div>
          )}
          {analysis.primeSuspects.length > 0 && (
            <div className='section'>
              <Typography variant='h4' gutterBottom>
                Suspect Summary
              </Typography>
              <SuspectsOverview suspects={analysis.primeSuspects} />
            </div>
          )} */}
          <div className='section'>
            <Typography variant='h4' gutterBottom>
              Analysis Components
            </Typography>
            <Tabs
              value={currentTab}
              onChange={handleTabChange}
              aria-label='Analysis components tabs'
              className='rounded-tabs'
            >
              <Tab
                className='rounded-tab'
                value={AnalysisComponentTabs.HEURISTIC}
                label={AnalysisComponentTabs.HEURISTIC}
              />
              <Tab
                className='rounded-tab'
                disabled
                value={AnalysisComponentTabs.NTH_SECTION}
                label={AnalysisComponentTabs.NTH_SECTION}
              />
              <Tab
                className='rounded-tab'
                disabled
                value={AnalysisComponentTabs.CULPRIT_VERIFICATION}
                label={AnalysisComponentTabs.CULPRIT_VERIFICATION}
              />
            </Tabs>
            <TabPanel value={currentTab} name={AnalysisComponentTabs.HEURISTIC}>
              {/* TODO: Show alert if there are no heuristic results yet */}
              <HeuristicAnalysisTable result={analysis!.heuristicResult} />
            </TabPanel>
            <TabPanel
              value={currentTab}
              name={AnalysisComponentTabs.NTH_SECTION}
            >
              {/* TODO: Replace with nth section analysis results */}
              Placeholder for nth section analysis details
            </TabPanel>
            <TabPanel
              value={currentTab}
              name={AnalysisComponentTabs.CULPRIT_VERIFICATION}
            >
              {/* TODO: Replace with culprit verification results */}
              Placeholder for culprit verification details
            </TabPanel>
          </div>
        </>
      )}
    </main>
  );
};
