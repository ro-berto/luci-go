// Copyright 2016 The LUCI Authors.
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

package swarmingimpl

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bazelbuild/remote-apis-sdks/go/pkg/digest"
	"github.com/google/uuid"
	"github.com/maruel/subcommands"

	"go.chromium.org/luci/common/api/swarming/swarming/v1"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/flag/flagenum"
	"go.chromium.org/luci/common/flag/stringlistflag"
	"go.chromium.org/luci/common/flag/stringmapflag"
	"go.chromium.org/luci/common/system/signals"
)

// CmdTrigger returns an object for the `trigger` subcommand.
func CmdTrigger(authFlags AuthFlags) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "trigger <options>",
		ShortDesc: "Triggers a Swarming task",
		LongDesc:  "Triggers a Swarming task.",
		CommandRun: func() subcommands.CommandRun {
			r := &triggerRun{}
			r.Init(authFlags)
			return r
		},
	}
}

// mapToArray converts a stringmapflag.Value into an array of
// swarming.SwarmingRpcsStringPair, sorted by key and then value.
func mapToArray(m stringmapflag.Value) []*swarming.SwarmingRpcsStringPair {
	a := make([]*swarming.SwarmingRpcsStringPair, 0, len(m))
	for k, v := range m {
		a = append(a, &swarming.SwarmingRpcsStringPair{Key: k, Value: v})
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].Key < a[j].Key ||
			(a[i].Key == a[j].Key && a[i].Value < a[j].Value)
	})
	return a
}

// listToStringListPairArray converts a stringlistflag.Flag into an array of
// swarming.SwarmingRpcsStringListPair, sorted by key.
func listToStringListPairArray(m stringlistflag.Flag) []*swarming.SwarmingRpcsStringListPair {
	prefixes := make(map[string][]string)
	for _, f := range m {
		kv := strings.SplitN(f, "=", 2)
		prefixes[kv[0]] = append(prefixes[kv[0]], kv[1])
	}

	a := make([]*swarming.SwarmingRpcsStringListPair, 0, len(prefixes))

	for key, value := range prefixes {
		a = append(a, &swarming.SwarmingRpcsStringListPair{
			Key:   key,
			Value: value,
		})
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].Key < a[j].Key
	})
	return a
}

// namePartFromDimensions creates a string from a map of dimensions that can
// be used as part of the task name.  The dimensions are first sorted as
// described in mapToArray().
func namePartFromDimensions(m stringmapflag.Value) string {
	a := mapToArray(m)
	pairs := make([]string, 0, len(a))
	for i := 0; i < len(a); i++ {
		pairs = append(pairs, fmt.Sprintf("%s=%s", a[i].Key, a[i].Value))
	}
	return strings.Join(pairs, "_")
}

type containmentType string

func (c *containmentType) String() string {
	return string(*c)
}

func (c *containmentType) Set(v string) error {
	return containmentChoices.FlagSet(c, v)
}

var containmentChoices = flagenum.Enum{
	"none":       containmentType("NONE"),
	"auto":       containmentType("AUTO"),
	"job_object": containmentType("JOB_OBJECT"),
}

type optionalDimension struct {
	kv         *swarming.SwarmingRpcsStringPair
	expiration int64
}

var _ flag.Value = (*optionalDimension)(nil)

// String implements the flag.Value interface.
func (f *optionalDimension) String() string {
	if f == nil || f.isEmpty() {
		return ""
	}
	return fmt.Sprintf("kv=%+v expiration=%d", *f.kv, f.expiration)
}

// Set implements the flag.Value interface.
func (f *optionalDimension) Set(s string) error {
	if s == "" {
		return nil
	}
	splits := strings.SplitN(s, "=", 2)

	if len(splits) != 2 {
		return errors.Reason("cannot find key in the optional dimension: %q", s).Err()
	}
	k := splits[0]
	valExp := splits[1]
	colon := strings.LastIndexByte(valExp, ':')
	if colon == -1 {
		return errors.Reason(`cannot find ":" between value and expiration in the optional dimension: %q`, valExp).Err()
	}
	exp, err := strconv.ParseInt(valExp[colon+1:], 10, 64)
	if err != nil {
		return errors.Reason("cannot parse the expiration in the optional dimension: %q", valExp).Err()
	}
	f.kv = &swarming.SwarmingRpcsStringPair{Key: k, Value: valExp[:colon]}
	f.expiration = exp
	return nil
}

func (f *optionalDimension) isEmpty() bool {
	return f.kv == nil
}

