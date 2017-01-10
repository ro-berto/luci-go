// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Package cfgclient contains service implementations for the LUCI configuration
// service defined in github.com/luci/luci-go/common/config.
//
// This defines an interface to the LUCI configuration service properties and
// files. The interface is designed to be used by services which handle user
// data, and has the ability to operate on behalf of authorities, either the
// service itself (privileged), on behalf of the user (delegation), or
// anonymously.
//
// This package also offers the concept of resolution, where a configuration
// value is transformed into a more versatile application format prior to being
// cached and/or returned. Resolution allows configuration data consumers to
// handle configuration data as native Go types instead of raw configuration
// service data.
//
// Configuration requests pass through the following layers:
// 1) A Backend, which is the configured configuration authority.
// 2) Cache resolution, which optionally transforms the data into an
//    application-specific cachable format.
// 3) A cache layer, which caches the data.
// 4) Value resolution, which transforms the cached data format from (2) into
//    a Go value.
// 5) The Go value is retuned to the user.
//
// Layers (2) and (4) are managed by the Resolver type, which is associated by
// the application with the underlying configuration data.
package cfgclient
