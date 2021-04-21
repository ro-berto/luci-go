// Copyright 2021 The LUCI Authors.
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

package lib

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/maruel/subcommands"

	"go.chromium.org/luci/cipd/client/cipd"
	"go.chromium.org/luci/cipd/client/cipd/ensure"
	"go.chromium.org/luci/cipd/client/cipd/template"
	clientswarming "go.chromium.org/luci/client/swarming"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/system/environ"
	"go.chromium.org/luci/common/system/signals"
	"go.chromium.org/luci/hardcoded/chromeinfra"
)

// CmdReproduce returns an object fo the `reproduce` subcommand.
func CmdReproduce(authFlags AuthFlags) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "reproduce -S <server> <task ID> ",
		ShortDesc: "reproduces a task locally",
		LongDesc:  "Fetches a TaskRequest and runs the same commands that were run on the bot.",
		CommandRun: func() subcommands.CommandRun {
			r := &reproduceRun{}
			r.init(authFlags)
			return r
		},
	}
}

type reproduceRun struct {
	commonFlags
	work string
	// cipdDownloader is used in testing to insert a mock CIPD downloader.
	cipdDownloader func(context.Context, string, map[string]ensure.PackageSlice) error
}

func (c *reproduceRun) init(authFlags AuthFlags) {
	c.commonFlags.Init(authFlags)

	c.Flags.StringVar(&c.work, "work", "work", "Directory to map the task input files into and execute the task.")
	c.cipdDownloader = downloadCIPDPackages
	// TODO(crbug.com/1188473): support cache and output directories.
}

func (c *reproduceRun) parse(args []string) error {
	if err := c.commonFlags.Parse(); err != nil {
		return err
	}
	if len(args) != 1 {
		return errors.Reason("must specify exactly one task id.").Err()
	}
	return nil
}

func (c *reproduceRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	if err := c.parse(args); err != nil {
		printError(a, err)
		return 1
	}
	if err := c.main(a, args, env); err != nil {
		printError(a, err)
		return 1
	}
	return 0
}

func (c *reproduceRun) main(a subcommands.Application, args []string, env subcommands.Env) error {
	ctx, cancel := context.WithCancel(c.defaultFlags.MakeLoggingContext(os.Stderr))
	defer cancel()
	defer signals.HandleInterrupt(cancel)()

	service, err := c.createSwarmingClient(ctx)
	if err != nil {
		return err
	}

	cmd, err := c.prepareTaskRequestEnvironment(ctx, args[0], service)
	if err != nil {
		return errors.Annotate(err, "failed to create command from task request").Err()
	}

	return c.executeTaskRequestCommand(cmd)
}

func (c *reproduceRun) executeTaskRequestCommand(cmd *exec.Cmd) error {
	if err := cmd.Start(); err != nil {
		return errors.Annotate(err, "failed to start command: %v", cmd).Err()
	}
	if err := cmd.Wait(); err != nil {
		return errors.Annotate(err, "failed to complete command: %v", cmd).Err()
	}
	return nil
}

