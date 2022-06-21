// Copyright 2015 The LUCI Authors.
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

// Package cipd implements client side of Chrome Infra Package Deployer.
//
// Binary package file format (in free form representation):
//   <binary package> := <zipped data>
//   <zipped data> := DeterministicZip(<all input files> + <manifest json>)
//   <manifest json> := File{
//     name: ".cipdpkg/manifest.json",
//     data: JSON({
//       "FormatVersion": "1",
//       "PackageName": <name of the package>
//     }),
//   }
//   DeterministicZip = zip archive with deterministic ordering of files and stripped timestamps
//
// Main package data (<zipped data> above) is deterministic, meaning its content
// depends only on inputs used to built it (byte to byte): contents and names of
// all files added to the package (plus 'executable' file mode bit) and
// a package name (and all other data in the manifest).
//
// Binary package data MUST NOT depend on a timestamp, hostname of machine that
// built it, revision of the source code it was built from, etc. All that
// information will be distributed as a separate metadata packet associated with
// the package when it gets uploaded to the server.
//
// TODO: expand more when there's server-side package data model (labels
// and stuff).
package cipd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/common/system/environ"
	"go.chromium.org/luci/grpc/prpc"

	api "go.chromium.org/luci/cipd/api/cipd/v1"
	"go.chromium.org/luci/cipd/client/cipd/deployer"
	"go.chromium.org/luci/cipd/client/cipd/digests"
	"go.chromium.org/luci/cipd/client/cipd/ensure"
	"go.chromium.org/luci/cipd/client/cipd/fs"
	"go.chromium.org/luci/cipd/client/cipd/internal"
	"go.chromium.org/luci/cipd/client/cipd/pkg"
	"go.chromium.org/luci/cipd/client/cipd/platform"
	"go.chromium.org/luci/cipd/client/cipd/plugin"
	"go.chromium.org/luci/cipd/client/cipd/reader"
	"go.chromium.org/luci/cipd/client/cipd/template"
	"go.chromium.org/luci/cipd/client/cipd/ui"
	"go.chromium.org/luci/cipd/common"
	"go.chromium.org/luci/cipd/version"
)

const (
	// CASFinalizationTimeout is how long to wait for CAS service to finalize
	// the upload in RegisterInstance.
	CASFinalizationTimeout = 5 * time.Minute

	// SetRefTimeout is how long to wait for an instance to be processed when
	// setting a ref in SetRefWhenReady.
	SetRefTimeout = 3 * time.Minute

	// TagAttachTimeout is how long to wait for an instance to be processed when
	// attaching tags in AttachTagsWhenReady.
	TagAttachTimeout = 3 * time.Minute

	// MetadataAttachTimeout is how long to wait for an instance to be processed
	// when attaching metadata in AttachMetadataWhenReady.
	MetadataAttachTimeout = 3 * time.Minute

	// DefaultParallelDownloads is a default value for CIPD_PARALLEL_DOWNLOADS.
	//
	// See ParallelDownloads client option.
	DefaultParallelDownloads = 4
)

// Environment variable definitions
const (
	EnvCacheDir            = "CIPD_CACHE_DIR"
	EnvHTTPUserAgentPrefix = "CIPD_HTTP_USER_AGENT_PREFIX"
	EnvMaxThreads          = "CIPD_MAX_THREADS"
	EnvParallelDownloads   = "CIPD_PARALLEL_DOWNLOADS"
	EnvAdmissionPlugin     = "CIPD_ADMISSION_PLUGIN"
	EnvCIPDServiceURL      = "CIPD_SERVICE_URL"
)

var (
	// ErrFinalizationTimeout is returned if CAS service can not finalize upload
	// fast enough.
	ErrFinalizationTimeout = errors.New("timeout while waiting for CAS service to finalize the upload", transient.Tag)

	// ErrBadUpload is returned when a package file is uploaded, but servers asks
	// us to upload it again.
	ErrBadUpload = errors.New("package file is uploaded, but servers asks us to upload it again", transient.Tag)

	// ErrProcessingTimeout is returned by SetRefWhenReady or AttachTagsWhenReady
	// if the instance processing on the backend takes longer than expected. Refs
	// and tags can be attached only to processed instances.
	ErrProcessingTimeout = errors.New("timeout while waiting for the instance to become ready", transient.Tag)

	// ErrDownloadError is returned by FetchInstance on download errors.
	ErrDownloadError = errors.New("failed to download the package file after multiple attempts", transient.Tag)

	// ErrUploadError is returned by RegisterInstance on upload errors.
	ErrUploadError = errors.New("failed to upload the package file after multiple attempts", transient.Tag)

	// ErrEnsurePackagesFailed is returned by EnsurePackages if something is not
	// right.
	ErrEnsurePackagesFailed = errors.New("failed to update the deployment")
)

var (
	// ClientPackage is a package with the CIPD client. Used during self-update.
	ClientPackage = "infra/tools/cipd/${platform}"
	// UserAgent is HTTP user agent string for CIPD client.
	UserAgent = "cipd 2.6.5"
)

func init() {
	ver, err := version.GetStartupVersion()
	if err != nil || ver.InstanceID == "" {
		return
	}
	UserAgent += fmt.Sprintf(" (%s@%s)", ver.PackageName, ver.InstanceID)
}

// UploadSession describes open CAS upload session.
type UploadSession struct {
	// ID identifies upload session in the backend.
	ID string
	// URL is where to upload the data to.
	URL string
}

// DescribeInstanceOpts is passed to DescribeInstance.
type DescribeInstanceOpts struct {
	DescribeRefs bool // if true, will fetch all refs pointing to the instance
	DescribeTags bool // if true, will fetch all tags attached to the instance
}

// Client provides high-level CIPD client interface. Thread safe.
type Client interface {
	// Close terminates plugins started by the client.
	//
	// Should be used to cleanly shutdown the client when using plugins.
	Close(ctx context.Context)

	// BeginBatch makes the client enter into a "batch mode".
	//
	// In this mode various cleanup and cache updates, usually performed right
	// away, are deferred until 'EndBatch' call.
	//
	// This is an optimization. Use it if you plan to call a bunch of Client
	// methods in a short amount of time (parallel or sequentially).
	//
	// Batches can be nested.
	BeginBatch(ctx context.Context)

	// EndBatch ends a batch started with BeginBatch.
	//
	// EndBatch does various delayed maintenance tasks (like cache updates, trash
	// cleanup and so on). This is best-effort operations, and thus this method
	// doesn't return an errors.
	//
	// See also BeginBatch doc for more details.
	EndBatch(ctx context.Context)

	// FetchACL returns a list of PackageACL objects (parent paths first).
	//
	// Together they define the access control list for the given package prefix.
	FetchACL(ctx context.Context, prefix string) ([]PackageACL, error)

	// ModifyACL applies a set of PackageACLChanges to a package prefix ACL.
	ModifyACL(ctx context.Context, prefix string, changes []PackageACLChange) error

	// FetchRoles returns all roles the caller has in the given package prefix.
	//
	// Understands roles inheritance, e.g. if the caller is OWNER, the return
	// value will list all roles implied by being an OWNER (e.g. READER, WRITER,
	// ...).
	FetchRoles(ctx context.Context, prefix string) ([]string, error)

	// ResolveVersion converts an instance ID, a tag or a ref into a concrete Pin.
	ResolveVersion(ctx context.Context, packageName, version string) (common.Pin, error)

	// RegisterInstance makes the package instance available for clients.
	//
	// It uploads the instance to the storage, waits until the storage verifies
	// its hash matches instance ID in 'pin', and then registers the package in
	// the repository, making it discoverable.
	//
	// 'pin' here should match the package body, otherwise CIPD backend will
	// reject the package. Either get it from builder.BuildInstance (when building
	// a new package) or from reader.CalculatePin (when uploading an existing
	// package file).
	//
	// 'timeout' specifies for how long to wait until the instance hash is
	// verified by the storage backend. If 0, default CASFinalizationTimeout will
	// be used.
	RegisterInstance(ctx context.Context, pin common.Pin, src pkg.Source, timeout time.Duration) error

	// DescribeInstance returns information about a package instance.
	//
	// May also be used as a simple instance presence check, if opts is nil. If
	// the request succeeds, then the instance exists.
	DescribeInstance(ctx context.Context, pin common.Pin, opts *DescribeInstanceOpts) (*InstanceDescription, error)

	// DescribeClient returns information about a CIPD client binary matching the
	// given client package pin.
	DescribeClient(ctx context.Context, pin common.Pin) (*ClientDescription, error)

	// SetRefWhenReady moves a ref to point to a package instance.
	SetRefWhenReady(ctx context.Context, ref string, pin common.Pin) error

	// AttachTagsWhenReady attaches tags to an instance.
	AttachTagsWhenReady(ctx context.Context, pin common.Pin, tags []string) error

	// AttachMetadataWhenReady attaches metadata to an instance.
	AttachMetadataWhenReady(ctx context.Context, pin common.Pin, md []Metadata) error

	// FetchPackageRefs returns information about all refs defined for a package.
	//
	// The returned list is sorted by modification timestamp (newest first).
	FetchPackageRefs(ctx context.Context, packageName string) ([]RefInfo, error)

	// FetchInstance downloads a package instance file from the repository.
	//
	// It verifies that the package hash matches pin.InstanceID.
	//
	// It returns a pkg.Source pointing to the raw package data. The caller must
	// close it when done.
	FetchInstance(ctx context.Context, pin common.Pin) (pkg.Source, error)

	// FetchInstanceTo downloads a package instance file into the given writer.
	//
	// This is roughly the same as getting a reader with 'FetchInstance' and
	// copying its data into the writer, except this call skips unnecessary temp
	// files if the client is not using cache.
	//
	// It verifies that the package hash matches pin.InstanceID, but does it while
	// writing to 'output', so expect to discard all data there if FetchInstanceTo
	// returns an error.
	FetchInstanceTo(ctx context.Context, pin common.Pin, output io.WriteSeeker) error

	// ListPackages returns a list packages and prefixes under the given prefix.
	ListPackages(ctx context.Context, prefix string, recursive, includeHidden bool) ([]string, error)

	// SearchInstances finds instances of some package with all given tags.
	//
	// Returns their concrete Pins. If the package doesn't exist at all, returns
	// empty slice and nil error.
	SearchInstances(ctx context.Context, packageName string, tags []string) (common.PinSlice, error)

	// ListInstances enumerates instances of a package, most recent first.
	//
	// Returns an object that can be used to fetch the listing, page by page.
	ListInstances(ctx context.Context, packageName string) (InstanceEnumerator, error)

	// FindDeployed returns a list of packages deployed to the site root.
	//
	// It just does a shallow examination of the metadata directory, without
	// paranoid checks that all installed packages are free from corruption.
	FindDeployed(ctx context.Context) (common.PinSliceBySubdir, error)

	// EnsurePackages installs, removes and updates packages in the site root.
	//
	// Given a description of what packages (and versions) should be installed it
	// will do all necessary actions to bring the state of the site root to the
	// desired one.
	//
	// If the update was only partially applied, returns both Actions and error.
	//
	// NOTE: You can repair the current deployment by passing in the output of
	// FindDeployed, plus an opts.Paranoia. Doing this with opts.DryRun=true
	// will check the integrity of the current deployment.
	EnsurePackages(ctx context.Context, pkgs common.PinSliceBySubdir, opts *EnsureOptions) (ActionMap, error)
}

