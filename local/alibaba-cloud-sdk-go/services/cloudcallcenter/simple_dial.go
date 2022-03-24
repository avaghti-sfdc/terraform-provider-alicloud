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

// SimpleDial invokes the cloudcallcenter.SimpleDial API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/simpledial.html
func (client *Client) SimpleDial(request *SimpleDialRequest) (response *SimpleDialResponse, err error) {
	response = CreateSimpleDialResponse()
	err = client.DoAction(request, response)
	return
}

// SimpleDialWithChan invokes the cloudcallcenter.SimpleDial API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/simpledial.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SimpleDialWithChan(request *SimpleDialRequest) (<-chan *SimpleDialResponse, <-chan error) {
	responseChan := make(chan *SimpleDialResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SimpleDial(request)
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

// SimpleDialWithCallback invokes the cloudcallcenter.SimpleDial API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/simpledial.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SimpleDialWithCallback(request *SimpleDialRequest, callback func(response *SimpleDialResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SimpleDialResponse
		var err error
		defer close(result)
		response, err = client.SimpleDial(request)
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

// SimpleDialRequest is the request struct for api SimpleDial
type SimpleDialRequest struct {
	*requests.RpcRequest
	Caller         string `position:"Query" name:"Caller"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	ContractFlowId string `position:"Query" name:"ContractFlowId"`
	Callee         string `position:"Query" name:"Callee"`
}

// SimpleDialResponse is the response struct for api SimpleDial
type SimpleDialResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	StatusCode     string `json:"StatusCode" xml:"StatusCode"`
	StatusDesc     string `json:"StatusDesc" xml:"StatusDesc"`
	TaskId         string `json:"TaskId" xml:"TaskId"`
	TimeStamp      string `json:"TimeStamp" xml:"TimeStamp"`
}

// CreateSimpleDialRequest creates a request to invoke SimpleDial API
func CreateSimpleDialRequest() (request *SimpleDialRequest) {
	request = &SimpleDialRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "SimpleDial", "", "")
	request.Method = requests.POST
	return
}

// CreateSimpleDialResponse creates a response to parse from SimpleDial response
func CreateSimpleDialResponse() (response *SimpleDialResponse) {
	response = &SimpleDialResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
