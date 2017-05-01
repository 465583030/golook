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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	repo "github.com/ottenwbe/golook/broker/repository"
	"github.com/ottenwbe/golook/broker/routing"
	golook "github.com/ottenwbe/golook/broker/runtime"
)

var _ = Describe("The system service", func() {

	It("stores and forwards the system's information", func() {
		tmp := routing.Router(broadCastRouter)
		routing.RunWithMockedRouter(&tmp, func() {
			sysUUID := golook.GolookSystem.UUID
			repo.GoLookRepository.DelSystem(sysUUID)
			s := SystemService{}

			s.Run()

			_, isSystemInRepo := repo.GoLookRepository.GetSystem(sysUUID)
			Expect(isSystemInRepo).To(BeTrue())
			Expect(routing.AccessMockedRouter(broadCastRouter).Visited).To(Equal(1))
		})
	})

})
