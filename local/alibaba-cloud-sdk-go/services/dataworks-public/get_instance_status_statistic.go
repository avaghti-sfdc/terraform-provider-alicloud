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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetInstanceStatusStatistic invokes the dataworks_public.GetInstanceStatusStatistic API synchronously
func (client *Client) GetInstanceStatusStatistic(request *GetInstanceStatusStatisticRequest) (response *GetInstanceStatusStatisticResponse, err error) {
	response = CreateGetInstanceStatusStatisticResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceStatusStatisticWithChan invokes the dataworks_public.GetInstanceStatusStatistic API asynchronously
func (client *Client) GetInstanceStatusStatisticWithChan(request *GetInstanceStatusStatisticRequest) (<-chan *GetInstanceStatusStatisticResponse, <-chan error) {
	responseChan := make(chan *GetInstanceStatusStatisticResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceStatusStatistic(request)
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

// GetInstanceStatusStatisticWithCallback invokes the dataworks_public.GetInstanceStatusStatistic API asynchronously
func (client *Client) GetInstanceStatusStatisticWithCallback(request *GetInstanceStatusStatisticRequest, callback func(response *GetInstanceStatusStatisticResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceStatusStatisticResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceStatusStatistic(request)
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

// GetInstanceStatusStatisticRequest is the request struct for api GetInstanceStatusStatistic
type GetInstanceStatusStatisticRequest struct {
	*requests.RpcRequest
	ProjectEnv string           `position:"Body" name:"ProjectEnv"`
	BizDate    string           `position:"Body" name:"BizDate"`
	ProjectId  requests.Integer `position:"Body" name:"ProjectId"`
}

// GetInstanceStatusStatisticResponse is the response struct for api GetInstanceStatusStatistic
type GetInstanceStatusStatisticResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	StatusCount StatusCount `json:"StatusCount" xml:"StatusCount"`
}

// CreateGetInstanceStatusStatisticRequest creates a request to invoke GetInstanceStatusStatistic API
func CreateGetInstanceStatusStatisticRequest() (request *GetInstanceStatusStatisticRequest) {
	request = &GetInstanceStatusStatisticRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetInstanceStatusStatistic", "", "")
	request.Method = requests.POST
	return
}

// CreateGetInstanceStatusStatisticResponse creates a response to parse from GetInstanceStatusStatistic response
func CreateGetInstanceStatusStatisticResponse() (response *GetInstanceStatusStatisticResponse) {
	response = &GetInstanceStatusStatisticResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
