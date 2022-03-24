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

// AnswerCall invokes the ccc.AnswerCall API synchronously
func (client *Client) AnswerCall(request *AnswerCallRequest) (response *AnswerCallResponse, err error) {
	response = CreateAnswerCallResponse()
	err = client.DoAction(request, response)
	return
}

// AnswerCallWithChan invokes the ccc.AnswerCall API asynchronously
func (client *Client) AnswerCallWithChan(request *AnswerCallRequest) (<-chan *AnswerCallResponse, <-chan error) {
	responseChan := make(chan *AnswerCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AnswerCall(request)
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

// AnswerCallWithCallback invokes the ccc.AnswerCall API asynchronously
func (client *Client) AnswerCallWithCallback(request *AnswerCallRequest, callback func(response *AnswerCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AnswerCallResponse
		var err error
		defer close(result)
		response, err = client.AnswerCall(request)
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

// AnswerCallRequest is the request struct for api AnswerCall
type AnswerCallRequest struct {
	*requests.RpcRequest
	UserId     string `position:"Query" name:"UserId"`
	DeviceId   string `position:"Query" name:"DeviceId"`
	JobId      string `position:"Query" name:"JobId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// AnswerCallResponse is the response struct for api AnswerCall
type AnswerCallResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateAnswerCallRequest creates a request to invoke AnswerCall API
func CreateAnswerCallRequest() (request *AnswerCallRequest) {
	request = &AnswerCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "AnswerCall", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAnswerCallResponse creates a response to parse from AnswerCall response
func CreateAnswerCallResponse() (response *AnswerCallResponse) {
	response = &AnswerCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