type triggerRun struct {
	commonFlags

	// Task properties.
	casInstance       string
	digest            string
	dimensions        stringmapflag.Value
	env               stringmapflag.Value
	envPrefix         stringlistflag.Flag
	idempotent        bool
	containmentType   containmentType
	namedCache        stringmapflag.Value
	hardTimeout       int64
	ioTimeout         int64
	gracePeriod       int64
	cipdPackage       stringmapflag.Value
	outputs           stringlistflag.Flag
	optionalDimension optionalDimension
	serviceAccount    string
	relativeCwd       string
	secretBytesPath   string

	// Task request.
	taskName       string
	priority       int64
	tags           stringlistflag.Flag
	user           string
	expiration     int64
	enableResultDB bool
	realm          string

	// Other.
	dumpJSON string
}

func (c *triggerRun) Init(authFlags AuthFlags) {
	c.commonFlags.Init(authFlags)
	// Task properties.
	c.Flags.StringVar(&c.casInstance, "cas-instance", "", "CAS instance (GCP). Format is \"projects/<project_id>/instances/<instance_id>\". Default is constructed from -server.")
	c.Flags.StringVar(&c.digest, "digest", "", "Digest of root directory uploaded to CAS `<Hash>/<Size>`.")
	c.Flags.Var(&c.dimensions, "dimension", "Dimension to select the right kind of bot. In the form of `key=value`")
	c.Flags.Var(&c.dimensions, "d", "Alias for -dimension.")
	c.Flags.Var(&c.env, "env", "Environment variables to set.")
	c.Flags.Var(&c.envPrefix, "env-prefix", "Environment prefixes to set.")
	c.Flags.BoolVar(&c.idempotent, "idempotent", false, "When set, the server will actively try to find a previous task with the same parameter and return this result instead if possible.")
	c.containmentType = "NONE"
	c.Flags.Var(&c.containmentType, "containment-type", "Specify which type of process containment to use. Choices are: "+containmentChoices.Choices())
	c.Flags.Int64Var(&c.hardTimeout, "hard-timeout", 60*60, "Seconds to allow the task to complete.")
	c.Flags.Int64Var(&c.ioTimeout, "io-timeout", 20*60, "Seconds to allow the task to be silent.")
	c.Flags.Int64Var(&c.gracePeriod, "grace-period", 30, "Seconds to wait after sending SIGBREAK.")
	c.Flags.Var(&c.cipdPackage, "cipd-package",
		"(repeatable) CIPD packages to install on the swarming bot. This takes a parameter of `[installdir:]pkgname=version`. "+
			"Using an empty version will remove the package. The installdir is optional and defaults to '.'.")
	c.Flags.Var(&c.namedCache, "named-cache", "This takes a parameter of `name=cachedir`.")
	c.Flags.Var(&c.outputs, "output", "(repeatable) Specify an output file or directory that can be retrieved via collect.")
	c.Flags.Var(&c.optionalDimension, "optional-dimension", "Format: <key>=<value>:<expiration>. See -expiration for the requirement.")
	c.Flags.StringVar(&c.relativeCwd, "relative-cwd", "", "Use this flag instead of the isolated 'relative_cwd'.")
	c.Flags.StringVar(&c.serviceAccount, "service-account", "",
		`Email of a service account to run the task as, or literal "bot" string to indicate that the task should use the same account the bot itself is using to authenticate to Swarming. Don't use task service accounts if not given (default).`)
	c.Flags.StringVar(&c.secretBytesPath, "secret-bytes-path", "", "Specify the secret bytes file path.")

	// Task request.
	c.Flags.StringVar(&c.taskName, "task-name", "", "Display name of the task. Defaults to <base_name>/<dimensions>/<isolated hash>/<timestamp> if an  isolated file is provided, if a hash is provided, it defaults to <user>/<dimensions>/<isolated hash>/<timestamp>")
	c.Flags.Int64Var(&c.priority, "priority", 200, "The lower value, the more important the task.")
	c.Flags.Var(&c.tags, "tag", "Tags to assign to the task. In the form of `key:value`.")
	c.Flags.StringVar(&c.user, "user", "", "User associated with the task. Defaults to authenticated user on the server.")
	c.Flags.Int64Var(&c.expiration, "expiration", 6*60*60, "Seconds to allow the task to be pending for a bot to run before this task request expires.")
	c.Flags.BoolVar(&c.enableResultDB, "enable-resultdb", false, "Enable ResultDB for this task.")
	c.Flags.StringVar(&c.realm, "realm", "", "Realm name for this task.")

	// Other.
	c.Flags.StringVar(&c.dumpJSON, "dump-json", "", "Dump details about the triggered task(s) to this file as json.")
}

