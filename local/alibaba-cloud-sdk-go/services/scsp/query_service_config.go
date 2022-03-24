package scsp

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

// QueryServiceConfig invokes the scsp.QueryServiceConfig API synchronously
func (client *Client) QueryServiceConfig(request *QueryServiceConfigRequest) (response *QueryServiceConfigResponse, err error) {
	response = CreateQueryServiceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// QueryServiceConfigWithChan invokes the scsp.QueryServiceConfig API asynchronously
func (client *Client) QueryServiceConfigWithChan(request *QueryServiceConfigRequest) (<-chan *QueryServiceConfigResponse, <-chan error) {
	responseChan := make(chan *QueryServiceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryServiceConfig(request)
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

// QueryServiceConfigWithCallback invokes the scsp.QueryServiceConfig API asynchronously
func (client *Client) QueryServiceConfigWithCallback(request *QueryServiceConfigRequest, callback func(response *QueryServiceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryServiceConfigResponse
		var err error
		defer close(result)
		response, err = client.QueryServiceConfig(request)
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

// QueryServiceConfigRequest is the request struct for api QueryServiceConfig
type QueryServiceConfigRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query"`
	ViewCode   string `position:"Query"`
	Parameters string `position:"Query"`
}

// QueryServiceConfigResponse is the response struct for api QueryServiceConfig
type QueryServiceConfigResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateQueryServiceConfigRequest creates a request to invoke QueryServiceConfig API
func CreateQueryServiceConfigRequest() (request *QueryServiceConfigRequest) {
	request = &QueryServiceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "QueryServiceConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateQueryServiceConfigResponse creates a response to parse from QueryServiceConfig response
func CreateQueryServiceConfigResponse() (response *QueryServiceConfigResponse) {
	response = &QueryServiceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
