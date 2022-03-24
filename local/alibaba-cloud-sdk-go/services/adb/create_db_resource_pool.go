package adb

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

// CreateDBResourcePool invokes the adb.CreateDBResourcePool API synchronously
func (client *Client) CreateDBResourcePool(request *CreateDBResourcePoolRequest) (response *CreateDBResourcePoolResponse, err error) {
	response = CreateCreateDBResourcePoolResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDBResourcePoolWithChan invokes the adb.CreateDBResourcePool API asynchronously
func (client *Client) CreateDBResourcePoolWithChan(request *CreateDBResourcePoolRequest) (<-chan *CreateDBResourcePoolResponse, <-chan error) {
	responseChan := make(chan *CreateDBResourcePoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBResourcePool(request)
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

// CreateDBResourcePoolWithCallback invokes the adb.CreateDBResourcePool API asynchronously
func (client *Client) CreateDBResourcePoolWithCallback(request *CreateDBResourcePoolRequest, callback func(response *CreateDBResourcePoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBResourcePoolResponse
		var err error
		defer close(result)
		response, err = client.CreateDBResourcePool(request)
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

// CreateDBResourcePoolRequest is the request struct for api CreateDBResourcePool
type CreateDBResourcePoolRequest struct {
	*requests.RpcRequest
	PoolName             string           `position:"Query" name:"PoolName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NodeNum              requests.Integer `position:"Query" name:"NodeNum"`
	QueryType            string           `position:"Query" name:"QueryType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateDBResourcePoolResponse is the response struct for api CreateDBResourcePool
type CreateDBResourcePoolResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDBResourcePoolRequest creates a request to invoke CreateDBResourcePool API
func CreateCreateDBResourcePoolRequest() (request *CreateDBResourcePoolRequest) {
	request = &CreateDBResourcePoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "CreateDBResourcePool", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDBResourcePoolResponse creates a response to parse from CreateDBResourcePool response
func CreateCreateDBResourcePoolResponse() (response *CreateDBResourcePoolResponse) {
	response = &CreateDBResourcePoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
