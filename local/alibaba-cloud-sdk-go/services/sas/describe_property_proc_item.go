package sas

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

// DescribePropertyProcItem invokes the sas.DescribePropertyProcItem API synchronously
func (client *Client) DescribePropertyProcItem(request *DescribePropertyProcItemRequest) (response *DescribePropertyProcItemResponse, err error) {
	response = CreateDescribePropertyProcItemResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePropertyProcItemWithChan invokes the sas.DescribePropertyProcItem API asynchronously
func (client *Client) DescribePropertyProcItemWithChan(request *DescribePropertyProcItemRequest) (<-chan *DescribePropertyProcItemResponse, <-chan error) {
	responseChan := make(chan *DescribePropertyProcItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePropertyProcItem(request)
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

// DescribePropertyProcItemWithCallback invokes the sas.DescribePropertyProcItem API asynchronously
func (client *Client) DescribePropertyProcItemWithCallback(request *DescribePropertyProcItemRequest, callback func(response *DescribePropertyProcItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePropertyProcItemResponse
		var err error
		defer close(result)
		response, err = client.DescribePropertyProcItem(request)
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

// DescribePropertyProcItemRequest is the request struct for api DescribePropertyProcItem
type DescribePropertyProcItemRequest struct {
	*requests.RpcRequest
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Name        string           `position:"Query" name:"Name"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	ForceFlush  requests.Boolean `position:"Query" name:"ForceFlush"`
}

// DescribePropertyProcItemResponse is the response struct for api DescribePropertyProcItem
type DescribePropertyProcItemResponse struct {
	*responses.BaseResponse
	RequestId     string             `json:"RequestId" xml:"RequestId"`
	PageInfo      PageInfo           `json:"PageInfo" xml:"PageInfo"`
	PropertyItems []PropertyProcItem `json:"PropertyItems" xml:"PropertyItems"`
}

// CreateDescribePropertyProcItemRequest creates a request to invoke DescribePropertyProcItem API
func CreateDescribePropertyProcItemRequest() (request *DescribePropertyProcItemRequest) {
	request = &DescribePropertyProcItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribePropertyProcItem", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePropertyProcItemResponse creates a response to parse from DescribePropertyProcItem response
func CreateDescribePropertyProcItemResponse() (response *DescribePropertyProcItemResponse) {
	response = &DescribePropertyProcItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
