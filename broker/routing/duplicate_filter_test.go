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
	golook "github.com/ottenwbe/golook/broker/runtime/core"
	"github.com/satori/go.uuid"
)

var _ = Describe("The duplicate filter", func() {
	It("detects dulicates from the same system.", func() {

		d := newDuplicateMap()

		source1 := Source{1, golook.GolookSystem.UUID}
		source2 := Source{1, golook.GolookSystem.UUID}

		Expect(d.CheckForDuplicates(source1)).To(BeFalse())
		Expect(d.CheckForDuplicates(source2)).To(BeTrue())
	})

	It("ignores duplicated ids from different systems", func() {

		d := newDuplicateMap()

		source1 := Source{1, uuid.NewV4().String()}
		source2 := Source{1, uuid.NewV4().String()}

		Expect(d.CheckForDuplicates(source1)).To(BeFalse())
		Expect(d.CheckForDuplicates(source2)).To(BeFalse())
	})

	It("prunes the duplicate filters to ensure that it does not grow beyond a maximum length.", func() {

		d := newDuplicateMap()
		saveMaxLen := maxDuplicateMapLen
		maxDuplicateMapLen = 2

		for i := 1; i < 2*maxDuplicateMapLen; i++ {
			d.CheckForDuplicates(Source{i, golook.GolookSystem.UUID})
		}

		Expect(len(d.filters[golook.GolookSystem.UUID])).To(Equal(maxDuplicateMapLen))

		maxDuplicateMapLen = saveMaxLen
	})

})
