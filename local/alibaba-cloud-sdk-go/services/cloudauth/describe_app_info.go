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

// DescribeAppInfo invokes the cloudauth.DescribeAppInfo API synchronously
func (client *Client) DescribeAppInfo(request *DescribeAppInfoRequest) (response *DescribeAppInfoResponse, err error) {
	response = CreateDescribeAppInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAppInfoWithChan invokes the cloudauth.DescribeAppInfo API asynchronously
func (client *Client) DescribeAppInfoWithChan(request *DescribeAppInfoRequest) (<-chan *DescribeAppInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeAppInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAppInfo(request)
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

// DescribeAppInfoWithCallback invokes the cloudauth.DescribeAppInfo API asynchronously
func (client *Client) DescribeAppInfoWithCallback(request *DescribeAppInfoRequest, callback func(response *DescribeAppInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAppInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeAppInfo(request)
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

// DescribeAppInfoRequest is the request struct for api DescribeAppInfo
type DescribeAppInfoRequest struct {
	*requests.RpcRequest
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Platform    string           `position:"Query" name:"Platform"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeAppInfoResponse is the response struct for api DescribeAppInfo
type DescribeAppInfoResponse struct {
	*responses.BaseResponse
	RequestId   string    `json:"RequestId" xml:"RequestId"`
	PageSize    int       `json:"PageSize" xml:"PageSize"`
	CurrentPage int       `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int       `json:"TotalCount" xml:"TotalCount"`
	AppInfoList []AppInfo `json:"AppInfoList" xml:"AppInfoList"`
}

// CreateDescribeAppInfoRequest creates a request to invoke DescribeAppInfo API
func CreateDescribeAppInfoRequest() (request *DescribeAppInfoRequest) {
	request = &DescribeAppInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "DescribeAppInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeAppInfoResponse creates a response to parse from DescribeAppInfo response
func CreateDescribeAppInfoResponse() (response *DescribeAppInfoResponse) {
	response = &DescribeAppInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
