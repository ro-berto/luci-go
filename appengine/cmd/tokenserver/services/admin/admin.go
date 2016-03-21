// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package admin contains implementation of various administrative procedures.
//
// Code defined here is either invoked by an administrator or by the service
// itself (via cron jobs or task queues).
package admin

import (
	"bytes"
	"crypto/x509"
	"fmt"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/luci/gae/service/datastore"
	"github.com/luci/gae/service/info"
	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/config"
	"github.com/luci/luci-go/common/errors"
	"github.com/luci/luci-go/common/logging"
	google_protobuf "github.com/luci/luci-go/common/proto/google"
	"github.com/luci/luci-go/common/stringset"

	"github.com/luci/luci-go/appengine/cmd/tokenserver/model"
	"github.com/luci/luci-go/common/api/tokenserver/v1"
)

// Server implements tokenserver.AdminServer RPC interface.
//
// It assumes authorization has happened already. Use DecoratedAdmin to plug it
// in.
type Server struct {
	// ConfigFactory returns instances of config.Interface on demand.
	ConfigFactory func(context.Context) (config.Interface, error)
}

// ImportConfig makes the server read its config from luci-config right now.
//
// Note that regularly configs are read in background each 5 min. ImportConfig
// can be used to force config reread immediately. It will block until configs
// are read.
func (s *Server) ImportConfig(c context.Context, _ *google_protobuf.Empty) (*tokenserver.ImportConfigResponse, error) {
	cfg, err := s.fetchConfigFile(c, "tokenserver.cfg")
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "can't read config file - %s", err)
	}
	logging.Infof(c, "Importing config at rev %s", cfg.Revision)

	// Read list of CAs.
	msg := tokenserver.TokenServerConfig{}
	if err = proto.UnmarshalText(cfg.Content, &msg); err != nil {
		return nil, grpc.Errorf(codes.Internal, "can't parse config file - %s", err)
	}

	// There should be no duplicates.
	seenCAs := stringset.New(len(msg.GetCertificateAuthority()))
	for _, ca := range msg.GetCertificateAuthority() {
		if seenCAs.Has(ca.Cn) {
			return nil, grpc.Errorf(codes.Internal, "duplicate entries in the config")
		}
		seenCAs.Add(ca.Cn)
	}

	// Add new CA datastore entries or update existing ones.
	wg := sync.WaitGroup{}
	me := errors.NewLazyMultiError(len(msg.GetCertificateAuthority()))
	for i, ca := range msg.GetCertificateAuthority() {
		wg.Add(1)
		go func(i int, ca *tokenserver.CertificateAuthorityConfig) {
			defer wg.Done()
			if err := s.importCA(c, ca, cfg.Revision); err != nil {
				logging.Errorf(c, "Failed to import %q: %s", ca.Cn, err)
				me.Assign(i, err)
			}
		}(i, ca)
	}
	wg.Wait()
	if err = me.Get(); err != nil {
		return nil, grpc.Errorf(codes.Internal, "can't import CA - %s", err)
	}

	// Find CAs that were removed from the config.
	toRemove := []string{}
	q := datastore.NewQuery("CA").Eq("Removed", false).KeysOnly(true)
	err = datastore.Get(c).Run(q, func(k *datastore.Key) {
		if !seenCAs.Has(k.StringID()) {
			toRemove = append(toRemove, k.StringID())
		}
	})
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "datastore error - %s", err)
	}

	// Mark them as inactive in the datastore.
	wg = sync.WaitGroup{}
	me = errors.NewLazyMultiError(len(toRemove))
	for i, name := range toRemove {
		wg.Add(1)
		go func(i int, name string) {
			defer wg.Done()
			if err := s.removeCA(c, name, cfg.Revision); err != nil {
				logging.Errorf(c, "Failed to remove %q: %s", name, err)
				me.Assign(i, err)
			}
		}(i, name)
	}
	wg.Wait()
	if err = me.Get(); err != nil {
		return nil, grpc.Errorf(codes.Internal, "datastore error - %s", err)
	}

	return &tokenserver.ImportConfigResponse{
		Revision: cfg.Revision,
	}, nil
}

