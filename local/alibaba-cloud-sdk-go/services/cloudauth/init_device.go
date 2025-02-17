package cloudauth

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

// InitDevice invokes the cloudauth.InitDevice API synchronously
func (client *Client) InitDevice(request *InitDeviceRequest) (response *InitDeviceResponse, err error) {
	response = CreateInitDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// InitDeviceWithChan invokes the cloudauth.InitDevice API asynchronously
func (client *Client) InitDeviceWithChan(request *InitDeviceRequest) (<-chan *InitDeviceResponse, <-chan error) {
	responseChan := make(chan *InitDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InitDevice(request)
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

// InitDeviceWithCallback invokes the cloudauth.InitDevice API asynchronously
func (client *Client) InitDeviceWithCallback(request *InitDeviceRequest, callback func(response *InitDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InitDeviceResponse
		var err error
		defer close(result)
		response, err = client.InitDevice(request)
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

// InitDeviceRequest is the request struct for api InitDevice
type InitDeviceRequest struct {
	*requests.RpcRequest
	Channel          string `position:"Query" name:"Channel"`
	BizData          string `position:"Query" name:"BizData"`
	Merchant         string `position:"Query" name:"Merchant"`
	AppVersion       string `position:"Query" name:"AppVersion"`
	DeviceToken      string `position:"Query" name:"DeviceToken"`
	CertifyId        string `position:"Query" name:"CertifyId"`
	WebUmidToken     string `position:"Body" name:"WebUmidToken"`
	OuterOrderNo     string `position:"Query" name:"OuterOrderNo"`
	ProduceNode      string `position:"Query" name:"ProduceNode"`
	UaToken          string `position:"Body" name:"UaToken"`
	ProductName      string `position:"Query" name:"ProductName"`
	CertifyPrincipal string `position:"Query" name:"CertifyPrincipal"`
	MetaInfo         string `position:"Query" name:"MetaInfo"`
}

// InitDeviceResponse is the response struct for api InitDevice
type InitDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Message      string       `json:"Message" xml:"Message"`
	Code         string       `json:"Code" xml:"Code"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateInitDeviceRequest creates a request to invoke InitDevice API
func CreateInitDeviceRequest() (request *InitDeviceRequest) {
	request = &InitDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "InitDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateInitDeviceResponse creates a response to parse from InitDevice response
func CreateInitDeviceResponse() (response *InitDeviceResponse) {
	response = &InitDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
