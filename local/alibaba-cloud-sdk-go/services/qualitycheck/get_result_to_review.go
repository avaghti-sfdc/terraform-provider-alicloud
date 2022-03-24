package qualitycheck

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

// GetResultToReview invokes the qualitycheck.GetResultToReview API synchronously
func (client *Client) GetResultToReview(request *GetResultToReviewRequest) (response *GetResultToReviewResponse, err error) {
	response = CreateGetResultToReviewResponse()
	err = client.DoAction(request, response)
	return
}

// GetResultToReviewWithChan invokes the qualitycheck.GetResultToReview API asynchronously
func (client *Client) GetResultToReviewWithChan(request *GetResultToReviewRequest) (<-chan *GetResultToReviewResponse, <-chan error) {
	responseChan := make(chan *GetResultToReviewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResultToReview(request)
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

// GetResultToReviewWithCallback invokes the qualitycheck.GetResultToReview API asynchronously
func (client *Client) GetResultToReviewWithCallback(request *GetResultToReviewRequest, callback func(response *GetResultToReviewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResultToReviewResponse
		var err error
		defer close(result)
		response, err = client.GetResultToReview(request)
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

// GetResultToReviewRequest is the request struct for api GetResultToReview
type GetResultToReviewRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetResultToReviewResponse is the response struct for api GetResultToReview
type GetResultToReviewResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetResultToReviewRequest creates a request to invoke GetResultToReview API
func CreateGetResultToReviewRequest() (request *GetResultToReviewRequest) {
	request = &GetResultToReviewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetResultToReview", "", "")
	request.Method = requests.POST
	return
}

// CreateGetResultToReviewResponse creates a response to parse from GetResultToReview response
func CreateGetResultToReviewResponse() (response *GetResultToReviewResponse) {
	response = &GetResultToReviewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
