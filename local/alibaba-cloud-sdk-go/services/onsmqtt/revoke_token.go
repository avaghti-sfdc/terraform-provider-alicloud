package onsmqtt

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

// RevokeToken invokes the onsmqtt.RevokeToken API synchronously
func (client *Client) RevokeToken(request *RevokeTokenRequest) (response *RevokeTokenResponse, err error) {
	response = CreateRevokeTokenResponse()
	err = client.DoAction(request, response)
	return
}

// RevokeTokenWithChan invokes the onsmqtt.RevokeToken API asynchronously
func (client *Client) RevokeTokenWithChan(request *RevokeTokenRequest) (<-chan *RevokeTokenResponse, <-chan error) {
	responseChan := make(chan *RevokeTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokeToken(request)
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

// RevokeTokenWithCallback invokes the onsmqtt.RevokeToken API asynchronously
func (client *Client) RevokeTokenWithCallback(request *RevokeTokenRequest, callback func(response *RevokeTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokeTokenResponse
		var err error
		defer close(result)
		response, err = client.RevokeToken(request)
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

// RevokeTokenRequest is the request struct for api RevokeToken
type RevokeTokenRequest struct {
	*requests.RpcRequest
	Token      string `position:"Query" name:"Token"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// RevokeTokenResponse is the response struct for api RevokeToken
type RevokeTokenResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRevokeTokenRequest creates a request to invoke RevokeToken API
func CreateRevokeTokenRequest() (request *RevokeTokenRequest) {
	request = &RevokeTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OnsMqtt", "2020-04-20", "RevokeToken", "onsmqtt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRevokeTokenResponse creates a response to parse from RevokeToken response
func CreateRevokeTokenResponse() (response *RevokeTokenResponse) {
	response = &RevokeTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
