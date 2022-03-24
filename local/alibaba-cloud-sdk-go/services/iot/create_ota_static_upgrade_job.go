package iot

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

// CreateOTAStaticUpgradeJob invokes the iot.CreateOTAStaticUpgradeJob API synchronously
func (client *Client) CreateOTAStaticUpgradeJob(request *CreateOTAStaticUpgradeJobRequest) (response *CreateOTAStaticUpgradeJobResponse, err error) {
	response = CreateCreateOTAStaticUpgradeJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOTAStaticUpgradeJobWithChan invokes the iot.CreateOTAStaticUpgradeJob API asynchronously
func (client *Client) CreateOTAStaticUpgradeJobWithChan(request *CreateOTAStaticUpgradeJobRequest) (<-chan *CreateOTAStaticUpgradeJobResponse, <-chan error) {
	responseChan := make(chan *CreateOTAStaticUpgradeJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOTAStaticUpgradeJob(request)
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

// CreateOTAStaticUpgradeJobWithCallback invokes the iot.CreateOTAStaticUpgradeJob API asynchronously
func (client *Client) CreateOTAStaticUpgradeJobWithCallback(request *CreateOTAStaticUpgradeJobRequest, callback func(response *CreateOTAStaticUpgradeJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOTAStaticUpgradeJobResponse
		var err error
		defer close(result)
		response, err = client.CreateOTAStaticUpgradeJob(request)
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

// CreateOTAStaticUpgradeJobRequest is the request struct for api CreateOTAStaticUpgradeJob
type CreateOTAStaticUpgradeJobRequest struct {
	*requests.RpcRequest
	RetryCount         requests.Integer                `position:"Query" name:"RetryCount"`
	TimeoutInMinutes   requests.Integer                `position:"Query" name:"TimeoutInMinutes"`
	NeedConfirm        requests.Boolean                `position:"Query" name:"NeedConfirm"`
	NeedPush           requests.Boolean                `position:"Query" name:"NeedPush"`
	IotInstanceId      string                          `position:"Query" name:"IotInstanceId"`
	TargetSelection    string                          `position:"Query" name:"TargetSelection"`
	ScheduleFinishTime requests.Integer                `position:"Query" name:"ScheduleFinishTime"`
	Tag                *[]CreateOTAStaticUpgradeJobTag `position:"Query" name:"Tag"  type:"Repeated"`
	GrayPercent        string                          `position:"Query" name:"GrayPercent"`
	DnListFileUrl      string                          `position:"Query" name:"DnListFileUrl"`
	FirmwareId         string                          `position:"Query" name:"FirmwareId"`
	ProductKey         string                          `position:"Query" name:"ProductKey"`
	RetryInterval      requests.Integer                `position:"Query" name:"RetryInterval"`
	SrcVersion         *[]string                       `position:"Query" name:"SrcVersion"  type:"Repeated"`
	ScheduleTime       requests.Integer                `position:"Query" name:"ScheduleTime"`
	OverwriteMode      requests.Integer                `position:"Query" name:"OverwriteMode"`
	ApiProduct         string                          `position:"Body" name:"ApiProduct"`
	ApiRevision        string                          `position:"Body" name:"ApiRevision"`
	MaximumPerMinute   requests.Integer                `position:"Query" name:"MaximumPerMinute"`
	TargetDeviceName   *[]string                       `position:"Query" name:"TargetDeviceName"  type:"Repeated"`
}

// CreateOTAStaticUpgradeJobTag is a repeated param struct in CreateOTAStaticUpgradeJobRequest
type CreateOTAStaticUpgradeJobTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateOTAStaticUpgradeJobResponse is the response struct for api CreateOTAStaticUpgradeJob
type CreateOTAStaticUpgradeJobResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateCreateOTAStaticUpgradeJobRequest creates a request to invoke CreateOTAStaticUpgradeJob API
func CreateCreateOTAStaticUpgradeJobRequest() (request *CreateOTAStaticUpgradeJobRequest) {
	request = &CreateOTAStaticUpgradeJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateOTAStaticUpgradeJob", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateOTAStaticUpgradeJobResponse creates a response to parse from CreateOTAStaticUpgradeJob response
func CreateCreateOTAStaticUpgradeJobResponse() (response *CreateOTAStaticUpgradeJobResponse) {
	response = &CreateOTAStaticUpgradeJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
