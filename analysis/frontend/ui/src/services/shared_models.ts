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

import { GrpcError, RpcCode } from '@chopsui/prpc-client';

// Contains data models shared between multiple services.

export interface ClusterId {
  algorithm: string;
  id: string;
}

export interface AssociatedBug {
  system: string;
  id: string;
  linkText: string;
  url: string;
}

export interface Changelist {
  // Gerrit hostname, e.g. "chromium-review.googlesource.com".
  host: string;

  // Change number, encoded as a string, e.g. "12345".
  change: string;

  // Patchset number, e.g. 1.
  patchset: number;
}

export function isRetriable(e: GrpcError) : boolean {
  // The following codes indicate transient errors that are retriable. See:
  // https://source.chromium.org/chromium/infra/infra/+/main:go/src/go.chromium.org/luci/grpc/grpcutil/errors.go;l=176?q=codeToStatus&type=cs
  switch (e.code) {
    case RpcCode.INTERNAL:
      return true;
    case RpcCode.UNKNOWN:
      return true;
    case RpcCode.UNAVAILABLE:
      return true;
  }
  return false;
}

export function prpcRetrier(_failureCount: number, e: Error) : boolean {
  if (e instanceof GrpcError && isRetriable(e)) {
    return true;
  }
  return false;
}
