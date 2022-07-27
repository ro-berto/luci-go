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

import { RpcCode } from '@chopsui/prpc-client';
import { assert } from 'chai';
import * as sinon from 'sinon';

import { PrpcClientExt } from './prpc_client_ext';

describe('PrpcClientExt', () => {
  it('should grab access token from getAccessToken', async () => {
    let accessToken = '1';
    const fetchStub = sinon.stub<[URL | RequestInfo, RequestInit | undefined], Promise<Response>>();
    fetchStub.onCall(0).resolves(new Response(")]}'\n{}", { headers: { 'X-Prpc-Grpc-Code': RpcCode.OK.toString() } }));
    fetchStub.onCall(1).resolves(new Response(")]}'\n{}", { headers: { 'X-Prpc-Grpc-Code': RpcCode.OK.toString() } }));

    const client = new PrpcClientExt({ fetchImpl: fetchStub }, () => accessToken);
    await client.call('service', 'method', {});
    const req1 = new Request(...fetchStub.getCall(0).args);
    assert.strictEqual(req1.headers.get('Authorization'), 'Bearer 1');

    accessToken = '2';
    await client.call('service', 'method', {});
    const req2 = new Request(...fetchStub.getCall(1).args);
    assert.strictEqual(req2.headers.get('Authorization'), 'Bearer 2');
  });

  it('should not override additional header', async () => {
    const accessToken = '1';
    const fetchStub = sinon.stub<[URL | RequestInfo, RequestInit | undefined], Promise<Response>>();
    fetchStub.onCall(0).resolves(new Response(")]}'\n{}", { headers: { 'X-Prpc-Grpc-Code': RpcCode.OK.toString() } }));
    fetchStub.onCall(1).resolves(new Response(")]}'\n{}", { headers: { 'X-Prpc-Grpc-Code': RpcCode.OK.toString() } }));

    const client = new PrpcClientExt({ fetchImpl: fetchStub }, () => accessToken);
    await client.call('service', 'method', {}, { Authorization: 'additional-header' });
    const req1 = new Request(...fetchStub.getCall(0).args);
    assert.strictEqual(req1.headers.get('Authorization'), 'additional-header');
  });
});
