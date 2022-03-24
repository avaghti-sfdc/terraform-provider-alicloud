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

// AddAITemplate invokes the vod.AddAITemplate API synchronously
func (client *Client) AddAITemplate(request *AddAITemplateRequest) (response *AddAITemplateResponse, err error) {
	response = CreateAddAITemplateResponse()
	err = client.DoAction(request, response)
	return
}

// AddAITemplateWithChan invokes the vod.AddAITemplate API asynchronously
func (client *Client) AddAITemplateWithChan(request *AddAITemplateRequest) (<-chan *AddAITemplateResponse, <-chan error) {
	responseChan := make(chan *AddAITemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddAITemplate(request)
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

// AddAITemplateWithCallback invokes the vod.AddAITemplate API asynchronously
func (client *Client) AddAITemplateWithCallback(request *AddAITemplateRequest, callback func(response *AddAITemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddAITemplateResponse
		var err error
		defer close(result)
		response, err = client.AddAITemplate(request)
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

// AddAITemplateRequest is the request struct for api AddAITemplate
type AddAITemplateRequest struct {
	*requests.RpcRequest
	TemplateConfig string `position:"Query" name:"TemplateConfig"`
	TemplateType   string `position:"Query" name:"TemplateType"`
	TemplateName   string `position:"Query" name:"TemplateName"`
}

// AddAITemplateResponse is the response struct for api AddAITemplate
type AddAITemplateResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
}

// CreateAddAITemplateRequest creates a request to invoke AddAITemplate API
func CreateAddAITemplateRequest() (request *AddAITemplateRequest) {
	request = &AddAITemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "AddAITemplate", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddAITemplateResponse creates a response to parse from AddAITemplate response
func CreateAddAITemplateResponse() (response *AddAITemplateResponse) {
	response = &AddAITemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