func (c *triggerRun) Parse(args []string) error {
	var err error
	if err := c.commonFlags.Parse(); err != nil {
		return err
	}

	// Validate options and args.
	if c.dimensions == nil {
		return errors.Reason("please at least specify one dimension").Err()
	}

	if len(args) == 0 {
		return errors.Reason("please specify command after '--'").Err()
	}

	if len(c.user) == 0 {
		c.user = os.Getenv(UserEnvVar)
	}

	return err
}

func (c *triggerRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	if err := c.Parse(args); err != nil {
		printError(a, err)
		return 1
	}
	if err := c.main(a, args, env); err != nil {
		printError(a, err)
		return 1
	}
	return 0
}

func (c *triggerRun) main(a subcommands.Application, args []string, env subcommands.Env) (rerr error) {
	start := time.Now()
	ctx, cancel := context.WithCancel(c.defaultFlags.MakeLoggingContext(os.Stderr))
	defer cancel()
	defer signals.HandleInterrupt(cancel)()

	request, err := c.processTriggerOptions(args, env)
	if err != nil {
		return errors.Annotate(err, "failed to process trigger options").Err()
	}

	service, err := c.createSwarmingClient(ctx)
	if err != nil {
		return err
	}

	result, err := service.NewTask(ctx, request)
	if err != nil {
		return err
	}

	if c.dumpJSON != "" {
		dump, err := os.Create(c.dumpJSON)
		if err != nil {
			return err
		}
		defer func() {
			if err := dump.Close(); rerr == nil {
				rerr = err
			}
		}()

		data := TriggerResults{Tasks: []*swarming.SwarmingRpcsTaskRequestMetadata{result}}
		b, err := json.MarshalIndent(&data, "", "  ")
		if err != nil {
			return errors.Annotate(err, "marshalling trigger result").Err()
		}

		_, err = dump.Write(b)
		if err != nil {
			return errors.Annotate(err, "writing json dump").Err()
		}

		if !c.defaultFlags.Quiet {
			fmt.Println("To collect results use:")
			fmt.Printf("  %s collect -server %s -output-dir out -task-summary-json summary.json -requests-json %s\n", os.Args[0], c.serverURL, c.dumpJSON)
		}
	} else if !c.defaultFlags.Quiet {
		fmt.Println("To collect results use:")
		fmt.Printf("  swarming collect -server %s -output-dir out -task-summary-json summary.json %s", c.serverURL, result.TaskId)
		fmt.Println()
	}

	if !c.defaultFlags.Quiet {
		fmt.Println("You can also see the task status in")
		fmt.Printf("  %s/task?id=%s\n", c.serverURL, result.TaskId)
	}

	duration := time.Since(start)
	log.Printf("Duration: %s\n", duration.Round(time.Millisecond))
	return nil
}

func (c *triggerRun) createTaskSliceForOptionalDimension(properties *swarming.SwarmingRpcsTaskProperties) (*swarming.SwarmingRpcsTaskSlice, error) {
	if c.optionalDimension.isEmpty() {
		return nil, nil
	}
	optDim := c.optionalDimension.kv
	exp := c.optionalDimension.expiration

	// Deep copy properties
	pj, err := properties.MarshalJSON()
	if err != nil {
		return nil, errors.Annotate(err, "failed to marshall properties").Err()
	}
	propsCpy := &swarming.SwarmingRpcsTaskProperties{}
	if err = json.Unmarshal(pj, propsCpy); err != nil {
		return nil, errors.Annotate(err, "failed to unmarshall properties").Err()
	}
	propsCpy.Dimensions = append(propsCpy.Dimensions, optDim)

	return &swarming.SwarmingRpcsTaskSlice{
		ExpirationSecs: exp,
		Properties:     propsCpy,
	}, nil
}

