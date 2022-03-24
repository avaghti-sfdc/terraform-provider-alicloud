package devops_rdc

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

// UpdatePipelineMember invokes the devops_rdc.UpdatePipelineMember API synchronously
func (client *Client) UpdatePipelineMember(request *UpdatePipelineMemberRequest) (response *UpdatePipelineMemberResponse, err error) {
	response = CreateUpdatePipelineMemberResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePipelineMemberWithChan invokes the devops_rdc.UpdatePipelineMember API asynchronously
func (client *Client) UpdatePipelineMemberWithChan(request *UpdatePipelineMemberRequest) (<-chan *UpdatePipelineMemberResponse, <-chan error) {
	responseChan := make(chan *UpdatePipelineMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePipelineMember(request)
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

// UpdatePipelineMemberWithCallback invokes the devops_rdc.UpdatePipelineMember API asynchronously
func (client *Client) UpdatePipelineMemberWithCallback(request *UpdatePipelineMemberRequest, callback func(response *UpdatePipelineMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePipelineMemberResponse
		var err error
		defer close(result)
		response, err = client.UpdatePipelineMember(request)
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

// UpdatePipelineMemberRequest is the request struct for api UpdatePipelineMember
type UpdatePipelineMemberRequest struct {
	*requests.RpcRequest
	RoleName   string           `position:"Body" name:"RoleName"`
	UserPk     string           `position:"Body" name:"UserPk"`
	UserId     string           `position:"Body" name:"UserId"`
	OrgId      string           `position:"Query" name:"OrgId"`
	PipelineId requests.Integer `position:"Query" name:"PipelineId"`
}

// UpdatePipelineMemberResponse is the response struct for api UpdatePipelineMember
type UpdatePipelineMemberResponse struct {
	*responses.BaseResponse
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Object       bool   `json:"Object" xml:"Object"`
}

// CreateUpdatePipelineMemberRequest creates a request to invoke UpdatePipelineMember API
func CreateUpdatePipelineMemberRequest() (request *UpdatePipelineMemberRequest) {
	request = &UpdatePipelineMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "UpdatePipelineMember", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdatePipelineMemberResponse creates a response to parse from UpdatePipelineMember response
func CreateUpdatePipelineMemberResponse() (response *UpdatePipelineMemberResponse) {
	response = &UpdatePipelineMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