// EnsureOptions is passed to Client.EnsurePackages.
//
// This can be used to control HOW Ensure does its actions.
type EnsureOptions struct {
	// Paranoia controls how much verification EnsurePackages does to verify that
	// the currently-installed packages are, in fact, correctly installed.
	//
	// If EnsurePackages determines that packages are NOT correctly installed, it
	// will attempt to fix them.
	//
	// Defaults to NotParanoid.
	// See the enum for more info.
	Paranoia ParanoidMode

	// DryRun, if true, will only check for changes and return them in the
	// ActionMap, but won't actually perform them.
	DryRun bool

	// Silent, if true, suppresses all logging output from EnsurePackages.
	Silent bool

	// OverrideInstallMode, if set, will be the install mode used for all
	// packages installed or repaired by this ensure operation. Note that
	// Paranoia needs to be at least CheckDeployed to change the install
	// mode of an already-deployed package.
	OverrideInstallMode pkg.InstallMode
}

// ClientOptions is passed to NewClient factory function.
//
// If you construct options manually, you almost certainly also need to call
// LoadFromEnv to load unset values from the process environment before passing
// options to NewClient.
type ClientOptions struct {
	// ServiceURL is root URL of the backend service.
	//
	// Default is ServiceURL const.
	ServiceURL string

	// Root is a site root directory.
	//
	// It is a directory where packages will be installed to. It also hosts
	// .cipd/* directory that tracks internal state of installed packages and
	// keeps various cache files. 'Root' can be an empty string if the client is
	// not going to be used to deploy or remove local packages.
	Root string

	// CacheDir is a directory for shared cache.
	//
	// If empty, instances are not cached and tags are cached inside the site
	// root. If both Root and CacheDir are empty, tag cache is disabled.
	CacheDir string

	// Versions is optional database of (pkg, version) => instance ID resolutions.
	//
	// If set, it will be used for all version resolutions done by the client.
	// The client won't be consulting (or updating) the tag cache and won't make
	// 'ResolveVersion' backend RPCs.
	//
	// This is primarily used to implement $ResolvedVersions ensure file feature.
	Versions ensure.VersionsFile

	// AnonymousClient is http.Client that doesn't attach authentication headers.
	//
	// Will be used when talking to the Google Storage. We use signed URLs that do
	// not require additional authentication.
	//
	// Default is http.DefaultClient.
	AnonymousClient *http.Client

	// AuthenticatedClient is http.Client that attaches authentication headers.
	//
	// Will be used when talking to the backend.
	//
	// Default is same as AnonymousClient (it will probably not work for most
	// packages, since the backend won't authorize an anonymous access).
	AuthenticatedClient *http.Client

	// MaxThreads defines how many threads to use when unzipping packages.
	//
	// If 0 or negative, will use all available CPUs.
	MaxThreads int

	// ParallelDownloads defines how many packages are allowed to be fetched
	// concurrently.
	//
	// Possible values:
	//   0: will use some default value (perhaps loading it from the environ).
	//  <0: will do fetching and unzipping completely serially.
	//   1: will fetch at most one package at once and unzip in parallel to that.
	//  >1: will fetch multiple packages at once and unzip in parallel to that.
	ParallelDownloads int

	// UserAgent is put into User-Agent HTTP header with each request.
	//
	// Default is UserAgent const.
	UserAgent string

	// LoginInstructions is appended to "permission denied" error messages.
	LoginInstructions string

	// PluginHost is a plugin system implementation to use.
	//
	// If not set, all plugin functionality will be ignored.
	PluginHost plugin.Host

	// AdmissionPlugin is the deployment admission plugin command line (if any).
	//
	// Will be started lazily when needed.
	AdmissionPlugin []string

	// Mocks used by tests.
	casMock     api.StorageClient
	repoMock    api.RepositoryClient
	storageMock storage
}

// LoadFromEnv loads supplied default values from an environment into opts.
//
// Uses the environment in the context via luci/common/system/environ library,
// falling back to the regular process environment as usual (so if you don't
// need to mess with the process environment, just pass any context, it would
// work fine).
func (opts *ClientOptions) LoadFromEnv(ctx context.Context) error {
	env := environ.FromCtx(ctx)

	if opts.CacheDir == "" {
		if v := env.Get(EnvCacheDir); v != "" {
			if !filepath.IsAbs(v) {
				return fmt.Errorf("bad %s %q: not an absolute path", EnvCacheDir, v)
			}
			opts.CacheDir = v
		}
	}
	if opts.MaxThreads == 0 {
		if v := env.Get(EnvMaxThreads); v != "" {
			maxThreads, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("bad %s %q: not an integer", EnvMaxThreads, v)
			}
			opts.MaxThreads = maxThreads
		}
	}
	if opts.ParallelDownloads == 0 {
		if v := env.Get(EnvParallelDownloads); v != "" {
			val, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("bad %s %q: not an integer", EnvParallelDownloads, v)
			}
			// CIPD_PARALLEL_DOWNLOADS == 0 means "no parallel work at all", this is
			// conveyed by negatives in opts.ParallelDownloads (because 0 was already
			// used to represent "use defaults").
			if val == 0 {
				opts.ParallelDownloads = -1
			} else {
				opts.ParallelDownloads = val
			}
		}
	}
	if opts.UserAgent == "" {
		if v := env.Get(EnvHTTPUserAgentPrefix); v != "" {
			opts.UserAgent = fmt.Sprintf("%s/%s", v, UserAgent)
		}
	}
	if opts.PluginHost != nil && len(opts.AdmissionPlugin) == 0 {
		if v := env.Get(EnvAdmissionPlugin); v != "" {
			if err := json.Unmarshal([]byte(v), &opts.AdmissionPlugin); err != nil {
				return fmt.Errorf("bad %s %q: not a valid JSON", EnvAdmissionPlugin, v)
			}
		}
	}
	if v := env.Get(EnvCIPDServiceURL); v != "" {
		opts.ServiceURL = v
	}
	return nil
}

