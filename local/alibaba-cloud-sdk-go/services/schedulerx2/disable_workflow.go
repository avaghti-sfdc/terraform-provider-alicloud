package schedulerx2

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

// DisableWorkflow invokes the schedulerx2.DisableWorkflow API synchronously
func (client *Client) DisableWorkflow(request *DisableWorkflowRequest) (response *DisableWorkflowResponse, err error) {
	response = CreateDisableWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

// DisableWorkflowWithChan invokes the schedulerx2.DisableWorkflow API asynchronously
func (client *Client) DisableWorkflowWithChan(request *DisableWorkflowRequest) (<-chan *DisableWorkflowResponse, <-chan error) {
	responseChan := make(chan *DisableWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableWorkflow(request)
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

// DisableWorkflowWithCallback invokes the schedulerx2.DisableWorkflow API asynchronously
func (client *Client) DisableWorkflowWithCallback(request *DisableWorkflowRequest, callback func(response *DisableWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableWorkflowResponse
		var err error
		defer close(result)
		response, err = client.DisableWorkflow(request)
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

// DisableWorkflowRequest is the request struct for api DisableWorkflow
type DisableWorkflowRequest struct {
	*requests.RpcRequest
	NamespaceSource string           `position:"Query" name:"NamespaceSource"`
	GroupId         string           `position:"Query" name:"GroupId"`
	Namespace       string           `position:"Query" name:"Namespace"`
	WorkflowId      requests.Integer `position:"Query" name:"WorkflowId"`
}

// DisableWorkflowResponse is the response struct for api DisableWorkflow
type DisableWorkflowResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDisableWorkflowRequest creates a request to invoke DisableWorkflow API
func CreateDisableWorkflowRequest() (request *DisableWorkflowRequest) {
	request = &DisableWorkflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "DisableWorkflow", "", "")
	request.Method = requests.GET
	return
}

// CreateDisableWorkflowResponse creates a response to parse from DisableWorkflow response
func CreateDisableWorkflowResponse() (response *DisableWorkflowResponse) {
	response = &DisableWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
