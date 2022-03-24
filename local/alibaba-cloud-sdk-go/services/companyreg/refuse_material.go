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

// RefuseMaterial invokes the companyreg.RefuseMaterial API synchronously
func (client *Client) RefuseMaterial(request *RefuseMaterialRequest) (response *RefuseMaterialResponse, err error) {
	response = CreateRefuseMaterialResponse()
	err = client.DoAction(request, response)
	return
}

// RefuseMaterialWithChan invokes the companyreg.RefuseMaterial API asynchronously
func (client *Client) RefuseMaterialWithChan(request *RefuseMaterialRequest) (<-chan *RefuseMaterialResponse, <-chan error) {
	responseChan := make(chan *RefuseMaterialResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefuseMaterial(request)
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

// RefuseMaterialWithCallback invokes the companyreg.RefuseMaterial API asynchronously
func (client *Client) RefuseMaterialWithCallback(request *RefuseMaterialRequest, callback func(response *RefuseMaterialResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefuseMaterialResponse
		var err error
		defer close(result)
		response, err = client.RefuseMaterial(request)
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

// RefuseMaterialRequest is the request struct for api RefuseMaterial
type RefuseMaterialRequest struct {
	*requests.RpcRequest
	Note  string `position:"Query" name:"Note"`
	BizId string `position:"Query" name:"BizId"`
}

// RefuseMaterialResponse is the response struct for api RefuseMaterial
type RefuseMaterialResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRefuseMaterialRequest creates a request to invoke RefuseMaterial API
func CreateRefuseMaterialRequest() (request *RefuseMaterialRequest) {
	request = &RefuseMaterialRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "RefuseMaterial", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRefuseMaterialResponse creates a response to parse from RefuseMaterial response
func CreateRefuseMaterialResponse() (response *RefuseMaterialResponse) {
	response = &RefuseMaterialResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
