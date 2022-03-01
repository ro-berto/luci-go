// Copyright 2020 The LUCI Authors.
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

package ledcmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/led/job"
)

// EditRecipeBundleOpts are user-provided options for the recipe bundling
// process.
type EditRecipeBundleOpts struct {
	// Path on disk to the repo to extract the recipes from. May be a subdirectory
	// of the repo, as long as `git rev-parse --show-toplevel` can find the root
	// of the repository.
	//
	// If empty, uses the current working directory.
	RepoDir string

	// Overrides is a mapping of recipe project id (e.g. "recipe_engine") to
	// a local path to a checkout of that repo (e.g. "/path/to/recipes-py.git").
	//
	// When the bundle is created, this local repo will be used instead of the
	// pinned version of this recipe project id. This is helpful for preparing
	// bundles which have code changes in multiple recipe repos.
	Overrides map[string]string

	// DebugSleep is the amount of time to wait after the recipe completes
	// execution (either success or failure). This is injected into the generated
	// recipe bundle as a 'sleep X' command after the invocation of the recipe
	// itself.
	DebugSleep time.Duration

	// PropertyOnly determines whether to pass the recipe bundle's CAS reference
	// as a property and preserve the executable and payload of the input job
	// rather than overwriting it.
	PropertyOnly bool
}

const (
	// RecipeDirectory is a very unfortunate constant which is here for
	// a combination of reasons:
	//   1) swarming doesn't allow you to 'checkout' an isolate relative to any
	//      path in the task (other than the task root). This means that
	//      whatever value we pick for EditRecipeBundle must be used EVERYWHERE
	//      the isolated hash is used.
	//   2) Currently the 'recipe_engine/led' module will blindly take the
	//      isolated input and 'inject' it into further uses of led. This module
	//      currently doesn't specify the checkout dir, relying on kitchen's
	//      default value of (you guessed it) "kitchen-checkout".
	//
	// In order to fix this (and it will need to be fixed for bbagent support):
	//   * The 'recipe_engine/led' module needs to accept 'checkout-dir' as
	//     a parameter in its input properties.
	//   * led needs to start passing the checkout dir to the led module's input
	//     properties.
	//   * `led edit` needs a way to manipulate the checkout directory in a job
	//   * The 'recipe_engine/led' module needs to set this in the job
	//     alongside the isolate hash when it's doing the injection.
	//
	// For now, we just hard-code it.
	//
	// TODO(crbug.com/1072117): Fix this, it's weird.
	RecipeDirectory = "kitchen-checkout"

	// A property that should be set to a boolean value. If true,
	// edit-recipe-bundle will set the "led_cas_recipe_bundle" property
	// instead of overwriting the build's payload.
	LEDBuilderIsBootstrappedProperty = "led_builder_is_bootstrapped"

	// In PropertyOnly mode or if the "led_builder_is_bootstrapped" property
	// of the build is true, this property will be set with the CAS digest
	// of the executable of the recipe bundle.
	CASRecipeBundleProperty = "led_cas_recipe_bundle"
)

// EditRecipeBundle overrides the recipe bundle in the given job with one
// located on disk.
//
// It isolates the recipes from the repository in the given working directory
// into the UserPayload under the directory "kitchen-checkout/". If there's an
// existing directory in the UserPayload at that location, it will be removed.
func EditRecipeBundle(ctx context.Context, authOpts auth.Options, jd *job.Definition, opts *EditRecipeBundleOpts) error {
	if jd.GetBuildbucket() == nil {
		return errors.New("ledcmd.EditRecipeBundle is only available for Buildbucket tasks")
	}

	if opts == nil {
		opts = &EditRecipeBundleOpts{}
	}

	recipesPy, err := findRecipesPy(ctx, opts.RepoDir)
	if err != nil {
		return err
	}
	logging.Debugf(ctx, "using recipes.py: %q", recipesPy)

	extraProperties := make(map[string]string)
	setRecipeBundleProperty := opts.PropertyOnly || jd.GetBuildbucket().GetBbagentArgs().GetBuild().GetInput().GetProperties().GetFields()[LEDBuilderIsBootstrappedProperty].GetBoolValue()
	if setRecipeBundleProperty {
		// In property-only mode, we want to leave the original payload as is
		// and just upload the recipe bundle as a brand new independent CAS
		// archive for the job's executable to download.
		bundlePath, err := ioutil.TempDir("", "led-recipe-bundle")
		if err != nil {
			return errors.Annotate(err, "creating temporary recipe bundle directory").Err()
		}
		if err := opts.prepBundle(ctx, opts.RepoDir, recipesPy, bundlePath); err != nil {
			return err
		}
		logging.Infof(ctx, "isolating recipes")
		casClient, err := newCASClient(ctx, authOpts, jd)
		if err != nil {
			return err
		}
		casRef, err := uploadToCas(ctx, casClient, bundlePath)
		if err != nil {
			return err
		}
		m := &jsonpb.Marshaler{OrigName: true}
		jsonCASRef, err := m.MarshalToString(casRef)
		if err != nil {
			return errors.Annotate(err, "encoding CAS user payload").Err()
		}
		extraProperties[CASRecipeBundleProperty] = jsonCASRef
	} else {
		if err := EditIsolated(ctx, authOpts, jd, func(ctx context.Context, dir string) error {
			bundlePath := filepath.Join(dir, RecipeDirectory)
			// Remove existing bundled recipes, if any. Ignore the error.
			os.RemoveAll(bundlePath)
			if err := opts.prepBundle(ctx, opts.RepoDir, recipesPy, bundlePath); err != nil {
				return err
			}
			logging.Infof(ctx, "isolating recipes")
			return nil
		}); err != nil {
			return err
		}
	}

	return jd.HighLevelEdit(func(je job.HighLevelEditor) {
		if setRecipeBundleProperty {
			je.Properties(extraProperties, false)
		} else {
			je.TaskPayloadSource("", "")
			je.TaskPayloadPath(RecipeDirectory)
		}
		if opts.DebugSleep != 0 {
			je.Env(map[string]string{
				"RECIPES_DEBUG_SLEEP": fmt.Sprintf("%f", opts.DebugSleep.Seconds()),
			})
		}
	})
}

