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

import { MoreVert, Settings, Upgrade } from '@mui/icons-material';
import { IconButton, ListItemIcon, ListItemText, Menu, MenuItem } from '@mui/material';
import { observer } from 'mobx-react-lite';
import { useState } from 'react';

import { useStore } from '../store';

export interface AppMenuProps {
  readonly container?: HTMLElement;
  readonly children?: JSX.Element;
}

export const AppMenu = observer(({ container, children }: AppMenuProps) => {
  const store = useStore();

  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null);
  const hasPendingUpdate = store.workbox.hasPendingUpdate;

  return (
    <>
      <IconButton
        onClick={(event) => setAnchorEl(event.currentTarget)}
        size="medium"
        color={hasPendingUpdate ? 'secondary' : 'default'}
        data-testid="menu-button"
      >
        {hasPendingUpdate ? <Upgrade /> : children ?? <MoreVert />}
      </IconButton>
      <Menu open={Boolean(anchorEl)} onClose={() => setAnchorEl(null)} anchorEl={anchorEl} container={container}>
        <MenuItem
          onClick={() => store.workbox.workbox?.messageSkipWaiting()}
          disabled={!hasPendingUpdate}
          title="A new version of the website is available. Click to update."
          aria-label="update website"
        >
          <ListItemIcon>
            <Upgrade />
          </ListItemIcon>
          <ListItemText>Update Website</ListItemText>
        </MenuItem>
        <MenuItem
          onClick={() => store.setShowSettingsDialog(true)}
          disabled={store.hasSettingsDialog === 0}
          title="Change settings specific to the page."
        >
          <ListItemIcon>
            <Settings />
          </ListItemIcon>
          <ListItemText>Page Settings</ListItemText>
        </MenuItem>
      </Menu>
    </>
  );
});
