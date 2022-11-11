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

// Executable secret-tool allows to generate and rotate secrets stored in
// Google Secret Manager and consumed by go.chromium.org/luci/server/secrets
// module.
//
// Is supports generation and manipulation of secrets that are:
//   - Randomly generated byte blobs.
//   - Password-like strings passed via terminal.
//   - Tink key sets serialized as JSON.
//
// By default it doesn't access secrets once they are stored. The set of active
// secrets is represented by individual GSM SecretVersion objects with aliases
// "current", "previous" and "next" pointing to them. The tool knows how to move
// these aliases to perform somewhat graceful rotations. When using Tink keys,
// the final key set used at runtime is assembled dynamically from keys stored
// in "current", "previous" and "next" SecretVersions.
//
// To generate a new secret, run e.g.
//
//	secret-tool create sm://<project>/root-secret -secret-type random-bytes-32
//	secret-tool create sm://<project>/tink-aead-primary -secret-type tink-aes256-gcm
//
// To rotate an existing secret (regardless of its type):
//
//	secret-tool rotation-begin sm://<project>/<name>
//	# wait several hours to make sure the new secret is cached everywhere
//	# confirm by looking at /chrome/infra/secrets/gsm/version metric
//	secret-tool rotation-end sm://<project>/<name>
package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"syscall"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/insecurecleartextkeyset"
	"github.com/google/tink/go/keyset"
	tinkpb "github.com/google/tink/go/proto/tink_go_proto"
	"github.com/maruel/subcommands"
	"golang.org/x/term"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/data/text"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/flag/fixflagpos"
	"go.chromium.org/luci/common/flag/flagenum"
	"go.chromium.org/luci/common/flag/stringmapflag"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"
	"go.chromium.org/luci/hardcoded/chromeinfra"
)

// TODO(vadimsh): Add a flag or something to instruct `rotation-begin` that it
// should read Tink keys the previous `current` value and append them to the
// new keyset. That way we can use old Tink keys for arbitrary long time, while
// still reading only 3 secrets at runtime. In the current implementation, a
// Tink key is completely forgotten after two rotations. If that's OK or not
// depends on how the key is used (and thus this behavior should be controllable
// by a flag or some kind of annotation).

////////////////////////////////////////////////////////////////////////////////
// CLI boilerplate.

var userError = errors.BoolTag{Key: errors.NewTagKey("user error")}

var authOpts = chromeinfra.SetDefaultAuthOptions(auth.Options{
	Scopes: []string{
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/userinfo.email",
	},
})

type flagState string

func (s flagState) shouldRegister() bool {
	return s != "" && s != "disable"
}

var (
	disableFlag  flagState = "disable"
	requireFlag  flagState = "require"
	optionalFlag flagState = "optional"
)

