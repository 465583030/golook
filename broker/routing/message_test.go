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

package routing

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ottenwbe/golook/utils"
)

var _ = Describe("The params", func() {
	It("can be unmarshalled from json.", func() {
		var expectedResult = "test message"

		m := Params(utils.MarshalSD(&expectedResult))

		var result string
		m.Unmarshal(&result)

		Expect(result).To(Equal(expectedResult))
	})
})

var _ = Describe("The encapsulated message", func() {
	It("should comprise a method name and the content after its creation", func() {
		m, err := NewRequestMessage(NilKey(), 1, "method", "msg")

		Expect(err).To(BeNil())
		Expect(m.Method).To(Equal("method"))
		Expect(len(m.Params)).ToNot(BeZero())
		Expect(string(m.Params)).To(ContainSubstring("msg"))
	})

	It("should support to get the encapsulated method", func() {
		m, err := NewRequestMessage(NilKey(), 1, "method", "msg")

		var s string
		m.GetEncapsulated(&s)

		Expect(err).To(BeNil())
		Expect(s).To(Equal("msg"))
	})
})

var _ = Describe("The response message", func() {
	It("should set created", func() {
		expectedParams := utils.MarshalSD("testParams")

		m, err := NewResponseMessage(Source{ID: 123}, "testParams")

		Expect(err).To(BeNil())
		Expect(m.RequestSrc.ID).To(Equal(123))
		Expect(m.Params).To(Equal(Params(expectedParams)))
	})

	It("should return an error when the params cannot be marshalled.", func() {
		_, err := NewResponseMessage(Source{ID: 123}, make(chan bool))
		Expect(err).ToNot(BeNil())
	})
})
