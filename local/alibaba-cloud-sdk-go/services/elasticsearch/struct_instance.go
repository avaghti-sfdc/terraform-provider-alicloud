package elasticsearch

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

// Instance is a nested struct in elasticsearch response
type Instance struct {
	DedicateMaster               bool                         `json:"dedicateMaster" xml:"dedicateMaster"`
	ServiceVpc                   bool                         `json:"serviceVpc" xml:"serviceVpc"`
	UpdatedAt                    string                       `json:"updatedAt" xml:"updatedAt"`
	NodeAmount                   int                          `json:"nodeAmount" xml:"nodeAmount"`
	InstanceId                   string                       `json:"instanceId" xml:"instanceId"`
	Status                       string                       `json:"status" xml:"status"`
	Description                  string                       `json:"description" xml:"description"`
	AdvancedDedicateMaster       bool                         `json:"advancedDedicateMaster" xml:"advancedDedicateMaster"`
	IsNewDeployment              string                       `json:"isNewDeployment" xml:"isNewDeployment"`
	EsVersion                    string                       `json:"esVersion" xml:"esVersion"`
	PostpaidServiceStatus        string                       `json:"postpaidServiceStatus" xml:"postpaidServiceStatus"`
	PaymentType                  string                       `json:"paymentType" xml:"paymentType"`
	ResourceGroupId              string                       `json:"resourceGroupId" xml:"resourceGroupId"`
	CreatedAt                    string                       `json:"createdAt" xml:"createdAt"`
	ExtendConfigs                []map[string]interface{}     `json:"extendConfigs" xml:"extendConfigs"`
	NodeSpec                     NodeSpec                     `json:"nodeSpec" xml:"nodeSpec"`
	ClientNodeConfiguration      ClientNodeConfiguration      `json:"clientNodeConfiguration" xml:"clientNodeConfiguration"`
	KibanaConfiguration          KibanaConfiguration          `json:"kibanaConfiguration" xml:"kibanaConfiguration"`
	NetworkConfig                NetworkConfig                `json:"networkConfig" xml:"networkConfig"`
	MasterConfiguration          MasterConfiguration          `json:"masterConfiguration" xml:"masterConfiguration"`
	ElasticDataNodeConfiguration ElasticDataNodeConfiguration `json:"elasticDataNodeConfiguration" xml:"elasticDataNodeConfiguration"`
	Tags                         []Tag                        `json:"tags" xml:"tags"`
}
