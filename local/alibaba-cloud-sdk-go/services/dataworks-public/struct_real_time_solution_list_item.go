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

// RealTimeSolutionListItem is a nested struct in dataworks_public response
type RealTimeSolutionListItem struct {
	ProcessId   int64  `json:"ProcessId" xml:"ProcessId"`
	ProcessName string `json:"ProcessName" xml:"ProcessName"`
	TaskStatus  string `json:"TaskStatus" xml:"TaskStatus"`
	TaskType    string `json:"TaskType" xml:"TaskType"`
	CreateTime  int64  `json:"CreateTime" xml:"CreateTime"`
}
