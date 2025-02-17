package ccc

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

// DataItem is a nested struct in ccc response
type DataItem struct {
	Name       string                            `json:"Name" xml:"Name"`
	StateCode  string                            `json:"StateCode" xml:"StateCode"`
	StartTime  int64                             `json:"StartTime" xml:"StartTime"`
	State      string                            `json:"State" xml:"State"`
	InstanceId string                            `json:"InstanceId" xml:"InstanceId"`
	Scope      string                            `json:"Scope" xml:"Scope"`
	StatsTime  int64                             `json:"StatsTime" xml:"StatsTime"`
	Duration   int64                             `json:"Duration" xml:"Duration"`
	Inbound    Inbound                           `json:"Inbound" xml:"Inbound"`
	Outbound   OutboundInListIntervalAgentReport `json:"Outbound" xml:"Outbound"`
	Overall    Overall                           `json:"Overall" xml:"Overall"`
}
