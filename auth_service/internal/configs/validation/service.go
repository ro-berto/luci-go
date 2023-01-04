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

package validation

import (
	"fmt"
	"net"
	"regexp"
	"sort"
	"strings"

	"google.golang.org/protobuf/encoding/prototext"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/auth_service/api/configspb"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/lhttp"
	"go.chromium.org/luci/config/validation"
	"go.chromium.org/luci/server/auth/service/protocol"
)

// ipAllowlistNameRE is the regular expression for IP Allowlist Names.
var ipAllowlistNameRE = regexp.MustCompile(`^[0-9A-Za-z_\-\+\.\ ]{2,200}$`)

// validateAllowlist validates an ip_allowlist.cfg file.
func validateAllowlist(ctx *validation.Context, configSet, path string, content []byte) error {
	cfg := configspb.IPAllowlistConfig{}

	if err := prototext.Unmarshal(content, &cfg); err != nil {
		ctx.Error(err)
		return nil
	}

	ctx.SetFile(path)
	if err := validateAllowlistCfg(&cfg); err != nil {
		ctx.Error(err)
	}

	if _, err := GetSubnets(cfg.GetIpAllowlists()); err != nil {
		ctx.Error(err)
	}
	return nil
}

func validateAllowlistCfg(cfg *configspb.IPAllowlistConfig) error {
	// Allowlist validation.
	allowlists := stringset.New(len(cfg.GetIpAllowlists()))
	for _, a := range cfg.GetIpAllowlists() {
		switch name := a.GetName(); {
		case !ipAllowlistNameRE.MatchString(name):
			return errors.New(fmt.Sprintf("invalid ip allowlist name %s", name))
		case allowlists.Has(name):
			return errors.New(fmt.Sprintf("ip allowlist is defined twice %s", name))
		default:
			allowlists.Add(name)
		}

		// Validate subnets, check that the format is valid.
		// Either in CIDR format or just a textual representation
		// of an IP.
		// e.g. "192.0.0.1", "127.0.0.1/23"
		for _, subnet := range a.Subnets {
			if strings.Contains(subnet, "/") {
				if _, _, err := net.ParseCIDR(subnet); err != nil {
					return err
				}
			} else {
				if ip := net.ParseIP(subnet); ip == nil {
					return errors.New(fmt.Sprintf("unable to parse ip for subnet: %s", subnet))
				}
			}
		}
	}

	// Assignment validation
	idents := stringset.New(len(cfg.GetAssignments()))
	for _, a := range cfg.GetAssignments() {
		ident := a.GetIdentity()
		alName := a.GetIpAllowlistName()

		// Checks if valid Identity.
		if _, err := identity.MakeIdentity(ident); err != nil {
			return err
		}
		if !allowlists.Has(alName) {
			return errors.New(fmt.Sprintf("unknown allowlist %s", alName))
		}
		if idents.Has(ident) {
			return errors.New(fmt.Sprintf("identity %s defined twice", ident))
		}
		idents.Add(ident)
	}

	return nil
}

func validateOAuth(ctx *validation.Context, configSet, path string, content []byte) error {
	cfg := configspb.OAuthConfig{}

	if err := prototext.Unmarshal(content, &cfg); err != nil {
		ctx.Error(err)
		return nil
	}

	ctx.SetFile(path)
	if cfg.GetTokenServerUrl() != "" {
		if _, err := lhttp.CheckURL(cfg.GetTokenServerUrl()); err != nil {
			ctx.Error(err)
		}
	}

	return nil
}

func validateSecurityCfg(ctx *validation.Context, configSet, path string, content []byte) error {
	ctx.SetFile(path)
	cfg := protocol.SecurityConfig{}

	if err := prototext.Unmarshal(content, &cfg); err != nil {
		ctx.Error(err)
		return nil
	}

	ctx.Enter("internal_service_regexp")
	for i, re := range cfg.GetInternalServiceRegexp() {
		ctx.Enter(fmt.Sprintf("# %d", i))
		if _, err := regexp.Compile(fmt.Sprintf("^%s$", re)); err != nil {
			ctx.Error(err)
		}
		ctx.Exit()
	}
	ctx.Exit()

	return nil
}

