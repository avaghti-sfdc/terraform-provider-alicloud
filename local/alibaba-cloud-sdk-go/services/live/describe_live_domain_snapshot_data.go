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

// DescribeLiveDomainSnapshotData invokes the live.DescribeLiveDomainSnapshotData API synchronously
func (client *Client) DescribeLiveDomainSnapshotData(request *DescribeLiveDomainSnapshotDataRequest) (response *DescribeLiveDomainSnapshotDataResponse, err error) {
	response = CreateDescribeLiveDomainSnapshotDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainSnapshotDataWithChan invokes the live.DescribeLiveDomainSnapshotData API asynchronously
func (client *Client) DescribeLiveDomainSnapshotDataWithChan(request *DescribeLiveDomainSnapshotDataRequest) (<-chan *DescribeLiveDomainSnapshotDataResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainSnapshotDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainSnapshotData(request)
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

// DescribeLiveDomainSnapshotDataWithCallback invokes the live.DescribeLiveDomainSnapshotData API asynchronously
func (client *Client) DescribeLiveDomainSnapshotDataWithCallback(request *DescribeLiveDomainSnapshotDataRequest, callback func(response *DescribeLiveDomainSnapshotDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainSnapshotDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainSnapshotData(request)
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

// DescribeLiveDomainSnapshotDataRequest is the request struct for api DescribeLiveDomainSnapshotData
type DescribeLiveDomainSnapshotDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveDomainSnapshotDataResponse is the response struct for api DescribeLiveDomainSnapshotData
type DescribeLiveDomainSnapshotDataResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	SnapshotDataInfos SnapshotDataInfos `json:"SnapshotDataInfos" xml:"SnapshotDataInfos"`
}

// CreateDescribeLiveDomainSnapshotDataRequest creates a request to invoke DescribeLiveDomainSnapshotData API
func CreateDescribeLiveDomainSnapshotDataRequest() (request *DescribeLiveDomainSnapshotDataRequest) {
	request = &DescribeLiveDomainSnapshotDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainSnapshotData", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainSnapshotDataResponse creates a response to parse from DescribeLiveDomainSnapshotData response
func CreateDescribeLiveDomainSnapshotDataResponse() (response *DescribeLiveDomainSnapshotDataResponse) {
	response = &DescribeLiveDomainSnapshotDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
