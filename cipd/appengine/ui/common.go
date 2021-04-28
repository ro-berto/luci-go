// Copyright 2018 The LUCI Authors.
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

// Package ui implements request handlers that serve user facing HTML pages.
package ui

import (
	"context"
	"net/url"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/cipd/appengine/impl"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/xsrf"
	"go.chromium.org/luci/server/router"
	"go.chromium.org/luci/server/templates"
)

var requestStateKey = "cipd/ui requestState"

// requestState is stored in the context of UI requests.
type requestState struct {
	startTime time.Time      // when the request handler started
	services  *impl.Services // the backend implementation to use
}

// state returns the current requestState.
func state(ctx context.Context) *requestState {
	return ctx.Value(&requestStateKey).(*requestState)
}

// InstallHandlers adds HTTP handlers that render HTML pages.
func InstallHandlers(srv *server.Server, svc *impl.Services, templatesPath string) {
	m := router.NewMiddlewareChain(
		func(c *router.Context, next router.Handler) {
			c.Context = context.WithValue(c.Context, &requestStateKey, &requestState{
				startTime: clock.Now(c.Context),
				services:  svc,
			})
			next(c)
		},
		templates.WithTemplates(prepareTemplates(&srv.Options, templatesPath)),
		auth.Authenticate(srv.CookieAuth),
	)
	srv.Routes.GET("/", m, renderErr(routeToPage))
	srv.Routes.GET("/p/*path", m, renderErr(routeToPage))
}

func prefixPageURL(pfx string) string {
	if pfx == "" {
		return "/"
	}
	return "/p/" + pfx
}

func packagePageURL(pkg, cursor string) string {
	p := "/p/" + pkg + "/+/"
	if cursor != "" {
		p += "?c=" + url.QueryEscape(cursor)
	}
	return p
}

func instancePageURL(pkg, ver string) string {
	return "/p/" + pkg + "/+/" + ver
}

// routeToPage routes to an appropriate page depending on the request URL.
func routeToPage(c *router.Context) error {
	path := c.Params.ByName("path")
	switch chunks := strings.SplitN(path, "/+/", 2); {
	case len(chunks) <= 1: // no '/+/' in path => prefix listing page
		return prefixListingPage(c, path)
	case len(chunks) == 2 && chunks[1] == "": // ends with '/+/' => package page
		return packagePage(c, chunks[0])
	case len(chunks) == 2: // has something after '/+/' => instance page
		return instancePage(c, chunks[0], chunks[1])
	default:
		return status.Errorf(codes.InvalidArgument, "malformed page URL")
	}
}

// prepareTemplates configures templates.Bundle used by all UI handlers.
//
// In particular it includes a set of default arguments passed to all templates.
func prepareTemplates(opts *server.Options, templatesPath string) *templates.Bundle {
	versionID := "unknown"
	if idx := strings.LastIndex(opts.ContainerImageID, ":"); idx != -1 {
		versionID = opts.ContainerImageID[idx+1:]
	}
	return &templates.Bundle{
		Loader:          templates.FileSystemLoader(templatesPath),
		DebugMode:       func(context.Context) bool { return !opts.Prod },
		DefaultTemplate: "base",
		DefaultArgs: func(ctx context.Context, e *templates.Extra) (templates.Args, error) {
			loginURL, err := auth.LoginURL(ctx, e.Request.URL.RequestURI())
			if err != nil {
				return nil, err
			}
			logoutURL, err := auth.LogoutURL(ctx, e.Request.URL.RequestURI())
			if err != nil {
				return nil, err
			}
			token, err := xsrf.Token(ctx)
			if err != nil {
				return nil, err
			}
			return templates.Args{
				"AppVersion":  versionID,
				"IsAnonymous": auth.CurrentIdentity(ctx) == identity.AnonymousIdentity,
				"User":        auth.CurrentUser(ctx),
				"LoginURL":    loginURL,
				"LogoutURL":   logoutURL,
				"XsrfToken":   token,
				"HandlerDuration": func() time.Duration {
					return clock.Now(ctx).Sub(state(ctx).startTime)
				},
			}, nil
		},
	}
}
