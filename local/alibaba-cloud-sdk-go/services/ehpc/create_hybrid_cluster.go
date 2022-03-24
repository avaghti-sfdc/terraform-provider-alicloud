package ehpc

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateHybridCluster invokes the ehpc.CreateHybridCluster API synchronously
func (client *Client) CreateHybridCluster(request *CreateHybridClusterRequest) (response *CreateHybridClusterResponse, err error) {
	response = CreateCreateHybridClusterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateHybridClusterWithChan invokes the ehpc.CreateHybridCluster API asynchronously
func (client *Client) CreateHybridClusterWithChan(request *CreateHybridClusterRequest) (<-chan *CreateHybridClusterResponse, <-chan error) {
	responseChan := make(chan *CreateHybridClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateHybridCluster(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateHybridClusterWithCallback invokes the ehpc.CreateHybridCluster API asynchronously
func (client *Client) CreateHybridClusterWithCallback(request *CreateHybridClusterRequest, callback func(response *CreateHybridClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateHybridClusterResponse
		var err error
		defer close(result)
		response, err = client.CreateHybridCluster(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateHybridClusterRequest is the request struct for api CreateHybridCluster
type CreateHybridClusterRequest struct {
	*requests.RpcRequest
	KeyPairName                 string                                  `position:"Query" name:"KeyPairName"`
	MultiOs                     requests.Boolean                        `position:"Query" name:"MultiOs"`
	SecurityGroupName           string                                  `position:"Query" name:"SecurityGroupName"`
	OnPremiseVolumeRemotePath   string                                  `position:"Query" name:"OnPremiseVolumeRemotePath"`
	ImageOwnerAlias             string                                  `position:"Query" name:"ImageOwnerAlias"`
	ResourceGroupId             string                                  `position:"Query" name:"ResourceGroupId"`
	Password                    string                                  `position:"Query" name:"Password"`
	ComputeSpotPriceLimit       requests.Float                          `position:"Query" name:"ComputeSpotPriceLimit"`
	OnPremiseVolumeLocalPath    string                                  `position:"Query" name:"OnPremiseVolumeLocalPath"`
	RemoteDirectory             string                                  `position:"Query" name:"RemoteDirectory"`
	ComputeSpotStrategy         string                                  `position:"Query" name:"ComputeSpotStrategy"`
	PostInstallScript           *[]CreateHybridClusterPostInstallScript `position:"Query" name:"PostInstallScript"  type:"Repeated"`
	VSwitchId                   string                                  `position:"Query" name:"VSwitchId"`
	Domain                      string                                  `position:"Query" name:"Domain"`
	Name                        string                                  `position:"Query" name:"Name"`
	VolumeId                    string                                  `position:"Query" name:"VolumeId"`
	ZoneId                      string                                  `position:"Query" name:"ZoneId"`
	ImageId                     string                                  `position:"Query" name:"ImageId"`
	ClientToken                 string                                  `position:"Query" name:"ClientToken"`
	EhpcVersion                 string                                  `position:"Query" name:"EhpcVersion"`
	SecurityGroupId             string                                  `position:"Query" name:"SecurityGroupId"`
	Description                 string                                  `position:"Query" name:"Description"`
	EcsOrderComputeInstanceType string                                  `position:"Query" name:"EcsOrder.Compute.InstanceType"`
	JobQueue                    string                                  `position:"Query" name:"JobQueue"`
	VolumeType                  string                                  `position:"Query" name:"VolumeType"`
	OnPremiseVolumeMountPoint   string                                  `position:"Query" name:"OnPremiseVolumeMountPoint"`
	OnPremiseVolumeProtocol     string                                  `position:"Query" name:"OnPremiseVolumeProtocol"`
	VolumeProtocol              string                                  `position:"Query" name:"VolumeProtocol"`
	ClientVersion               string                                  `position:"Query" name:"ClientVersion"`
	OsTag                       string                                  `position:"Query" name:"OsTag"`
	Nodes                       *[]CreateHybridClusterNodes             `position:"Query" name:"Nodes"  type:"Repeated"`
	Application                 *[]CreateHybridClusterApplication       `position:"Query" name:"Application"  type:"Repeated"`
	VpcId                       string                                  `position:"Query" name:"VpcId"`
	VolumeMountpoint            string                                  `position:"Query" name:"VolumeMountpoint"`
	SchedulerPreInstall         requests.Boolean                        `position:"Query" name:"SchedulerPreInstall"`
	Location                    string                                  `position:"Query" name:"Location"`
}

// CreateHybridClusterPostInstallScript is a repeated param struct in CreateHybridClusterRequest
type CreateHybridClusterPostInstallScript struct {
	Args string `name:"Args"`
	Url  string `name:"Url"`
}

// CreateHybridClusterNodes is a repeated param struct in CreateHybridClusterRequest
type CreateHybridClusterNodes struct {
	IpAddress     string `name:"IpAddress"`
	HostName      string `name:"HostName"`
	Role          string `name:"Role"`
	AccountType   string `name:"AccountType"`
	SchedulerType string `name:"SchedulerType"`
}

// CreateHybridClusterApplication is a repeated param struct in CreateHybridClusterRequest
type CreateHybridClusterApplication struct {
	Tag string `name:"Tag"`
}

// CreateHybridClusterResponse is the response struct for api CreateHybridCluster
type CreateHybridClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ClusterId string `json:"ClusterId" xml:"ClusterId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateCreateHybridClusterRequest creates a request to invoke CreateHybridCluster API
func CreateCreateHybridClusterRequest() (request *CreateHybridClusterRequest) {
	request = &CreateHybridClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "CreateHybridCluster", "", "")
	request.Method = requests.GET
	return
}

// CreateCreateHybridClusterResponse creates a response to parse from CreateHybridCluster response
func CreateCreateHybridClusterResponse() (response *CreateHybridClusterResponse) {
	response = &CreateHybridClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
