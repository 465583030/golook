//Copyright 2016-2017 Beate Ottenwälder
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package communication

import (
	"fmt"

	. "github.com/ottenwbe/golook/broker/models"

	"github.com/ottenwbe/golook/utils"
	log "github.com/sirupsen/logrus"
	"github.com/ybbus/jsonrpc"
)

type (
	/*JsonRpcClientStub implements the RPCClients interface*/
	JsonRpcClientStub struct {
		serverUrl string
		c         *jsonrpc.RPCClient
	}
	/*JsonRPCReturnType implements the EncapsulatedValues interface*/
	JsonRPCReturnType struct {
		response *jsonrpc.RPCResponse
	}
)

func newJsonRPCClient(url string, port int) RPCClient {

	serverUrl := fmt.Sprintf("http://%s", url)
	if port >= 0 {
		serverUrl = fmt.Sprintf("%s:%d", serverUrl, port)
	}

	return &JsonRpcClientStub{
		serverUrl: serverUrl,
		c:         jsonrpc.NewJsonRPCClient(fmt.Sprintf("%s/rpc", serverUrl)),
	}
}

/*
Call executes a RPC call with the given method name and the given parameters.
It returns a generic return value. Clients can retrieve the result by calling the Unmarshal method on the result.
*/
func (lc *JsonRpcClientStub) Call(method string, parameters interface{}) (EncapsulatedValues, error) {

	log.WithField("method", method).WithField("com", "jsonrpc").Debugf("Making a call for %s with %s", method, utils.MarshalSD(parameters))

	response, err := lc.c.Call(method, parameters)
	if err != nil {
		return nil, err
	}

	r := &JsonRPCReturnType{response: response}
	log.WithField("method", method).WithField("com", "jsonrpc").Debugf("Getting a response for %s with %s", method, r.response)

	return r, nil
}

/*
URL returns the url of the RPC server to which this client connects
*/
func (lc *JsonRpcClientStub) URL() string {
	return lc.serverUrl
}

/*
Unmarshal allows callers of the RPC client to unmarshal the result retrieved by the client
*/
func (rt *JsonRPCReturnType) Unmarshal(v interface{}) error {
	return rt.response.GetObject(v)
}
