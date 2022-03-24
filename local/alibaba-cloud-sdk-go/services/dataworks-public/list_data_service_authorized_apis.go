package dataworks_public

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

// ListDataServiceAuthorizedApis invokes the dataworks_public.ListDataServiceAuthorizedApis API synchronously
func (client *Client) ListDataServiceAuthorizedApis(request *ListDataServiceAuthorizedApisRequest) (response *ListDataServiceAuthorizedApisResponse, err error) {
	response = CreateListDataServiceAuthorizedApisResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataServiceAuthorizedApisWithChan invokes the dataworks_public.ListDataServiceAuthorizedApis API asynchronously
func (client *Client) ListDataServiceAuthorizedApisWithChan(request *ListDataServiceAuthorizedApisRequest) (<-chan *ListDataServiceAuthorizedApisResponse, <-chan error) {
	responseChan := make(chan *ListDataServiceAuthorizedApisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataServiceAuthorizedApis(request)
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

// ListDataServiceAuthorizedApisWithCallback invokes the dataworks_public.ListDataServiceAuthorizedApis API asynchronously
func (client *Client) ListDataServiceAuthorizedApisWithCallback(request *ListDataServiceAuthorizedApisRequest, callback func(response *ListDataServiceAuthorizedApisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataServiceAuthorizedApisResponse
		var err error
		defer close(result)
		response, err = client.ListDataServiceAuthorizedApis(request)
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

// ListDataServiceAuthorizedApisRequest is the request struct for api ListDataServiceAuthorizedApis
type ListDataServiceAuthorizedApisRequest struct {
	*requests.RpcRequest
	ApiNameKeyword string           `position:"Body" name:"ApiNameKeyword"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	TenantId       requests.Integer `position:"Body" name:"TenantId"`
	ProjectId      requests.Integer `position:"Body" name:"ProjectId"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
}

// ListDataServiceAuthorizedApisResponse is the response struct for api ListDataServiceAuthorizedApis
type ListDataServiceAuthorizedApisResponse struct {
	*responses.BaseResponse
	ErrorCode      string                              `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string                              `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int                                 `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string                              `json:"RequestId" xml:"RequestId"`
	Success        bool                                `json:"Success" xml:"Success"`
	Data           DataInListDataServiceAuthorizedApis `json:"Data" xml:"Data"`
}

// CreateListDataServiceAuthorizedApisRequest creates a request to invoke ListDataServiceAuthorizedApis API
func CreateListDataServiceAuthorizedApisRequest() (request *ListDataServiceAuthorizedApisRequest) {
	request = &ListDataServiceAuthorizedApisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListDataServiceAuthorizedApis", "", "")
	request.Method = requests.POST
	return
}

// CreateListDataServiceAuthorizedApisResponse creates a response to parse from ListDataServiceAuthorizedApis response
func CreateListDataServiceAuthorizedApisResponse() (response *ListDataServiceAuthorizedApisResponse) {
	response = &ListDataServiceAuthorizedApisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