// FetchCRL makes the server fetch a CRL for some CA.
func (s *Server) FetchCRL(c context.Context, r *tokenserver.FetchCRLRequest) (*tokenserver.FetchCRLResponse, error) {
	ds := datastore.Get(c)

	// Grab a corresponding CA entity. It contains URL of CRL to fetch.
	ca := &model.CA{CN: r.Cn}
	switch err := ds.Get(ca); {
	case err == datastore.ErrNoSuchEntity:
		return nil, grpc.Errorf(codes.NotFound, "no such CA %q", ca.CN)
	case err != nil:
		return nil, grpc.Errorf(codes.Internal, "datastore error - %s", err)
	}

	// Grab CRL URL from the CA config.
	cfg, err := ca.ParseConfig()
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "broken CA config in the datastore - %s", err)
	}
	if cfg.CrlUrl == "" {
		return nil, grpc.Errorf(codes.NotFound, "CA %q doesn't have CRL defined", ca.CN)
	}

	// Grab info about last processed CRL, if any.
	crl := &model.CRL{Parent: ds.KeyForObj(ca)}
	if err = ds.Get(crl); err != nil && err != datastore.ErrNoSuchEntity {
		return nil, grpc.Errorf(codes.Internal, "datastore error - %s", err)
	}

	// Fetch latest CRL blob.
	logging.Infof(c, "Fetching CRL for %q from %s", ca.CN, cfg.CrlUrl)
	knownETag := crl.LastFetchETag
	if r.Force {
		knownETag = ""
	}
	fetchCtx, _ := clock.WithTimeout(c, time.Minute)
	crlDer, newEtag, err := fetchCRL(fetchCtx, cfg, knownETag)
	switch {
	case errors.IsTransient(err):
		return nil, grpc.Errorf(codes.Internal, "transient error when fetching CRL - %s", err)
	case err != nil:
		return nil, grpc.Errorf(codes.Unknown, "can't fetch CRL - %s", err)
	}

	// No changes?
	if knownETag != "" && knownETag == newEtag {
		logging.Infof(c, "No changes to CRL (etag is %s), skipping", knownETag)
	} else {
		logging.Infof(c, "Fetched CRL size is %d bytes, etag is %s", len(crlDer), newEtag)
		crl, err = validateAndStoreCRL(c, crlDer, newEtag, ca, crl)
		switch {
		case errors.IsTransient(err):
			return nil, grpc.Errorf(codes.Internal, "transient error when storing CRL - %s", err)
		case err != nil:
			return nil, grpc.Errorf(codes.Unknown, "bad CRL - %s", err)
		}
	}

	return &tokenserver.FetchCRLResponse{CrlStatus: crl.GetStatusProto()}, nil
}

// GetCAStatus returns configuration of some CA defined in the config.
func (s *Server) GetCAStatus(c context.Context, r *tokenserver.GetCAStatusRequest) (*tokenserver.GetCAStatusResponse, error) {
	ds := datastore.Get(c)

	// Entities to fetch.
	ca := model.CA{CN: r.Cn}
	crl := model.CRL{Parent: ds.KeyForObj(&ca)}

	// Fetch them at the same revision. It is fine if CRL is not there yet. Don't
	// bother doing it in parallel: GetCAStatus is used only by admins, manually.
	err := ds.RunInTransaction(func(c context.Context) error {
		ds := datastore.Get(c)
		if err := ds.Get(&ca); err != nil {
			return err // can be ErrNoSuchEntity
		}
		if err := ds.Get(&crl); err != nil && err != datastore.ErrNoSuchEntity {
			return err // only transient errors
		}
		return nil
	}, nil)
	switch {
	case err == datastore.ErrNoSuchEntity:
		return &tokenserver.GetCAStatusResponse{}, nil
	case err != nil:
		return nil, grpc.Errorf(codes.Internal, "datastore error - %s", err)
	}

	cfgMsg, err := ca.ParseConfig()
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "broken config in the datastore - %s", err)
	}

	return &tokenserver.GetCAStatusResponse{
		Config:     cfgMsg,
		Cert:       dumpPEM(ca.Cert, "CERTIFICATE"),
		Removed:    ca.Removed,
		Ready:      ca.Ready,
		AddedRev:   ca.AddedRev,
		UpdatedRev: ca.UpdatedRev,
		RemovedRev: ca.RemovedRev,
		CrlStatus:  crl.GetStatusProto(),
	}, nil
}