func (c *triggerRun) processTriggerOptions(commands []string, env subcommands.Env) (*swarming.SwarmingRpcsNewTaskRequest, error) {
	if c.taskName == "" {
		c.taskName = fmt.Sprintf("%s/%s", c.user, namePartFromDimensions(c.dimensions))
	}

	var secretBytesEnc string
	if c.secretBytesPath != "" {
		secretBytes, err := os.ReadFile(c.secretBytesPath)
		if err != nil {
			return nil, errors.Annotate(err, "failed to read secret bytes from %s", c.secretBytesPath).Err()
		}
		secretBytesEnc = base64.StdEncoding.EncodeToString(secretBytes)
	}

	var CASRef *swarming.SwarmingRpcsCASReference
	if c.digest != "" {
		d, err := digest.NewFromString(c.digest)
		if err != nil {
			return nil, errors.Annotate(err, "invalid digest: %s", c.digest).Err()
		}

		casInstance := c.casInstance
		if casInstance == "" {
			// infer cas instance from server URL.
			u, err := url.Parse(c.serverURL)
			if err != nil {
				return nil, errors.Annotate(err, "invalid server url: %s", c.serverURL).Err()
			}

			const appspot = ".appspot.com"
			if !strings.HasSuffix(u.Host, appspot) {
				return nil, errors.Reason("server url should have '%s' suffix: %s", appspot, c.serverURL).Err()
			}

			casInstance = "projects/" + strings.TrimSuffix(u.Host, appspot) + "/instances/default_instance"
		}

		CASRef = &swarming.SwarmingRpcsCASReference{
			CasInstance: casInstance,
			Digest: &swarming.SwarmingRpcsDigest{
				Hash:      d.Hash,
				SizeBytes: d.Size,

				// 0 is valid value for SizeBytes.
				ForceSendFields: []string{"SizeBytes"},
			},
		}
	}

	properties := swarming.SwarmingRpcsTaskProperties{
		Command:              commands,
		RelativeCwd:          c.relativeCwd,
		Dimensions:           mapToArray(c.dimensions),
		Env:                  mapToArray(c.env),
		EnvPrefixes:          listToStringListPairArray(c.envPrefix),
		ExecutionTimeoutSecs: c.hardTimeout,
		GracePeriodSecs:      c.gracePeriod,
		Idempotent:           c.idempotent,
		CasInputRoot:         CASRef,
		Outputs:              c.outputs,
		IoTimeoutSecs:        c.ioTimeout,
		Containment: &swarming.SwarmingRpcsContainment{
			ContainmentType: string(c.containmentType),
		},
		SecretBytes: secretBytesEnc,
	}

	if len(c.cipdPackage) > 0 {
		var pkgs []*swarming.SwarmingRpcsCipdPackage
		for k, v := range c.cipdPackage {
			s := strings.SplitN(k, ":", 2)
			pkg := swarming.SwarmingRpcsCipdPackage{
				PackageName: s[len(s)-1],
				Version:     v,
			}
			if len(s) > 1 {
				pkg.Path = s[0]
			}
			pkgs = append(pkgs, &pkg)
		}

		sort.Slice(pkgs, func(i, j int) bool {
			pi, pj := pkgs[i], pkgs[j]
			if pi.PackageName != pj.PackageName {
				return pi.PackageName < pj.PackageName
			}
			if pi.Version != pj.Version {
				return pi.Version < pj.Version
			}
			return pi.Path < pj.Path
		})

		properties.CipdInput = &swarming.SwarmingRpcsCipdInput{Packages: pkgs}
	}

	for name, path := range c.namedCache {
		properties.Caches = append(properties.Caches,
			&swarming.SwarmingRpcsCacheEntry{
				Name: name,
				Path: path,
			},
		)
	}

	sort.Slice(properties.Caches, func(i, j int) bool {
		ci, cj := properties.Caches[i], properties.Caches[j]
		if ci.Name != cj.Name {
			return ci.Name < cj.Name
		}
		return ci.Path < cj.Path
	})

	randomUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Annotate(err, "failed to get random UUID").Err()
	}

	var taskSlices []*swarming.SwarmingRpcsTaskSlice
	taskSlice, err := c.createTaskSliceForOptionalDimension(&properties)
	if err != nil {
		return nil, errors.Annotate(err, "failed to createTaskSliceForOptionalDimension").Err()
	}
	baseExpiration := c.expiration
	if taskSlice != nil {
		taskSlices = append(taskSlices, taskSlice)

		baseExpiration -= taskSlice.ExpirationSecs
		if baseExpiration < 60 {
			baseExpiration = 60
		}
	}
	taskSlices = append(taskSlices, &swarming.SwarmingRpcsTaskSlice{
		ExpirationSecs: baseExpiration,
		Properties:     &properties,
	})

	return &swarming.SwarmingRpcsNewTaskRequest{
		Name:           c.taskName,
		ParentTaskId:   env[TaskIDEnvVar].Value,
		Priority:       c.priority,
		ServiceAccount: c.serviceAccount,
		Tags:           c.tags,
		TaskSlices:     taskSlices,
		User:           c.user,
		RequestUuid:    randomUUID.String(),
		Resultdb: &swarming.SwarmingRpcsResultDBCfg{
			Enable: c.enableResultDB,
		},
		Realm: c.realm,
	}, nil
}