// NewClient initializes CIPD client object.
//
// Note that when constructing ClientOptions, you almost certainly also need
// to call LoadFromEnv to load unset values from the process environment.
func NewClient(opts ClientOptions) (Client, error) {
	if opts.AnonymousClient == nil {
		opts.AnonymousClient = http.DefaultClient
	}
	if opts.AuthenticatedClient == nil {
		opts.AuthenticatedClient = opts.AnonymousClient
	}
	if opts.UserAgent == "" {
		opts.UserAgent = UserAgent
	}
	if opts.MaxThreads <= 0 {
		opts.MaxThreads = runtime.NumCPU()
	}

	// Validate and normalize service URL.
	if opts.ServiceURL == "" {
		return nil, fmt.Errorf("ServiceURL is required")
	}
	parsed, err := url.Parse(opts.ServiceURL)
	if err != nil {
		return nil, fmt.Errorf("not a valid URL %q: %s", opts.ServiceURL, err)
	}
	if parsed.Path != "" && parsed.Path != "/" {
		return nil, fmt.Errorf("expecting a root URL, not %q", opts.ServiceURL)
	}
	opts.ServiceURL = fmt.Sprintf("%s://%s", parsed.Scheme, parsed.Host)

	prpcC := &prpc.Client{
		C:    opts.AuthenticatedClient,
		Host: parsed.Host,
		Options: &prpc.Options{
			UserAgent: opts.UserAgent,
			Insecure:  parsed.Scheme == "http", // for testing with local dev server
			Retry: func() retry.Iterator {
				return &retry.ExponentialBackoff{
					Limited: retry.Limited{
						Delay:   time.Second,
						Retries: 10,
					},
				}
			},
		},
	}

	cas := opts.casMock
	if cas == nil {
		cas = api.NewStorageClient(prpcC)
	}
	repo := opts.repoMock
	if repo == nil {
		repo = api.NewRepositoryClient(prpcC)
	}
	s := opts.storageMock
	if s == nil {
		s = &storageImpl{
			chunkSize: uploadChunkSize,
			userAgent: opts.UserAgent,
			client:    opts.AnonymousClient,
		}
	}

	if opts.PluginHost != nil {
		err := opts.PluginHost.Initialize(plugin.Config{
			ServiceURL: opts.ServiceURL,
			Repository: repo,
		})
		if err != nil {
			return nil, errors.Annotate(err, "failed to initialize the plugin host").Err()
		}
	}

	client := &clientImpl{
		ClientOptions: opts,
		cas:           cas,
		repo:          repo,
		storage:       s,
		deployer:      deployer.New(opts.Root),
		pluginHost:    opts.PluginHost,
	}

	if len(opts.AdmissionPlugin) != 0 {
		if client.pluginHost == nil {
			return nil, errors.Reason("using admission plugins requires configuring the plugin system").Err()
		}
		client.pluginAdmission, err = client.pluginHost.NewAdmissionPlugin(opts.AdmissionPlugin)
		if err != nil {
			return nil, errors.Annotate(err, "failed to initialize the admission plugin").Err()
		}
	}

	return client, nil
}

// MaybeUpdateClient will update the client binary at clientExe (given as
// a native path) to targetVersion if it's out of date (based on its hash).
//
// This update is done from the "infra/tools/cipd/${platform}" package, see
// ClientPackage. The function will use the given ClientOptions to figure out
// how to establish a connection with the backend. Its Root and CacheDir values
// are ignored (values derived from clientExe are used instead).
//
// If given 'digests' is not nil, will make sure the hash of the downloaded
// client binary is in 'digests'.
//
// Note that this function make sense only in a context of a default CIPD CLI
// client. Other binaries that link to cipd package should not use it, they'll
// be "updated" to the CIPD client binary.
func MaybeUpdateClient(ctx context.Context, opts ClientOptions, targetVersion, clientExe string, digests *digests.ClientDigestsFile) (common.Pin, error) {
	if err := common.ValidateInstanceVersion(targetVersion); err != nil {
		return common.Pin{}, err
	}

	opts.Root = filepath.Dir(clientExe)
	opts.CacheDir = filepath.Join(opts.Root, ".cipd_client_cache")

	client, err := NewClient(opts)
	if err != nil {
		return common.Pin{}, err
	}
	defer client.Close(ctx)

	impl := client.(*clientImpl)

	fs := fs.NewFileSystem(opts.Root, filepath.Join(opts.CacheDir, "trash"))
	defer fs.CleanupTrash(ctx)

	pin, err := impl.maybeUpdateClient(ctx, fs, targetVersion, clientExe, digests)
	if err == nil {
		impl.ensureClientVersionInfo(ctx, fs, pin, clientExe)
	}
	return pin, err
}

type clientImpl struct {
	ClientOptions

	// pRPC API clients.
	cas  api.StorageClient
	repo api.RepositoryClient

	// batchLock protects guts of by BeginBatch/EndBatch implementation.
	batchLock    sync.Mutex
	batchNesting int
	batchPending map[batchAwareOp]struct{}

	// storage knows how to upload and download raw binaries using signed URLs.
	storage storage
	// deployer knows how to install packages to local file system. Thread safe.
	deployer deployer.Deployer

	// tagCache is a file-system based cache of resolved tags.
	tagCache     *internal.TagCache
	tagCacheInit sync.Once

	// Plugin system.
	pluginHost      plugin.Host            // nil if disabled
	pluginAdmission plugin.AdmissionPlugin // nil if disabled
}

type batchAwareOp int

const (
	batchAwareOpSaveTagCache batchAwareOp = iota
	batchAwareOpCleanupTrash
	batchAwareOpClearAdmissionCache
)

// See https://golang.org/ref/spec#Method_expressions
var batchAwareOps = map[batchAwareOp]func(*clientImpl, context.Context){
	batchAwareOpSaveTagCache:        (*clientImpl).saveTagCache,
	batchAwareOpCleanupTrash:        (*clientImpl).cleanupTrash,
	batchAwareOpClearAdmissionCache: (*clientImpl).clearAdmissionCache,
}

func (c *clientImpl) saveTagCache(ctx context.Context) {
	if c.tagCache != nil {
		if err := c.tagCache.Save(ctx); err != nil {
			logging.Warningf(ctx, "Failed to save tag cache: %s", err)
		}
	}
}

func (c *clientImpl) cleanupTrash(ctx context.Context) {
	if f := c.deployer.FS(); f != nil {
		f.EnsureDirectoryGone(ctx, filepath.Join(f.Root(), fs.SiteServiceDir, "tmp"))
		f.CleanupTrash(ctx)
	}
}

func (c *clientImpl) clearAdmissionCache(ctx context.Context) {
	if c.pluginAdmission != nil {
		c.pluginAdmission.ClearCache()
	}
}

// getTagCache lazy-initializes tagCache and returns it.
//
// May return nil if tag cache is disabled.
func (c *clientImpl) getTagCache() *internal.TagCache {
	c.tagCacheInit.Do(func() {
		var dir string
		switch {
		case c.CacheDir != "":
			dir = c.CacheDir
		case c.Root != "":
			dir = filepath.Join(c.Root, fs.SiteServiceDir)
		default:
			return
		}
		parsed, err := url.Parse(c.ServiceURL)
		if err != nil {
			panic(err) // the URL has been validated in NewClient already
		}
		c.tagCache = internal.NewTagCache(fs.NewFileSystem(dir, ""), parsed.Host)
	})
	return c.tagCache
}

