package rtc

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

// DescribeMPULayoutList invokes the rtc.DescribeMPULayoutList API synchronously
func (client *Client) DescribeMPULayoutList(request *DescribeMPULayoutListRequest) (response *DescribeMPULayoutListResponse, err error) {
	response = CreateDescribeMPULayoutListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMPULayoutListWithChan invokes the rtc.DescribeMPULayoutList API asynchronously
func (client *Client) DescribeMPULayoutListWithChan(request *DescribeMPULayoutListRequest) (<-chan *DescribeMPULayoutListResponse, <-chan error) {
	responseChan := make(chan *DescribeMPULayoutListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMPULayoutList(request)
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

// DescribeMPULayoutListWithCallback invokes the rtc.DescribeMPULayoutList API asynchronously
func (client *Client) DescribeMPULayoutListWithCallback(request *DescribeMPULayoutListRequest, callback func(response *DescribeMPULayoutListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMPULayoutListResponse
		var err error
		defer close(result)
		response, err = client.DescribeMPULayoutList(request)
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

// DescribeMPULayoutListRequest is the request struct for api DescribeMPULayoutList
type DescribeMPULayoutListRequest struct {
	*requests.RpcRequest
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
}

// DescribeMPULayoutListResponse is the response struct for api DescribeMPULayoutList
type DescribeMPULayoutListResponse struct {
	*responses.BaseResponse
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	LayoutIds LayoutIdsInDescribeMPULayoutList `json:"LayoutIds" xml:"LayoutIds"`
}

// CreateDescribeMPULayoutListRequest creates a request to invoke DescribeMPULayoutList API
func CreateDescribeMPULayoutListRequest() (request *DescribeMPULayoutListRequest) {
	request = &DescribeMPULayoutListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeMPULayoutList", "rtc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMPULayoutListResponse creates a response to parse from DescribeMPULayoutList response
func CreateDescribeMPULayoutListResponse() (response *DescribeMPULayoutListResponse) {
	response = &DescribeMPULayoutListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
