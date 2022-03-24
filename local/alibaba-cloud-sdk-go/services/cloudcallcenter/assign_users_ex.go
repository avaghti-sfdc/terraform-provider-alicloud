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

// AssignUsersEx invokes the cloudcallcenter.AssignUsersEx API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/assignusersex.html
func (client *Client) AssignUsersEx(request *AssignUsersExRequest) (response *AssignUsersExResponse, err error) {
	response = CreateAssignUsersExResponse()
	err = client.DoAction(request, response)
	return
}

// AssignUsersExWithChan invokes the cloudcallcenter.AssignUsersEx API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/assignusersex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssignUsersExWithChan(request *AssignUsersExRequest) (<-chan *AssignUsersExResponse, <-chan error) {
	responseChan := make(chan *AssignUsersExResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssignUsersEx(request)
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

// AssignUsersExWithCallback invokes the cloudcallcenter.AssignUsersEx API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/assignusersex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssignUsersExWithCallback(request *AssignUsersExRequest, callback func(response *AssignUsersExResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssignUsersExResponse
		var err error
		defer close(result)
		response, err = client.AssignUsersEx(request)
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

// AssignUsersExRequest is the request struct for api AssignUsersEx
type AssignUsersExRequest struct {
	*requests.RpcRequest
	SkillLevel   *[]string `position:"Query" name:"SkillLevel"  type:"Repeated"`
	InstanceId   string    `position:"Query" name:"InstanceId"`
	UserRamInfo  *[]string `position:"Query" name:"userRamInfo"  type:"Repeated"`
	RoleId       *[]string `position:"Query" name:"RoleId"  type:"Repeated"`
	SkillGroupId *[]string `position:"Query" name:"SkillGroupId"  type:"Repeated"`
}

// AssignUsersExResponse is the response struct for api AssignUsersEx
type AssignUsersExResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateAssignUsersExRequest creates a request to invoke AssignUsersEx API
func CreateAssignUsersExRequest() (request *AssignUsersExRequest) {
	request = &AssignUsersExRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "AssignUsersEx", "", "")
	request.Method = requests.POST
	return
}

// CreateAssignUsersExResponse creates a response to parse from AssignUsersEx response
func CreateAssignUsersExResponse() (response *AssignUsersExResponse) {
	response = &AssignUsersExResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
