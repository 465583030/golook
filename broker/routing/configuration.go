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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	peers           = "routing.peers"
	duplicateLength = "routing.duplicateMap.length"
)

/*
ApplyConfiguration applies the configuration of the routing layer
*/
func ApplyConfiguration() {
	log.Debug("Configure routing layer")
	defaultPeers = viper.GetStringSlice(peers)
	maxDuplicateMapLen = viper.GetInt(duplicateLength)
}

/*
InitConfiguration initializes the configuration of the routing layer
*/
func InitConfiguration() {
	viper.SetDefault(peers, []string{""})
	viper.SetDefault(duplicateLength, 100)
}