func (c *reproduceRun) prepareTaskRequestEnvironment(ctx context.Context, taskID string, service swarmingService) (*exec.Cmd, error) {
	tr, err := service.GetTaskRequest(ctx, taskID)
	if err != nil {
		return nil, errors.Annotate(err, "failed to get task request: %s", taskID).Err()
	}
	// In practice, later slices are less likely to assume that there is a named cache
	// that is not available locally.
	properties := tr.TaskSlices[len(tr.TaskSlices)-1].Properties

	workdir := c.work
	if properties.RelativeCwd != "" {
		workdir = filepath.Join(workdir, properties.RelativeCwd)
	}
	if err := prepareDir(workdir); err != nil {
		return nil, err
	}

	// Set environment variables.
	cmdEnvMap := environ.FromCtx(ctx)
	for _, env := range properties.Env {
		if env.Value == "" {
			cmdEnvMap.Remove(env.Key)
		} else {
			cmdEnvMap.Set(env.Key, env.Value)
		}
	}

	// Set environment prefixes.
	for _, prefix := range properties.EnvPrefixes {
		paths := make([]string, 0, len(prefix.Value)+1)
		for _, value := range prefix.Value {
			paths = append(paths, filepath.Clean(filepath.Join(workdir, value)))
		}
		cur, ok := cmdEnvMap.Get(prefix.Key)
		if ok {
			paths = append(paths, cur)
		}
		cmdEnvMap.Set(prefix.Key, strings.Join(paths, string(os.PathListSeparator)))
	}

	// Download input files.
	if properties.InputsRef != nil && properties.InputsRef.Isolated != "" && properties.CasInputRoot != nil {
		return nil, errors.Reason("fetched TaskRequest has files from Isolate and RBE-CAS").Err()
	}

	// Support isolated input in task request.
	if properties.InputsRef != nil && properties.InputsRef.Isolated != "" {
		if _, err := service.GetFilesFromIsolate(ctx, workdir, properties.InputsRef); err != nil {
			return nil, errors.Annotate(err, "failed to fetch files from isolate").Err()
		}
	}

	// Support RBE-CAS input in task request.
	if properties.CasInputRoot != nil {
		cascli, err := c.authFlags.NewCASClient(ctx, properties.CasInputRoot.CasInstance)
		if err != nil {
			return nil, errors.Annotate(err, "failed to fetch RBE-CAS client").Err()
		}
		if _, err := service.GetFilesFromCAS(ctx, workdir, cascli, properties.CasInputRoot); err != nil {
			return nil, errors.Annotate(err, "failed to fetched friles from RBE-CAS").Err()
		}
	}

	// Support CIPD package download in task request.
	if properties.CipdInput != nil {
		packages := properties.CipdInput.Packages
		slicesByPath := map[string]ensure.PackageSlice{}
		for _, pkg := range packages {
			path := pkg.Path
			// CIPD deals with 'root' as ''.
			if path == "." {
				path = ""
			}
			if _, ok := slicesByPath[path]; !ok {
				slicesByPath[path] = make(ensure.PackageSlice, 0, len(packages))
			}
			slicesByPath[path] = append(
				slicesByPath[path], ensure.PackageDef{UnresolvedVersion: pkg.Version, PackageTemplate: pkg.PackageName})
		}

		if err := c.cipdDownloader(ctx, workdir, slicesByPath); err != nil {
			return nil, err
		}
	}

	// Create a Comand that can run the task request.
	processedCmds, err := clientswarming.ProcessCommand(ctx, properties.Command, workdir, "")
	if err != nil {
		return nil, errors.Annotate(err, "failed to process command in properties").Err()
	}
	cmd := exec.CommandContext(ctx, processedCmds[0], processedCmds[1:]...)
	cmd.Env = cmdEnvMap.Sorted()
	cmd.Dir = workdir
	return cmd, nil
}

func downloadCIPDPackages(ctx context.Context, workdir string, slicesByPath map[string]ensure.PackageSlice) error {
	// Create CIPD client.
	opts := cipd.ClientOptions{
		Root:       workdir,
		ServiceURL: chromeinfra.CIPDServiceURL,
	}
	client, err := cipd.NewClient(opts)
	if err != nil {
		return errors.Annotate(err, "failed to create CIPD client").Err()
	}
	defer client.Close(ctx)

	// Resolve versions.
	resolver := cipd.Resolver{Client: client}
	resolved, err := resolver.Resolve(
		ctx, &ensure.File{ServiceURL: chromeinfra.CIPDServiceURL, PackagesBySubdir: slicesByPath}, template.DefaultExpander())
	if err != nil {
		return errors.Annotate(err, "failed to resolve CIPD package versions").Err()
	}

	// Download packages.
	if _, err := client.EnsurePackages(ctx, resolved.PackagesBySubdir, resolved.ParanoidMode, 1, false); err != nil {
		return errors.Annotate(err, "failed to install or update CIPD packages").Err()
	}
	return nil

}

func prepareDir(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return errors.Annotate(err, "failed to remove directory: %s", dir).Err()
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return errors.Annotate(err, "failed to create directory: %s", dir).Err()
	}
	return nil
}
