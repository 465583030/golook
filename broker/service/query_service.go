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
	. "github.com/ottenwbe/golook/broker/models"
	. "github.com/ottenwbe/golook/broker/routing"
)

type QueryService interface {
	MakeFileQuery(searchString string) interface{}
}

func NewQueryService() QueryService {
	return &queryService{}
}

type queryService struct{}

func (*queryService) MakeFileQuery(searchString string) interface{} {
	fq := FileQueryTransfer{SearchString: searchString}
	return systemIndex.Route(SysKey(), FILE_QUERY, fq)
}