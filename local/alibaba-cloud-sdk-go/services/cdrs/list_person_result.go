package cdrs

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

// ListPersonResult invokes the cdrs.ListPersonResult API synchronously
func (client *Client) ListPersonResult(request *ListPersonResultRequest) (response *ListPersonResultResponse, err error) {
	response = CreateListPersonResultResponse()
	err = client.DoAction(request, response)
	return
}

// ListPersonResultWithChan invokes the cdrs.ListPersonResult API asynchronously
func (client *Client) ListPersonResultWithChan(request *ListPersonResultRequest) (<-chan *ListPersonResultResponse, <-chan error) {
	responseChan := make(chan *ListPersonResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPersonResult(request)
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

// ListPersonResultWithCallback invokes the cdrs.ListPersonResult API asynchronously
func (client *Client) ListPersonResultWithCallback(request *ListPersonResultRequest, callback func(response *ListPersonResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPersonResultResponse
		var err error
		defer close(result)
		response, err = client.ListPersonResult(request)
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

// ListPersonResultRequest is the request struct for api ListPersonResult
type ListPersonResultRequest struct {
	*requests.RpcRequest
	Profession string           `position:"Body" name:"Profession"`
	Schema     string           `position:"Body" name:"Schema"`
	CorpId     string           `position:"Body" name:"CorpId"`
	Gender     string           `position:"Body" name:"Gender"`
	EndTime    string           `position:"Body" name:"EndTime"`
	StartTime  string           `position:"Body" name:"StartTime"`
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
	Age        string           `position:"Body" name:"Age"`
}

// ListPersonResultResponse is the response struct for api ListPersonResult
type ListPersonResultResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListPersonResultRequest creates a request to invoke ListPersonResult API
func CreateListPersonResultRequest() (request *ListPersonResultRequest) {
	request = &ListPersonResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListPersonResult", "", "")
	request.Method = requests.POST
	return
}

// CreateListPersonResultResponse creates a response to parse from ListPersonResult response
func CreateListPersonResultResponse() (response *ListPersonResultResponse) {
	response = &ListPersonResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
