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

package config

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/config"
	"go.chromium.org/luci/config/cfgclient"
	"go.chromium.org/luci/config/validation"
	"go.chromium.org/luci/server/router"

	gce "go.chromium.org/luci/gce/api/config/v1"
	"go.chromium.org/luci/gce/api/projects/v1"
	"go.chromium.org/luci/gce/appengine/rpc"
)

// projectsFile is the name of the projects config file.
const projectsFile = "projects.cfg"

// vmsFile is the name of the VMs config file.
const vmsFile = "vms.cfg"

// Config encapsulates the service config.
type Config struct {
	revision string
	Projects *projects.Configs
	VMs      *gce.Configs
}

// prjKey is the key to a projects.ProjectsServer in the context.
var prjKey = "prj"

// withProjServer returns a new context with the given projects.ProjectsServer
// installed.
func withProjServer(c context.Context, srv projects.ProjectsServer) context.Context {
	return context.WithValue(c, &prjKey, srv)
}

// getProjServer returns the projects.ProjectsServer installed in the current
// context.
func getProjServer(c context.Context) projects.ProjectsServer {
	return c.Value(&prjKey).(projects.ProjectsServer)
}

// vmsKey is the key to a gce.ConfigurationServer in the context.
var vmsKey = "vms"

// withVMsServer returns a new context with the given gce.ConfigurationServer
// installed.
func withVMsServer(c context.Context, srv gce.ConfigurationServer) context.Context {
	return context.WithValue(c, &vmsKey, srv)
}

// getVMsServer returns the gce.ConfigurationServer installed in the current
// context.
func getVMsServer(c context.Context) gce.ConfigurationServer {
	return c.Value(&vmsKey).(gce.ConfigurationServer)
}

// fetch fetches configs from the config service.
func fetch(c context.Context) (*Config, error) {
	cli := cfgclient.Client(c)
	rev := ""
	vms := &gce.Configs{}
	switch vmsCfg, err := cli.GetConfig(c, "services/${appid}", vmsFile, false); {
	case err == config.ErrNoConfig:
		logging.Debugf(c, "%q not found", vmsFile)
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch %q", vmsFile).Err()
	default:
		rev = vmsCfg.Revision
		logging.Debugf(c, "found %q revision %s", vmsFile, vmsCfg.Revision)
		if err := proto.UnmarshalText(vmsCfg.Content, vms); err != nil {
			return nil, errors.Annotate(err, "failed to load %q", vmsFile).Err()
		}
	}
	prjs := &projects.Configs{}
	switch prjsCfg, err := cli.GetConfig(c, "services/${appid}", projectsFile, false); {
	case err == config.ErrNoConfig:
		logging.Debugf(c, "%q not found", projectsFile)
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch %q", projectsFile).Err()
	default:
		logging.Debugf(c, "found %q revision %s", projectsFile, prjsCfg.Revision)
		if rev != "" && prjsCfg.Revision != rev {
			return nil, errors.Reason("config revision mismatch").Err()
		}
		if err := proto.UnmarshalText(prjsCfg.Content, prjs); err != nil {
			return nil, errors.Annotate(err, "failed to load %q", projectsFile).Err()
		}
	}
	return &Config{
		revision: rev,
		Projects: prjs,
		VMs:      vms,
	}, nil
}

// validate validates configs.
func validate(c context.Context, cfg *Config) error {
	v := &validation.Context{Context: c}
	v.SetFile(projectsFile)
	cfg.Projects.Validate(v)
	v.SetFile(vmsFile)
	cfg.VMs.Validate(v)
	return v.Finalize()
}

// deref dereferences VMs metadata by fetching referenced files.
func deref(c context.Context, cfg *Config) error {
	// Cache fetched files.
	fileMap := make(map[string]string)
	cli := cfgclient.Client(c)
	for _, v := range cfg.VMs.GetVms() {
		for i, m := range v.GetAttributes().Metadata {
			if m.GetFromFile() != "" {
				parts := strings.SplitN(m.GetFromFile(), ":", 2)
				if len(parts) < 2 {
					return errors.Reason("metadata from file must be in key:value form").Err()
				}
				file := parts[1]
				if _, ok := fileMap[file]; !ok {
					fileCfg, err := cli.GetConfig(c, "services/${appid}", file, false)
					if err != nil {
						return errors.Annotate(err, "failed to fetch %q", file).Err()
					}
					logging.Debugf(c, "found %q revision %s", file, fileCfg.Revision)
					if fileCfg.Revision != cfg.revision {
						return errors.Reason("config revision mismatch %q", fileCfg.Revision).Err()
					}
					fileMap[file] = fileCfg.Content
				}
				// fileMap[file] definitely exists.
				key := parts[0]
				val := fileMap[file]
				v.Attributes.Metadata[i].Metadata = &gce.Metadata_FromText{
					FromText: fmt.Sprintf("%s:%s", key, val),
				}
			}
		}
	}
	return nil
}

