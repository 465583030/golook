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
package client

import "testing"

func TestHost(t *testing.T) {
	if Host() != "localhost" {
		t.Log("Default configuration is not set")
	}
}

func TestServerPort(t *testing.T) {
	if ServerPort() == 8080 {
		t.Log("Default port is not set")
	}
}