////////////////////////////////////////////////////////////////////////////////

// fetchConfigFile fetches a file from this services' config set.
func (s *Server) fetchConfigFile(c context.Context, path string) (*config.Config, error) {
	logging.Infof(c, "Reading %q", path)
	c, _ = context.WithTimeout(c, 30*time.Second) // URL fetch deadline
	cfg, err := s.ConfigFactory(c)
	if err != nil {
		return nil, err
	}
	return cfg.GetConfig("services/"+info.Get(c).AppID(), path, false)
}

// importCA imports CA definition from the config (or updates an existing one).
func (s *Server) importCA(c context.Context, ca *tokenserver.CertificateAuthorityConfig, rev string) error {
	// Read CA certificate file, convert it to der.
	caCfg, err := s.fetchConfigFile(c, ca.CertPath)
	if err != nil {
		return err
	}
	certDer, err := parsePEM(caCfg.Content, "CERTIFICATE")
	if err != nil {
		return fmt.Errorf("bad PEM - %s", err)
	}

	// Check the certificate makes sense.
	cert, err := x509.ParseCertificate(certDer)
	if err != nil {
		return fmt.Errorf("bad cert - %s", err)
	}
	if !cert.IsCA {
		return fmt.Errorf("not a CA cert")
	}
	if cert.Subject.CommonName != ca.Cn {
		return fmt.Errorf("bad CN in the certificate, expecting %q, got %q", ca.Cn, cert.Subject.CommonName)
	}

	// Serialize the config back to proto to store it in the entity.
	cfgBlob, err := proto.Marshal(ca)
	if err != nil {
		return err
	}

	// Create or update the entity.
	return datastore.Get(c).RunInTransaction(func(c context.Context) error {
		ds := datastore.Get(c)
		existing := model.CA{CN: ca.Cn}
		err := ds.Get(&existing)
		if err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}
		// New one?
		if err == datastore.ErrNoSuchEntity {
			logging.Infof(c, "Adding new CA %q", ca.Cn)
			return ds.Put(&model.CA{
				CN:         ca.Cn,
				Config:     cfgBlob,
				Cert:       certDer,
				AddedRev:   rev,
				UpdatedRev: rev,
			})
		}
		// Exists already? Check whether we should update it.
		if !existing.Removed &&
			bytes.Equal(existing.Config, cfgBlob) &&
			bytes.Equal(existing.Cert, certDer) {
			return nil
		}
		logging.Infof(c, "Updating CA %q", ca.Cn)
		existing.Config = cfgBlob
		existing.Cert = certDer
		existing.Removed = false
		existing.UpdatedRev = rev
		existing.RemovedRev = ""
		return ds.Put(&existing)
	}, nil)
}

// removeCA marks the CA in the datastore as removed.
func (s *Server) removeCA(c context.Context, name string, rev string) error {
	return datastore.Get(c).RunInTransaction(func(c context.Context) error {
		ds := datastore.Get(c)
		existing := model.CA{CN: name}
		if err := ds.Get(&existing); err != nil {
			return err
		}
		if existing.Removed {
			return nil
		}
		logging.Infof(c, "Removing CA %q", name)
		existing.Removed = true
		existing.RemovedRev = rev
		return ds.Put(&existing)
	}, nil)
}
