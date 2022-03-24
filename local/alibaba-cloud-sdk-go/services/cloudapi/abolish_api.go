package cloudapi

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

// AbolishApi invokes the cloudapi.AbolishApi API synchronously
// api document: https://help.aliyun.com/api/cloudapi/abolishapi.html
func (client *Client) AbolishApi(request *AbolishApiRequest) (response *AbolishApiResponse, err error) {
	response = CreateAbolishApiResponse()
	err = client.DoAction(request, response)
	return
}

// AbolishApiWithChan invokes the cloudapi.AbolishApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/abolishapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AbolishApiWithChan(request *AbolishApiRequest) (<-chan *AbolishApiResponse, <-chan error) {
	responseChan := make(chan *AbolishApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AbolishApi(request)
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

// AbolishApiWithCallback invokes the cloudapi.AbolishApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/abolishapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AbolishApiWithCallback(request *AbolishApiRequest, callback func(response *AbolishApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AbolishApiResponse
		var err error
		defer close(result)
		response, err = client.AbolishApi(request)
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

// AbolishApiRequest is the request struct for api AbolishApi
type AbolishApiRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	GroupId       string `position:"Query" name:"GroupId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// AbolishApiResponse is the response struct for api AbolishApi
type AbolishApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAbolishApiRequest creates a request to invoke AbolishApi API
func CreateAbolishApiRequest() (request *AbolishApiRequest) {
	request = &AbolishApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "AbolishApi", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAbolishApiResponse creates a response to parse from AbolishApi response
func CreateAbolishApiResponse() (response *AbolishApiResponse) {
	response = &AbolishApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