// normalize normalizes VMs durations by converting them to seconds, and sets
// output-only properties.
func normalize(c context.Context, cfg *Config) error {
	for _, p := range cfg.Projects.GetProject() {
		p.Revision = cfg.revision
	}
	for _, v := range cfg.VMs.GetVms() {
		for _, ch := range v.Amount.GetChange() {
			if err := ch.Length.Normalize(); err != nil {
				return errors.Annotate(err, "failed to normalize %q", v.Prefix).Err()
			}
		}
		amt, err := v.ComputeAmount(v.GetCurrentAmount(), clock.Now(c))
		if err != nil {
			return errors.Annotate(err, "failed to normalize %q", v.Prefix).Err()
		}
		v.CurrentAmount = amt
		if err := v.Lifetime.Normalize(); err != nil {
			return errors.Annotate(err, "failed to normalize %q", v.Prefix).Err()
		}
		v.Revision = cfg.revision
		if err := v.Timeout.Normalize(); err != nil {
			return errors.Annotate(err, "failed to normalize %q", v.Prefix).Err()
		}
	}
	return nil
}

// syncVMs synchronizes the given validated VM configs.
func syncVMs(c context.Context, vms []*gce.Config) error {
	// Fetch existing configs.
	srv := getVMsServer(c)
	rsp, err := srv.List(c, &gce.ListRequest{})
	if err != nil {
		return errors.Annotate(err, "failed to fetch VMs configs").Err()
	}
	// Track the revision of each config.
	revs := make(map[string]string, len(rsp.Configs))
	for _, v := range rsp.Configs {
		revs[v.Prefix] = v.Revision
	}
	logging.Debugf(c, "fetched %d VMs configs", len(rsp.Configs))

	// Update configs to new revisions.
	ens := &gce.EnsureRequest{}
	for _, v := range vms {
		rev, ok := revs[v.Prefix]
		delete(revs, v.Prefix)
		if ok && rev == v.Revision {
			continue
		}
		ens.Id = v.Prefix
		ens.Config = v
		if _, err := srv.Ensure(c, ens); err != nil {
			return errors.Annotate(err, "failed to ensure VMs config %q", ens.Id).Err()
		}
	}

	// Delete unreferenced configs.
	del := &gce.DeleteRequest{}
	for id := range revs {
		del.Id = id
		if _, err := srv.Delete(c, del); err != nil {
			return errors.Annotate(err, "failed to delete VMs config %q", del.Id).Err()
		}
		logging.Debugf(c, "deleted VMs config %q", del.Id)
	}
	return nil
}

// syncPrjs synchronizes the given validated project configs.
func syncPrjs(c context.Context, prjs []*projects.Config) error {
	// Fetch existing configs.
	srv := getProjServer(c)
	rsp, err := srv.List(c, &projects.ListRequest{})
	if err != nil {
		return errors.Annotate(err, "failed to fetch project configs").Err()
	}
	// Track the revision of each config.
	revs := make(map[string]string, len(rsp.Projects))
	for _, p := range rsp.Projects {
		revs[p.Project] = p.Revision
	}
	logging.Debugf(c, "fetched %d project configs", len(rsp.Projects))

	// Update configs to new revisions.
	ens := &projects.EnsureRequest{}
	for _, p := range prjs {
		rev, ok := revs[p.Project]
		delete(revs, p.Project)
		if ok && rev == p.Revision {
			continue
		}
		ens.Id = p.Project
		ens.Project = p
		if _, err := srv.Ensure(c, ens); err != nil {
			return errors.Annotate(err, "failed to ensure project config %q", ens.Id).Err()
		}
	}

	// Delete unreferenced configs.
	del := &projects.DeleteRequest{}
	for id := range revs {
		del.Id = id
		if _, err := srv.Delete(c, del); err != nil {
			return errors.Annotate(err, "failed to delete project config %q", del.Id).Err()
		}
		logging.Debugf(c, "deleted project config %q", del.Id)
	}
	return nil
}

// sync synchronizes the given validated configs.
func sync(c context.Context, cfg *Config) error {
	if err := syncVMs(c, cfg.VMs.GetVms()); err != nil {
		return errors.Annotate(err, "failed to sync VMs configs").Err()
	}
	if err := syncPrjs(c, cfg.Projects.GetProject()); err != nil {
		return errors.Annotate(err, "failed to sync project configs").Err()
	}
	return nil
}

// Import fetches and validates configs from the config service.
func Import(c context.Context) error {
	cfg, err := fetch(c)
	if err != nil {
		return errors.Annotate(err, "failed to fetch configs").Err()
	}

	// Deref before validating. VMs may be invalid until metadata from file is imported.
	if err := deref(c, cfg); err != nil {
		return errors.Annotate(err, "failed to dereference files").Err()
	}

	if err := validate(c, cfg); err != nil {
		return errors.Annotate(err, "invalid configs").Err()
	}

	if err := normalize(c, cfg); err != nil {
		return errors.Annotate(err, "failed to normalize configs").Err()
	}

	if err := sync(c, cfg); err != nil {
		return errors.Annotate(err, "failed to synchronize configs").Err()
	}
	return nil
}

// importHandler imports the config from the config service.
func importHandler(c *router.Context) {
	c.Writer.Header().Set("Content-Type", "text/plain")

	if err := Import(c.Context); err != nil {
		errors.Log(c.Context, err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

// InstallHandlers installs HTTP request handlers into the given router.
func InstallHandlers(r *router.Router, mw router.MiddlewareChain) {
	mw = mw.Extend(func(c *router.Context, next router.Handler) {
		// Install the services.
		c.Context = withProjServer(c.Context, &rpc.Projects{})
		c.Context = withVMsServer(c.Context, &rpc.Config{})
		next(c)
	})
	r.GET("/internal/cron/import-config", mw, importHandler)
}