func validateImportsCfg(ctx *validation.Context, configSet, path string, content []byte) error {
	ctx.SetFile(path)
	cfg := configspb.GroupImporterConfig{}
	urlErr := errors.New("url field required")
	ctx.Enter("validating imports.cfg")
	defer ctx.Exit()

	if err := prototext.Unmarshal(content, &cfg); err != nil {
		ctx.Error(err)
	}

	ctx.Enter("checking tarball URLs...")
	for _, tb := range cfg.GetTarball() {
		if tb.Url == "" {
			ctx.Error(urlErr)
		}
	}
	ctx.Exit()

	ctx.Enter("checking plainlist URLs...")
	for _, pl := range cfg.GetPlainlist() {
		if pl.Url == "" {
			ctx.Error(urlErr)
		}
	}
	ctx.Exit()

	ctx.Enter("validating tarball_upload names...")
	tarballUploadNames := make(map[string]bool)
	for _, entry := range cfg.GetTarballUpload() {
		entryName := entry.GetName()
		if entryName == "" {
			ctx.Error(errors.New("Some tarball_upload entry doesn't have a name"))
		}

		if tarballUploadNames[entryName] {
			ctx.Errorf("tarball_upload entry %s is specified twice", entryName)
		}
		tarballUploadNames[entryName] = true

		authorizedUploader := entry.GetAuthorizedUploader()
		if authorizedUploader == nil {
			ctx.Errorf("authorized_uploader is required in tarball_upload entry %s", entryName)
		}

		for _, email := range authorizedUploader {
			_, err := identity.MakeIdentity(fmt.Sprintf("user:%s", email))
			if err != nil {
				ctx.Error(err)
			}
		}
	}
	ctx.Exit()

	ctx.Enter("validating systems")
	seenSystems := make(map[string]bool)
	seenSystems["external"] = true
	for _, entry := range cfg.GetTarball() {
		title := fmt.Sprintf(`"tarball" entry with URL %q`, entry.GetUrl())
		if err := validateSystems(entry.GetSystems(), seenSystems, title); err != nil {
			ctx.Error(err)
		}
	}
	for _, entry := range cfg.GetTarballUpload() {
		title := fmt.Sprintf(`"tarball_upload" entry with name %q`, entry.GetName())
		if err := validateSystems(entry.GetSystems(), seenSystems, title); err != nil {
			ctx.Error(err)
		}
	}
	ctx.Exit()

	ctx.Enter("validating plainlist groups")
	seenGroups := make(map[string]bool)
	for _, entry := range cfg.GetPlainlist() {
		group := entry.GetGroup()
		if group == "" {
			ctx.Errorf(`"plainlist" entry %q needs a "group" field`, entry.GetUrl())
		}
		if seenGroups[group] {
			ctx.Errorf(`the group %q is imported twice`, group)
		}
		seenGroups[group] = true
	}
	ctx.Exit()

	return nil
}

func validateSystems(systems []string, seenSystems map[string]bool, title string) error {
	if systems == nil {
		return errors.New(fmt.Sprintf(`%s needs a "systems" field`, title))
	}
	twice := []string{}
	for _, system := range systems {
		if seenSystems[system] {
			twice = append(twice, system)
		} else {
			seenSystems[system] = true
		}
	}
	if len(twice) > 0 {
		sort.Strings(twice)
		return errors.New(fmt.Sprintf("%s is specifying a duplicated system(s): %v", title, twice))
	}
	return nil
}

// GetSubnets validates the includes of all allowlists and generates a map {allowlistName: []subnets}.
func GetSubnets(allowlists []*configspb.IPAllowlistConfig_IPAllowlist) (map[string][]string, error) {
	allowlistsByName := make(map[string]*configspb.IPAllowlistConfig_IPAllowlist, len(allowlists))
	for _, al := range allowlists {
		allowlistsByName[al.GetName()] = al
	}

	subnetMap := make(map[string][]string, len(allowlists))
	for _, al := range allowlists {
		subnets, err := getSubnetsRecursive(al, make([]string, 0, len(allowlists)), allowlistsByName, subnetMap)
		if err != nil {
			return nil, err
		}
		subnetMap[al.GetName()] = subnets
	}
	return subnetMap, nil
}

// getSubnetsRecursive does a depth first search traversal to find all transitively included subnets for a given allowlist.
func getSubnetsRecursive(al *configspb.IPAllowlistConfig_IPAllowlist, visiting []string, allowlistsByName map[string]*configspb.IPAllowlistConfig_IPAllowlist, subnetMap map[string][]string) ([]string, error) {
	alName := al.GetName()

	// If we've already seen this allowlist before.
	if val, ok := subnetMap[alName]; ok {
		return val, nil
	}

	// Cycle check.
	if contains(visiting, alName) {
		errorCycle := fmt.Sprintf("%s -> %s", strings.Join(visiting, " -> "), alName)
		return nil, errors.New(fmt.Sprintf("IP allowlist is part of an included cycle %s", errorCycle))
	}

	visiting = append(visiting, alName)
	subnets := stringset.NewFromSlice(al.GetSubnets()...)
	for _, inc := range al.GetIncludes() {
		val, ok := allowlistsByName[inc]
		if !ok {
			return nil, errors.New(fmt.Sprintf("IP Allowlist contains unknown allowlist %s", inc))
		}

		resolved, err := getSubnetsRecursive(val, visiting, allowlistsByName, subnetMap)
		if err != nil {
			return nil, err
		}
		subnets.AddAll(resolved)
	}
	return subnets.ToSortedSlice(), nil
}

func contains(s []string, val string) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}
