package vod

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

// GetAICaptionExtractionJobs invokes the vod.GetAICaptionExtractionJobs API synchronously
func (client *Client) GetAICaptionExtractionJobs(request *GetAICaptionExtractionJobsRequest) (response *GetAICaptionExtractionJobsResponse, err error) {
	response = CreateGetAICaptionExtractionJobsResponse()
	err = client.DoAction(request, response)
	return
}

// GetAICaptionExtractionJobsWithChan invokes the vod.GetAICaptionExtractionJobs API asynchronously
func (client *Client) GetAICaptionExtractionJobsWithChan(request *GetAICaptionExtractionJobsRequest) (<-chan *GetAICaptionExtractionJobsResponse, <-chan error) {
	responseChan := make(chan *GetAICaptionExtractionJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAICaptionExtractionJobs(request)
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

// GetAICaptionExtractionJobsWithCallback invokes the vod.GetAICaptionExtractionJobs API asynchronously
func (client *Client) GetAICaptionExtractionJobsWithCallback(request *GetAICaptionExtractionJobsRequest, callback func(response *GetAICaptionExtractionJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAICaptionExtractionJobsResponse
		var err error
		defer close(result)
		response, err = client.GetAICaptionExtractionJobs(request)
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

// GetAICaptionExtractionJobsRequest is the request struct for api GetAICaptionExtractionJobs
type GetAICaptionExtractionJobsRequest struct {
	*requests.RpcRequest
	JobIds string `position:"Query" name:"JobIds"`
}

// GetAICaptionExtractionJobsResponse is the response struct for api GetAICaptionExtractionJobs
type GetAICaptionExtractionJobsResponse struct {
	*responses.BaseResponse
	RequestId                  string                   `json:"RequestId" xml:"RequestId"`
	AICaptionExtractionJobList []AICaptionExtractionJob `json:"AICaptionExtractionJobList" xml:"AICaptionExtractionJobList"`
}

// CreateGetAICaptionExtractionJobsRequest creates a request to invoke GetAICaptionExtractionJobs API
func CreateGetAICaptionExtractionJobsRequest() (request *GetAICaptionExtractionJobsRequest) {
	request = &GetAICaptionExtractionJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetAICaptionExtractionJobs", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAICaptionExtractionJobsResponse creates a response to parse from GetAICaptionExtractionJobs response
func CreateGetAICaptionExtractionJobsResponse() (response *GetAICaptionExtractionJobsResponse) {
	response = &GetAICaptionExtractionJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
