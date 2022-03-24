package dms_enterprise

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

// RevokeUserPermission invokes the dms_enterprise.RevokeUserPermission API synchronously
func (client *Client) RevokeUserPermission(request *RevokeUserPermissionRequest) (response *RevokeUserPermissionResponse, err error) {
	response = CreateRevokeUserPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// RevokeUserPermissionWithChan invokes the dms_enterprise.RevokeUserPermission API asynchronously
func (client *Client) RevokeUserPermissionWithChan(request *RevokeUserPermissionRequest) (<-chan *RevokeUserPermissionResponse, <-chan error) {
	responseChan := make(chan *RevokeUserPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokeUserPermission(request)
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

// RevokeUserPermissionWithCallback invokes the dms_enterprise.RevokeUserPermission API asynchronously
func (client *Client) RevokeUserPermissionWithCallback(request *RevokeUserPermissionRequest, callback func(response *RevokeUserPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokeUserPermissionResponse
		var err error
		defer close(result)
		response, err = client.RevokeUserPermission(request)
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

// RevokeUserPermissionRequest is the request struct for api RevokeUserPermission
type RevokeUserPermissionRequest struct {
	*requests.RpcRequest
	PermTypes    string           `position:"Query" name:"PermTypes"`
	UserAccessId string           `position:"Query" name:"UserAccessId"`
	DsType       string           `position:"Query" name:"DsType"`
	UserId       string           `position:"Query" name:"UserId"`
	Tid          requests.Integer `position:"Query" name:"Tid"`
	DbId         string           `position:"Query" name:"DbId"`
	TableId      string           `position:"Query" name:"TableId"`
	Logic        requests.Boolean `position:"Query" name:"Logic"`
	TableName    string           `position:"Query" name:"TableName"`
}

// RevokeUserPermissionResponse is the response struct for api RevokeUserPermission
type RevokeUserPermissionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateRevokeUserPermissionRequest creates a request to invoke RevokeUserPermission API
func CreateRevokeUserPermissionRequest() (request *RevokeUserPermissionRequest) {
	request = &RevokeUserPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "RevokeUserPermission", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRevokeUserPermissionResponse creates a response to parse from RevokeUserPermission response
func CreateRevokeUserPermissionResponse() (response *RevokeUserPermissionResponse) {
	response = &RevokeUserPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
