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

package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/flag/stringlistflag"
	"go.chromium.org/luci/common/flag/stringmapflag"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"
	"go.chromium.org/luci/common/proto/protoc"
	"go.chromium.org/luci/common/system/exitcode"
)

var (
	verbose          = flag.Bool("verbose", false, "print debug messages to stderr")
	protoImportPaths = stringlistflag.Flag{}
	goModules        = stringlistflag.Flag{}
	pathMap          = stringmapflag.Value{}
	withDiscovery    = flag.Bool(
		"discovery", true,
		"generate pb.discovery.go file")
	descFile = flag.String(
		"desc",
		"",
		"write FileDescriptorSet file containing all the the .proto files and their transitive dependencies",
	)
	disableGRPC = flag.Bool(
		"disable-grpc", false,
		"disable grpc and prpc stubs generation, implies -discovery=false",
	)
	useGRPCPlugin = flag.Bool(
		"use-grpc-plugin", false,
		"use protoc-gen-go-grpc to generate gRPC stubs instead of protoc-gen-go",
	)
)

func run(ctx context.Context, inputDir string) error {
	// Stage all requested Go modules under a single root.
	inputs, err := protoc.StageGoInputs(ctx, inputDir, goModules, protoImportPaths)
	if err != nil {
		return err
	}
	defer inputs.Cleanup()

	// Prep a path to the generated descriptors file.
	descPath := *descFile
	if descPath == "" {
		tmpDir, err := ioutil.TempDir("", "")
		if err != nil {
			return err
		}
		defer os.RemoveAll(tmpDir)
		descPath = filepath.Join(tmpDir, "package.desc")
	}

	// Compile all .proto files.
	err = protoc.Compile(ctx, &protoc.CompileParams{
		Inputs:                 inputs,
		OutputDescriptorSet:    descPath,
		GoEnabled:              true,
		GoPackageMap:           pathMap,
		GoDeprecatedGRPCPlugin: !*disableGRPC && !*useGRPCPlugin,
		GoGRPCEnabled:          !*disableGRPC && *useGRPCPlugin,
	})
	if err != nil {
		return err
	}

	// protoc-gen-go puts generated files based on go_package option, rooting them
	// in the inputs.OutputDir. We can't generally guess the Go package name just
	// based on proto file names, but we can extract it from the generated
	// descriptor.
	//
	// Doc:
	// https://developers.google.com/protocol-buffers/docs/reference/go-generated
	descSet, rawDesc, err := loadDescriptorSet(descPath)
	if err != nil {
		return errors.Annotate(err, "failed to load the descriptor set with generated files").Err()
	}

	generatedDesc := make([]*descriptorpb.FileDescriptorProto, 0, len(inputs.ProtoFiles))
	goPackages := stringset.New(0)

	// Since we use --include_imports, there may be a lot of descriptors in the
	// set. Visit only ones we care about.
	for _, protoFile := range inputs.ProtoFiles {
		fileDesc := descSet[path.Join(inputs.ProtoPackage, protoFile)]
		if fileDesc == nil {
			return errors.Reason("descriptor for %q is unexpectedly absent", protoFile).Err()
		}
		generatedDesc = append(generatedDesc, fileDesc)

		// "go_package" option is required now.
		goPackage := fileDesc.Options.GetGoPackage()
		if goPackage == "" {
			return errors.Reason("file %q has no go_package option set, it is required", protoFile).Err()
		}
		// Convert e.g. "foo/bar;pkgname" => "foo/bar".
		if idx := strings.LastIndex(goPackage, ";"); idx != -1 {
			goPackage = goPackage[:idx]
		}
		goPackages.Add(goPackage)

		// A file that protoc must have generated for us.
		goFile := filepath.Join(
			inputs.OutputDir,
			filepath.FromSlash(goPackage),
			strings.TrimSuffix(protoFile, ".proto")+".pb.go",
		)
		if _, err := os.Stat(goFile); err != nil {
			return errors.Reason("could not find *.pb.go file generated from %q, is go_package option correct?", protoFile).Err()
		}

		// Transform .go files by adding pRPC stubs after gPRC stubs. Code generated
		// by protoc-gen-go-grpc plugin doesn't need this, since it uses interfaces
		// in the generated code (that pRPC implements) instead of concrete gRPC
		// types.
		if !*disableGRPC && !*useGRPCPlugin {
			var t transformer
			if err := t.transformGoFile(goFile); err != nil {
				return errors.Annotate(err, "could not transform %q", goFile).Err()
			}
		}

		// _test.proto's should go into the test package.
		if strings.HasSuffix(protoFile, "_test.proto") {
			newName := strings.TrimSuffix(goFile, ".go") + "_test.go"
			if err := os.Rename(goFile, newName); err != nil {
				return err
			}
		}
	}

	if !*disableGRPC && *withDiscovery {
		// We support generating a discovery file only when all generated *.pb.go
		// ended up in the same Go package. Otherwise it's not clear what package to
		// put the pb.discovery.go into.
		if goPackages.Len() != 1 {
			return errors.Reason(
				"cannot generate pb.discovery.go: generated *.pb.go files are in multiple packages %v",
				goPackages.ToSortedSlice(),
			).Err()
		}
		goPkg := goPackages.ToSlice()[0]
		out := filepath.Join(
			inputs.OutputDir,
			filepath.FromSlash(goPkg),
			"pb.discovery.go",
		)
		if err := genDiscoveryFile(out, goPkg, generatedDesc, rawDesc); err != nil {
			return err
		}
	}

	return nil
}

// loadDescriptorSet reads and parses FileDescriptorSet proto.
//
// Returns it as a map: *.proto path in the registry => FileDescriptorProto,
// as well as a raw byte blob.
func loadDescriptorSet(path string) (map[string]*descriptorpb.FileDescriptorProto, []byte, error) {
	blob, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	set := &descriptorpb.FileDescriptorSet{}
	if proto.Unmarshal(blob, set); err != nil {
		return nil, nil, err
	}
	mapping := make(map[string]*descriptorpb.FileDescriptorProto, len(set.File))
	for _, f := range set.File {
		mapping[f.GetName()] = f
	}
	return mapping, blob, nil
}

func setupLogging(ctx context.Context) context.Context {
	lvl := logging.Warning
	if *verbose {
		lvl = logging.Debug
	}
	return logging.SetLevel(gologger.StdConfig.Use(context.Background()), lvl)
}

func usage() {
	fmt.Fprintln(os.Stderr,
		`Compiles all .proto files in a directory to .go with grpc+prpc support.
usage: cproto [flags] [dir]

Flags:`)
	flag.PrintDefaults()
}

func main() {
	flag.Var(
		&protoImportPaths,
		"proto-path",
		"additional proto import paths; "+
			"May be relative to CWD; "+
			"May be specified multiple times.")
	flag.Var(
		&goModules,
		"go-module",
		"make protos in the given module available in proto import path. "+
			"May be specified multiple times.")
	flag.Var(
		&pathMap,
		"map-package",
		"maps a proto path to a go package name. "+
			"May be specified multiple times.")
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() > 1 {
		flag.Usage()
		os.Exit(1)
	}
	dir := "."
	if flag.NArg() == 1 {
		dir = flag.Arg(0)
	}

	ctx := setupLogging(context.Background())
	if err := run(ctx, dir); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		exitCode := 1
		if rc, ok := exitcode.Get(err); ok {
			exitCode = rc
		}
		os.Exit(exitCode)
	}
}
