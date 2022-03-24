package sddp

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

// DescribeAccounts invokes the sddp.DescribeAccounts API synchronously
// api document: https://help.aliyun.com/api/sddp/describeaccounts.html
func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
	response = CreateDescribeAccountsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAccountsWithChan invokes the sddp.DescribeAccounts API asynchronously
// api document: https://help.aliyun.com/api/sddp/describeaccounts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAccountsWithChan(request *DescribeAccountsRequest) (<-chan *DescribeAccountsResponse, <-chan error) {
	responseChan := make(chan *DescribeAccountsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAccounts(request)
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

// DescribeAccountsWithCallback invokes the sddp.DescribeAccounts API asynchronously
// api document: https://help.aliyun.com/api/sddp/describeaccounts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAccountsWithCallback(request *DescribeAccountsRequest, callback func(response *DescribeAccountsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAccountsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAccounts(request)
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

// DescribeAccountsRequest is the request struct for api DescribeAccounts
type DescribeAccountsRequest struct {
	*requests.RpcRequest
	ProductCode string           `position:"Query" name:"ProductCode"`
	LoginName   string           `position:"Query" name:"LoginName"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	ColumnId    string           `position:"Query" name:"ColumnId"`
	PackageId   string           `position:"Query" name:"PackageId"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	DepartId    requests.Integer `position:"Query" name:"DepartId"`
	OperationId requests.Integer `position:"Query" name:"OperationId"`
	TableId     string           `position:"Query" name:"TableId"`
	Lang        string           `position:"Query" name:"Lang"`
	Key         string           `position:"Query" name:"Key"`
	QueryType   requests.Integer `position:"Query" name:"QueryType"`
}

// DescribeAccountsResponse is the response struct for api DescribeAccounts
type DescribeAccountsResponse struct {
	*responses.BaseResponse
	RequestId   string    `json:"RequestId" xml:"RequestId"`
	PageSize    int       `json:"PageSize" xml:"PageSize"`
	CurrentPage int       `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int       `json:"TotalCount" xml:"TotalCount"`
	Items       []Account `json:"Items" xml:"Items"`
}

// CreateDescribeAccountsRequest creates a request to invoke DescribeAccounts API
func CreateDescribeAccountsRequest() (request *DescribeAccountsRequest) {
	request = &DescribeAccountsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeAccounts", "sddp", "openAPI")
	return
}

// CreateDescribeAccountsResponse creates a response to parse from DescribeAccounts response
func CreateDescribeAccountsResponse() (response *DescribeAccountsResponse) {
	response = &DescribeAccountsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
