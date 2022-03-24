package dyvmsapi

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

// SubmitHotlineTransferRegister invokes the dyvmsapi.SubmitHotlineTransferRegister API synchronously
func (client *Client) SubmitHotlineTransferRegister(request *SubmitHotlineTransferRegisterRequest) (response *SubmitHotlineTransferRegisterResponse, err error) {
	response = CreateSubmitHotlineTransferRegisterResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitHotlineTransferRegisterWithChan invokes the dyvmsapi.SubmitHotlineTransferRegister API asynchronously
func (client *Client) SubmitHotlineTransferRegisterWithChan(request *SubmitHotlineTransferRegisterRequest) (<-chan *SubmitHotlineTransferRegisterResponse, <-chan error) {
	responseChan := make(chan *SubmitHotlineTransferRegisterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitHotlineTransferRegister(request)
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

// SubmitHotlineTransferRegisterWithCallback invokes the dyvmsapi.SubmitHotlineTransferRegister API asynchronously
func (client *Client) SubmitHotlineTransferRegisterWithCallback(request *SubmitHotlineTransferRegisterRequest, callback func(response *SubmitHotlineTransferRegisterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitHotlineTransferRegisterResponse
		var err error
		defer close(result)
		response, err = client.SubmitHotlineTransferRegister(request)
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

// SubmitHotlineTransferRegisterRequest is the request struct for api SubmitHotlineTransferRegister
type SubmitHotlineTransferRegisterRequest struct {
	*requests.RpcRequest
	OperatorIdentityCard     string                                                   `position:"Query" name:"OperatorIdentityCard"`
	ResourceOwnerId          requests.Integer                                         `position:"Query" name:"ResourceOwnerId"`
	OperatorMail             string                                                   `position:"Query" name:"OperatorMail"`
	HotlineNumber            string                                                   `position:"Query" name:"HotlineNumber"`
	TransferPhoneNumberInfos *[]SubmitHotlineTransferRegisterTransferPhoneNumberInfos `position:"Query" name:"TransferPhoneNumberInfos"  type:"Repeated"`
	OperatorMobileVerifyCode string                                                   `position:"Query" name:"OperatorMobileVerifyCode"`
	Agreement                string                                                   `position:"Query" name:"Agreement"`
	QualificationId          string                                                   `position:"Query" name:"QualificationId"`
	ResourceOwnerAccount     string                                                   `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId                  requests.Integer                                         `position:"Query" name:"OwnerId"`
	OperatorMobile           string                                                   `position:"Query" name:"OperatorMobile"`
	OperatorMailVerifyCode   string                                                   `position:"Query" name:"OperatorMailVerifyCode"`
	OperatorName             string                                                   `position:"Query" name:"OperatorName"`
}

// SubmitHotlineTransferRegisterTransferPhoneNumberInfos is a repeated param struct in SubmitHotlineTransferRegisterRequest
type SubmitHotlineTransferRegisterTransferPhoneNumberInfos struct {
	PhoneNumber          string `name:"PhoneNumber"`
	PhoneNumberOwnerName string `name:"PhoneNumberOwnerName"`
	IdentityCard         string `name:"IdentityCard"`
}

// SubmitHotlineTransferRegisterResponse is the response struct for api SubmitHotlineTransferRegister
type SubmitHotlineTransferRegisterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      int64  `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateSubmitHotlineTransferRegisterRequest creates a request to invoke SubmitHotlineTransferRegister API
func CreateSubmitHotlineTransferRegisterRequest() (request *SubmitHotlineTransferRegisterRequest) {
	request = &SubmitHotlineTransferRegisterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "SubmitHotlineTransferRegister", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitHotlineTransferRegisterResponse creates a response to parse from SubmitHotlineTransferRegister response
func CreateSubmitHotlineTransferRegisterResponse() (response *SubmitHotlineTransferRegisterResponse) {
	response = &SubmitHotlineTransferRegisterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
