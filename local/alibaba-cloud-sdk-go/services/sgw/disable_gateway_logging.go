package sgw

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

// DisableGatewayLogging invokes the sgw.DisableGatewayLogging API synchronously
func (client *Client) DisableGatewayLogging(request *DisableGatewayLoggingRequest) (response *DisableGatewayLoggingResponse, err error) {
	response = CreateDisableGatewayLoggingResponse()
	err = client.DoAction(request, response)
	return
}

// DisableGatewayLoggingWithChan invokes the sgw.DisableGatewayLogging API asynchronously
func (client *Client) DisableGatewayLoggingWithChan(request *DisableGatewayLoggingRequest) (<-chan *DisableGatewayLoggingResponse, <-chan error) {
	responseChan := make(chan *DisableGatewayLoggingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableGatewayLogging(request)
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

// DisableGatewayLoggingWithCallback invokes the sgw.DisableGatewayLogging API asynchronously
func (client *Client) DisableGatewayLoggingWithCallback(request *DisableGatewayLoggingRequest, callback func(response *DisableGatewayLoggingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableGatewayLoggingResponse
		var err error
		defer close(result)
		response, err = client.DisableGatewayLogging(request)
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

// DisableGatewayLoggingRequest is the request struct for api DisableGatewayLogging
type DisableGatewayLoggingRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// DisableGatewayLoggingResponse is the response struct for api DisableGatewayLogging
type DisableGatewayLoggingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDisableGatewayLoggingRequest creates a request to invoke DisableGatewayLogging API
func CreateDisableGatewayLoggingRequest() (request *DisableGatewayLoggingRequest) {
	request = &DisableGatewayLoggingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DisableGatewayLogging", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableGatewayLoggingResponse creates a response to parse from DisableGatewayLogging response
func CreateDisableGatewayLoggingResponse() (response *DisableGatewayLoggingResponse) {
	response = &DisableGatewayLoggingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
