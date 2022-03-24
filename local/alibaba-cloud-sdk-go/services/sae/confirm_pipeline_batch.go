package sae

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

// ConfirmPipelineBatch invokes the sae.ConfirmPipelineBatch API synchronously
func (client *Client) ConfirmPipelineBatch(request *ConfirmPipelineBatchRequest) (response *ConfirmPipelineBatchResponse, err error) {
	response = CreateConfirmPipelineBatchResponse()
	err = client.DoAction(request, response)
	return
}

// ConfirmPipelineBatchWithChan invokes the sae.ConfirmPipelineBatch API asynchronously
func (client *Client) ConfirmPipelineBatchWithChan(request *ConfirmPipelineBatchRequest) (<-chan *ConfirmPipelineBatchResponse, <-chan error) {
	responseChan := make(chan *ConfirmPipelineBatchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfirmPipelineBatch(request)
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

// ConfirmPipelineBatchWithCallback invokes the sae.ConfirmPipelineBatch API asynchronously
func (client *Client) ConfirmPipelineBatchWithCallback(request *ConfirmPipelineBatchRequest, callback func(response *ConfirmPipelineBatchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfirmPipelineBatchResponse
		var err error
		defer close(result)
		response, err = client.ConfirmPipelineBatch(request)
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

// ConfirmPipelineBatchRequest is the request struct for api ConfirmPipelineBatch
type ConfirmPipelineBatchRequest struct {
	*requests.RoaRequest
	Confirm    requests.Boolean `position:"Query" name:"Confirm"`
	PipelineId string           `position:"Query" name:"PipelineId"`
}

// ConfirmPipelineBatchResponse is the response struct for api ConfirmPipelineBatch
type ConfirmPipelineBatchResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateConfirmPipelineBatchRequest creates a request to invoke ConfirmPipelineBatch API
func CreateConfirmPipelineBatchRequest() (request *ConfirmPipelineBatchRequest) {
	request = &ConfirmPipelineBatchRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "ConfirmPipelineBatch", "/pop/v1/sam/changeorder/ConfirmPipelineBatch", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateConfirmPipelineBatchResponse creates a response to parse from ConfirmPipelineBatch response
func CreateConfirmPipelineBatchResponse() (response *ConfirmPipelineBatchResponse) {
	response = &ConfirmPipelineBatchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
