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

// DescribeSubscriptionInstanceStatus invokes the dts.DescribeSubscriptionInstanceStatus API synchronously
func (client *Client) DescribeSubscriptionInstanceStatus(request *DescribeSubscriptionInstanceStatusRequest) (response *DescribeSubscriptionInstanceStatusResponse, err error) {
	response = CreateDescribeSubscriptionInstanceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSubscriptionInstanceStatusWithChan invokes the dts.DescribeSubscriptionInstanceStatus API asynchronously
func (client *Client) DescribeSubscriptionInstanceStatusWithChan(request *DescribeSubscriptionInstanceStatusRequest) (<-chan *DescribeSubscriptionInstanceStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeSubscriptionInstanceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSubscriptionInstanceStatus(request)
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

// DescribeSubscriptionInstanceStatusWithCallback invokes the dts.DescribeSubscriptionInstanceStatus API asynchronously
func (client *Client) DescribeSubscriptionInstanceStatusWithCallback(request *DescribeSubscriptionInstanceStatusRequest, callback func(response *DescribeSubscriptionInstanceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSubscriptionInstanceStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeSubscriptionInstanceStatus(request)
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

// DescribeSubscriptionInstanceStatusRequest is the request struct for api DescribeSubscriptionInstanceStatus
type DescribeSubscriptionInstanceStatusRequest struct {
	*requests.RpcRequest
	SubscriptionInstanceId string `position:"Query" name:"SubscriptionInstanceId"`
	OwnerId                string `position:"Query" name:"OwnerId"`
	AccountId              string `position:"Query" name:"AccountId"`
}

// DescribeSubscriptionInstanceStatusResponse is the response struct for api DescribeSubscriptionInstanceStatus
type DescribeSubscriptionInstanceStatusResponse struct {
	*responses.BaseResponse
	BeginTimestamp           string                                                 `json:"BeginTimestamp" xml:"BeginTimestamp"`
	ConsumptionCheckpoint    string                                                 `json:"ConsumptionCheckpoint" xml:"ConsumptionCheckpoint"`
	ConsumptionClient        string                                                 `json:"ConsumptionClient" xml:"ConsumptionClient"`
	EndTimestamp             string                                                 `json:"EndTimestamp" xml:"EndTimestamp"`
	ErrMessage               string                                                 `json:"ErrMessage" xml:"ErrMessage"`
	PayType                  string                                                 `json:"PayType" xml:"PayType"`
	RequestId                string                                                 `json:"RequestId" xml:"RequestId"`
	Status                   string                                                 `json:"Status" xml:"Status"`
	SubscribeTopic           string                                                 `json:"SubscribeTopic" xml:"SubscribeTopic"`
	SubscriptionInstanceID   string                                                 `json:"SubscriptionInstanceID" xml:"SubscriptionInstanceID"`
	SubscriptionInstanceName string                                                 `json:"SubscriptionInstanceName" xml:"SubscriptionInstanceName"`
	ErrCode                  string                                                 `json:"ErrCode" xml:"ErrCode"`
	Success                  string                                                 `json:"Success" xml:"Success"`
	ErrorMessage             string                                                 `json:"ErrorMessage" xml:"ErrorMessage"`
	TaskId                   string                                                 `json:"TaskId" xml:"TaskId"`
	SourceEndpoint           SourceEndpoint                                         `json:"SourceEndpoint" xml:"SourceEndpoint"`
	SubscriptionDataType     SubscriptionDataType                                   `json:"SubscriptionDataType" xml:"SubscriptionDataType"`
	SubscriptionHost         SubscriptionHost                                       `json:"SubscriptionHost" xml:"SubscriptionHost"`
	SubscriptionObject       SubscriptionObjectInDescribeSubscriptionInstanceStatus `json:"SubscriptionObject" xml:"SubscriptionObject"`
}

// CreateDescribeSubscriptionInstanceStatusRequest creates a request to invoke DescribeSubscriptionInstanceStatus API
func CreateDescribeSubscriptionInstanceStatusRequest() (request *DescribeSubscriptionInstanceStatusRequest) {
	request = &DescribeSubscriptionInstanceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeSubscriptionInstanceStatus", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSubscriptionInstanceStatusResponse creates a response to parse from DescribeSubscriptionInstanceStatus response
func CreateDescribeSubscriptionInstanceStatusResponse() (response *DescribeSubscriptionInstanceStatusResponse) {
	response = &DescribeSubscriptionInstanceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
