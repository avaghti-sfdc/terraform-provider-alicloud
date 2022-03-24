package ivision

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

// DescribeFaceGroups invokes the ivision.DescribeFaceGroups API synchronously
func (client *Client) DescribeFaceGroups(request *DescribeFaceGroupsRequest) (response *DescribeFaceGroupsResponse, err error) {
	response = CreateDescribeFaceGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFaceGroupsWithChan invokes the ivision.DescribeFaceGroups API asynchronously
func (client *Client) DescribeFaceGroupsWithChan(request *DescribeFaceGroupsRequest) (<-chan *DescribeFaceGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeFaceGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFaceGroups(request)
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

// DescribeFaceGroupsWithCallback invokes the ivision.DescribeFaceGroups API asynchronously
func (client *Client) DescribeFaceGroupsWithCallback(request *DescribeFaceGroupsRequest, callback func(response *DescribeFaceGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFaceGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeFaceGroups(request)
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

// DescribeFaceGroupsRequest is the request struct for api DescribeFaceGroups
type DescribeFaceGroupsRequest struct {
	*requests.RpcRequest
	NextPageToken string           `position:"Query" name:"NextPageToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	ShowLog       string           `position:"Query" name:"ShowLog"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeFaceGroupsResponse is the response struct for api DescribeFaceGroups
type DescribeFaceGroupsResponse struct {
	*responses.BaseResponse
	RequestId     string  `json:"RequestId" xml:"RequestId"`
	CurrentPage   int64   `json:"CurrentPage" xml:"CurrentPage"`
	PageSize      int64   `json:"PageSize" xml:"PageSize"`
	NextPageToken string  `json:"NextPageToken" xml:"NextPageToken"`
	TotalNum      int64   `json:"TotalNum" xml:"TotalNum"`
	Groups        []Group `json:"Groups" xml:"Groups"`
}

// CreateDescribeFaceGroupsRequest creates a request to invoke DescribeFaceGroups API
func CreateDescribeFaceGroupsRequest() (request *DescribeFaceGroupsRequest) {
	request = &DescribeFaceGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivision", "2019-03-08", "DescribeFaceGroups", "ivision", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeFaceGroupsResponse creates a response to parse from DescribeFaceGroups response
func CreateDescribeFaceGroupsResponse() (response *DescribeFaceGroupsResponse) {
	response = &DescribeFaceGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
