package voicenavigator

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

// QueryConversations invokes the voicenavigator.QueryConversations API synchronously
func (client *Client) QueryConversations(request *QueryConversationsRequest) (response *QueryConversationsResponse, err error) {
	response = CreateQueryConversationsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryConversationsWithChan invokes the voicenavigator.QueryConversations API asynchronously
func (client *Client) QueryConversationsWithChan(request *QueryConversationsRequest) (<-chan *QueryConversationsResponse, <-chan error) {
	responseChan := make(chan *QueryConversationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryConversations(request)
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

// QueryConversationsWithCallback invokes the voicenavigator.QueryConversations API asynchronously
func (client *Client) QueryConversationsWithCallback(request *QueryConversationsRequest, callback func(response *QueryConversationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryConversationsResponse
		var err error
		defer close(result)
		response, err = client.QueryConversations(request)
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

// QueryConversationsRequest is the request struct for api QueryConversations
type QueryConversationsRequest struct {
	*requests.RpcRequest
	BeginTimeLeftRange  requests.Integer `position:"Query" name:"BeginTimeLeftRange"`
	PageNumber          requests.Integer `position:"Query" name:"PageNumber"`
	CallingNumber       string           `position:"Query" name:"CallingNumber"`
	InstanceId          string           `position:"Query" name:"InstanceId"`
	BeginTimeRightRange requests.Integer `position:"Query" name:"BeginTimeRightRange"`
	PageSize            requests.Integer `position:"Query" name:"PageSize"`
}

// QueryConversationsResponse is the response struct for api QueryConversations
type QueryConversationsResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	TotalCount    int64          `json:"TotalCount" xml:"TotalCount"`
	PageNumber    int            `json:"PageNumber" xml:"PageNumber"`
	PageSize      int            `json:"PageSize" xml:"PageSize"`
	Conversations []Conversation `json:"Conversations" xml:"Conversations"`
}

// CreateQueryConversationsRequest creates a request to invoke QueryConversations API
func CreateQueryConversationsRequest() (request *QueryConversationsRequest) {
	request = &QueryConversationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "QueryConversations", "voicebot", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryConversationsResponse creates a response to parse from QueryConversations response
func CreateQueryConversationsResponse() (response *QueryConversationsResponse) {
	response = &QueryConversationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
