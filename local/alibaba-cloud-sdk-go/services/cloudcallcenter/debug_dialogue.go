package cloudcallcenter

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

// DebugDialogue invokes the cloudcallcenter.DebugDialogue API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugdialogue.html
func (client *Client) DebugDialogue(request *DebugDialogueRequest) (response *DebugDialogueResponse, err error) {
	response = CreateDebugDialogueResponse()
	err = client.DoAction(request, response)
	return
}

// DebugDialogueWithChan invokes the cloudcallcenter.DebugDialogue API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugdialogue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugDialogueWithChan(request *DebugDialogueRequest) (<-chan *DebugDialogueResponse, <-chan error) {
	responseChan := make(chan *DebugDialogueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DebugDialogue(request)
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

// DebugDialogueWithCallback invokes the cloudcallcenter.DebugDialogue API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugdialogue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugDialogueWithCallback(request *DebugDialogueRequest, callback func(response *DebugDialogueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DebugDialogueResponse
		var err error
		defer close(result)
		response, err = client.DebugDialogue(request)
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

// DebugDialogueRequest is the request struct for api DebugDialogue
type DebugDialogueRequest struct {
	*requests.RpcRequest
	CallId        string           `position:"Query" name:"CallId"`
	CalledNumber  string           `position:"Query" name:"CalledNumber"`
	CallType      string           `position:"Query" name:"CallType"`
	ScenarioId    string           `position:"Query" name:"ScenarioId"`
	TaskId        string           `position:"Query" name:"TaskId"`
	Utterance     string           `position:"Query" name:"Utterance"`
	SurveyId      string           `position:"Query" name:"SurveyId"`
	ActionParams  string           `position:"Query" name:"ActionParams"`
	CallingNumber string           `position:"Query" name:"CallingNumber"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	NewSession    requests.Boolean `position:"Query" name:"NewSession"`
}

// DebugDialogueResponse is the response struct for api DebugDialogue
type DebugDialogueResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Feedback       Feedback `json:"Feedback" xml:"Feedback"`
}

// CreateDebugDialogueRequest creates a request to invoke DebugDialogue API
func CreateDebugDialogueRequest() (request *DebugDialogueRequest) {
	request = &DebugDialogueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DebugDialogue", "", "")
	request.Method = requests.POST
	return
}

// CreateDebugDialogueResponse creates a response to parse from DebugDialogue response
func CreateDebugDialogueResponse() (response *DebugDialogueResponse) {
	response = &DebugDialogueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
