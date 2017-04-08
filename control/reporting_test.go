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
package control

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ottenwbe/golook/routing"
)

var _ = Describe("The report service", func() {
	It("should call the golook routing with a given file", func() {
		runWithMockedGolookClient(func() {
			ReportFile(FILE_NAME)
			Expect(routing.GolookClient.(*MockGolookClient).visitDoPostFile).To(BeTrue())
		})
	})

	It("should NOT call the golook routing with a non existing file", func() {
		runWithMockedGolookClient(func() {
			ReportFile(FILE_NAME + ".abc")
			Expect(routing.GolookClient.(*MockGolookClient).visitDoPostFile).To(BeFalse())
		})
	})

	It("should call the golook routing for a given folder", func() {
		runWithMockedGolookClient(func() {
			ReportFolder(FOLDER_NAME)
			//Expect(routing.GolookClient.(*MockGolookClient).visitDoPostFile).To(BeTrue())
		})
	})

	It("should NOT call the golook routing with a non existing file", func() {
		runWithMockedGolookClient(func() {
			ReportFolder("no_folder")
			Expect(routing.GolookClient.(*MockGolookClient).visitDoPostFile).To(BeFalse())
		})
	})

	It("should call the golook routing with files from existing folder which replace reported files", func() {
		runWithMockedGolookClient(func() {
			ReportFolderR(FOLDER_NAME)
			Expect(routing.GolookClient.(*MockGolookClient).visitDoPutFiles).To(BeTrue())
		})
	})

	It("should NOT call the golook routing with files from existing folder when folder does not exist", func() {
		runWithMockedGolookClient(func() {
			ReportFolderR("no_folder")
			Expect(routing.GolookClient.(*MockGolookClient).visitDoPutFiles).To(BeFalse())
		})
	})
})