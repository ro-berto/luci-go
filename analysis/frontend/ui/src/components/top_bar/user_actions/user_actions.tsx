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

import {
  Link as RouterLink,
} from 'react-router-dom';

import HelpIcon from '@mui/icons-material/Help';
import FeedbackIcon from '@mui/icons-material/Feedback';
import Box from '@mui/material/Box';
import IconButton from '@mui/material/IconButton';
import Tooltip from '@mui/material/Tooltip';
import LoginButton from '@/components/top_bar/user_actions/login_button/login_button';
import UserProfileButton from '@/components/top_bar/user_actions/user_profile_button/user_profile_button';


const UserActions = () => {
  return (
    <Box sx={{ flexGrow: 0, display: 'flex', alignItems: 'center' }}>
      <Tooltip title="Help">
        <IconButton
          component={RouterLink}
          to="/help"
          sx={{ color: 'white' }}>
          <HelpIcon />
        </IconButton>
      </Tooltip>
      <Tooltip title="Send feedback">
        <IconButton
          href="https://goto.google.com/luci-analysis-bug"
          target="_blank"
          sx={{ color: 'white' }}>
          <FeedbackIcon />
        </IconButton>
      </Tooltip>
      {
        window.isAnonymous ?
        (
          <LoginButton />
        ) :
        (
          <UserProfileButton />
        )
      }
    </Box>
  );
};

export default UserActions;