func main() {
	os.Exit(subcommands.Run(&cli.Application{
		Name:  "secret-tool",
		Title: "Tool for creating and rotating secrets stored in Google Secret Manager and used by LUCI servers.",
		Context: func(ctx context.Context) context.Context {
			return (&gologger.LoggerConfig{
				Out:    os.Stderr,
				Format: `%{color}%{message}%{color:reset}`,
			}).Use(ctx)
		},
		Commands: []*subcommands.Command{
			subcommands.CmdHelp,

			authcli.SubcommandLogin(authOpts, "login", false),
			authcli.SubcommandLogout(authOpts, "logout", false),

			{
				UsageLine: "create sm://<project>/<name> -secret-type <type>",
				ShortDesc: "creates a new secret",
				LongDesc: text.Doc(fmt.Sprintf(`
					Creates a new secret populating its value based on <type>.

					Supported types:
%s

					All Tink keysets are stored as clear text JSONPB.
				`, generatorsHelp("					  * "))),
				CommandRun: func() subcommands.CommandRun {
					return initCommand(&commandRun{
						exec:           (*commandRun).cmdCreate,
						secretTypeFlag: requireFlag,
						aliasesFlag:    disableFlag,
						forceFlag:      optionalFlag,
					})
				},
			},

			{
				UsageLine: "inspect sm://<project>/<name>",
				ShortDesc: "shows the current state of a secret",
				LongDesc: text.Doc(`
					Shows the current state of a secret, in particular values of aliases
					denoting the current, previous and next versions of the secret.
				`),
				CommandRun: func() subcommands.CommandRun {
					return initCommand(&commandRun{
						exec:           (*commandRun).cmdInspect,
						secretTypeFlag: disableFlag,
						aliasesFlag:    disableFlag,
						forceFlag:      disableFlag,
					})
				},
			},

			{
				UsageLine: "set-aliases sm://<project>/<name>",
				ShortDesc: "moves version aliases on the secret",
				LongDesc: text.Doc(`
					Moves the version aliases. This should rarely be used directly, only
					as a way to immediately return to some previous state.

					For rotations, use rotation-begin and rotation-end subcommands which
					move version aliases as well.

					Aliases not mentioned in the flags are left untouched. To delete an
					alias, use "-alias <name>=0".
				`),
				CommandRun: func() subcommands.CommandRun {
					return initCommand(&commandRun{
						exec:           (*commandRun).cmdSetAliases,
						secretTypeFlag: disableFlag,
						aliasesFlag:    requireFlag,
						forceFlag:      disableFlag,
					})
				},
			},

			{
				UsageLine: "rotation-begin sm://<project>/<name>",
				ShortDesc: "generates a new version of the secret and designates it as next",
				LongDesc: text.Doc(`
					This starts the secret rotation process by generating a new version
					of the secret and moving "next" alias to point to it (keeping all
					other aliases intact). This allows the processes that use the secret
					to precache the new version, before it is actually used.

					To finish the rotation, call rotation-end at some later time when
					all processes picked up the new secret. How long it takes depends
					on the service configuration and can measure in hours.
				`),
				CommandRun: func() subcommands.CommandRun {
					return initCommand(&commandRun{
						exec:           (*commandRun).cmdRotationBegin,
						secretTypeFlag: optionalFlag,
						aliasesFlag:    disableFlag,
						forceFlag:      optionalFlag,
					})
				},
			},

			{
				UsageLine: "rotation-end sm://<project>/<name>",
				ShortDesc: "finishes the rotation started with rotation-begin",
				LongDesc: text.Doc(`
					This finishes the rotation started with rotation-begin by
					updating aliases as:
						previous := current
						current := next

					This should be done once all processes cached the "next" version
					of the secret. After this command finishes, this version will be used
					as "current" (i.e. used for encryption, signing, etc).

					Note that this completely evicts the old "previous" value (which is
					a leftover from the previous rotation). Be careful when doing
					rotations back to back.
				`),
				CommandRun: func() subcommands.CommandRun {
					return initCommand(&commandRun{
						exec:           (*commandRun).cmdRotationEnd,
						secretTypeFlag: disableFlag,
						aliasesFlag:    disableFlag,
						forceFlag:      disableFlag,
					})
				},
			},
		},
	}, fixflagpos.FixSubcommands(os.Args[1:])))
}

type commandRun struct {
	subcommands.CommandRunBase
	authFlags authcli.Flags

	secretTypeFlag flagState // controls presence of -secret-type flag
	aliasesFlag    flagState // controls presence of -alias flag
	forceFlag      flagState // controls presence of -force flag

	gsm        *secretmanager.Client // GSM client
	project    string                // GCP project with the secret
	secret     string                // name of the secret
	secretRef  string                // full name of the secret for GSM
	secretGen  secretGenerator       // parsed -secret-type or nil if wasn't set
	aliasesRaw stringmapflag.Value   // raw collected -alias flag values
	aliases    map[string]int64      // parsed -alias flags
	force      bool                  // parsed -force flag

	exec func(*commandRun, context.Context) error // method to call to execute the command
}

func initCommand(c *commandRun) *commandRun {
	c.authFlags.Register(&c.Flags, authOpts)
	if c.secretTypeFlag.shouldRegister() {
		c.Flags.Var(&c.secretGen, "secret-type", "What kind of secret value to generate.")
	}
	if c.aliasesFlag.shouldRegister() {
		c.Flags.Var(&c.aliasesRaw, "alias", "A name=version pair indicating an alias.")
	}
	if c.forceFlag.shouldRegister() {
		c.Flags.BoolVar(&c.force, "force", false, "Ignore safeguards and apply the change.")
	}
	return c
}

