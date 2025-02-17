package dts

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

// DescribeSubscriptionMeta invokes the dts.DescribeSubscriptionMeta API synchronously
func (client *Client) DescribeSubscriptionMeta(request *DescribeSubscriptionMetaRequest) (response *DescribeSubscriptionMetaResponse, err error) {
	response = CreateDescribeSubscriptionMetaResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSubscriptionMetaWithChan invokes the dts.DescribeSubscriptionMeta API asynchronously
func (client *Client) DescribeSubscriptionMetaWithChan(request *DescribeSubscriptionMetaRequest) (<-chan *DescribeSubscriptionMetaResponse, <-chan error) {
	responseChan := make(chan *DescribeSubscriptionMetaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSubscriptionMeta(request)
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

// DescribeSubscriptionMetaWithCallback invokes the dts.DescribeSubscriptionMeta API asynchronously
func (client *Client) DescribeSubscriptionMetaWithCallback(request *DescribeSubscriptionMetaRequest, callback func(response *DescribeSubscriptionMetaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSubscriptionMetaResponse
		var err error
		defer close(result)
		response, err = client.DescribeSubscriptionMeta(request)
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

// DescribeSubscriptionMetaRequest is the request struct for api DescribeSubscriptionMeta
type DescribeSubscriptionMetaRequest struct {
	*requests.RpcRequest
	Topics             string `position:"Query" name:"Topics"`
	Sid                string `position:"Query" name:"Sid"`
	SubMigrationJobIds string `position:"Query" name:"SubMigrationJobIds"`
	DtsInstanceId      string `position:"Query" name:"DtsInstanceId"`
}

// DescribeSubscriptionMetaResponse is the response struct for api DescribeSubscriptionMeta
type DescribeSubscriptionMetaResponse struct {
	*responses.BaseResponse
	ErrCode              string                     `json:"ErrCode" xml:"ErrCode"`
	ErrMessage           string                     `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode       string                     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId            string                     `json:"RequestId" xml:"RequestId"`
	Success              string                     `json:"Success" xml:"Success"`
	SubscriptionMetaList []SubscriptionMetaListItem `json:"SubscriptionMetaList" xml:"SubscriptionMetaList"`
}

// CreateDescribeSubscriptionMetaRequest creates a request to invoke DescribeSubscriptionMeta API
func CreateDescribeSubscriptionMetaRequest() (request *DescribeSubscriptionMetaRequest) {
	request = &DescribeSubscriptionMetaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeSubscriptionMeta", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSubscriptionMetaResponse creates a response to parse from DescribeSubscriptionMeta response
func CreateDescribeSubscriptionMetaResponse() (response *DescribeSubscriptionMetaResponse) {
	response = &DescribeSubscriptionMetaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
