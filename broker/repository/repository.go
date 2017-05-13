//Copyright 2016-2017 Beate OttenwäldersysUUID
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

package repositories

import (
	"github.com/ottenwbe/golook/broker/models"
	golook "github.com/ottenwbe/golook/broker/runtime/core"
)

/*
Repository is the interface for all implementations of repositories that store files and system information
*/
type Repository interface {
	StoreSystem(sysUUID string, system *golook.System) bool
	GetSystem(sysUUID string) (*golook.System, bool)
	GetSystems() map[string]*golook.System
	DelSystem(sysUUID string) *golook.System
	UpdateFiles(sysUUID string, files map[string]map[string]*models.File) bool
	FindSystemAndFiles(findString string) map[string][]*models.File
	GetFiles(sysUUID string) map[string]map[string]*models.File
}
