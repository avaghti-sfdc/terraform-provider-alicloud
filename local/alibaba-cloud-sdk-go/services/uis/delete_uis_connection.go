package uis

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

// DeleteUisConnection invokes the uis.DeleteUisConnection API synchronously
// api document: https://help.aliyun.com/api/uis/deleteuisconnection.html
func (client *Client) DeleteUisConnection(request *DeleteUisConnectionRequest) (response *DeleteUisConnectionResponse, err error) {
	response = CreateDeleteUisConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUisConnectionWithChan invokes the uis.DeleteUisConnection API asynchronously
// api document: https://help.aliyun.com/api/uis/deleteuisconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUisConnectionWithChan(request *DeleteUisConnectionRequest) (<-chan *DeleteUisConnectionResponse, <-chan error) {
	responseChan := make(chan *DeleteUisConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUisConnection(request)
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

// DeleteUisConnectionWithCallback invokes the uis.DeleteUisConnection API asynchronously
// api document: https://help.aliyun.com/api/uis/deleteuisconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUisConnectionWithCallback(request *DeleteUisConnectionRequest, callback func(response *DeleteUisConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUisConnectionResponse
		var err error
		defer close(result)
		response, err = client.DeleteUisConnection(request)
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

// DeleteUisConnectionRequest is the request struct for api DeleteUisConnection
type DeleteUisConnectionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UisConnectionId      string           `position:"Query" name:"UisConnectionId"`
	UisNodeId            string           `position:"Query" name:"UisNodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteUisConnectionResponse is the response struct for api DeleteUisConnection
type DeleteUisConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteUisConnectionRequest creates a request to invoke DeleteUisConnection API
func CreateDeleteUisConnectionRequest() (request *DeleteUisConnectionRequest) {
	request = &DeleteUisConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Uis", "2018-08-21", "DeleteUisConnection", "uis", "openAPI")
	return
}

// CreateDeleteUisConnectionResponse creates a response to parse from DeleteUisConnection response
func CreateDeleteUisConnectionResponse() (response *DeleteUisConnectionResponse) {
	response = &DeleteUisConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
