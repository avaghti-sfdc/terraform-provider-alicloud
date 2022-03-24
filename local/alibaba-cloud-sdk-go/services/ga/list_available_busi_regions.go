package ga

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

// ListAvailableBusiRegions invokes the ga.ListAvailableBusiRegions API synchronously
func (client *Client) ListAvailableBusiRegions(request *ListAvailableBusiRegionsRequest) (response *ListAvailableBusiRegionsResponse, err error) {
	response = CreateListAvailableBusiRegionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAvailableBusiRegionsWithChan invokes the ga.ListAvailableBusiRegions API asynchronously
func (client *Client) ListAvailableBusiRegionsWithChan(request *ListAvailableBusiRegionsRequest) (<-chan *ListAvailableBusiRegionsResponse, <-chan error) {
	responseChan := make(chan *ListAvailableBusiRegionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAvailableBusiRegions(request)
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

// ListAvailableBusiRegionsWithCallback invokes the ga.ListAvailableBusiRegions API asynchronously
func (client *Client) ListAvailableBusiRegionsWithCallback(request *ListAvailableBusiRegionsRequest, callback func(response *ListAvailableBusiRegionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAvailableBusiRegionsResponse
		var err error
		defer close(result)
		response, err = client.ListAvailableBusiRegions(request)
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

// ListAvailableBusiRegionsRequest is the request struct for api ListAvailableBusiRegions
type ListAvailableBusiRegionsRequest struct {
	*requests.RpcRequest
	AcceleratorId string `position:"Query" name:"AcceleratorId"`
}

// ListAvailableBusiRegionsResponse is the response struct for api ListAvailableBusiRegions
type ListAvailableBusiRegionsResponse struct {
	*responses.BaseResponse
	RequestId string        `json:"RequestId" xml:"RequestId"`
	Regions   []RegionsItem `json:"Regions" xml:"Regions"`
}

// CreateListAvailableBusiRegionsRequest creates a request to invoke ListAvailableBusiRegions API
func CreateListAvailableBusiRegionsRequest() (request *ListAvailableBusiRegionsRequest) {
	request = &ListAvailableBusiRegionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "ListAvailableBusiRegions", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAvailableBusiRegionsResponse creates a response to parse from ListAvailableBusiRegions response
func CreateListAvailableBusiRegionsResponse() (response *ListAvailableBusiRegionsResponse) {
	response = &ListAvailableBusiRegionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
