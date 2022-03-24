package r_kvstore

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

// EnableAdditionalBandwidth invokes the r_kvstore.EnableAdditionalBandwidth API synchronously
func (client *Client) EnableAdditionalBandwidth(request *EnableAdditionalBandwidthRequest) (response *EnableAdditionalBandwidthResponse, err error) {
	response = CreateEnableAdditionalBandwidthResponse()
	err = client.DoAction(request, response)
	return
}

// EnableAdditionalBandwidthWithChan invokes the r_kvstore.EnableAdditionalBandwidth API asynchronously
func (client *Client) EnableAdditionalBandwidthWithChan(request *EnableAdditionalBandwidthRequest) (<-chan *EnableAdditionalBandwidthResponse, <-chan error) {
	responseChan := make(chan *EnableAdditionalBandwidthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableAdditionalBandwidth(request)
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

// EnableAdditionalBandwidthWithCallback invokes the r_kvstore.EnableAdditionalBandwidth API asynchronously
func (client *Client) EnableAdditionalBandwidthWithCallback(request *EnableAdditionalBandwidthRequest, callback func(response *EnableAdditionalBandwidthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableAdditionalBandwidthResponse
		var err error
		defer close(result)
		response, err = client.EnableAdditionalBandwidth(request)
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

// EnableAdditionalBandwidthRequest is the request struct for api EnableAdditionalBandwidth
type EnableAdditionalBandwidthRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CouponNo             string           `position:"Query" name:"CouponNo"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	SourceBiz            string           `position:"Query" name:"SourceBiz"`
	NodeId               string           `position:"Query" name:"NodeId"`
	OrderTimeLength      string           `position:"Query" name:"OrderTimeLength"`
	AutoRenewPeriod      requests.Integer `position:"Query" name:"AutoRenewPeriod"`
	Product              string           `position:"Query" name:"Product"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string           `position:"Query" name:"Bandwidth"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	AutoRenew            requests.Boolean `position:"Query" name:"AutoRenew"`
	Category             string           `position:"Query" name:"Category"`
}

// EnableAdditionalBandwidthResponse is the response struct for api EnableAdditionalBandwidth
type EnableAdditionalBandwidthResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateEnableAdditionalBandwidthRequest creates a request to invoke EnableAdditionalBandwidth API
func CreateEnableAdditionalBandwidthRequest() (request *EnableAdditionalBandwidthRequest) {
	request = &EnableAdditionalBandwidthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "EnableAdditionalBandwidth", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableAdditionalBandwidthResponse creates a response to parse from EnableAdditionalBandwidth response
func CreateEnableAdditionalBandwidthResponse() (response *EnableAdditionalBandwidthResponse) {
	response = &EnableAdditionalBandwidthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