func (c *commandRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	ctx := cli.GetContext(a, c, env)

	if len(args) != 1 {
		logging.Errorf(ctx, "Expecting exactly one positional argument: the secret path as sm://<project>/<name>.")
		return 1
	}
	secretPath := args[0]

	// Parse the secret reference into its components.
	if !strings.HasPrefix(secretPath, "sm://") {
		logging.Errorf(ctx, "Only sm:// secrets are supported.")
		return 1
	}
	parts := strings.Split(strings.TrimPrefix(secretPath, "sm://"), "/")
	if len(parts) != 2 {
		logging.Errorf(ctx, "Expecting full secret reference as sm://<project>/<name>.")
		return 1
	}
	c.project, c.secret = parts[0], parts[1]
	c.secretRef = fmt.Sprintf("projects/%s/secrets/%s", c.project, c.secret)

	// Check flags.
	if c.secretTypeFlag == requireFlag && c.secretGen.name == "" {
		logging.Errorf(ctx, "Missing required flag -secret-type.")
		return 1
	}
	if c.aliasesFlag.shouldRegister() {
		c.aliases = make(map[string]int64, len(c.aliasesRaw))
		for alias, ver := range c.aliasesRaw {
			verInt, err := strconv.ParseInt(ver, 10, 64)
			if err != nil {
				logging.Errorf(ctx, "Bad -alias flag %s=%s: the version is not an integer.", alias, ver)
				return 1
			}
			c.aliases[alias] = verInt
		}
		if len(c.aliases) == 0 && c.aliasesFlag == requireFlag {
			logging.Errorf(ctx, "At least one -alias <name>=<version> flag is required.")
			return 1
		}
	}

	// Setup the GSM client.
	authOpts, err := c.authFlags.Options()
	if err != nil {
		logging.Errorf(ctx, "Bad auth options: %s.", err)
		return 1
	}
	switch ts, err := auth.NewAuthenticator(ctx, auth.SilentLogin, authOpts).TokenSource(); {
	case err == auth.ErrLoginRequired:
		logging.Errorf(ctx, "Need to login first. Run `auth-login` subcommand.")
		return 1
	case err != nil:
		errors.Log(ctx, err)
		return 1
	default:
		c.gsm, err = secretmanager.NewClient(
			ctx,
			option.WithTokenSource(ts),
		)
		if err != nil {
			errors.Log(ctx, err)
			return 1
		}
	}

	if err = c.exec(c, ctx); err != nil {
		if userError.In(err) {
			logging.Errorf(ctx, "%s", err)
		} else {
			errors.Log(ctx, err)
		}
		return 1
	}
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// Helpers that work with the secret selected in `c.secretRef`.

const secretTypeLabel = "luci-secret"

// parseVersion extracts int64 version from SecretVersion name.
func parseVersion(versionName string) (int64, error) {
	idx := strings.LastIndex(versionName, "/")
	if idx == -1 {
		return 0, errors.Reason("unexpected version name format %q", versionName).Err()
	}
	ver, err := strconv.ParseInt(versionName[idx+1:], 10, 64)
	if err != nil {
		return 0, errors.Reason("unexpected version name format %q", versionName).Err()
	}
	if ver == 0 {
		return 0, errors.Reason("the version is unexpectedly 0").Err()
	}
	return ver, nil
}

// secretMetadata fetches metadata about the secret.
func (c *commandRun) secretMetadata(ctx context.Context) (*secretmanagerpb.Secret, error) {
	secret, err := c.gsm.GetSecret(ctx, &secretmanagerpb.GetSecretRequest{
		Name: c.secretRef,
	})
	if err != nil {
		return nil, errors.Annotate(err, "failed to fetch the secret metadata").Err()
	}
	return secret, nil
}

// generateNewVersion generates new secret blob, adds it as SecretVersion, and
// returns its version number.
func (c *commandRun) generateNewVersion(ctx context.Context) (int64, error) {
	logging.Infof(ctx, "Creating and storing the secret of type %q...", c.secretGen.name)
	secretBlob, err := c.secretGen.gen(ctx)
	if err != nil {
		return 0, errors.Annotate(err, "failed to generate the secret of type %q", c.secretGen.name).Err()
	}
	version, err := c.gsm.AddSecretVersion(ctx, &secretmanagerpb.AddSecretVersionRequest{
		Parent: c.secretRef,
		Payload: &secretmanagerpb.SecretPayload{
			Data: secretBlob,
		},
	})
	if err != nil {
		return 0, errors.Annotate(err, "failed to add the new secret version").Err()
	}
	ver, err := parseVersion(version.Name)
	if err != nil {
		return 0, err
	}
	logging.Infof(ctx, "Created the new secret version %d.", ver)
	return ver, nil
}

// latestVersion resolves "latest" into an int64 version.
func (c *commandRun) latestVersion(ctx context.Context) (int64, error) {
	version, err := c.gsm.GetSecretVersion(ctx, &secretmanagerpb.GetSecretVersionRequest{
		Name: fmt.Sprintf("%s/versions/latest", c.secretRef),
	})
	if err != nil {
		return 0, errors.Annotate(err, "failed to resolve the latest version").Err()
	}
	return parseVersion(version.Name)
}

// overrideAliases overrides *all* aliases.
func (c *commandRun) overrideAliases(ctx context.Context, etag string, aliases map[string]int64) (*secretmanagerpb.Secret, error) {
	logging.Infof(ctx, "Updating aliases...")
	secret, err := c.gsm.UpdateSecret(ctx, &secretmanagerpb.UpdateSecretRequest{
		Secret: &secretmanagerpb.Secret{
			Name:           c.secretRef,
			Etag:           etag,
			VersionAliases: aliases,
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"version_aliases"},
		},
	})
	if err != nil {
		return nil, errors.Annotate(err, "failed to set version aliases to %v", aliases).Err()
	}
	return secret, nil
}

