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
package server

import (
	"encoding/json"
	"io"
	"time"
)

type File struct {
	Name     string    `json:"name"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Accessed time.Time `json:"accessed"`
}

type System struct {
	Name  string `json:"name"`
	OS    string `json:"os"`
	IP    string `json:"ip"`
	UUID  string `json:"uuid"`
	Files []File `json:"files"`
}

func DecodeFiles(fileReader io.Reader) ([]File, error) {
	files := make([]File, 0)
	err := json.NewDecoder(fileReader).Decode(&files)
	return files, err
}

func DecodeFile(fileReader io.Reader) (File, error) {
	var file File
	err := json.NewDecoder(fileReader).Decode(&file)
	return file, err
}

func DecodeSystem(sysReader io.Reader) (System, error) {
	var sys System
	err := json.NewDecoder(sysReader).Decode(&sys)
	return sys, err
}
