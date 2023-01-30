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

// Command svcdec stands for 'service decorator'. It reads a service interface
// XYZServer generated by protoc and generates DecoratedXYZ with this structure:
//
//	type DecoratedXYZ struct {
//	  Service XYZServer
//	  Prelude func(ctx context.Context, methodName string, req proto.Message) (context.Context, error)
//	}
//
// DecoratedXYZ has the same methods as XYZServer: they call Prelude before
// forwarding the call to the corresponding XYZServer method.
//
// svcdec is designed to be run through go generate:
//
//	//go:generate svcdec -type GreetServer
package main

import (
	"context"
	"os"
	"strings"

	"go.chromium.org/luci/grpc/internal/svctool"
)

func generate(ctx context.Context, a *svctool.GeneratorArgs) error {
	args := templateArgs{
		PackageName:  a.PackageName,
		ExtraImports: a.ExtraImports,
	}
	for _, svc := range a.Services {
		args.Services = append(args.Services, &service{
			Service:    svc,
			StructName: "Decorated" + strings.TrimSuffix(svc.TypeName, "Server"),
		})
	}

	// Execute template.
	return tmpl.Execute(a.Out, args)
}

func tool() *svctool.Tool {
	return &svctool.Tool{Name: "svcdec", OutputFilenameSuffix: "dec"}
}

func main() {
	tool().Main(os.Args[1:], generate)
}
