package dataworks_public

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

// ListDISyncTasks invokes the dataworks_public.ListDISyncTasks API synchronously
func (client *Client) ListDISyncTasks(request *ListDISyncTasksRequest) (response *ListDISyncTasksResponse, err error) {
	response = CreateListDISyncTasksResponse()
	err = client.DoAction(request, response)
	return
}

// ListDISyncTasksWithChan invokes the dataworks_public.ListDISyncTasks API asynchronously
func (client *Client) ListDISyncTasksWithChan(request *ListDISyncTasksRequest) (<-chan *ListDISyncTasksResponse, <-chan error) {
	responseChan := make(chan *ListDISyncTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDISyncTasks(request)
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

// ListDISyncTasksWithCallback invokes the dataworks_public.ListDISyncTasks API asynchronously
func (client *Client) ListDISyncTasksWithCallback(request *ListDISyncTasksRequest, callback func(response *ListDISyncTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDISyncTasksResponse
		var err error
		defer close(result)
		response, err = client.ListDISyncTasks(request)
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

// ListDISyncTasksRequest is the request struct for api ListDISyncTasks
type ListDISyncTasksRequest struct {
	*requests.RpcRequest
	TaskType   string           `position:"Query" name:"TaskType"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	ProjectId  requests.Integer `position:"Query" name:"ProjectId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListDISyncTasksResponse is the response struct for api ListDISyncTasks
type ListDISyncTasksResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Success   bool     `json:"Success" xml:"Success"`
	TaskList  TaskList `json:"TaskList" xml:"TaskList"`
}

// CreateListDISyncTasksRequest creates a request to invoke ListDISyncTasks API
func CreateListDISyncTasksRequest() (request *ListDISyncTasksRequest) {
	request = &ListDISyncTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListDISyncTasks", "", "")
	request.Method = requests.POST
	return
}

// CreateListDISyncTasksResponse creates a response to parse from ListDISyncTasks response
func CreateListDISyncTasksResponse() (response *ListDISyncTasksResponse) {
	response = &ListDISyncTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
