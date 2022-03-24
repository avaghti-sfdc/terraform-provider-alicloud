package dataworks_public

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
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// InfluencesItem is a nested struct in dataworks_public response
type InfluencesItem struct {
	BaselineId   int64  `json:"BaselineId" xml:"BaselineId"`
	Bizdate      int64  `json:"Bizdate" xml:"Bizdate"`
	InGroupId    int    `json:"InGroupId" xml:"InGroupId"`
	BaselineName string `json:"BaselineName" xml:"BaselineName"`
	Owner        string `json:"Owner" xml:"Owner"`
	Status       string `json:"Status" xml:"Status"`
	ProjectId    int64  `json:"ProjectId" xml:"ProjectId"`
	Priority     int    `json:"Priority" xml:"Priority"`
	Buffer       int64  `json:"Buffer" xml:"Buffer"`
}
