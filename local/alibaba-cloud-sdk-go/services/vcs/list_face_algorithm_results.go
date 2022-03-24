package vcs

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

// ListFaceAlgorithmResults invokes the vcs.ListFaceAlgorithmResults API synchronously
func (client *Client) ListFaceAlgorithmResults(request *ListFaceAlgorithmResultsRequest) (response *ListFaceAlgorithmResultsResponse, err error) {
	response = CreateListFaceAlgorithmResultsResponse()
	err = client.DoAction(request, response)
	return
}

// ListFaceAlgorithmResultsWithChan invokes the vcs.ListFaceAlgorithmResults API asynchronously
func (client *Client) ListFaceAlgorithmResultsWithChan(request *ListFaceAlgorithmResultsRequest) (<-chan *ListFaceAlgorithmResultsResponse, <-chan error) {
	responseChan := make(chan *ListFaceAlgorithmResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFaceAlgorithmResults(request)
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

// ListFaceAlgorithmResultsWithCallback invokes the vcs.ListFaceAlgorithmResults API asynchronously
func (client *Client) ListFaceAlgorithmResultsWithCallback(request *ListFaceAlgorithmResultsRequest, callback func(response *ListFaceAlgorithmResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFaceAlgorithmResultsResponse
		var err error
		defer close(result)
		response, err = client.ListFaceAlgorithmResults(request)
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

// ListFaceAlgorithmResultsRequest is the request struct for api ListFaceAlgorithmResults
type ListFaceAlgorithmResultsRequest struct {
	*requests.RpcRequest
	AlgorithmType string `position:"Body" name:"AlgorithmType"`
	CorpId        string `position:"Body" name:"CorpId"`
	EndTime       string `position:"Body" name:"EndTime"`
	StartTime     string `position:"Body" name:"StartTime"`
	PageNumber    string `position:"Body" name:"PageNumber"`
	DataSourceId  string `position:"Body" name:"DataSourceId"`
	PageSize      string `position:"Body" name:"PageSize"`
}

// ListFaceAlgorithmResultsResponse is the response struct for api ListFaceAlgorithmResults
type ListFaceAlgorithmResultsResponse struct {
	*responses.BaseResponse
	Code      string                         `json:"Code" xml:"Code"`
	Message   string                         `json:"Message" xml:"Message"`
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	Data      DataInListFaceAlgorithmResults `json:"Data" xml:"Data"`
}

// CreateListFaceAlgorithmResultsRequest creates a request to invoke ListFaceAlgorithmResults API
func CreateListFaceAlgorithmResultsRequest() (request *ListFaceAlgorithmResultsRequest) {
	request = &ListFaceAlgorithmResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "ListFaceAlgorithmResults", "", "")
	request.Method = requests.POST
	return
}

// CreateListFaceAlgorithmResultsResponse creates a response to parse from ListFaceAlgorithmResults response
func CreateListFaceAlgorithmResultsResponse() (response *ListFaceAlgorithmResultsResponse) {
	response = &ListFaceAlgorithmResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
