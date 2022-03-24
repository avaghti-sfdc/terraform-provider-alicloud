package quotas

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

// ListProductQuotaDimensions invokes the quotas.ListProductQuotaDimensions API synchronously
func (client *Client) ListProductQuotaDimensions(request *ListProductQuotaDimensionsRequest) (response *ListProductQuotaDimensionsResponse, err error) {
	response = CreateListProductQuotaDimensionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListProductQuotaDimensionsWithChan invokes the quotas.ListProductQuotaDimensions API asynchronously
func (client *Client) ListProductQuotaDimensionsWithChan(request *ListProductQuotaDimensionsRequest) (<-chan *ListProductQuotaDimensionsResponse, <-chan error) {
	responseChan := make(chan *ListProductQuotaDimensionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProductQuotaDimensions(request)
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

// ListProductQuotaDimensionsWithCallback invokes the quotas.ListProductQuotaDimensions API asynchronously
func (client *Client) ListProductQuotaDimensionsWithCallback(request *ListProductQuotaDimensionsRequest, callback func(response *ListProductQuotaDimensionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProductQuotaDimensionsResponse
		var err error
		defer close(result)
		response, err = client.ListProductQuotaDimensions(request)
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

// ListProductQuotaDimensionsRequest is the request struct for api ListProductQuotaDimensions
type ListProductQuotaDimensionsRequest struct {
	*requests.RpcRequest
	ProductCode   string           `position:"Body" name:"ProductCode"`
	NextToken     string           `position:"Body" name:"NextToken"`
	MaxResults    requests.Integer `position:"Body" name:"MaxResults"`
	QuotaCategory string           `position:"Body" name:"QuotaCategory"`
}

// ListProductQuotaDimensionsResponse is the response struct for api ListProductQuotaDimensions
type ListProductQuotaDimensionsResponse struct {
	*responses.BaseResponse
	TotalCount      int                   `json:"TotalCount" xml:"TotalCount"`
	RequestId       string                `json:"RequestId" xml:"RequestId"`
	NextToken       string                `json:"NextToken" xml:"NextToken"`
	MaxResults      int                   `json:"MaxResults" xml:"MaxResults"`
	QuotaDimensions []QuotaDimensionsItem `json:"QuotaDimensions" xml:"QuotaDimensions"`
}

// CreateListProductQuotaDimensionsRequest creates a request to invoke ListProductQuotaDimensions API
func CreateListProductQuotaDimensionsRequest() (request *ListProductQuotaDimensionsRequest) {
	request = &ListProductQuotaDimensionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quotas", "2020-05-10", "ListProductQuotaDimensions", "quotas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListProductQuotaDimensionsResponse creates a response to parse from ListProductQuotaDimensions response
func CreateListProductQuotaDimensionsResponse() (response *ListProductQuotaDimensionsResponse) {
	response = &ListProductQuotaDimensionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
