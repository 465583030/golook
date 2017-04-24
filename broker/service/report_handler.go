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
	. "github.com/ottenwbe/golook/broker/repository"
	. "github.com/ottenwbe/golook/broker/utils"
	log "github.com/sirupsen/logrus"
)

const (
	FILE_REPORT = "file_report"
)

func handleFileReport(params interface{}) interface{} {
	var fileMessage PeerFileReport
	if err := UnmarshalB(params, &fileMessage); err == nil {
		GoLookRepository.StoreFiles(fileMessage.System, fileMessage.Files)
	} else {
		log.WithError(err).Error("Could not handle file report.")
	}
	return nil
}

func init() {
	systemIndex.HandlerFunction(FILE_REPORT, handleFileReport)
}