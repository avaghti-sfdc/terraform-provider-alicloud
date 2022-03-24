package edas

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

// InsertRole invokes the edas.InsertRole API synchronously
func (client *Client) InsertRole(request *InsertRoleRequest) (response *InsertRoleResponse, err error) {
	response = CreateInsertRoleResponse()
	err = client.DoAction(request, response)
	return
}

// InsertRoleWithChan invokes the edas.InsertRole API asynchronously
func (client *Client) InsertRoleWithChan(request *InsertRoleRequest) (<-chan *InsertRoleResponse, <-chan error) {
	responseChan := make(chan *InsertRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertRole(request)
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

// InsertRoleWithCallback invokes the edas.InsertRole API asynchronously
func (client *Client) InsertRoleWithCallback(request *InsertRoleRequest, callback func(response *InsertRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertRoleResponse
		var err error
		defer close(result)
		response, err = client.InsertRole(request)
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

// InsertRoleRequest is the request struct for api InsertRole
type InsertRoleRequest struct {
	*requests.RoaRequest
	RoleName   string `position:"Query" name:"RoleName"`
	ActionData string `position:"Query" name:"ActionData"`
}

// InsertRoleResponse is the response struct for api InsertRole
type InsertRoleResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RoleId    int    `json:"RoleId" xml:"RoleId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateInsertRoleRequest creates a request to invoke InsertRole API
func CreateInsertRoleRequest() (request *InsertRoleRequest) {
	request = &InsertRoleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "InsertRole", "/pop/v5/account/create_role", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInsertRoleResponse creates a response to parse from InsertRole response
func CreateInsertRoleResponse() (response *InsertRoleResponse) {
	response = &InsertRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
