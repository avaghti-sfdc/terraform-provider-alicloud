package vod

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

// GetAppInfos invokes the vod.GetAppInfos API synchronously
func (client *Client) GetAppInfos(request *GetAppInfosRequest) (response *GetAppInfosResponse, err error) {
	response = CreateGetAppInfosResponse()
	err = client.DoAction(request, response)
	return
}

// GetAppInfosWithChan invokes the vod.GetAppInfos API asynchronously
func (client *Client) GetAppInfosWithChan(request *GetAppInfosRequest) (<-chan *GetAppInfosResponse, <-chan error) {
	responseChan := make(chan *GetAppInfosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAppInfos(request)
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

// GetAppInfosWithCallback invokes the vod.GetAppInfos API asynchronously
func (client *Client) GetAppInfosWithCallback(request *GetAppInfosRequest, callback func(response *GetAppInfosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAppInfosResponse
		var err error
		defer close(result)
		response, err = client.GetAppInfos(request)
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

// GetAppInfosRequest is the request struct for api GetAppInfos
type GetAppInfosRequest struct {
	*requests.RpcRequest
	ResourceRealOwnerId requests.Integer `position:"Query" name:"ResourceRealOwnerId"`
	AppIds              string           `position:"Query" name:"AppIds"`
}

// GetAppInfosResponse is the response struct for api GetAppInfos
type GetAppInfosResponse struct {
	*responses.BaseResponse
	RequestId      string    `json:"RequestId" xml:"RequestId"`
	NonExistAppIds []string  `json:"NonExistAppIds" xml:"NonExistAppIds"`
	AppInfoList    []AppInfo `json:"AppInfoList" xml:"AppInfoList"`
}

// CreateGetAppInfosRequest creates a request to invoke GetAppInfos API
func CreateGetAppInfosRequest() (request *GetAppInfosRequest) {
	request = &GetAppInfosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetAppInfos", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAppInfosResponse creates a response to parse from GetAppInfos response
func CreateGetAppInfosResponse() (response *GetAppInfosResponse) {
	response = &GetAppInfosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
