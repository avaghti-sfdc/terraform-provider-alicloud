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

// GetContactInfoByOutboundTaskId invokes the cloudcallcenter.GetContactInfoByOutboundTaskId API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcontactinfobyoutboundtaskid.html
func (client *Client) GetContactInfoByOutboundTaskId(request *GetContactInfoByOutboundTaskIdRequest) (response *GetContactInfoByOutboundTaskIdResponse, err error) {
	response = CreateGetContactInfoByOutboundTaskIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetContactInfoByOutboundTaskIdWithChan invokes the cloudcallcenter.GetContactInfoByOutboundTaskId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcontactinfobyoutboundtaskid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetContactInfoByOutboundTaskIdWithChan(request *GetContactInfoByOutboundTaskIdRequest) (<-chan *GetContactInfoByOutboundTaskIdResponse, <-chan error) {
	responseChan := make(chan *GetContactInfoByOutboundTaskIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetContactInfoByOutboundTaskId(request)
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

// GetContactInfoByOutboundTaskIdWithCallback invokes the cloudcallcenter.GetContactInfoByOutboundTaskId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcontactinfobyoutboundtaskid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetContactInfoByOutboundTaskIdWithCallback(request *GetContactInfoByOutboundTaskIdRequest, callback func(response *GetContactInfoByOutboundTaskIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetContactInfoByOutboundTaskIdResponse
		var err error
		defer close(result)
		response, err = client.GetContactInfoByOutboundTaskId(request)
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

// GetContactInfoByOutboundTaskIdRequest is the request struct for api GetContactInfoByOutboundTaskId
type GetContactInfoByOutboundTaskIdRequest struct {
	*requests.RpcRequest
	InstanceId     string `position:"Query" name:"InstanceId"`
	OutboundTaskId string `position:"Query" name:"OutboundTaskId"`
	SkillGroupId   string `position:"Query" name:"SkillGroupId"`
}

// GetContactInfoByOutboundTaskIdResponse is the response struct for api GetContactInfoByOutboundTaskId
type GetContactInfoByOutboundTaskIdResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Success         bool            `json:"Success" xml:"Success"`
	Code            string          `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	HttpStatusCode  int             `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ContactIdentity ContactIdentity `json:"ContactIdentity" xml:"ContactIdentity"`
}

// CreateGetContactInfoByOutboundTaskIdRequest creates a request to invoke GetContactInfoByOutboundTaskId API
func CreateGetContactInfoByOutboundTaskIdRequest() (request *GetContactInfoByOutboundTaskIdRequest) {
	request = &GetContactInfoByOutboundTaskIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetContactInfoByOutboundTaskId", "", "")
	request.Method = requests.POST
	return
}

// CreateGetContactInfoByOutboundTaskIdResponse creates a response to parse from GetContactInfoByOutboundTaskId response
func CreateGetContactInfoByOutboundTaskIdResponse() (response *GetContactInfoByOutboundTaskIdResponse) {
	response = &GetContactInfoByOutboundTaskIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
