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

// RunsItem is a nested struct in dataworks_public response
type RunsItem struct {
	AbsTime           int64  `json:"AbsTime" xml:"AbsTime"`
	BeginCast         int64  `json:"BeginCast" xml:"BeginCast"`
	BeginRunningTime  int64  `json:"BeginRunningTime" xml:"BeginRunningTime"`
	BeginWaitResTime  int64  `json:"BeginWaitResTime" xml:"BeginWaitResTime"`
	BeginWaitTimeTime int64  `json:"BeginWaitTimeTime" xml:"BeginWaitTimeTime"`
	Bizdate           int64  `json:"Bizdate" xml:"Bizdate"`
	CycTime           int64  `json:"CycTime" xml:"CycTime"`
	EndCast           int64  `json:"EndCast" xml:"EndCast"`
	FinishTime        int64  `json:"FinishTime" xml:"FinishTime"`
	InGroupId         int    `json:"InGroupId" xml:"InGroupId"`
	InstanceId        int64  `json:"InstanceId" xml:"InstanceId"`
	NodeId            int64  `json:"NodeId" xml:"NodeId"`
	NodeName          string `json:"NodeName" xml:"NodeName"`
	Owner             string `json:"Owner" xml:"Owner"`
	ProjectId         int64  `json:"ProjectId" xml:"ProjectId"`
	Status            string `json:"Status" xml:"Status"`
}