// overrideLabels overrides *all* labels.
func (c *commandRun) overrideLabels(ctx context.Context, etag string, labels map[string]string) (*secretmanagerpb.Secret, error) {
	logging.Infof(ctx, "Updating labels...")
	secret, err := c.gsm.UpdateSecret(ctx, &secretmanagerpb.UpdateSecretRequest{
		Secret: &secretmanagerpb.Secret{
			Name:   c.secretRef,
			Etag:   etag,
			Labels: labels,
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"labels"},
		},
	})
	if err != nil {
		return nil, errors.Annotate(err, "failed to set labels to %v", labels).Err()
	}
	return secret, nil
}

// printAliasMap prints information about current aliases.
func (c *commandRun) printAliasMap(ctx context.Context, title string, secret *secretmanagerpb.Secret, showSwitchCmd bool) {
	aliases := []string{"current", "previous", "next"}

	logging.Infof(ctx, "%s:", title)
	for _, alias := range aliases {
		if ver := secret.VersionAliases[alias]; ver != 0 {
			logging.Infof(ctx, "  %s = %d", alias, secret.VersionAliases[alias])
		}
	}
	if len(secret.VersionAliases) == 0 {
		logging.Infof(ctx, "  <no aliases>")
	}

	// Generate a command to help to hop into this state.
	if showSwitchCmd {
		switchCmd := []string{
			"secret-tool",
			"set-aliases",
			fmt.Sprintf("sm://%s/%s", c.project, c.secret),
		}
		for _, alias := range aliases {
			switchCmd = append(switchCmd, "-alias", fmt.Sprintf("%s=%d", alias, secret.VersionAliases[alias]))
		}
		logging.Infof(ctx, "Command to immediately switch to this state if necessary:")
		logging.Infof(ctx, "  $ %s", strings.Join(switchCmd, " "))
	}
}

// printSecretMetadata prints some information about the secret to stdout.
func (c *commandRun) printSecretMetadata(ctx context.Context, secret *secretmanagerpb.Secret) {
	secretType := "<unknown>"
	if typ := secret.Labels[secretTypeLabel]; typ != "" {
		secretType = typ
	}
	logging.Infof(ctx, "Secret type: %s", secretType)
	c.printAliasMap(ctx, "Aliases", secret, true)
}

////////////////////////////////////////////////////////////////////////////////
// "create" implementation.