// instanceCache returns an instance cache to download packages into.
//
// Returns a new object each time. Multiple InstanceCache objects may perhaps
// share the same underlying cache directory if used concurrently (just like two
// different CIPD processes share it).
//
// This is a heavy object that may spawn multiple goroutines inside. Must be
// closed with Close() when done working with it.
func (c *clientImpl) instanceCache(ctx context.Context) (*internal.InstanceCache, error) {
	var cacheDir string
	var tmp bool

	if c.CacheDir != "" {
		// This is a persistent global cache (not a temp one).
		cacheDir = filepath.Join(c.CacheDir, "instances")
	} else {
		// This is going to be a temporary cache that self-destructs.
		tmp = true

		if c.Root != "" {
			// Create the root tmp directory in the site root guts.
			tmpDir, err := c.deployer.FS().EnsureDirectory(ctx, filepath.Join(c.Root, fs.SiteServiceDir, "tmp"))
			if err != nil {
				return nil, err
			}
			// An inside it create a unique directory for the new InstanceCache.
			// Multiple temp caches must not reuse the same directory or they'll
			// interfere with one another when deleting instances or cleaning them up
			// when closing.
			if cacheDir, err = ioutil.TempDir(tmpDir, "dl_"); err != nil {
				return nil, err
			}
		} else {
			// When not using a site root, just create the directory in /tmp.
			var err error
			if cacheDir, err = ioutil.TempDir("", "cipd_dl_"); err != nil {
				return nil, err
			}
		}
	}

	// Since 0 is used as "use defaults" indicator, we have to use negatives to
	// represent "do not do anything in parallel at all" (by passing 0 to
	// the InstanceCache).
	parallelDownloads := c.ParallelDownloads
	switch {
	case parallelDownloads == 0:
		parallelDownloads = DefaultParallelDownloads
	case parallelDownloads < 0:
		parallelDownloads = 0
	}

	cache := &internal.InstanceCache{
		FS:                fs.NewFileSystem(cacheDir, ""),
		Tmp:               tmp,
		Fetcher:           c.remoteFetchInstance,
		ParallelDownloads: parallelDownloads,
	}
	cache.Launch(ctx) // start background download goroutines
	return cache, nil
}

func (c *clientImpl) Close(ctx context.Context) {
	if c.pluginAdmission != nil {
		c.pluginAdmission.Close(ctx)
	}
	if c.pluginHost != nil {
		c.pluginHost.Close(ctx)
	}
}

func (c *clientImpl) BeginBatch(ctx context.Context) {
	c.batchLock.Lock()
	defer c.batchLock.Unlock()
	c.batchNesting++
}

func (c *clientImpl) EndBatch(ctx context.Context) {
	c.batchLock.Lock()
	defer c.batchLock.Unlock()
	if c.batchNesting <= 0 {
		panic("EndBatch called without corresponding BeginBatch")
	}
	if c.batchNesting--; c.batchNesting == 0 {
		for op := range c.batchPending {
			batchAwareOps[op](c, ctx)
		}
		c.batchPending = nil
	}
}

func (c *clientImpl) doBatchAwareOp(ctx context.Context, op batchAwareOp) {
	c.batchLock.Lock()
	defer c.batchLock.Unlock()
	if c.batchNesting == 0 {
		// Not inside a batch, execute right now.
		batchAwareOps[op](c, ctx)
	} else {
		// Schedule to execute when 'EndBatch' is called.
		if c.batchPending == nil {
			c.batchPending = make(map[batchAwareOp]struct{}, 1)
		}
		c.batchPending[op] = struct{}{}
	}
}

func (c *clientImpl) FetchACL(ctx context.Context, prefix string) ([]PackageACL, error) {
	if _, err := common.ValidatePackagePrefix(prefix); err != nil {
		return nil, err
	}

	resp, err := c.repo.GetInheritedPrefixMetadata(ctx, &api.PrefixRequest{
		Prefix: prefix,
	}, expectedCodes)
	if err != nil {
		return nil, c.humanErr(err)
	}

	return prefixMetadataToACLs(resp), nil
}

func (c *clientImpl) ModifyACL(ctx context.Context, prefix string, changes []PackageACLChange) error {
	if _, err := common.ValidatePackagePrefix(prefix); err != nil {
		return err
	}

	// Fetch existing metadata, if any.
	meta, err := c.repo.GetPrefixMetadata(ctx, &api.PrefixRequest{
		Prefix: prefix,
	}, expectedCodes)
	if code := status.Code(err); code != codes.OK && code != codes.NotFound {
		return c.humanErr(err)
	}

	// Construct new empty metadata for codes.NotFound.
	if meta == nil {
		meta = &api.PrefixMetadata{Prefix: prefix}
	}

	// Apply mutations.
	if dirty, err := mutateACLs(meta, changes); !dirty || err != nil {
		return err
	}

	// Store the new metadata. This call will check meta.Fingerprint.
	_, err = c.repo.UpdatePrefixMetadata(ctx, meta, expectedCodes)
	return c.humanErr(err)
}

func (c *clientImpl) FetchRoles(ctx context.Context, prefix string) ([]string, error) {
	if _, err := common.ValidatePackagePrefix(prefix); err != nil {
		return nil, err
	}

	resp, err := c.repo.GetRolesInPrefix(ctx, &api.PrefixRequest{
		Prefix: prefix,
	}, expectedCodes)
	if err != nil {
		return nil, c.humanErr(err)
	}

	out := make([]string, len(resp.Roles))
	for i, r := range resp.Roles {
		out[i] = r.Role.String()
	}
	return out, nil
}

func (c *clientImpl) ListPackages(ctx context.Context, prefix string, recursive, includeHidden bool) ([]string, error) {
	if _, err := common.ValidatePackagePrefix(prefix); err != nil {
		return nil, err
	}

	resp, err := c.repo.ListPrefix(ctx, &api.ListPrefixRequest{
		Prefix:        prefix,
		Recursive:     recursive,
		IncludeHidden: includeHidden,
	}, expectedCodes)
	if err != nil {
		return nil, c.humanErr(err)
	}

	listing := resp.Packages
	for _, pfx := range resp.Prefixes {
		listing = append(listing, pfx+"/")
	}
	sort.Strings(listing)
	return listing, nil
}

func (c *clientImpl) ResolveVersion(ctx context.Context, packageName, version string) (common.Pin, error) {
	if err := common.ValidatePackageName(packageName); err != nil {
		return common.Pin{}, err
	}

	// Is it instance ID already? Then it is already resolved.
	if common.ValidateInstanceID(version, common.AnyHash) == nil {
		return common.Pin{PackageName: packageName, InstanceID: version}, nil
	}
	if err := common.ValidateInstanceVersion(version); err != nil {
		return common.Pin{}, err
	}

	// Use the preresolved version if configured to do so. Do NOT fallback to
	// the backend calls. A missing version is an error.
	if c.Versions != nil {
		return c.Versions.ResolveVersion(packageName, version)
	}

	// Use a local cache when resolving tags to avoid round trips to the backend
	// when calling same 'cipd ensure' command again and again.
	var cache *internal.TagCache
	if common.ValidateInstanceTag(version) == nil {
		cache = c.getTagCache() // note: may be nil if the cache is disabled
	}
	if cache != nil {
		cached, err := cache.ResolveTag(ctx, packageName, version)
		if err != nil {
			logging.Warningf(ctx, "Could not query tag cache: %s", err)
		}
		if cached.InstanceID != "" {
			logging.Debugf(ctx, "Tag cache hit for %s:%s - %s", packageName, version, cached.InstanceID)
			return cached, nil
		}
	}

	// Either resolving a ref, or a tag cache miss? Hit the backend.
	resp, err := c.repo.ResolveVersion(ctx, &api.ResolveVersionRequest{
		Package: packageName,
		Version: version,
	}, expectedCodes)
	if err != nil {
		return common.Pin{}, c.humanErr(err)
	}
	pin := common.Pin{
		PackageName: packageName,
		InstanceID:  common.ObjectRefToInstanceID(resp.Instance),
	}

	// If was resolving a tag, store it in the cache.
	if cache != nil {
		if err := cache.AddTag(ctx, pin, version); err != nil {
			logging.Warningf(ctx, "Could not add tag to the cache")
		}
		c.doBatchAwareOp(ctx, batchAwareOpSaveTagCache)
	}

	return pin, nil
}

// ensureClientVersionInfo is called only with the specially constructed client,
// see MaybeUpdateClient function.
func (c *clientImpl) ensureClientVersionInfo(ctx context.Context, fs fs.FileSystem, pin common.Pin, clientExe string) {
	expect, err := json.Marshal(version.Info{
		PackageName: pin.PackageName,
		InstanceID:  pin.InstanceID,
	})
	if err != nil {
		// Should never occur; only error could be if version.Info is not JSON
		// serializable.
		logging.WithError(err).Errorf(ctx, "Unable to generate version file content")
		return
	}

	verFile := version.GetVersionFile(clientExe)
	if data, err := ioutil.ReadFile(verFile); err == nil && bytes.Equal(expect, data) {
		return // up to date
	}

	// There was an error reading the existing version file, or its content does
	// not match. Proceed with EnsureFile.
	err = fs.EnsureFile(ctx, verFile, func(of *os.File) error {
		_, err := of.Write(expect)
		return err
	})
	if err != nil {
		logging.WithError(err).Warningf(ctx, "Unable to update version info %q", verFile)
	}
}

