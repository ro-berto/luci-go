package main

// This file generates descriptor.html. Run itself:
//go:generate go run gen.go

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func run() error {
	descFile, err := ioutil.TempFile("", "desc")
	if err != nil {
		return err
	}

	protoc := exec.Command(
		"protoc",
		"--descriptor_set_out="+descFile.Name(),
		"--include_source_info",
		"types.proto")
	protoc.Stdout = os.Stdout
	protoc.Stderr = os.Stderr
	if err := protoc.Run(); err != nil {
		return fmt.Errorf("protoc: %s", err)
	}

	descBytes, err := ioutil.ReadAll(descFile)
	if err != nil {
		return err
	}
	if len(descBytes) == 0 {
		return fmt.Errorf("desc file was read as empty")
	}

	desc := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(descBytes, desc); err != nil {
		return fmt.Errorf("could not read descriptor file: %s", err)
	}

	out, err := os.Create("descriptor.html")
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(out, `<!--
  Copyright 2016 The LUCI Authors.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
  -->

<!-- This file is generated by gen.go -->

<script>
var testDescriptor = `)
	if err != nil {
		return err
	}

	m := jsonpb.Marshaler{}
	m.Indent = "  "
	if err := m.Marshal(out, desc); err != nil {
		return err
	}

	_, err = fmt.Fprintln(out, ";\n</script>")
	return err
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