func (c *commandRun) cmdCreate(ctx context.Context) error {
	logging.Infof(ctx, "Creating the secret...")
	secret, err := c.gsm.CreateSecret(ctx, &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", c.project),
		SecretId: c.secret,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{},
			},
			Labels: map[string]string{
				secretTypeLabel: c.secretGen.name,
			},
		},
	})

	if status.Code(err) == codes.AlreadyExists {
		// Check if it is just an empty container that doesn't have any versions.
		// This is an allowed use case. The container may be create by Terraform.
		iter := c.gsm.ListSecretVersions(ctx, &secretmanagerpb.ListSecretVersionsRequest{
			Parent:   c.secretRef,
			PageSize: 1,
		})
		switch _, err := iter.Next(); {
		case err == iterator.Done:
			logging.Infof(ctx, "The secret already exists and has no versions, proceeding...")
		case err == nil:
			return errors.New(
				"This secret already exists and has versions. "+
					"If you want to rotate it use rotation-begin and rotation-end subcommands.", userError)
		default:
			return errors.Annotate(err, "failed to check if the secret has any versions").Err()
		}
		// Verify the type label is set, update if not.
		secret, err = c.secretMetadata(ctx)
		if err != nil {
			return err
		}
		existingType := secret.Labels[secretTypeLabel]
		if existingType != c.secretGen.name {
			if existingType != "" && !c.force {
				return errors.Reason(
					"The secret already exists and its type is set to %q (not %q, as requested). "+
						"Pass -force to override the type.", existingType, c.secretGen.name).Tag(userError).Err()
			}
			if existingType != "" {
				logging.Warningf(ctx, "Overriding the secret type %q => %q.", existingType, c.secretGen.name)
			}
			if secret.Labels == nil {
				secret.Labels = map[string]string{}
			}
			secret.Labels[secretTypeLabel] = c.secretGen.name
			secret, err = c.overrideLabels(ctx, secret.Etag, secret.Labels)
			if err != nil {
				return err
			}
		}
	} else if err != nil {
		return errors.Annotate(err, "failed to create the secret").Err()
	}

	added, err := c.generateNewVersion(ctx)
	if err != nil {
		return err
	}

	secret, err = c.overrideAliases(ctx, secret.Etag, map[string]int64{
		"current":  added,
		"previous": added,
		"next":     added,
	})
	if err != nil {
		return err
	}
	c.printSecretMetadata(ctx, secret)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// "inspect" implementation.

