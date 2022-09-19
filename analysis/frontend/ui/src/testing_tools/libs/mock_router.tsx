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
/* eslint-disable valid-jsdoc */
/* eslint-disable @typescript-eslint/no-explicit-any */

import { FC } from 'react';
import {
  QueryClient,
  QueryClientProvider,
} from 'react-query';
import {
  BrowserRouter as Router,
  Route,
  Routes,
} from 'react-router-dom';

import {
  render,
  RenderResult,
} from '@testing-library/react';

/**
 * Renders a component wrapped with a mock Router.
 *
 * @param ui The component to render.
 * @param route The route that the current component is at, defaults to '/'.
 * @return The render result.
 */
export const renderWithRouter = (

    ui: React.ReactElement<any, string | React.JSXElementConstructor<any>>,
    route = '/',
): RenderResult => {
  window.history.pushState({}, 'Test page', route);

  return render(ui, { wrapper: Router });
};

/**
 * Renders a component with a mock router and a mock query client.
 *
 * @param ui The UI component to render.
 * @param route The route that the current component is at, defaults to '/'.
 * @param routeDefinition The definition of the current route,
 *                        useful for getting route params.
 * @return The render result.
 */
export const renderWithRouterAndClient = (
    ui: React.ReactElement,
    route = '/',
    routeDefinition = '',
) => {
  const wrapper: FC = ({ children }) => {
    return (
      <Router >
        <Routes>
          <Route
            path={routeDefinition ? routeDefinition : route}
            element={children}/>
        </Routes>
      </Router>
    );
  };
  window.history.pushState({}, 'Test page', route);
  const client = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  });
  return render(
      <QueryClientProvider client={client}>{ui}</QueryClientProvider>,
      {
        wrapper,
      });
};
