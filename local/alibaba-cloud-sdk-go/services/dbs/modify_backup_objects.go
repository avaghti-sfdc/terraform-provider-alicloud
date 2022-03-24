package dbs

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

// ModifyBackupObjects invokes the dbs.ModifyBackupObjects API synchronously
func (client *Client) ModifyBackupObjects(request *ModifyBackupObjectsRequest) (response *ModifyBackupObjectsResponse, err error) {
	response = CreateModifyBackupObjectsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBackupObjectsWithChan invokes the dbs.ModifyBackupObjects API asynchronously
func (client *Client) ModifyBackupObjectsWithChan(request *ModifyBackupObjectsRequest) (<-chan *ModifyBackupObjectsResponse, <-chan error) {
	responseChan := make(chan *ModifyBackupObjectsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBackupObjects(request)
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

// ModifyBackupObjectsWithCallback invokes the dbs.ModifyBackupObjects API asynchronously
func (client *Client) ModifyBackupObjectsWithCallback(request *ModifyBackupObjectsRequest, callback func(response *ModifyBackupObjectsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBackupObjectsResponse
		var err error
		defer close(result)
		response, err = client.ModifyBackupObjects(request)
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

// ModifyBackupObjectsRequest is the request struct for api ModifyBackupObjects
type ModifyBackupObjectsRequest struct {
	*requests.RpcRequest
	ClientToken   string `position:"Query" name:"ClientToken"`
	BackupPlanId  string `position:"Query" name:"BackupPlanId"`
	BackupObjects string `position:"Query" name:"BackupObjects"`
	OwnerId       string `position:"Query" name:"OwnerId"`
}

// ModifyBackupObjectsResponse is the response struct for api ModifyBackupObjects
type ModifyBackupObjectsResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	BackupPlanId   string `json:"BackupPlanId" xml:"BackupPlanId"`
	NeedPrecheck   bool   `json:"NeedPrecheck" xml:"NeedPrecheck"`
}

// CreateModifyBackupObjectsRequest creates a request to invoke ModifyBackupObjects API
func CreateModifyBackupObjectsRequest() (request *ModifyBackupObjectsRequest) {
	request = &ModifyBackupObjectsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "ModifyBackupObjects", "cbs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyBackupObjectsResponse creates a response to parse from ModifyBackupObjects response
func CreateModifyBackupObjectsResponse() (response *ModifyBackupObjectsResponse) {
	response = &ModifyBackupObjectsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
