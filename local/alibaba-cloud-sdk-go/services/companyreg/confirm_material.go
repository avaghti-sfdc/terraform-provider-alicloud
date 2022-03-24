package companyreg

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

// ConfirmMaterial invokes the companyreg.ConfirmMaterial API synchronously
func (client *Client) ConfirmMaterial(request *ConfirmMaterialRequest) (response *ConfirmMaterialResponse, err error) {
	response = CreateConfirmMaterialResponse()
	err = client.DoAction(request, response)
	return
}

// ConfirmMaterialWithChan invokes the companyreg.ConfirmMaterial API asynchronously
func (client *Client) ConfirmMaterialWithChan(request *ConfirmMaterialRequest) (<-chan *ConfirmMaterialResponse, <-chan error) {
	responseChan := make(chan *ConfirmMaterialResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfirmMaterial(request)
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

// ConfirmMaterialWithCallback invokes the companyreg.ConfirmMaterial API asynchronously
func (client *Client) ConfirmMaterialWithCallback(request *ConfirmMaterialRequest, callback func(response *ConfirmMaterialResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfirmMaterialResponse
		var err error
		defer close(result)
		response, err = client.ConfirmMaterial(request)
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

// ConfirmMaterialRequest is the request struct for api ConfirmMaterial
type ConfirmMaterialRequest struct {
	*requests.RpcRequest
	UserOtherList string `position:"Query" name:"UserOtherList"`
	BizId         string `position:"Query" name:"BizId"`
}

// ConfirmMaterialResponse is the response struct for api ConfirmMaterial
type ConfirmMaterialResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfirmMaterialRequest creates a request to invoke ConfirmMaterial API
func CreateConfirmMaterialRequest() (request *ConfirmMaterialRequest) {
	request = &ConfirmMaterialRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "ConfirmMaterial", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfirmMaterialResponse creates a response to parse from ConfirmMaterial response
func CreateConfirmMaterialResponse() (response *ConfirmMaterialResponse) {
	response = &ConfirmMaterialResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
