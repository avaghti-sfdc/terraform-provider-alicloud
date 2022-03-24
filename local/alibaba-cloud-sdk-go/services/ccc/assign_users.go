package ccc

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

// AssignUsers invokes the ccc.AssignUsers API synchronously
func (client *Client) AssignUsers(request *AssignUsersRequest) (response *AssignUsersResponse, err error) {
	response = CreateAssignUsersResponse()
	err = client.DoAction(request, response)
	return
}

// AssignUsersWithChan invokes the ccc.AssignUsers API asynchronously
func (client *Client) AssignUsersWithChan(request *AssignUsersRequest) (<-chan *AssignUsersResponse, <-chan error) {
	responseChan := make(chan *AssignUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssignUsers(request)
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

// AssignUsersWithCallback invokes the ccc.AssignUsers API asynchronously
func (client *Client) AssignUsersWithCallback(request *AssignUsersRequest, callback func(response *AssignUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssignUsersResponse
		var err error
		defer close(result)
		response, err = client.AssignUsers(request)
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

// AssignUsersRequest is the request struct for api AssignUsers
type AssignUsersRequest struct {
	*requests.RpcRequest
	RamIdList      string `position:"Query" name:"RamIdList"`
	RoleId         string `position:"Query" name:"RoleId"`
	WorkMode       string `position:"Query" name:"WorkMode"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	SkillLevelList string `position:"Query" name:"SkillLevelList"`
}

// AssignUsersResponse is the response struct for api AssignUsers
type AssignUsersResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	Data           string `json:"Data" xml:"Data"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Sync           string `json:"Sync" xml:"Sync"`
	WorkflowId     string `json:"WorkflowId" xml:"WorkflowId"`
}

// CreateAssignUsersRequest creates a request to invoke AssignUsers API
func CreateAssignUsersRequest() (request *AssignUsersRequest) {
	request = &AssignUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "AssignUsers", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssignUsersResponse creates a response to parse from AssignUsers response
func CreateAssignUsersResponse() (response *AssignUsersResponse) {
	response = &AssignUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
