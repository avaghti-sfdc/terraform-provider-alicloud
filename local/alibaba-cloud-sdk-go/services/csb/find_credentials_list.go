package csb

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

// FindCredentialsList invokes the csb.FindCredentialsList API synchronously
// api document: https://help.aliyun.com/api/csb/findcredentialslist.html
func (client *Client) FindCredentialsList(request *FindCredentialsListRequest) (response *FindCredentialsListResponse, err error) {
	response = CreateFindCredentialsListResponse()
	err = client.DoAction(request, response)
	return
}

// FindCredentialsListWithChan invokes the csb.FindCredentialsList API asynchronously
// api document: https://help.aliyun.com/api/csb/findcredentialslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindCredentialsListWithChan(request *FindCredentialsListRequest) (<-chan *FindCredentialsListResponse, <-chan error) {
	responseChan := make(chan *FindCredentialsListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindCredentialsList(request)
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

// FindCredentialsListWithCallback invokes the csb.FindCredentialsList API asynchronously
// api document: https://help.aliyun.com/api/csb/findcredentialslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindCredentialsListWithCallback(request *FindCredentialsListRequest, callback func(response *FindCredentialsListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindCredentialsListResponse
		var err error
		defer close(result)
		response, err = client.FindCredentialsList(request)
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

// FindCredentialsListRequest is the request struct for api FindCredentialsList
type FindCredentialsListRequest struct {
	*requests.RpcRequest
	CsbId     requests.Integer `position:"Query" name:"CsbId"`
	PageNum   requests.Integer `position:"Query" name:"PageNum"`
	GroupName string           `position:"Query" name:"GroupName"`
}

// FindCredentialsListResponse is the response struct for api FindCredentialsList
type FindCredentialsListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateFindCredentialsListRequest creates a request to invoke FindCredentialsList API
func CreateFindCredentialsListRequest() (request *FindCredentialsListRequest) {
	request = &FindCredentialsListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "FindCredentialsList", "", "")
	request.Method = requests.GET
	return
}

// CreateFindCredentialsListResponse creates a response to parse from FindCredentialsList response
func CreateFindCredentialsListResponse() (response *FindCredentialsListResponse) {
	response = &FindCredentialsListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
