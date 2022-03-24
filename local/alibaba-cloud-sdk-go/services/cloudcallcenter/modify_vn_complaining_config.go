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

// ModifyVnComplainingConfig invokes the cloudcallcenter.ModifyVnComplainingConfig API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvncomplainingconfig.html
func (client *Client) ModifyVnComplainingConfig(request *ModifyVnComplainingConfigRequest) (response *ModifyVnComplainingConfigResponse, err error) {
	response = CreateModifyVnComplainingConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVnComplainingConfigWithChan invokes the cloudcallcenter.ModifyVnComplainingConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvncomplainingconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnComplainingConfigWithChan(request *ModifyVnComplainingConfigRequest) (<-chan *ModifyVnComplainingConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyVnComplainingConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVnComplainingConfig(request)
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

// ModifyVnComplainingConfigWithCallback invokes the cloudcallcenter.ModifyVnComplainingConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvncomplainingconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnComplainingConfigWithCallback(request *ModifyVnComplainingConfigRequest, callback func(response *ModifyVnComplainingConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVnComplainingConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyVnComplainingConfig(request)
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

// ModifyVnComplainingConfigRequest is the request struct for api ModifyVnComplainingConfig
type ModifyVnComplainingConfigRequest struct {
	*requests.RpcRequest
	Utterances        *[]string `position:"Query" name:"Utterances"  type:"Repeated"`
	FinalAction       string    `position:"Query" name:"FinalAction"`
	InstanceId        string    `position:"Query" name:"InstanceId"`
	FinalActionParams string    `position:"Query" name:"FinalActionParams"`
	Prompt            string    `position:"Query" name:"Prompt"`
}

// ModifyVnComplainingConfigResponse is the response struct for api ModifyVnComplainingConfig
type ModifyVnComplainingConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVnComplainingConfigRequest creates a request to invoke ModifyVnComplainingConfig API
func CreateModifyVnComplainingConfigRequest() (request *ModifyVnComplainingConfigRequest) {
	request = &ModifyVnComplainingConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyVnComplainingConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateModifyVnComplainingConfigResponse creates a response to parse from ModifyVnComplainingConfig response
func CreateModifyVnComplainingConfigResponse() (response *ModifyVnComplainingConfigResponse) {
	response = &ModifyVnComplainingConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
