package airec

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

// QueryExceptionHistory invokes the airec.QueryExceptionHistory API synchronously
func (client *Client) QueryExceptionHistory(request *QueryExceptionHistoryRequest) (response *QueryExceptionHistoryResponse, err error) {
	response = CreateQueryExceptionHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// QueryExceptionHistoryWithChan invokes the airec.QueryExceptionHistory API asynchronously
func (client *Client) QueryExceptionHistoryWithChan(request *QueryExceptionHistoryRequest) (<-chan *QueryExceptionHistoryResponse, <-chan error) {
	responseChan := make(chan *QueryExceptionHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryExceptionHistory(request)
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

// QueryExceptionHistoryWithCallback invokes the airec.QueryExceptionHistory API asynchronously
func (client *Client) QueryExceptionHistoryWithCallback(request *QueryExceptionHistoryRequest, callback func(response *QueryExceptionHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryExceptionHistoryResponse
		var err error
		defer close(result)
		response, err = client.QueryExceptionHistory(request)
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

// QueryExceptionHistoryRequest is the request struct for api QueryExceptionHistory
type QueryExceptionHistoryRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"instanceId"`
	EndTime    requests.Integer `position:"Query" name:"endTime"`
	StartTime  requests.Integer `position:"Query" name:"startTime"`
	Type       string           `position:"Query" name:"type"`
}

// QueryExceptionHistoryResponse is the response struct for api QueryExceptionHistory
type QueryExceptionHistoryResponse struct {
	*responses.BaseResponse
	Code      string                 `json:"code" xml:"code"`
	Message   string                 `json:"message" xml:"message"`
	RequestId string                 `json:"requestId" xml:"requestId"`
	Result    map[string]interface{} `json:"result" xml:"result"`
}

// CreateQueryExceptionHistoryRequest creates a request to invoke QueryExceptionHistory API
func CreateQueryExceptionHistoryRequest() (request *QueryExceptionHistoryRequest) {
	request = &QueryExceptionHistoryRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "QueryExceptionHistory", "/v2/openapi/instances/[instanceId]/sync-reports/exception-history", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryExceptionHistoryResponse creates a response to parse from QueryExceptionHistory response
func CreateQueryExceptionHistoryResponse() (response *QueryExceptionHistoryResponse) {
	response = &QueryExceptionHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
