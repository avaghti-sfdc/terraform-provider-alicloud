package unimkt

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

// AdSlot is a nested struct in unimkt response
type AdSlot struct {
	CreateTime            int64  `json:"CreateTime" xml:"CreateTime"`
	ModifyTime            int64  `json:"ModifyTime" xml:"ModifyTime"`
	TenantId              string `json:"TenantId" xml:"TenantId"`
	MediaName             string `json:"MediaName" xml:"MediaName"`
	MediaId               string `json:"MediaId" xml:"MediaId"`
	AdSlotId              string `json:"AdSlotId" xml:"AdSlotId"`
	AdSlotName            string `json:"AdSlotName" xml:"AdSlotName"`
	AdSlotType            string `json:"AdSlotType" xml:"AdSlotType"`
	AdSlotTemplateId      string `json:"AdSlotTemplateId" xml:"AdSlotTemplateId"`
	AdSlotStatus          string `json:"AdSlotStatus" xml:"AdSlotStatus"`
	AdSlotCorporateStatus string `json:"AdSlotCorporateStatus" xml:"AdSlotCorporateStatus"`
	ExtInfo               string `json:"ExtInfo" xml:"ExtInfo"`
	BlockingRule          string `json:"BlockingRule" xml:"BlockingRule"`
	InspireScene          string `json:"InspireScene" xml:"InspireScene"`
	Version               int64  `json:"Version" xml:"Version"`
}
