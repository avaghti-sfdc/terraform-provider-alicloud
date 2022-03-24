package opensearch

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

// ResultInModifyAppGroup is a nested struct in opensearch response
type ResultInModifyAppGroup struct {
	Id                                string `json:"id" xml:"id"`
	Name                              string `json:"name" xml:"name"`
	CurrentVersion                    string `json:"currentVersion" xml:"currentVersion"`
	SwitchedTime                      int    `json:"switchedTime" xml:"switchedTime"`
	ChargingWay                       int    `json:"chargingWay" xml:"chargingWay"`
	Type                              string `json:"type" xml:"type"`
	ProjectId                         string `json:"projectId" xml:"projectId"`
	ChargeType                        string `json:"chargeType" xml:"chargeType"`
	ExpireOn                          string `json:"expireOn" xml:"expireOn"`
	InstanceId                        string `json:"instanceId" xml:"instanceId"`
	CommodityCode                     string `json:"commodityCode" xml:"commodityCode"`
	ProcessingOrderId                 string `json:"processingOrderId" xml:"processingOrderId"`
	FirstRankAlgoDeploymentId         int    `json:"firstRankAlgoDeploymentId" xml:"firstRankAlgoDeploymentId"`
	SecondRankAlgoDeploymentId        int    `json:"secondRankAlgoDeploymentId" xml:"secondRankAlgoDeploymentId"`
	PendingSecondRankAlgoDeploymentId int    `json:"pendingSecondRankAlgoDeploymentId" xml:"pendingSecondRankAlgoDeploymentId"`
	Description                       string `json:"description" xml:"description"`
	Produced                          int    `json:"produced" xml:"produced"`
	LockedByExpiration                int    `json:"lockedByExpiration" xml:"lockedByExpiration"`
	HasPendingQuotaReviewTask         int    `json:"hasPendingQuotaReviewTask" xml:"hasPendingQuotaReviewTask"`
	Created                           int    `json:"created" xml:"created"`
	Updated                           int    `json:"updated" xml:"updated"`
	Status                            string `json:"status" xml:"status"`
	LockMode                          string `json:"lockMode" xml:"lockMode"`
	Domain                            string `json:"domain" xml:"domain"`
	Quota                             Quota  `json:"quota" xml:"quota"`
}
