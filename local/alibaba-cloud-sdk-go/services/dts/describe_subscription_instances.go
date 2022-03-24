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

// DescribeSubscriptionInstances invokes the dts.DescribeSubscriptionInstances API synchronously
func (client *Client) DescribeSubscriptionInstances(request *DescribeSubscriptionInstancesRequest) (response *DescribeSubscriptionInstancesResponse, err error) {
	response = CreateDescribeSubscriptionInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSubscriptionInstancesWithChan invokes the dts.DescribeSubscriptionInstances API asynchronously
func (client *Client) DescribeSubscriptionInstancesWithChan(request *DescribeSubscriptionInstancesRequest) (<-chan *DescribeSubscriptionInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeSubscriptionInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSubscriptionInstances(request)
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

// DescribeSubscriptionInstancesWithCallback invokes the dts.DescribeSubscriptionInstances API asynchronously
func (client *Client) DescribeSubscriptionInstancesWithCallback(request *DescribeSubscriptionInstancesRequest, callback func(response *DescribeSubscriptionInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSubscriptionInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeSubscriptionInstances(request)
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

// DescribeSubscriptionInstancesRequest is the request struct for api DescribeSubscriptionInstances
type DescribeSubscriptionInstancesRequest struct {
	*requests.RpcRequest
	ClientToken              string                              `position:"Query" name:"ClientToken"`
	InstFilterRegion         string                              `position:"Query" name:"InstFilterRegion"`
	PageNum                  requests.Integer                    `position:"Query" name:"PageNum"`
	OwnerId                  string                              `position:"Query" name:"OwnerId"`
	AccountId                string                              `position:"Query" name:"AccountId"`
	PageSize                 requests.Integer                    `position:"Query" name:"PageSize"`
	SubscriptionInstanceName string                              `position:"Query" name:"SubscriptionInstanceName"`
	Tag                      *[]DescribeSubscriptionInstancesTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// DescribeSubscriptionInstancesTag is a repeated param struct in DescribeSubscriptionInstancesRequest
type DescribeSubscriptionInstancesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeSubscriptionInstancesResponse is the response struct for api DescribeSubscriptionInstances
type DescribeSubscriptionInstancesResponse struct {
	*responses.BaseResponse
	ErrCode               string                `json:"ErrCode" xml:"ErrCode"`
	ErrMessage            string                `json:"ErrMessage" xml:"ErrMessage"`
	PageNumber            int                   `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount       int                   `json:"PageRecordCount" xml:"PageRecordCount"`
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	Success               string                `json:"Success" xml:"Success"`
	TotalRecordCount      int64                 `json:"TotalRecordCount" xml:"TotalRecordCount"`
	SubscriptionInstances SubscriptionInstances `json:"SubscriptionInstances" xml:"SubscriptionInstances"`
}

// CreateDescribeSubscriptionInstancesRequest creates a request to invoke DescribeSubscriptionInstances API
func CreateDescribeSubscriptionInstancesRequest() (request *DescribeSubscriptionInstancesRequest) {
	request = &DescribeSubscriptionInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeSubscriptionInstances", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSubscriptionInstancesResponse creates a response to parse from DescribeSubscriptionInstances response
func CreateDescribeSubscriptionInstancesResponse() (response *DescribeSubscriptionInstancesResponse) {
	response = &DescribeSubscriptionInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