// maybeUpdateClient is called only with the specially constructed client, see
// MaybeUpdateClient function.
func (c *clientImpl) maybeUpdateClient(ctx context.Context, fs fs.FileSystem,
	targetVersion, clientExe string, digests *digests.ClientDigestsFile) (common.Pin, error) {

	// currentHashMatches calculates the existing client binary hash and compares
	// it to 'obj'.
	currentHashMatches := func(obj *api.ObjectRef) (yep bool, err error) {
		hash, err := common.NewHash(obj.HashAlgo)
		if err != nil {
			return false, err
		}
		file, err := os.Open(clientExe)
		if err != nil {
			return false, err
		}
		defer file.Close()
		if _, err := io.Copy(hash, file); err != nil {
			return false, err
		}
		return common.HexDigest(hash) == obj.HexDigest, nil
	}

	c.BeginBatch(ctx)
	defer c.EndBatch(ctx)

	// Resolve the client version to a pin, to be able to later grab URL to the
	// binary by querying info for that pin.
	var pin common.Pin
	clientPackage, err := template.DefaultExpander().Expand(ClientPackage)
	if err != nil {
		return common.Pin{}, err // shouldn't be happening in reality
	}
	if pin, err = c.ResolveVersion(ctx, clientPackage, targetVersion); err != nil {
		return common.Pin{}, err
	}

	// The name of the client binary inside the client CIPD package. Acts only as
	// a key inside the extracted refs cache, nothing more, so can technically be
	// arbitrary.
	clientFileName := "cipd"
	if platform.CurrentOS() == "windows" {
		clientFileName = "cipd.exe"
	}

	// rememberClientRef populates the extracted refs cache.
	rememberClientRef := func(pin common.Pin, ref *api.ObjectRef) {
		if cache := c.getTagCache(); cache != nil {
			cache.AddExtractedObjectRef(ctx, pin, clientFileName, ref)
			c.doBatchAwareOp(ctx, batchAwareOpSaveTagCache)
		}
	}

	// Look up the hash corresponding to the pin in the extracted refs cache. See
	// rememberClientRef calls below for where it is stored initially. A cache
	// miss is fine, we'll reach to the backend to get the hash. A warm cache
	// allows skipping RPCs to the backend on a "happy path", when the client is
	// already up-to-date.
	var clientRef *api.ObjectRef
	if cache := c.getTagCache(); cache != nil {
		if clientRef, err = cache.ResolveExtractedObjectRef(ctx, pin, clientFileName); err != nil {
			return common.Pin{}, err
		}
	}

	// If not using the tags cache or it is cold, ask the backend for an expected
	// client ref. Note that we do it even if we have 'digests' file available,
	// to handle the case when 'digests' file is stale (which can happen when
	// updating the client version file).
	var info *ClientDescription
	if clientRef == nil {
		if info, err = c.DescribeClient(ctx, pin); err != nil {
			return common.Pin{}, err
		}
		clientRef = info.Digest
	}

	// If using pinned client digests, make sure the hash reported by the backend
	// is mentioned there. In most cases a mismatch means the pinned digests file
	// is just stale. The mismatch can also happen if the backend is compromised
	// or the client package was forcefully replaced (this should never really
	// happen...).
	if digests != nil {
		plat := platform.CurrentPlatform()
		switch pinnedRef := digests.ClientRef(plat); {
		case pinnedRef == nil:
			return common.Pin{}, fmt.Errorf("there's no supported hash for %q in CIPD *.digests file", plat)
		case !digests.Contains(plat, clientRef):
			return common.Pin{}, fmt.Errorf(
				"the CIPD client hash reported by the backend (%s) is not in *.digests file, "+
					"if you changed CIPD client version recently most likely the *.digests "+
					"file is just stale and needs to be regenerated via 'cipd selfupdate-roll ...'",
				clientRef.HexDigest)
		default:
			clientRef = pinnedRef // pick the best supported hash algo from *.digests
		}
	}

	// Is the client binary already up-to-date (has the expected hash)?
	switch yep, err := currentHashMatches(clientRef); {
	case err != nil:
		return common.Pin{}, err // can't read clientExe
	case yep:
		// If we had to fetch the expected hash, store it in the cache to avoid
		// fetching it again. Don't do it if we read it from the cache initially (to
		// skip unnecessary cache write).
		if info != nil {
			rememberClientRef(pin, clientRef)
		}
		return pin, nil
	}

	if targetVersion == pin.InstanceID {
		logging.Infof(ctx, "Updating CIPD client to %s", pin)
	} else {
		logging.Infof(ctx, "Updating CIPD client to %s (%s)", pin, targetVersion)
	}

	// Grab the signed URL of the client binary if we haven't done so already.
	if info == nil {
		if info, err = c.DescribeClient(ctx, pin); err != nil {
			return common.Pin{}, err
		}
	}

	// Here we know for sure that the current binary has wrong hash (most likely
	// it is outdated). Fetch the new binary, verifying its hash matches the one
	// we expect.
	err = c.installClient(
		ctx, fs,
		common.MustNewHash(clientRef.HashAlgo),
		info.SignedURL,
		clientExe,
		clientRef.HexDigest)
	if err != nil {
		// Either a download error or hash mismatch.
		return common.Pin{}, errors.Annotate(err, "when updating the CIPD client to %q", targetVersion).Err()
	}

	// The new fetched binary is valid.
	rememberClientRef(pin, clientRef)
	return pin, nil
}

func (c *clientImpl) RegisterInstance(ctx context.Context, pin common.Pin, src pkg.Source, timeout time.Duration) (err error) {
	if timeout == 0 {
		timeout = CASFinalizationTimeout
	}

	ctx, done := ui.NewActivity(ctx, nil, "")
	defer func() {
		if err != nil {
			logging.Errorf(ctx, "Instance registration failed: %s", err)
		}
		done()
	}()

	// attemptToRegister calls RegisterInstance RPC and logs the result.
	attemptToRegister := func() (*api.UploadOperation, error) {
		logging.Infof(ctx, "Registering %s", pin)
		resp, err := c.repo.RegisterInstance(ctx, &api.Instance{
			Package:  pin.PackageName,
			Instance: common.InstanceIDToObjectRef(pin.InstanceID),
		}, expectedCodes)
		if err != nil {
			return nil, c.humanErr(err)
		}
		switch resp.Status {
		case api.RegistrationStatus_REGISTERED:
			logging.Infof(ctx, "Instance %s was successfully registered", pin)
			return nil, nil
		case api.RegistrationStatus_ALREADY_REGISTERED:
			logging.Infof(
				ctx, "Instance %s is already registered by %s on %s",
				pin, resp.Instance.RegisteredBy,
				resp.Instance.RegisteredTs.AsTime().Local())
			return nil, nil
		case api.RegistrationStatus_NOT_UPLOADED:
			return resp.UploadOp, nil
		default:
			return nil, fmt.Errorf("unrecognized package registration status %s", resp.Status)
		}
	}

	// Attempt to register. May be asked to actually upload the file first.
	uploadOp, err := attemptToRegister()
	switch {
	case err != nil:
		return err
	case uploadOp == nil:
		return nil // no need to upload, the instance is registered
	}

	// The backend asked us to upload the data to CAS. Do it.
	if err := c.storage.upload(ctx, uploadOp.UploadUrl, io.NewSectionReader(src, 0, src.Size())); err != nil {
		return err
	}
	if err := c.finalizeUpload(ctx, uploadOp.OperationId, timeout); err != nil {
		return err
	}
	logging.Infof(ctx, "Successfully uploaded and verified %s", pin)

	// Try the registration again now that the file is uploaded to CAS. It should
	// succeed.
	switch uploadOp, err := attemptToRegister(); {
	case uploadOp != nil:
		return ErrBadUpload // welp, the upload didn't work for some reason, give up
	default:
		return err
	}
}

// finalizeUpload repeatedly calls FinishUpload RPC until server reports that
// the uploaded file has been verified.
func (c *clientImpl) finalizeUpload(ctx context.Context, opID string, timeout time.Duration) error {
	ctx, cancel := clock.WithTimeout(ctx, timeout)
	defer cancel()

	sleep := time.Second
	for {
		select {
		case <-ctx.Done():
			return ErrFinalizationTimeout
		default:
		}

		op, err := c.cas.FinishUpload(ctx, &api.FinishUploadRequest{
			UploadOperationId: opID,
		})
		switch {
		case status.Code(err) == codes.DeadlineExceeded:
			continue // this may be short RPC deadline, try again
		case err != nil:
			return c.humanErr(err)
		case op.Status == api.UploadStatus_PUBLISHED:
			return nil // verified!
		case op.Status == api.UploadStatus_ERRORED:
			return errors.New(op.ErrorMessage) // fatal verification error
		case op.Status == api.UploadStatus_UPLOADING || op.Status == api.UploadStatus_VERIFYING:
			logging.Infof(ctx, "Verifying...")
			if clock.Sleep(clock.Tag(ctx, "cipd-sleeping"), sleep).Incomplete() {
				return ErrFinalizationTimeout
			}
			if sleep < 10*time.Second {
				sleep += 500 * time.Millisecond
			}
		default:
			return fmt.Errorf("unrecognized upload operation status %s", op.Status)
		}
	}
}

