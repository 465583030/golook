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
	"reflect"
)

var _ = Describe("The file services", func() {
	It("create a pair of report and query service", func() {
		OpenFileServices(BroadcastFiles)
		r, q := ReportService, QueryService

		Expect(r).ToNot(BeNil())
		Expect(q).ToNot(BeNil())

		Expect(reflect.TypeOf(r).String()).To(Equal(reflect.TypeOf(&broadcastReportService{}).String()))
		Expect(reflect.TypeOf(q).String()).To(Equal(reflect.TypeOf(&localQueryService{}).String()))
	})
})
