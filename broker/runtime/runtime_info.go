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
package runtime

import (
	"encoding/json"
	"github.com/sirupsen/logrus"

	. "github.com/ottenwbe/golook/broker/models"
)

const (
	GOLOOK_NAME = "golook"
	VERSION     = "v0.1.0-dev"
)

func NewAppInfo() *AppInfo {
	return &AppInfo{
		App:     GOLOOK_NAME,
		Version: VERSION,
		System:  NewSystem(),
	}
}

func EncodeAppInfo(info *AppInfo) string {
	b, err := json.Marshal(info)
	if err != nil {
		logrus.WithError(err).Error("Could not encode app info.")
		return "{}"
	}
	return string(b)
}