func (c *clientImpl) DescribeInstance(ctx context.Context, pin common.Pin, opts *DescribeInstanceOpts) (*InstanceDescription, error) {
	if err := common.ValidatePin(pin, common.AnyHash); err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &DescribeInstanceOpts{}
	}

	resp, err := c.repo.DescribeInstance(ctx, &api.DescribeInstanceRequest{
		Package:      pin.PackageName,
		Instance:     common.InstanceIDToObjectRef(pin.InstanceID),
		DescribeRefs: opts.DescribeRefs,
		DescribeTags: opts.DescribeTags,
	}, expectedCodes)
	if err != nil {
		return nil, c.humanErr(err)
	}

	return apiDescToInfo(resp), nil
}

func (c *clientImpl) DescribeClient(ctx context.Context, pin common.Pin) (*ClientDescription, error) {
	if err := common.ValidatePin(pin, common.AnyHash); err != nil {
		return nil, err
	}

	resp, err := c.repo.DescribeClient(ctx, &api.DescribeClientRequest{
		Package:  pin.PackageName,
		Instance: common.InstanceIDToObjectRef(pin.InstanceID),
	}, expectedCodes)
	if err != nil {
		return nil, c.humanErr(err)
	}

	return apiClientDescToInfo(resp), nil
}

func (c *clientImpl) SetRefWhenReady(ctx context.Context, ref string, pin common.Pin) error {
	if err := common.ValidatePackageRef(ref); err != nil {
		return err
	}
	if err := common.ValidatePin(pin, common.AnyHash); err != nil {
		return err
	}

	ctx, done := ui.NewActivity(ctx, nil, "")
	defer done()

	logging.Infof(ctx, "Setting ref of %s: %q => %q", pin.PackageName, ref, pin.InstanceID)

	err := c.retryUntilReady(ctx, SetRefTimeout, func(ctx context.Context) error {
		_, err := c.repo.CreateRef(ctx, &api.Ref{
			Name:     ref,
			Package:  pin.PackageName,
			Instance: common.InstanceIDToObjectRef(pin.InstanceID),
		}, expectedCodes)
		return err
	})

	switch err {
	case nil:
		logging.Infof(ctx, "Ref %q is set", ref)
	case ErrProcessingTimeout:
		logging.Errorf(ctx, "Failed to set ref: deadline exceeded")
	default:
		logging.Errorf(ctx, "Failed to set ref: %s", err)
	}
	return err
}

func (c *clientImpl) AttachTagsWhenReady(ctx context.Context, pin common.Pin, tags []string) error {
	if err := common.ValidatePin(pin, common.AnyHash); err != nil {
		return err
	}
	if len(tags) == 0 {
		return nil
	}

	apiTags := make([]*api.Tag, len(tags))
	for i, t := range tags {
		var err error
		if apiTags[i], err = common.ParseInstanceTag(t); err != nil {
			return err
		}
	}

	ctx, done := ui.NewActivity(ctx, nil, "")
	defer done()

	if len(tags) == 1 {
		logging.Infof(ctx, "Attaching tag %q", tags[0])
	} else {
		for _, t := range tags {
			logging.Infof(ctx, "Will attach tag %q", t)
		}
		logging.Infof(ctx, "Attaching tags")
	}

	err := c.retryUntilReady(ctx, TagAttachTimeout, func(ctx context.Context) error {
		_, err := c.repo.AttachTags(ctx, &api.AttachTagsRequest{
			Package:  pin.PackageName,
			Instance: common.InstanceIDToObjectRef(pin.InstanceID),
			Tags:     apiTags,
		}, expectedCodes)
		return err
	})

	switch err {
	case nil:
		if len(tags) == 1 {
			logging.Infof(ctx, "Tag %q was attached", tags[0])
		} else {
			logging.Infof(ctx, "All tags were attached")
		}
	case ErrProcessingTimeout:
		logging.Errorf(ctx, "Failed to attach tags: deadline exceeded")
	default:
		logging.Errorf(ctx, "Failed to attach tags: %s", err)
	}
	return err
}

func (c *clientImpl) AttachMetadataWhenReady(ctx context.Context, pin common.Pin, md []Metadata) error {
	if err := common.ValidatePin(pin, common.AnyHash); err != nil {
		return err
	}
	if len(md) == 0 {
		return nil
	}

	apiMD := make([]*api.InstanceMetadata, len(md))
	for i, m := range md {
		if err := common.ValidateInstanceMetadataKey(m.Key); err != nil {
			return err
		}
		if err := common.ValidateInstanceMetadataLen(len(m.Value)); err != nil {
			return errors.Annotate(err, "bad metadata %q", m.Key).Err()
		}
		if err := common.ValidateContentType(m.ContentType); err != nil {
			return errors.Annotate(err, "bad metadata %q", m.Key).Err()
		}
		apiMD[i] = &api.InstanceMetadata{
			Key:         m.Key,
			Value:       m.Value,
			ContentType: m.ContentType,
		}
	}

	ctx, done := ui.NewActivity(ctx, nil, "")
	defer done()

	if len(md) == 1 {
		logging.Infof(ctx, "Attaching metadata %q", md[0].Key)
	} else {
		for _, m := range md {
			logging.Infof(ctx, "Will attach metadata %q", m.Key)
		}
		logging.Infof(ctx, "Attaching metadata")
	}

	err := c.retryUntilReady(ctx, MetadataAttachTimeout, func(ctx context.Context) error {
		_, err := c.repo.AttachMetadata(ctx, &api.AttachMetadataRequest{
			Package:  pin.PackageName,
			Instance: common.InstanceIDToObjectRef(pin.InstanceID),
			Metadata: apiMD,
		}, expectedCodes)
		return err
	})

	switch err {
	case nil:
		if len(md) == 1 {
			logging.Infof(ctx, "Metadata %q was attached", md[0].Key)
		} else {
			logging.Infof(ctx, "Metadata was attached")
		}
	case ErrProcessingTimeout:
		logging.Errorf(ctx, "Failed to attach metadata: deadline exceeded")
	default:
		logging.Errorf(ctx, "Failed to attach metadata: %s", err)
	}
	return err
}

// How long to wait between retries in retryUntilReady.
const retryDelay = 5 * time.Second

// retryUntilReady calls the callback and retries on FailedPrecondition errors,
// which indicate that the instance is not ready yet (still being processed by
// the backend).
func (c *clientImpl) retryUntilReady(ctx context.Context, timeout time.Duration, cb func(context.Context) error) error {
	ctx, cancel := clock.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return ErrProcessingTimeout
		default:
		}

		switch err := cb(ctx); {
		case err == nil:
			return nil
		case status.Code(err) == codes.DeadlineExceeded:
			continue // this may be short RPC deadline, try again
		case status.Code(err) == codes.FailedPrecondition: // the instance is not ready
			logging.Warningf(ctx, "Not ready: %s", c.humanErr(err))
			if clock.Sleep(clock.Tag(ctx, "cipd-sleeping"), retryDelay).Incomplete() {
				return ErrProcessingTimeout
			}
		default:
			return c.humanErr(err)
		}
	}
}

func (c *clientImpl) SearchInstances(ctx context.Context, packageName string, tags []string) (common.PinSlice, error) {
	if err := common.ValidatePackageName(packageName); err != nil {
		return nil, err
	}
	if len(tags) == 0 {
		return nil, errors.New("at least one tag is required")
	}

	apiTags := make([]*api.Tag, len(tags))
	for i, t := range tags {
		var err error
		if apiTags[i], err = common.ParseInstanceTag(t); err != nil {
			return nil, err
		}
	}

	resp, err := c.repo.SearchInstances(ctx, &api.SearchInstancesRequest{
		Package:  packageName,
		Tags:     apiTags,
		PageSize: 1000, // TODO(vadimsh): Support pagination on the c.
	}, expectedCodes)
	switch {
	case status.Code(err) == codes.NotFound: // no such package => no instances
		return nil, nil
	case err != nil:
		return nil, c.humanErr(err)
	}

	out := make(common.PinSlice, len(resp.Instances))
	for i, inst := range resp.Instances {
		out[i] = apiInstanceToInfo(inst).Pin
	}
	if resp.NextPageToken != "" {
		logging.Warningf(ctx, "Truncating the result only to first %d instances", len(resp.Instances))
	}
	return out, nil
}

