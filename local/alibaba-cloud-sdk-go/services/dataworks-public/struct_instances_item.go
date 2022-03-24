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

// InstancesItem is a nested struct in dataworks_public response
type InstancesItem struct {
	ModifyTime        int64  `json:"ModifyTime" xml:"ModifyTime"`
	ParamValues       string `json:"ParamValues" xml:"ParamValues"`
	DagId             int64  `json:"DagId" xml:"DagId"`
	DagType           string `json:"DagType" xml:"DagType"`
	FinishTime        int64  `json:"FinishTime" xml:"FinishTime"`
	BeginRunningTime  int64  `json:"BeginRunningTime" xml:"BeginRunningTime"`
	CycTime           int64  `json:"CycTime" xml:"CycTime"`
	BeginWaitResTime  int64  `json:"BeginWaitResTime" xml:"BeginWaitResTime"`
	CreateTime        int64  `json:"CreateTime" xml:"CreateTime"`
	InstanceId        int64  `json:"InstanceId" xml:"InstanceId"`
	BizDate           int64  `json:"BizDate" xml:"BizDate"`
	NodeName          string `json:"NodeName" xml:"NodeName"`
	NodeId            int64  `json:"NodeId" xml:"NodeId"`
	BeginWaitTimeTime int64  `json:"BeginWaitTimeTime" xml:"BeginWaitTimeTime"`
	ProjectId         int64  `json:"ProjectId" xml:"ProjectId"`
	Status            string `json:"Status" xml:"Status"`
	TaskType          string `json:"TaskType" xml:"TaskType"`
}
