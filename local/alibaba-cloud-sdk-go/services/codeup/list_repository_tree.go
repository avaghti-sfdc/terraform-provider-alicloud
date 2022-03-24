package codeup

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

// ListRepositoryTree invokes the codeup.ListRepositoryTree API synchronously
func (client *Client) ListRepositoryTree(request *ListRepositoryTreeRequest) (response *ListRepositoryTreeResponse, err error) {
	response = CreateListRepositoryTreeResponse()
	err = client.DoAction(request, response)
	return
}

// ListRepositoryTreeWithChan invokes the codeup.ListRepositoryTree API asynchronously
func (client *Client) ListRepositoryTreeWithChan(request *ListRepositoryTreeRequest) (<-chan *ListRepositoryTreeResponse, <-chan error) {
	responseChan := make(chan *ListRepositoryTreeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRepositoryTree(request)
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

// ListRepositoryTreeWithCallback invokes the codeup.ListRepositoryTree API asynchronously
func (client *Client) ListRepositoryTreeWithCallback(request *ListRepositoryTreeRequest, callback func(response *ListRepositoryTreeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRepositoryTreeResponse
		var err error
		defer close(result)
		response, err = client.ListRepositoryTree(request)
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

// ListRepositoryTreeRequest is the request struct for api ListRepositoryTree
type ListRepositoryTreeRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	Path           string           `position:"Query" name:"Path"`
	SubUserId      string           `position:"Query" name:"SubUserId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	Type           string           `position:"Query" name:"Type"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
	RefName        string           `position:"Query" name:"RefName"`
}

// ListRepositoryTreeResponse is the response struct for api ListRepositoryTree
type ListRepositoryTreeResponse struct {
	*responses.BaseResponse
	RequestId    string                           `json:"RequestId" xml:"RequestId"`
	ErrorCode    string                           `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool                             `json:"Success" xml:"Success"`
	ErrorMessage string                           `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       []ResultItemInListRepositoryTree `json:"Result" xml:"Result"`
}

// CreateListRepositoryTreeRequest creates a request to invoke ListRepositoryTree API
func CreateListRepositoryTreeRequest() (request *ListRepositoryTreeRequest) {
	request = &ListRepositoryTreeRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "ListRepositoryTree", "/api/v4/projects/[ProjectId]/repository/tree", "", "")
	request.Method = requests.GET
	return
}

// CreateListRepositoryTreeResponse creates a response to parse from ListRepositoryTree response
func CreateListRepositoryTreeResponse() (response *ListRepositoryTreeResponse) {
	response = &ListRepositoryTreeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