func (c *clientImpl) ListInstances(ctx context.Context, packageName string) (InstanceEnumerator, error) {
	if err := common.ValidatePackageName(packageName); err != nil {
		return nil, err
	}
	return &instanceEnumeratorImpl{
		fetch: func(ctx context.Context, limit int, cursor string) ([]InstanceInfo, string, error) {
			resp, err := c.repo.ListInstances(ctx, &api.ListInstancesRequest{
				Package:   packageName,
				PageSize:  int32(limit),
				PageToken: cursor,
			}, expectedCodes)
			if err != nil {
				return nil, "", c.humanErr(err)
			}
			instances := make([]InstanceInfo, len(resp.Instances))
			for i, inst := range resp.Instances {
				instances[i] = apiInstanceToInfo(inst)
			}
			return instances, resp.NextPageToken, nil
		},
	}, nil
}

func (c *clientImpl) FetchPackageRefs(ctx context.Context, packageName string) ([]RefInfo, error) {
	if err := common.ValidatePackageName(packageName); err != nil {
		return nil, err
	}

	resp, err := c.repo.ListRefs(ctx, &api.ListRefsRequest{
		Package: packageName,
	}, expectedCodes)
	if err != nil {
		return nil, c.humanErr(err)
	}

	refs := make([]RefInfo, len(resp.Refs))
	for i, r := range resp.Refs {
		refs[i] = apiRefToInfo(r)
	}
	return refs, nil
}

func (c *clientImpl) FetchInstance(ctx context.Context, pin common.Pin) (pkg.Source, error) {
	if err := common.ValidatePin(pin, common.KnownHash); err != nil {
		return nil, err
	}

	ctx, done := ui.NewActivity(ctx, nil, "")
	defer done()

	cache, err := c.instanceCache(ctx)
	if err != nil {
		return nil, err
	}
	defer cache.Close(ctx)

	cache.RequestInstances([]*internal.InstanceRequest{
		{Context: ctx, Pin: pin},
	})
	res := cache.WaitInstance()

	if res.Err != nil {
		return nil, res.Err
	}
	return res.Source, nil
}

func (c *clientImpl) FetchInstanceTo(ctx context.Context, pin common.Pin, output io.WriteSeeker) error {
	if err := common.ValidatePin(pin, common.KnownHash); err != nil {
		return err
	}

	ctx, done := ui.NewActivity(ctx, nil, "")
	defer done()

	// Deal with no-cache situation first, it is simple - just fetch the instance
	// into the 'output'.
	if c.CacheDir == "" {
		return c.remoteFetchInstance(ctx, pin, output)
	}

	// If using the cache, always fetch into the cache first, and then copy data
	// from the cache into the output.
	input, err := c.FetchInstance(ctx, pin)
	if err != nil {
		return err
	}
	defer func() {
		if err := input.Close(ctx, false); err != nil {
			logging.Warningf(ctx, "Failed to close the instance file: %s", err)
		}
	}()

	logging.Infof(ctx, "Copying the instance into the final destination...")
	if _, err = io.Copy(output, io.NewSectionReader(input, 0, input.Size())); err != nil {
		logging.Errorf(ctx, "Failed to copy the instance file: %s", err)
		return err
	}
	logging.Infof(ctx, "Successfully copied the instance %s", pin)
	return nil
}

// remoteFetchInstance fetches the package file into 'output' and verifies its
// hash along the way. Assumes 'pin' is already validated.
func (c *clientImpl) remoteFetchInstance(ctx context.Context, pin common.Pin, output io.WriteSeeker) (err error) {
	startTS := clock.Now(ctx)
	defer func() {
		if err != nil {
			logging.Errorf(ctx, "Failed to fetch %s: %s", pin, err)
		} else {
			logging.Infof(ctx, "Fetched %s in %.1fs", pin, clock.Now(ctx).Sub(startTS).Seconds())
		}
	}()

	objRef := common.InstanceIDToObjectRef(pin.InstanceID)

	logging.Infof(ctx, "Resolving fetch URL for %s", pin)
	resp, err := c.repo.GetInstanceURL(ctx, &api.GetInstanceURLRequest{
		Package:  pin.PackageName,
		Instance: objRef,
	}, expectedCodes)
	if err != nil {
		return c.humanErr(err)
	}

	hash := common.MustNewHash(objRef.HashAlgo)
	if err = c.storage.download(ctx, resp.SignedUrl, output, hash); err != nil {
		return
	}

	// Make sure we fetched what we've asked for.
	if digest := common.HexDigest(hash); objRef.HexDigest != digest {
		err = fmt.Errorf("package hash mismatch: expecting %q, got %q", objRef.HexDigest, digest)
	}
	return
}

func (c *clientImpl) FindDeployed(ctx context.Context) (common.PinSliceBySubdir, error) {
	return c.deployer.FindDeployed(ctx)
}