func logCmd(ctx context.Context, inDir string, arg0 string, args ...string) *exec.Cmd {
	ret := exec.CommandContext(ctx, arg0, args...)
	ret.Dir = inDir
	logging.Debugf(ctx, "Running (from %q) - %s %v", inDir, arg0, args)
	return ret
}

func cmdErr(cmd *exec.Cmd, err error, reason string) error {
	if err == nil {
		return nil
	}
	var outErr string
	if ee, ok := err.(*exec.ExitError); ok {
		outErr = strings.TrimSpace(string(ee.Stderr))
		if len(outErr) > 128 {
			outErr = outErr[:128] + "..."
		}
	} else {
		outErr = err.Error()
	}
	return errors.Annotate(err, "running %q: %s: %s", strings.Join(cmd.Args, " "), reason, outErr).Err()
}

func appendText(path, fmtStr string, items ...interface{}) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, fmtStr, items...)
	return err
}

func (opts *EditRecipeBundleOpts) prepBundle(ctx context.Context, inDir, recipesPy, toDirectory string) error {
	logging.Infof(ctx, "bundling recipes")
	args := []string{
		recipesPy,
	}
	if logging.GetLevel(ctx) < logging.Info {
		args = append(args, "-v")
	}
	for projID, path := range opts.Overrides {
		args = append(args, "-O", fmt.Sprintf("%s=%s", projID, path))
	}
	args = append(args, "bundle", "--destination", filepath.Join(toDirectory))

	// Always prefer python3 to python
	python, err := exec.LookPath("python3")
	if err != nil {
		python, err = exec.LookPath("python")
	}
	if err != nil {
		return errors.Annotate(err, "unable to find python3 or python in $PATH").Err()
	}

	cmd := logCmd(ctx, inDir, python, args...)
	if logging.GetLevel(ctx) < logging.Info {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmdErr(cmd, cmd.Run(), "creating bundle")
}

// findRecipesPy locates the current repo's `recipes.py`. It does this by:
//   * invoking git to find the repo root
//   * loading the recipes.cfg at infra/config/recipes.cfg
//   * stat'ing the recipes.py implied by the recipes_path in that cfg file.
//
// Failure will return an error.
//
// On success, the absolute path to recipes.py is returned.
func findRecipesPy(ctx context.Context, inDir string) (string, error) {
	cmd := logCmd(ctx, inDir, "git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err = cmdErr(cmd, err, "finding git repo"); err != nil {
		return "", err
	}

	repoRoot := strings.TrimSpace(string(out))

	pth := filepath.Join(repoRoot, "infra", "config", "recipes.cfg")
	switch st, err := os.Stat(pth); {
	case err != nil:
		return "", errors.Annotate(err, "reading recipes.cfg").Err()

	case !st.Mode().IsRegular():
		return "", errors.Reason("%q is not a regular file", pth).Err()
	}

	type recipesJSON struct {
		RecipesPath string `json:"recipes_path"`
	}
	rj := &recipesJSON{}

	f, err := os.Open(pth)
	if err != nil {
		return "", errors.Reason("reading recipes.cfg: %q", pth).Err()
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(rj); err != nil {
		return "", errors.Reason("parsing recipes.cfg: %q", pth).Err()
	}

	return filepath.Join(
		repoRoot, filepath.FromSlash(rj.RecipesPath), "recipes.py"), nil
}
