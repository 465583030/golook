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
package cmd

import (
	"fmt"
	"github.com/ottenwbe/golook/broker/api"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: fmt.Sprint("Print some information about the API."),
	Long:  fmt.Sprint("Print the supported versions of the API."),
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(fmt.Sprintf("[%s]", api.GolookAPIVersion))
	},
}

func init() {
	RootCmd.AddCommand(apiCmd)
}