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

package service

import (
	"github.com/ottenwbe/golook/broker/routing"
	golook "github.com/ottenwbe/golook/broker/runtime/core"
)

/*
router dispatches messages. To this end, it implements an embedded routing.router.
router informs the embedded routing.router about possible candidates for peers by registering to the system service's callbacks.
*/
type router struct {
	routing.Router
}

func newRouter(name string, routerType routing.RouterType) *router {
	r := &router{routing.NewRouter(name, routerType)}
	routing.ActivateRouter(r)
	changedSystemCallbacks.Add(name, r.handleNewSystem)
	delSystemCallbacks.Add(name, r.handleDelSystem)
	return r
}

func (r *router) close() {
	routing.DeactivateRouter(r)
	changedSystemCallbacks.Delete(r.Router.Name())
}

func (r *router) handleDelSystem(uuid string, system *golook.System) {
	r.DelPeer(routing.NewKey(system.UUID))
}

func (r *router) handleNewSystem(uuid string, systems map[string]*golook.System) {
	for _, s := range systems {
		// avoid self references
		if s.UUID != golook.GolookSystem.UUID {
			r.NewPeer(routing.NewKey(s.UUID), s.IP)
		}
	}
}
