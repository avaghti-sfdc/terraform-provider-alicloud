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

// SendShortMessage invokes the cloudcallcenter.SendShortMessage API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/sendshortmessage.html
func (client *Client) SendShortMessage(request *SendShortMessageRequest) (response *SendShortMessageResponse, err error) {
	response = CreateSendShortMessageResponse()
	err = client.DoAction(request, response)
	return
}

// SendShortMessageWithChan invokes the cloudcallcenter.SendShortMessage API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/sendshortmessage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SendShortMessageWithChan(request *SendShortMessageRequest) (<-chan *SendShortMessageResponse, <-chan error) {
	responseChan := make(chan *SendShortMessageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendShortMessage(request)
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

// SendShortMessageWithCallback invokes the cloudcallcenter.SendShortMessage API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/sendshortmessage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SendShortMessageWithCallback(request *SendShortMessageRequest, callback func(response *SendShortMessageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendShortMessageResponse
		var err error
		defer close(result)
		response, err = client.SendShortMessage(request)
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

// SendShortMessageRequest is the request struct for api SendShortMessage
type SendShortMessageRequest struct {
	*requests.RpcRequest
	SmsUpExtendCode string `position:"Query" name:"SmsUpExtendCode"`
	SignName        string `position:"Query" name:"SignName"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	PhoneNumbers    string `position:"Query" name:"PhoneNumbers"`
	OutId           string `position:"Query" name:"OutId"`
	TemplateCode    string `position:"Query" name:"TemplateCode"`
	TemplateParam   string `position:"Query" name:"TemplateParam"`
}

// SendShortMessageResponse is the response struct for api SendShortMessage
type SendShortMessageResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	BizId          string `json:"BizId" xml:"BizId"`
}

// CreateSendShortMessageRequest creates a request to invoke SendShortMessage API
func CreateSendShortMessageRequest() (request *SendShortMessageRequest) {
	request = &SendShortMessageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "SendShortMessage", "", "")
	request.Method = requests.POST
	return
}

// CreateSendShortMessageResponse creates a response to parse from SendShortMessage response
func CreateSendShortMessageResponse() (response *SendShortMessageResponse) {
	response = &SendShortMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
