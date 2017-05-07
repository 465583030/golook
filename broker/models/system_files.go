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

package models

import (
	golook "github.com/ottenwbe/golook/broker/runtime/core"
)

/*
SystemFiles is a wrapper around files and systems in a common data structure
*/
type SystemFiles struct {
	System *golook.System   `json:"system"`
	Files  map[string]*File `json:"files"`
}