func (c *clientImpl) EnsurePackages(ctx context.Context, allPins common.PinSliceBySubdir, opts *EnsureOptions) (aMap ActionMap, err error) {
	if err = allPins.Validate(common.AnyHash); err != nil {
		return
	}

	realOpts := EnsureOptions{}
	if opts != nil {
		realOpts = *opts
	}
	if realOpts.Paranoia == "" {
		realOpts.Paranoia = NotParanoid
	} else if err = realOpts.Paranoia.Validate(); err != nil {
		return
	}

	c.BeginBatch(ctx)
	defer c.EndBatch(ctx)

	// Enumerate existing packages.
	existing, err := c.deployer.FindDeployed(ctx)
	if err != nil {
		return
	}

	// Figure out what needs to be updated and deleted, log it.
	aMap = buildActionPlan(allPins, existing, c.makeRepairChecker(ctx, realOpts.OverrideInstallMode, realOpts.Paranoia))
	if len(aMap) == 0 {
		if !realOpts.Silent {
			logging.Debugf(ctx, "Everything is up-to-date.")
		}
		return
	}

	// TODO(iannucci): ensure that no packages cross root boundaries
	if !realOpts.Silent {
		aMap.Log(ctx, false)
	}
	if realOpts.DryRun {
		if !realOpts.Silent {
			logging.Infof(ctx, "Dry run, not actually doing anything.")
		}
		return
	}

	hasErrors := false
	reportActionErr := func(ctx context.Context, a pinAction, err error) {
		subdir := ""
		if a.subdir != "" {
			subdir = fmt.Sprintf(" in %q", a.subdir)
		}
		logging.Errorf(ctx, "Failed to %s %s%s: %s", a.action, a.pin.PackageName, subdir, err)
		aMap[a.subdir].Errors = append(aMap[a.subdir].Errors, ActionError{
			Action: a.action,
			Pin:    a.pin,
			Error:  JSONError{err},
		})
		hasErrors = true
	}

	// Need a cache to fetch packages into.
	cache, err := c.instanceCache(ctx)
	if err != nil {
		return nil, err
	}
	defer cache.Close(ctx)

	// Figure out what pins we need to fetch and what to do with them once
	// they are fetched. Collect a list of packages to delete and "relink".
	perPinActions := aMap.perPinActions()

	// Enqueue deployment admission checks if have the plugin enabled. They will
	// be consulted later before unzipping fetched instances. This is just an
	// optimization to do checks in parallel with fetching and installing.
	if c.pluginAdmission != nil {
		defer c.doBatchAwareOp(ctx, batchAwareOpClearAdmissionCache)
		for _, a := range perPinActions.updates {
			c.pluginAdmission.CheckAdmission(a.pin)
		}
	}

	// Group all activities together so they get different IDs.
	activities := &ui.ActivityGroup{}

	// The state carried through the fetch task queue. Describes what needs to be
	// done once a package is fetched.
	type pinActionsState struct {
		ctx      context.Context    // the installation activity context
		done     context.CancelFunc // called once the unzipping is done
		pin      common.Pin         // the pin we are fetching
		updates  []pinAction        // what to do with it when we get it
		attempts int                // incremented on a retry after detecting a corruption
	}

	// Start fetching all packages we will need.
	reqs := make([]*internal.InstanceRequest, len(perPinActions.updates))
	for i, a := range perPinActions.updates {
		fetchCtx, fetchDone := ui.NewActivity(ctx, activities, "fetch")
		unzipCtx, unzipDone := ui.NewActivity(ctx, activities, "unzip")
		reqs[i] = &internal.InstanceRequest{
			Context: fetchCtx,
			Done:    fetchDone,
			Pin:     a.pin,
			Open:    true, // want pkg.Instance, not just pkg.Source
			State: pinActionsState{
				ctx:     unzipCtx,
				done:    unzipDone,
				pin:     a.pin,
				updates: a.updates,
			},
		}
	}
	cache.RequestInstances(reqs)

	// While packages are being fetched, do maintenance operations that do not
	// require package data (removals and restoration of broken symlinks).
	if len(perPinActions.maintenance) > 0 {
		ctx, done := ui.NewActivity(ctx, activities, "cleanup")
		for _, a := range perPinActions.maintenance {
			var err error
			switch a.action {
			case ActionRemove:
				err = c.deployer.RemoveDeployed(ctx, a.subdir, a.pin.PackageName)
			case ActionRelink:
				err = c.deployer.RepairDeployed(ctx, a.subdir, a.pin, realOpts.OverrideInstallMode, c.MaxThreads, deployer.RepairParams{
					ToRelink: a.repairPlan.ToRelink,
				})
			default:
				// The invariant of perPinActions().
				panic(fmt.Sprintf("impossible maintenance action %s", a.action))
			}
			if err != nil {
				reportActionErr(ctx, a, err)
			}
		}
		logging.Infof(ctx, "All cleanups done")
		done()
	}

	// As soon as some package data is fetched, perform all installations, updates
	// and repairs that needed it.
	for cache.HasPendingFetches() {
		res := cache.WaitInstance()
		state := res.State.(pinActionsState)
		unzipCtx := state.ctx // the installation activity context
		unzipDone := state.done
		deployErr := res.Err

		// Check if we are even allowed to install this package. Note that
		// CheckAdmission results are cached internally and it is fine to call
		// it multiple times with the same pin (which may happen if we are
		// refetching a corrupted package).
		if deployErr == nil && c.pluginAdmission != nil {
			admErr := c.pluginAdmission.CheckAdmission(state.pin).Wait(unzipCtx)
			if admErr != nil {
				if status, ok := status.FromError(admErr); ok && status.Code() == codes.FailedPrecondition {
					deployErr = errors.Reason("not admitted: %s", status.Message()).Err()
				} else {
					deployErr = errors.Annotate(admErr, "admission check failed").Err()
				}
			}
		}

		// Keep installing stuff as long as it keeps working (no errors).
		actionIdx := 0
		for deployErr == nil && actionIdx < len(state.updates) {
			switch a := state.updates[actionIdx]; a.action {
			case ActionInstall:
				_, deployErr = c.deployer.DeployInstance(unzipCtx, a.subdir, res.Instance, realOpts.OverrideInstallMode, c.MaxThreads)
			case ActionRepair:
				deployErr = c.deployer.RepairDeployed(unzipCtx, a.subdir, state.pin, realOpts.OverrideInstallMode, c.MaxThreads, deployer.RepairParams{
					Instance:   res.Instance,
					ToRedeploy: a.repairPlan.ToRedeploy,
					ToRelink:   a.repairPlan.ToRelink,
				})
			default:
				// The invariant of perPinActions().
				panic(fmt.Sprintf("impossible update action %s", a.action))
			}
			if deployErr == nil {
				actionIdx++
			}
		}

		// Close the instance, marking it as corrupted if necessary. Note it may be
		// nil if res.Err above was non-nil.
		corruption := reader.IsCorruptionError(deployErr)
		if res.Instance != nil {
			res.Instance.Close(unzipCtx, corruption)
		}

		// If we've got a corrupted package, ask the cache to refetch it. We'll
		// resume operations from where we left (redoing the last failed one again).
		// Note that the cache should have removed the bad package from its
		// internals as soon as we closed it as corrupted above.
		//
		// Do it no more than once.
		if corruption && state.attempts < 1 {
			logging.Errorf(unzipCtx, "Failed to unzip %s, will refetch: %s", state.pin.PackageName, deployErr)
			unzipDone()

			refetchCtx, refetchDone := ui.NewActivity(ctx, activities, "refetch")
			reunzipCtx, reunzipDone := ui.NewActivity(ctx, activities, "reunzip")

			cache.RequestInstances([]*internal.InstanceRequest{
				{
					Context: refetchCtx,
					Done:    refetchDone,
					Pin:     state.pin,
					Open:    true,
					State: pinActionsState{
						ctx:      reunzipCtx,
						done:     reunzipDone,
						pin:      state.pin,
						updates:  state.updates[actionIdx:],
						attempts: state.attempts + 1,
					},
				},
			})
			continue
		}

		// If we are here, we are done with this pin (either installed or gave up).
		// Mark all unfinished actions as failed if necessary.
		if deployErr != nil {
			for ; actionIdx < len(state.updates); actionIdx++ {
				reportActionErr(unzipCtx, state.updates[actionIdx], deployErr)
			}
		}
		unzipDone()
	}

	// Opportunistically cleanup the trash left from previous installs.
	c.doBatchAwareOp(ctx, batchAwareOpCleanupTrash)

	if !hasErrors {
		logging.Infof(ctx, "All changes applied.")
	} else {
		err = ErrEnsurePackagesFailed
	}
	return
}

// makeRepairChecker returns a function that decided whether we should attempt
// to repair an already installed package.
//
// The implementation depends on selected paranoia mode.
func (c *clientImpl) makeRepairChecker(ctx context.Context, OverrideInstallMode pkg.InstallMode, paranoia ParanoidMode) repairCB {
	if paranoia == NotParanoid {
		return func(string, common.Pin) *RepairPlan { return nil }
	}

	return func(subdir string, pin common.Pin) *RepairPlan {
		switch state, err := c.deployer.CheckDeployed(ctx, subdir, pin.PackageName, paranoia, WithoutManifest); {
		case err != nil:
			// This error is probably non-recoverable, but we'll try anyway and
			// properly fail later.
			return &RepairPlan{
				NeedsReinstall:  true,
				ReinstallReason: fmt.Sprintf("failed to check the package state: %s", err),
			}
		case !state.Deployed:
			// This should generally not happen. Can probably happen if two clients
			// are messing with same .cipd/* directory concurrently.
			return &RepairPlan{
				NeedsReinstall:  true,
				ReinstallReason: "the package is not deployed at all",
			}
		case state.Pin.InstanceID != pin.InstanceID:
			// Same here.
			return &RepairPlan{
				NeedsReinstall:  true,
				ReinstallReason: fmt.Sprintf("expected to see instance %q, but saw %q", pin.InstanceID, state.Pin.InstanceID),
			}
		case OverrideInstallMode != "" && state.ActualInstallMode != OverrideInstallMode:
			// This package is installed at the right version, but not with the
			// requested override install mode.
			return &RepairPlan{
				NeedsReinstall:  true,
				ReinstallReason: fmt.Sprintf("the package is not deployed in %s-mode (override)", OverrideInstallMode),
			}
		case OverrideInstallMode == "" && state.ActualInstallMode != state.InstallMode:
			// This package is installed at the right version, but not with its
			// native install mode (and we haven't been provided with an override).
			return &RepairPlan{
				NeedsReinstall:  true,
				ReinstallReason: fmt.Sprintf("the package is not deployed in %s-mode (intended)", state.InstallMode),
			}
		case len(state.ToRedeploy) != 0 || len(state.ToRelink) != 0:
			// Have some corrupted files that need to be repaired.
			return &RepairPlan{
				ToRedeploy: state.ToRedeploy,
				ToRelink:   state.ToRelink,
			}
		default:
			return nil // the package needs no repairs
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
// pRPC error handling.

// gRPC errors that may be returned by api.RepositoryClient that we recognize
// and handle ourselves. They will not be logged by the pRPC library.
var expectedCodes = prpc.ExpectedCode(
	codes.Aborted,
	codes.AlreadyExists,
	codes.Canceled,
	codes.FailedPrecondition,
	codes.NotFound,
	codes.PermissionDenied,
)

// humanErr takes gRPC errors and returns a human readable error that can be
// presented in the CLI.
//
// It basically strips scary looking gRPC framing around the error message.
func (c *clientImpl) humanErr(err error) error {
	if err != nil {
		if status, ok := status.FromError(err); ok {
			if status.Code() == codes.PermissionDenied && c.LoginInstructions != "" {
				return errors.Reason("%s, %s", status.Message(), c.LoginInstructions).Err()
			}
			return errors.New(status.Message())
		}
	}
	return err
}
