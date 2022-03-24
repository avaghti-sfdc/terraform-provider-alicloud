package config

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

// GetDiscoveredResourceSummary invokes the config.GetDiscoveredResourceSummary API synchronously
func (client *Client) GetDiscoveredResourceSummary(request *GetDiscoveredResourceSummaryRequest) (response *GetDiscoveredResourceSummaryResponse, err error) {
	response = CreateGetDiscoveredResourceSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// GetDiscoveredResourceSummaryWithChan invokes the config.GetDiscoveredResourceSummary API asynchronously
func (client *Client) GetDiscoveredResourceSummaryWithChan(request *GetDiscoveredResourceSummaryRequest) (<-chan *GetDiscoveredResourceSummaryResponse, <-chan error) {
	responseChan := make(chan *GetDiscoveredResourceSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDiscoveredResourceSummary(request)
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

// GetDiscoveredResourceSummaryWithCallback invokes the config.GetDiscoveredResourceSummary API asynchronously
func (client *Client) GetDiscoveredResourceSummaryWithCallback(request *GetDiscoveredResourceSummaryRequest, callback func(response *GetDiscoveredResourceSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDiscoveredResourceSummaryResponse
		var err error
		defer close(result)
		response, err = client.GetDiscoveredResourceSummary(request)
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

// GetDiscoveredResourceSummaryRequest is the request struct for api GetDiscoveredResourceSummary
type GetDiscoveredResourceSummaryRequest struct {
	*requests.RpcRequest
	MultiAccount requests.Boolean `position:"Query" name:"MultiAccount"`
	MemberId     requests.Integer `position:"Query" name:"MemberId"`
}

// GetDiscoveredResourceSummaryResponse is the response struct for api GetDiscoveredResourceSummary
type GetDiscoveredResourceSummaryResponse struct {
	*responses.BaseResponse
	RequestId                 string                    `json:"RequestId" xml:"RequestId"`
	DiscoveredResourceSummary DiscoveredResourceSummary `json:"DiscoveredResourceSummary" xml:"DiscoveredResourceSummary"`
}

// CreateGetDiscoveredResourceSummaryRequest creates a request to invoke GetDiscoveredResourceSummary API
func CreateGetDiscoveredResourceSummaryRequest() (request *GetDiscoveredResourceSummaryRequest) {
	request = &GetDiscoveredResourceSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "GetDiscoveredResourceSummary", "Config", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetDiscoveredResourceSummaryResponse creates a response to parse from GetDiscoveredResourceSummary response
func CreateGetDiscoveredResourceSummaryResponse() (response *GetDiscoveredResourceSummaryResponse) {
	response = &GetDiscoveredResourceSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
