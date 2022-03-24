package vod

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

// DescribePlayUserTotal invokes the vod.DescribePlayUserTotal API synchronously
func (client *Client) DescribePlayUserTotal(request *DescribePlayUserTotalRequest) (response *DescribePlayUserTotalResponse, err error) {
	response = CreateDescribePlayUserTotalResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePlayUserTotalWithChan invokes the vod.DescribePlayUserTotal API asynchronously
func (client *Client) DescribePlayUserTotalWithChan(request *DescribePlayUserTotalRequest) (<-chan *DescribePlayUserTotalResponse, <-chan error) {
	responseChan := make(chan *DescribePlayUserTotalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePlayUserTotal(request)
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

// DescribePlayUserTotalWithCallback invokes the vod.DescribePlayUserTotal API asynchronously
func (client *Client) DescribePlayUserTotalWithCallback(request *DescribePlayUserTotalRequest, callback func(response *DescribePlayUserTotalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePlayUserTotalResponse
		var err error
		defer close(result)
		response, err = client.DescribePlayUserTotal(request)
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

// DescribePlayUserTotalRequest is the request struct for api DescribePlayUserTotal
type DescribePlayUserTotalRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribePlayUserTotalResponse is the response struct for api DescribePlayUserTotal
type DescribePlayUserTotalResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	UserPlayStatisTotals UserPlayStatisTotals `json:"UserPlayStatisTotals" xml:"UserPlayStatisTotals"`
}

// CreateDescribePlayUserTotalRequest creates a request to invoke DescribePlayUserTotal API
func CreateDescribePlayUserTotalRequest() (request *DescribePlayUserTotalRequest) {
	request = &DescribePlayUserTotalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribePlayUserTotal", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePlayUserTotalResponse creates a response to parse from DescribePlayUserTotal response
func CreateDescribePlayUserTotalResponse() (response *DescribePlayUserTotalResponse) {
	response = &DescribePlayUserTotalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
