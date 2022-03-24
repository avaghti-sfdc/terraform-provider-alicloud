package opensearch

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

// GetValidationError invokes the opensearch.GetValidationError API synchronously
func (client *Client) GetValidationError(request *GetValidationErrorRequest) (response *GetValidationErrorResponse, err error) {
	response = CreateGetValidationErrorResponse()
	err = client.DoAction(request, response)
	return
}

// GetValidationErrorWithChan invokes the opensearch.GetValidationError API asynchronously
func (client *Client) GetValidationErrorWithChan(request *GetValidationErrorRequest) (<-chan *GetValidationErrorResponse, <-chan error) {
	responseChan := make(chan *GetValidationErrorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetValidationError(request)
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

// GetValidationErrorWithCallback invokes the opensearch.GetValidationError API asynchronously
func (client *Client) GetValidationErrorWithCallback(request *GetValidationErrorRequest, callback func(response *GetValidationErrorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetValidationErrorResponse
		var err error
		defer close(result)
		response, err = client.GetValidationError(request)
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

// GetValidationErrorRequest is the request struct for api GetValidationError
type GetValidationErrorRequest struct {
	*requests.RoaRequest
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
	ErrorCode        string `position:"Query" name:"errorCode"`
}

// GetValidationErrorResponse is the response struct for api GetValidationError
type GetValidationErrorResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"requestId" xml:"requestId"`
	Result    []map[string]interface{} `json:"result" xml:"result"`
}

// CreateGetValidationErrorRequest creates a request to invoke GetValidationError API
func CreateGetValidationErrorRequest() (request *GetValidationErrorRequest) {
	request = &GetValidationErrorRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "GetValidationError", "/v4/openapi/app-groups/[appGroupIdentity]/algorithm/data/validation-error", "opensearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetValidationErrorResponse creates a response to parse from GetValidationError response
func CreateGetValidationErrorResponse() (response *GetValidationErrorResponse) {
	response = &GetValidationErrorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