func (c *commandRun) cmdInspect(ctx context.Context) error {
	secret, err := c.secretMetadata(ctx)
	if err != nil {
		return err
	}
	c.printSecretMetadata(ctx, secret)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// "set-aliases" implementation.

func (c *commandRun) cmdSetAliases(ctx context.Context) error {
	secret, err := c.secretMetadata(ctx)
	if err != nil {
		return err
	}

	c.printAliasMap(ctx, "Current aliases", secret, true)

	for k, v := range c.aliases {
		if v == 0 {
			delete(secret.VersionAliases, k)
		} else {
			secret.VersionAliases[k] = v
		}
	}

	secret, err = c.overrideAliases(ctx, secret.Etag, secret.VersionAliases)
	if err != nil {
		return err
	}

	c.printAliasMap(ctx, "Updated aliases", secret, false)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// "rotation-begin" implementation.

func (c *commandRun) cmdRotationBegin(ctx context.Context) error {
	secret, err := c.secretMetadata(ctx)
	if err != nil {
		return err
	}

	// Figure out how to generate the new secret, populate c.secretGen.
	existingType := secret.Labels[secretTypeLabel]
	if existingType == "" {
		if c.secretGen.name == "" {
			return errors.New("This secret is not annotated with a type, pass -secret-type explicitly.", userError)
		}
		existingType = c.secretGen.name
	}
	if c.secretGen.name == "" {
		typ, ok := secretTypes[existingType]
		if !ok {
			return errors.Reason(
				"The secret is annotated with unrecognized type %q. "+
					"You may need to pass -secret-type and -force flags to override, but be careful.",
				existingType).Tag(userError).Err()
		}
		c.secretGen = typ.(secretGenerator)
	}
	if c.secretGen.name != existingType {
		if !c.force && !c.secretGen.compatible.Has(existingType) {
			return errors.Reason(
				"Can't change the secret type from %q to %q. Types are incompatible. "+
					"If you really need this change, pass -force flag. This is dangerous.",
				existingType, c.secretGen.name,
			).Tag(userError).Err()
		}
		logging.Warningf(ctx, "Overriding the secret type %q => %q.", existingType, c.secretGen.name)
		labels := secret.Labels
		if labels == nil {
			labels = map[string]string{}
		}
		labels[secretTypeLabel] = c.secretGen.name
		if secret, err = c.overrideLabels(ctx, secret.Etag, labels); err != nil {
			return err
		}
	}

	// Legacy secrets use "latest" as "current", and it is a magical alias that
	// needs to be resolved via an RPC.
	current := secret.VersionAliases["current"]
	if current == 0 {
		current, err = c.latestVersion(ctx)
		if err != nil {
			return err
		}
		logging.Infof(ctx, "This looks like a legacy secret without \"current\" alias.")
		logging.Infof(ctx, "The current value of \"latest\" (%d) will be set as \"current\".", current)
	}

	// Abort if already rotating.
	if next := secret.VersionAliases["next"]; next != 0 && next != current {
		c.printAliasMap(ctx, "Current aliases", secret, false)
		return errors.New("Looks like a rotation is already in progress.", userError)
	}

	c.printAliasMap(ctx, "Aliases prior to the starting rotation", secret, true)

	// Create the new version.
	next, err := c.generateNewVersion(ctx)
	if err != nil {
		return err
	}

	// Update the alias map. Don't touch existing aliases, including "previous".
	secret.VersionAliases["current"] = current
	secret.VersionAliases["next"] = next
	secret, err = c.overrideAliases(ctx, secret.Etag, secret.VersionAliases)
	if err != nil {
		return err
	}

	c.printAliasMap(ctx, "Aliases now", secret, false)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// "rotation-end" implementation.

func (c *commandRun) cmdRotationEnd(ctx context.Context) error {
	secret, err := c.secretMetadata(ctx)
	if err != nil {
		return err
	}

	c.printAliasMap(ctx, "Current aliases", secret, true)

	current := secret.VersionAliases["current"]
	next := secret.VersionAliases["next"]
	if current == 0 || next == 0 || current == next {
		return errors.New("There's no rotation in progress.", userError)
	}

	secret.VersionAliases["previous"] = current
	secret.VersionAliases["current"] = next

	secret, err = c.overrideAliases(ctx, secret.Etag, secret.VersionAliases)
	if err != nil {
		return err
	}

	c.printAliasMap(ctx, "Updated aliases", secret, false)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Secret generators registry.

type secretGenerator struct {
	name       string
	help       string
	compatible stringset.Set // types that can upgraded from
	gen        func(context.Context) ([]byte, error)
}

var secretTypes = flagenum.Enum{
	// populated in init()
}

func (gen *secretGenerator) Set(v string) error {
	return secretTypes.FlagSet(gen, v)
}

func (gen *secretGenerator) String() string {
	return gen.name
}

func registerGenerator(name, help string, compatible []string, gen func(context.Context) ([]byte, error)) {
	compatibleTypes := stringset.NewFromSlice(compatible...)
	compatibleTypes.Add(name)
	secretTypes[name] = secretGenerator{
		name:       name,
		help:       help,
		compatible: compatibleTypes,
		gen:        gen,
	}
}

func generatorsHelp(padding string) string {
	lines := make([]string, 0, len(secretTypes))
	for _, gen := range secretTypes {
		gen := gen.(secretGenerator)
		lines = append(lines, fmt.Sprintf("%s%s: %s", padding, gen.name, gen.help))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}

func generateTinkKey(template *tinkpb.KeyTemplate) ([]byte, error) {
	kh, err := keyset.NewHandle(template)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	if err = insecurecleartextkeyset.Write(kh, keyset.NewJSONWriter(buf)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

////////////////////////////////////////////////////////////////////////////////
// Supported secret generators.

func init() {
	registerGenerator(
		"random-bytes-32",
		"a random 32 byte blob",
		nil,
		func(context.Context) ([]byte, error) {
			blob := make([]byte, 32)
			_, err := rand.Read(blob)
			return blob, err
		},
	)

	registerGenerator(
		"password",
		"read a secret from the terminal as a password",
		nil,
		func(ctx context.Context) ([]byte, error) {
			fmt.Printf("Type the secret value and hit Enter: ")
			return term.ReadPassword(int(syscall.Stdin))
		},
	)

	registerGenerator(
		"tink-aes256-gcm",
		"a generated Tink keyset with AES256 GCM key used for AEAD",
		nil,
		func(ctx context.Context) ([]byte, error) {
			return generateTinkKey(aead.AES256GCMKeyTemplate())
		},
	)
}
