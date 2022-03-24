package baas

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

// DescribeAntChainCertificateApplications invokes the baas.DescribeAntChainCertificateApplications API synchronously
// api document: https://help.aliyun.com/api/baas/describeantchaincertificateapplications.html
func (client *Client) DescribeAntChainCertificateApplications(request *DescribeAntChainCertificateApplicationsRequest) (response *DescribeAntChainCertificateApplicationsResponse, err error) {
	response = CreateDescribeAntChainCertificateApplicationsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAntChainCertificateApplicationsWithChan invokes the baas.DescribeAntChainCertificateApplications API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchaincertificateapplications.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainCertificateApplicationsWithChan(request *DescribeAntChainCertificateApplicationsRequest) (<-chan *DescribeAntChainCertificateApplicationsResponse, <-chan error) {
	responseChan := make(chan *DescribeAntChainCertificateApplicationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAntChainCertificateApplications(request)
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

// DescribeAntChainCertificateApplicationsWithCallback invokes the baas.DescribeAntChainCertificateApplications API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchaincertificateapplications.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainCertificateApplicationsWithCallback(request *DescribeAntChainCertificateApplicationsRequest, callback func(response *DescribeAntChainCertificateApplicationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAntChainCertificateApplicationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAntChainCertificateApplications(request)
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

// DescribeAntChainCertificateApplicationsRequest is the request struct for api DescribeAntChainCertificateApplications
type DescribeAntChainCertificateApplicationsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
	AntChainId string           `position:"Body" name:"AntChainId"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
	Status     string           `position:"Body" name:"Status"`
}

// DescribeAntChainCertificateApplicationsResponse is the response struct for api DescribeAntChainCertificateApplications
type DescribeAntChainCertificateApplicationsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeAntChainCertificateApplicationsRequest creates a request to invoke DescribeAntChainCertificateApplications API
func CreateDescribeAntChainCertificateApplicationsRequest() (request *DescribeAntChainCertificateApplicationsRequest) {
	request = &DescribeAntChainCertificateApplicationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeAntChainCertificateApplications", "baas", "openAPI")
	return
}

// CreateDescribeAntChainCertificateApplicationsResponse creates a response to parse from DescribeAntChainCertificateApplications response
func CreateDescribeAntChainCertificateApplicationsResponse() (response *DescribeAntChainCertificateApplicationsResponse) {
	response = &DescribeAntChainCertificateApplicationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
