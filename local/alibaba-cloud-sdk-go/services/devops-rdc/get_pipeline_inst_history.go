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

// GetPipelineInstHistory invokes the devops_rdc.GetPipelineInstHistory API synchronously
func (client *Client) GetPipelineInstHistory(request *GetPipelineInstHistoryRequest) (response *GetPipelineInstHistoryResponse, err error) {
	response = CreateGetPipelineInstHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// GetPipelineInstHistoryWithChan invokes the devops_rdc.GetPipelineInstHistory API asynchronously
func (client *Client) GetPipelineInstHistoryWithChan(request *GetPipelineInstHistoryRequest) (<-chan *GetPipelineInstHistoryResponse, <-chan error) {
	responseChan := make(chan *GetPipelineInstHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPipelineInstHistory(request)
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

// GetPipelineInstHistoryWithCallback invokes the devops_rdc.GetPipelineInstHistory API asynchronously
func (client *Client) GetPipelineInstHistoryWithCallback(request *GetPipelineInstHistoryRequest, callback func(response *GetPipelineInstHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPipelineInstHistoryResponse
		var err error
		defer close(result)
		response, err = client.GetPipelineInstHistory(request)
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

// GetPipelineInstHistoryRequest is the request struct for api GetPipelineInstHistory
type GetPipelineInstHistoryRequest struct {
	*requests.RpcRequest
	EndTime    string           `position:"Body" name:"EndTime"`
	StartTime  string           `position:"Body" name:"StartTime"`
	UserPk     string           `position:"Body" name:"UserPk"`
	OrgId      string           `position:"Body" name:"OrgId"`
	PipelineId requests.Integer `position:"Body" name:"PipelineId"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
	PageStart  requests.Integer `position:"Body" name:"PageStart"`
}

// GetPipelineInstHistoryResponse is the response struct for api GetPipelineInstHistory
type GetPipelineInstHistoryResponse struct {
	*responses.BaseResponse
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateGetPipelineInstHistoryRequest creates a request to invoke GetPipelineInstHistory API
func CreateGetPipelineInstHistoryRequest() (request *GetPipelineInstHistoryRequest) {
	request = &GetPipelineInstHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "GetPipelineInstHistory", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPipelineInstHistoryResponse creates a response to parse from GetPipelineInstHistory response
func CreateGetPipelineInstHistoryResponse() (response *GetPipelineInstHistoryResponse) {
	response = &GetPipelineInstHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
