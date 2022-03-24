package live

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

// DescribeHtmlResource invokes the live.DescribeHtmlResource API synchronously
func (client *Client) DescribeHtmlResource(request *DescribeHtmlResourceRequest) (response *DescribeHtmlResourceResponse, err error) {
	response = CreateDescribeHtmlResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHtmlResourceWithChan invokes the live.DescribeHtmlResource API asynchronously
func (client *Client) DescribeHtmlResourceWithChan(request *DescribeHtmlResourceRequest) (<-chan *DescribeHtmlResourceResponse, <-chan error) {
	responseChan := make(chan *DescribeHtmlResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHtmlResource(request)
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

// DescribeHtmlResourceWithCallback invokes the live.DescribeHtmlResource API asynchronously
func (client *Client) DescribeHtmlResourceWithCallback(request *DescribeHtmlResourceRequest, callback func(response *DescribeHtmlResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHtmlResourceResponse
		var err error
		defer close(result)
		response, err = client.DescribeHtmlResource(request)
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

// DescribeHtmlResourceRequest is the request struct for api DescribeHtmlResource
type DescribeHtmlResourceRequest struct {
	*requests.RpcRequest
	HtmlUrl        string           `position:"Query" name:"htmlUrl"`
	CasterId       string           `position:"Query" name:"CasterId"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	HtmlResourceId string           `position:"Query" name:"HtmlResourceId"`
}

// DescribeHtmlResourceResponse is the response struct for api DescribeHtmlResource
type DescribeHtmlResourceResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	HtmlResource HtmlResource `json:"HtmlResource" xml:"HtmlResource"`
}

// CreateDescribeHtmlResourceRequest creates a request to invoke DescribeHtmlResource API
func CreateDescribeHtmlResourceRequest() (request *DescribeHtmlResourceRequest) {
	request = &DescribeHtmlResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeHtmlResource", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeHtmlResourceResponse creates a response to parse from DescribeHtmlResource response
func CreateDescribeHtmlResourceResponse() (response *DescribeHtmlResourceResponse) {
	response = &DescribeHtmlResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
