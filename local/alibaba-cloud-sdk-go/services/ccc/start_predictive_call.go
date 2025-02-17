package ccc

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

// StartPredictiveCall invokes the ccc.StartPredictiveCall API synchronously
func (client *Client) StartPredictiveCall(request *StartPredictiveCallRequest) (response *StartPredictiveCallResponse, err error) {
	response = CreateStartPredictiveCallResponse()
	err = client.DoAction(request, response)
	return
}

// StartPredictiveCallWithChan invokes the ccc.StartPredictiveCall API asynchronously
func (client *Client) StartPredictiveCallWithChan(request *StartPredictiveCallRequest) (<-chan *StartPredictiveCallResponse, <-chan error) {
	responseChan := make(chan *StartPredictiveCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartPredictiveCall(request)
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

// StartPredictiveCallWithCallback invokes the ccc.StartPredictiveCall API asynchronously
func (client *Client) StartPredictiveCallWithCallback(request *StartPredictiveCallRequest, callback func(response *StartPredictiveCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartPredictiveCallResponse
		var err error
		defer close(result)
		response, err = client.StartPredictiveCall(request)
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

// StartPredictiveCallRequest is the request struct for api StartPredictiveCall
type StartPredictiveCallRequest struct {
	*requests.RpcRequest
	ContactFlowId        string           `position:"Query" name:"ContactFlowId"`
	Callee               string           `position:"Query" name:"Callee"`
	ContactFlowVariables string           `position:"Query" name:"ContactFlowVariables"`
	Tags                 string           `position:"Query" name:"Tags"`
	TimeoutSeconds       requests.Integer `position:"Query" name:"TimeoutSeconds"`
	Caller               string           `position:"Query" name:"Caller"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// StartPredictiveCallResponse is the response struct for api StartPredictiveCall
type StartPredictiveCallResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateStartPredictiveCallRequest creates a request to invoke StartPredictiveCall API
func CreateStartPredictiveCallRequest() (request *StartPredictiveCallRequest) {
	request = &StartPredictiveCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "StartPredictiveCall", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartPredictiveCallResponse creates a response to parse from StartPredictiveCall response
func CreateStartPredictiveCallResponse() (response *StartPredictiveCallResponse) {
	response = &StartPredictiveCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
