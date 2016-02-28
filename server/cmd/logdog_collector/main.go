// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/luci/luci-go/common/auth"
	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/errors"
	"github.com/luci/luci-go/common/gcloud/pubsub"
	"github.com/luci/luci-go/common/gcloud/pubsub/ackbuffer"
	"github.com/luci/luci-go/common/gcloud/pubsub/subscriber"
	log "github.com/luci/luci-go/common/logging"
	"github.com/luci/luci-go/server/internal/logdog/collector"
	"github.com/luci/luci-go/server/internal/logdog/collector/coordinator"
	"github.com/luci/luci-go/server/internal/logdog/service"
	"golang.org/x/net/context"
)

var (
	errInvalidConfig = errors.New("invalid configuration")
)

// application is the Collector application state.
type application struct {
	*service.Service

	// shutdownCtx is a Context that will be cancelled if our application
	// receives a shutdown signal.
	shutdownCtx context.Context
}

// run is the main execution function.
func (a *application) runCollector() error {
	cfg := a.Config()
	ccfg := cfg.GetCollector()
	if ccfg == nil {
		return errors.New("no collector configuration")
	}

	pscfg := cfg.GetTransport().GetPubsub()
	if pscfg == nil {
		return errors.New("missing Pub/Sub configuration")
	}

	// Our Subscription must be a valid one.
	sub := pubsub.NewSubscription(pscfg.Project, pscfg.Subscription)
	if err := sub.Validate(); err != nil {
		return fmt.Errorf("invalid Pub/Sub subscription %q: %v", sub, err)
	}

	// New PubSub instance with the authenticated client.
	psClient, err := a.AuthenticatedClient(func(o *auth.Options) {
		o.Scopes = pubsub.SubscriberScopes
	})
	if err != nil {
		log.WithError(err).Errorf(a, "Failed to create Pub/Sub client.")
		return err
	}

	// Create a retrying Pub/Sub client.
	ps := &pubsub.Retry{
		Connection: pubsub.NewConnection(psClient),
		Callback: func(err error, d time.Duration) {
			log.Fields{
				log.ErrorKey: err,
				"delay":      d,
			}.Warningf(a, "Transient error encountered; retrying...")
		},
	}

	exists, err := ps.SubExists(a, sub)
	if err != nil {
		log.Fields{
			log.ErrorKey:   err,
			"subscription": sub,
		}.Errorf(a, "Could not confirm Pub/Sub subscription.")
		return errInvalidConfig
	}
	if !exists {
		log.Fields{
			"subscription": sub,
		}.Errorf(a, "Subscription does not exist.")
		return errInvalidConfig
	}
	log.Fields{
		"subscription": sub,
	}.Infof(a, "Successfully validated Pub/Sub subscription.")

	// Initialize our Storage.
	s, err := a.Storage()
	if err != nil {
		log.WithError(err).Errorf(a, "Failed to get storage instance.")
		return err
	}
	defer s.Close()

	// Application shutdown will now operate by cancelling the Collector's
	// shutdown Context.
	shutdownCtx, shutdownFunc := context.WithCancel(a)
	a.SetShutdownFunc(shutdownFunc)
	defer a.SetShutdownFunc(nil)

	// Start an ACK buffer so that we can batch ACKs.
	ab := ackbuffer.New(a, ackbuffer.Config{
		Ack: ackbuffer.NewACK(ps, sub, 0),
	})
	defer ab.CloseAndFlush()

	// Initialize our Collector service object using a caching Coordinator
	// interface.
	coord := coordinator.NewCoordinator(a.Coordinator())
	coord = coordinator.NewCache(coord, int(ccfg.StateCacheSize), ccfg.StateCacheExpiration.Duration())

	coll := collector.Collector{
		Coordinator:        coord,
		Storage:            s,
		MaxParallelBundles: int(ccfg.Workers),
		MaxIngestWorkers:   int(ccfg.Workers),
	}
	defer coll.Close()

	// Execute our main Subscriber loop. It will run until the supplied Context
	// is cancelled.
	clk := clock.Get(a)
	engine := subscriber.Subscriber{
		S: subscriber.NewSource(ps, sub, 0),
		A: ab,

		PullWorkers:    int(ccfg.TransportWorkers),
		HandlerWorkers: int(ccfg.Workers),
	}
	engine.Run(shutdownCtx, func(msg *pubsub.Message) bool {
		startTime := clk.Now()

		ctx := log.SetField(a, "messageID", msg.ID)
		log.Fields{
			"ackID": msg.AckID,
			"size":  len(msg.Data),
		}.Infof(ctx, "Received Pub/Sub Message.")

		err := coll.Process(ctx, msg.Data)

		duration := clk.Now().Sub(startTime)
		switch {
		case errors.IsTransient(err):
			// Do not consume
			log.Fields{
				log.ErrorKey: err,
				"duration":   duration,
			}.Warningf(ctx, "TRANSIENT error ingesting Pub/Sub message.")
			return false

		case err == nil:
			log.Fields{
				"ackID":    msg.AckID,
				"size":     len(msg.Data),
				"duration": duration,
			}.Infof(ctx, "Message successfully processed; ACKing.")
			return true

		default:
			log.Fields{
				"ackID":    msg.AckID,
				"size":     len(msg.Data),
				"duration": duration,
			}.Errorf(ctx, "Non-transient error ingesting Pub/Sub message; ACKing.")
			return true
		}
	})

	log.Debugf(a, "Collector finished.")
	return nil
}

// mainImpl is the Main implementaion, and returns the application return code
// as an integer.
func mainImpl() int {
	a := application{
		Service: service.New(context.Background()),
	}

	fs := flag.FlagSet{}
	a.AddFlags(&fs)

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Errorf(log.SetError(a, err), "Failed to parse command-line.")
		return 1
	}

	// Run our configured application instance.
	var rc int
	if err := a.Run(a.runCollector); err != nil {
		log.Errorf(log.SetError(a, err), "Application execution failed.")
		return 1
	}
	log.Infof(log.SetField(a, "returnCode", rc), "Terminating.")
	return 0
}

// Entry point.
func main() {
	os.Exit(mainImpl())
}
