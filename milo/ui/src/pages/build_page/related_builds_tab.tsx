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

import { MobxLitElement } from '@adobe/lit-mobx';
import createCache from '@emotion/cache';
import { CacheProvider, EmotionCache } from '@emotion/react';
import { ChevronRight, ExpandMore } from '@mui/icons-material';
import {
  Box,
  Collapse,
  Icon,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography,
} from '@mui/material';
import { customElement } from 'lit/decorators.js';
import { makeObservable, observable } from 'mobx';
import { observer } from 'mobx-react-lite';
import { createContext, useContext, useEffect, useState } from 'react';
import { createRoot, Root } from 'react-dom/client';

import '../../components/dot_spinner';
import { DotSpinner } from '../../components/dot_spinner';
import { GA_ACTIONS, GA_CATEGORIES, trackEvent } from '../../libs/analytics_utils';
import { getURLPathForBuild, getURLPathForBuilder, getURLPathForProject } from '../../libs/build_utils';
import { BUILD_STATUS_CLASS_MAP, BUILD_STATUS_DISPLAY_MAP, BUILD_STATUS_ICON_MAP } from '../../libs/constants';
import { consumer } from '../../libs/context';
import { renderMarkdown } from '../../libs/markdown_utils';
import { displayDuration, NUMERIC_TIME_FORMAT } from '../../libs/time_utils';
import { consumeStore, StoreInstance, StoreProvider, useStore } from '../../store';
import { BuildStateInstance } from '../../store/build_state';
import { ExpandableEntriesState, ExpandableEntriesStateInstance } from '../../store/expandable_entries_state';
import colorClasses from '../../styles/color_classes.css';
import commonStyle from '../../styles/common_style.css';

const TableStateContext = createContext<ExpandableEntriesStateInstance>(ExpandableEntriesState.create());

interface RelatedBuildsTableRowProps {
  readonly index: number;
  readonly build: BuildStateInstance;
}

const RelatedBuildsTableRow = observer(({ index, build }: RelatedBuildsTableRowProps) => {
  const tableState = useContext(TableStateContext);

  const expanded = tableState.isExpanded(build.data.id);

  return (
    <>
      <TableRow
        sx={{
          backgroundColor: index % 2 === 0 ? 'var(--block-background-color)' : '',
          '& > td': { borderBottom: 'unset' },
        }}
      >
        <TableCell>
          <IconButton aria-label="toggle-row" size="small" onClick={() => tableState.toggle(build.data.id, !expanded)}>
            {expanded ? <ExpandMore /> : <ChevronRight />}
          </IconButton>
        </TableCell>
        <TableCell>
          <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Material+Icons&display=block" />
          <Icon
            className={BUILD_STATUS_CLASS_MAP[build.data.status]}
            title={BUILD_STATUS_DISPLAY_MAP[build.data.status]}
          >
            {BUILD_STATUS_ICON_MAP[build.data.status]}
          </Icon>
        </TableCell>
        <TableCell>
          <a href={getURLPathForProject(build.data.builder.project)}>{build.data.builder.project}</a>/
          {build.data.builder.bucket}/
          <a href={getURLPathForBuilder(build.data.builder)}>{build.data.builder.builder}</a>/
          <a href={getURLPathForBuild(build.data)}>{build.data.number ?? 'b' + build.data.id}</a>
        </TableCell>
        <TableCell>{build.createTime.toFormat(NUMERIC_TIME_FORMAT)}</TableCell>
        <TableCell>{displayDuration(build.pendingDuration) || 'N/A'}</TableCell>
        <TableCell>{(build.executionDuration && displayDuration(build.executionDuration)) || 'N/A'}</TableCell>
      </TableRow>
      <TableRow>
        <TableCell colSpan={6} sx={{ p: 0 }}>
          <Collapse in={expanded} timeout="auto">
            <Box
              className={`${BUILD_STATUS_CLASS_MAP[build.data.status]}-bg`}
              sx={{
                padding: '0 10px',
                clear: 'both',
                overflowWrap: 'break-word',
                '& pre': {
                  whiteSpace: 'pre-wrap',
                  overflowWrap: 'break-word',
                  fontSize: '12px',
                },
                '& *': {
                  marginBlock: '10px',
                },
              }}
              dangerouslySetInnerHTML={{ __html: renderMarkdown(build.data.summaryMarkdown || 'No Summary.') }}
            ></Box>
          </Collapse>
        </TableCell>
      </TableRow>
    </>
  );
});

export const RelatedBuildsTab = observer(() => {
  const store = useStore();
  const [tableState] = useState(() => ExpandableEntriesState.create());

  useEffect(() => {
    store.setSelectedTabId('related-builds');
    trackEvent(GA_CATEGORIES.RELATED_BUILD_TAB, GA_ACTIONS.TAB_VISITED, window.location.href);
  }, []);

  if (!store.buildPage.build || !store.buildPage.relatedBuilds) {
    return (
      <Box sx={{ p: 1, color: 'var(--active-text-color' }}>
        Loading <DotSpinner />
      </Box>
    );
  }

  if (!store.buildPage.relatedBuilds.length) {
    return <Box sx={{ p: 1 }}>No other builds found with the same buildset</Box>;
  }

  return (
    <Box>
      <Box sx={{ p: 2 }}>
        <Typography variant="h6">Other builds with the same buildset</Typography>
        <ul>
          {store.buildPage.build.buildSets.map((bs) => (
            <li key={bs}>{bs}</li>
          ))}
        </ul>
      </Box>
      <TableStateContext.Provider value={tableState}>
        <Table size="small">
          <TableHead>
            <TableRow>
              <TableCell>
                <IconButton
                  aria-label="expand-all-rows"
                  size="small"
                  onClick={() => {
                    tableState.toggleAll(!tableState.defaultExpanded);
                  }}
                >
                  {tableState.defaultExpanded ? <ExpandMore /> : <ChevronRight />}
                </IconButton>
              </TableCell>
              <TableCell>Status</TableCell>
              <TableCell>Build</TableCell>
              <TableCell>Create Time</TableCell>
              <TableCell>Pending</TableCell>
              <TableCell>Duration</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {store.buildPage.relatedBuilds.map((b, i) => (
              <RelatedBuildsTableRow key={b.data.id} index={i} build={b} />
            ))}
          </TableBody>
        </Table>
      </TableStateContext.Provider>
    </Box>
  );
});

@customElement('milo-related-builds-tab')
@consumer
export class BuildPageRelatedBuildsTabElement extends MobxLitElement {
  @observable.ref @consumeStore() store!: StoreInstance;

  private readonly cache: EmotionCache;
  private readonly parent: HTMLDivElement;
  private readonly root: Root;

  constructor() {
    super();
    makeObservable(this);
    this.parent = document.createElement('div');
    const child = document.createElement('div');
    this.root = createRoot(child);
    this.parent.appendChild(child);
    this.cache = createCache({
      key: 'milo-related-builds-tab',
      container: this.parent,
    });
  }

  protected render() {
    this.root.render(
      <CacheProvider value={this.cache}>
        <StoreProvider value={this.store}>
          <RelatedBuildsTab />
        </StoreProvider>
      </CacheProvider>
    );
    return this.parent;
  }

  static styles = [commonStyle, colorClasses];
}
