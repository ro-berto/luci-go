// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package openid

import (
	"testing"

	"github.com/luci/luci-go/server/settings"
	"golang.org/x/net/context"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSettings(t *testing.T) {
	Convey("Works", t, func() {
		c := context.Background()
		c = settings.Use(c, settings.New(&settings.MemoryStorage{}))

		cfg, err := FetchSettings(c)
		So(err, ShouldBeNil)
		So(cfg, ShouldResemble, &Settings{})

		store := Settings{
			DiscoveryURL: "http://discovery",
			ClientID:     "client_id",
			ClientSecret: "client_secret",
			RedirectURI:  "http://redirect",
		}
		So(StoreSettings(c, &store, "who", "why"), ShouldBeNil)

		cfg, err = FetchSettings(c)
		So(err, ShouldBeNil)
		So(cfg, ShouldResemble, &store)
	})
}
