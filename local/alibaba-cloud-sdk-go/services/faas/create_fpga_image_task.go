package faas

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

// CreateFpgaImageTask invokes the faas.CreateFpgaImageTask API synchronously
// api document: https://help.aliyun.com/api/faas/createfpgaimagetask.html
func (client *Client) CreateFpgaImageTask(request *CreateFpgaImageTaskRequest) (response *CreateFpgaImageTaskResponse, err error) {
	response = CreateCreateFpgaImageTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFpgaImageTaskWithChan invokes the faas.CreateFpgaImageTask API asynchronously
// api document: https://help.aliyun.com/api/faas/createfpgaimagetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFpgaImageTaskWithChan(request *CreateFpgaImageTaskRequest) (<-chan *CreateFpgaImageTaskResponse, <-chan error) {
	responseChan := make(chan *CreateFpgaImageTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFpgaImageTask(request)
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

// CreateFpgaImageTaskWithCallback invokes the faas.CreateFpgaImageTask API asynchronously
// api document: https://help.aliyun.com/api/faas/createfpgaimagetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFpgaImageTaskWithCallback(request *CreateFpgaImageTaskRequest, callback func(response *CreateFpgaImageTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFpgaImageTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateFpgaImageTask(request)
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

// CreateFpgaImageTaskRequest is the request struct for api CreateFpgaImageTask
type CreateFpgaImageTaskRequest struct {
	*requests.RpcRequest
	ToolsVersion  string           `position:"Query" name:"ToolsVersion"`
	OssEndpoint   string           `position:"Query" name:"OssEndpoint"`
	Description   string           `position:"Query" name:"Description"`
	KeyId         string           `position:"Query" name:"KeyId"`
	Tags          string           `position:"Query" name:"Tags"`
	Bucket        string           `position:"Query" name:"Bucket"`
	ShellUniqueId string           `position:"Query" name:"ShellUniqueId"`
	Encryption    requests.Boolean `position:"Query" name:"Encryption"`
	Name          string           `position:"Query" name:"Name"`
	FpgaType      string           `position:"Query" name:"FpgaType"`
	Email         string           `position:"Query" name:"Email"`
	Object        string           `position:"Query" name:"Object"`
}

// CreateFpgaImageTaskResponse is the response struct for api CreateFpgaImageTask
type CreateFpgaImageTaskResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	FpgaImageUniqueId string `json:"FpgaImageUniqueId" xml:"FpgaImageUniqueId"`
	Name              string `json:"Name" xml:"Name"`
	CreateTime        string `json:"CreateTime" xml:"CreateTime"`
	Description       string `json:"Description" xml:"Description"`
	ShellUniqueId     string `json:"ShellUniqueId" xml:"ShellUniqueId"`
	State             string `json:"State" xml:"State"`
}

// CreateCreateFpgaImageTaskRequest creates a request to invoke CreateFpgaImageTask API
func CreateCreateFpgaImageTaskRequest() (request *CreateFpgaImageTaskRequest) {
	request = &CreateFpgaImageTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2020-02-17", "CreateFpgaImageTask", "faas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateFpgaImageTaskResponse creates a response to parse from CreateFpgaImageTask response
func CreateCreateFpgaImageTaskResponse() (response *CreateFpgaImageTaskResponse) {
	response = &CreateFpgaImageTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
