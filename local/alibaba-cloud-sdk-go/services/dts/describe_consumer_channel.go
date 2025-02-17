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

// DescribeConsumerChannel invokes the dts.DescribeConsumerChannel API synchronously
func (client *Client) DescribeConsumerChannel(request *DescribeConsumerChannelRequest) (response *DescribeConsumerChannelResponse, err error) {
	response = CreateDescribeConsumerChannelResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConsumerChannelWithChan invokes the dts.DescribeConsumerChannel API asynchronously
func (client *Client) DescribeConsumerChannelWithChan(request *DescribeConsumerChannelRequest) (<-chan *DescribeConsumerChannelResponse, <-chan error) {
	responseChan := make(chan *DescribeConsumerChannelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConsumerChannel(request)
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

// DescribeConsumerChannelWithCallback invokes the dts.DescribeConsumerChannel API asynchronously
func (client *Client) DescribeConsumerChannelWithCallback(request *DescribeConsumerChannelRequest, callback func(response *DescribeConsumerChannelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConsumerChannelResponse
		var err error
		defer close(result)
		response, err = client.DescribeConsumerChannel(request)
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

// DescribeConsumerChannelRequest is the request struct for api DescribeConsumerChannel
type DescribeConsumerChannelRequest struct {
	*requests.RpcRequest
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ParentChannelId string           `position:"Query" name:"ParentChannelId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	DtsJobId        string           `position:"Query" name:"DtsJobId"`
	DtsInstanceId   string           `position:"Query" name:"DtsInstanceId"`
}

// DescribeConsumerChannelResponse is the response struct for api DescribeConsumerChannel
type DescribeConsumerChannelResponse struct {
	*responses.BaseResponse
	ErrCode          string            `json:"ErrCode" xml:"ErrCode"`
	ErrMessage       string            `json:"ErrMessage" xml:"ErrMessage"`
	RequestId        string            `json:"RequestId" xml:"RequestId"`
	Success          string            `json:"Success" xml:"Success"`
	HttpStatusCode   string            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PageNumber       int               `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int               `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int64             `json:"TotalRecordCount" xml:"TotalRecordCount"`
	ConsumerChannels []ConsumerChannel `json:"ConsumerChannels" xml:"ConsumerChannels"`
}

// CreateDescribeConsumerChannelRequest creates a request to invoke DescribeConsumerChannel API
func CreateDescribeConsumerChannelRequest() (request *DescribeConsumerChannelRequest) {
	request = &DescribeConsumerChannelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeConsumerChannel", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeConsumerChannelResponse creates a response to parse from DescribeConsumerChannel response
func CreateDescribeConsumerChannelResponse() (response *DescribeConsumerChannelResponse) {
	response = &DescribeConsumerChannelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
