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

// GetNumLocation invokes the scsp.GetNumLocation API synchronously
func (client *Client) GetNumLocation(request *GetNumLocationRequest) (response *GetNumLocationResponse, err error) {
	response = CreateGetNumLocationResponse()
	err = client.DoAction(request, response)
	return
}

// GetNumLocationWithChan invokes the scsp.GetNumLocation API asynchronously
func (client *Client) GetNumLocationWithChan(request *GetNumLocationRequest) (<-chan *GetNumLocationResponse, <-chan error) {
	responseChan := make(chan *GetNumLocationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNumLocation(request)
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

// GetNumLocationWithCallback invokes the scsp.GetNumLocation API asynchronously
func (client *Client) GetNumLocationWithCallback(request *GetNumLocationRequest, callback func(response *GetNumLocationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNumLocationResponse
		var err error
		defer close(result)
		response, err = client.GetNumLocation(request)
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

// GetNumLocationRequest is the request struct for api GetNumLocation
type GetNumLocationRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query"`
	InstanceId  string `position:"Query"`
	PhoneNum    string `position:"Query"`
}

// GetNumLocationResponse is the response struct for api GetNumLocation
type GetNumLocationResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           string `json:"Data" xml:"Data"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateGetNumLocationRequest creates a request to invoke GetNumLocation API
func CreateGetNumLocationRequest() (request *GetNumLocationRequest) {
	request = &GetNumLocationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "GetNumLocation", "", "")
	request.Method = requests.GET
	return
}

// CreateGetNumLocationResponse creates a response to parse from GetNumLocation response
func CreateGetNumLocationResponse() (response *GetNumLocationResponse) {
	response = &GetNumLocationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
