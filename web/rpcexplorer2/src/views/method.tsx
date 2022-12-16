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

import { useParams } from 'react-router-dom';

import { useGlobals } from '../context/globals';

const Method = () => {
  const { serviceName, methodName } = useParams();
  const { descriptors, oauthClient } = useGlobals();

  const svc = descriptors.service(serviceName ?? 'unknown');
  if (svc === null) {
    return <p>No such service</p>;
  }
  const method = svc.method(methodName ?? 'unknown');
  if (method === null) {
    return <p>No such method</p>;
  }

  return (
    <>
      <p>A method {methodName} of {serviceName}</p>
      <p>{method.help}</p>
      <p>OAuth client ID: {oauthClient.clientId}</p>
    </>
  );
};

export default Method;
