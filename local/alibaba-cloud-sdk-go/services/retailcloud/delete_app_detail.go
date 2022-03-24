package retailcloud

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

// DeleteAppDetail invokes the retailcloud.DeleteAppDetail API synchronously
func (client *Client) DeleteAppDetail(request *DeleteAppDetailRequest) (response *DeleteAppDetailResponse, err error) {
	response = CreateDeleteAppDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAppDetailWithChan invokes the retailcloud.DeleteAppDetail API asynchronously
func (client *Client) DeleteAppDetailWithChan(request *DeleteAppDetailRequest) (<-chan *DeleteAppDetailResponse, <-chan error) {
	responseChan := make(chan *DeleteAppDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAppDetail(request)
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

// DeleteAppDetailWithCallback invokes the retailcloud.DeleteAppDetail API asynchronously
func (client *Client) DeleteAppDetailWithCallback(request *DeleteAppDetailRequest, callback func(response *DeleteAppDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAppDetailResponse
		var err error
		defer close(result)
		response, err = client.DeleteAppDetail(request)
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

// DeleteAppDetailRequest is the request struct for api DeleteAppDetail
type DeleteAppDetailRequest struct {
	*requests.RpcRequest
	AppId requests.Integer `position:"Query" name:"AppId"`
	Force requests.Boolean `position:"Query" name:"Force"`
}

// DeleteAppDetailResponse is the response struct for api DeleteAppDetail
type DeleteAppDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDeleteAppDetailRequest creates a request to invoke DeleteAppDetail API
func CreateDeleteAppDetailRequest() (request *DeleteAppDetailRequest) {
	request = &DeleteAppDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "DeleteAppDetail", "", "")
	request.Method = requests.GET
	return
}

// CreateDeleteAppDetailResponse creates a response to parse from DeleteAppDetail response
func CreateDeleteAppDetailResponse() (response *DeleteAppDetailResponse) {
	response = &DeleteAppDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
