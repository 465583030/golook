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
package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The configuration", func() {
	Context("default", func() {
		It("should return the default host", func() {
			Expect(Host()).To(Equal("http://127.0.0.1"))
		})
		It("should return the default port", func() {
			Expect(ServerPort()).To(Equal(8080))
		})
		It("should return the default detatch modus", func() {
			Expect(RunDetatched()).To(BeFalse())
		})
	})
})
