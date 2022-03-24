package schedulerx2

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

// JobInstanceDetailsItem is a nested struct in schedulerx2 response
type JobInstanceDetailsItem struct {
	InstanceId   int64  `json:"InstanceId" xml:"InstanceId"`
	JobId        int64  `json:"JobId" xml:"JobId"`
	Status       int    `json:"Status" xml:"Status"`
	StartTime    string `json:"StartTime" xml:"StartTime"`
	EndTime      string `json:"EndTime" xml:"EndTime"`
	ScheduleTime string `json:"ScheduleTime" xml:"ScheduleTime"`
	DataTime     string `json:"DataTime" xml:"DataTime"`
	Executor     string `json:"Executor" xml:"Executor"`
	WorkAddr     string `json:"WorkAddr" xml:"WorkAddr"`
	Result       string `json:"Result" xml:"Result"`
	Progress     string `json:"Progress" xml:"Progress"`
	TimeType     int    `json:"TimeType" xml:"TimeType"`
	TriggerType  int    `json:"TriggerType" xml:"TriggerType"`
}
