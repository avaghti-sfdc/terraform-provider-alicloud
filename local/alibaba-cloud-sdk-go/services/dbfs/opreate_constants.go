package dbfs

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

// OpreateConstants invokes the dbfs.OpreateConstants API synchronously
func (client *Client) OpreateConstants(request *OpreateConstantsRequest) (response *OpreateConstantsResponse, err error) {
	response = CreateOpreateConstantsResponse()
	err = client.DoAction(request, response)
	return
}

// OpreateConstantsWithChan invokes the dbfs.OpreateConstants API asynchronously
func (client *Client) OpreateConstantsWithChan(request *OpreateConstantsRequest) (<-chan *OpreateConstantsResponse, <-chan error) {
	responseChan := make(chan *OpreateConstantsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpreateConstants(request)
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

// OpreateConstantsWithCallback invokes the dbfs.OpreateConstants API asynchronously
func (client *Client) OpreateConstantsWithCallback(request *OpreateConstantsRequest, callback func(response *OpreateConstantsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpreateConstantsResponse
		var err error
		defer close(result)
		response, err = client.OpreateConstants(request)
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

// OpreateConstantsRequest is the request struct for api OpreateConstants
type OpreateConstantsRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	ConstantsData string           `position:"Query" name:"ConstantsData"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// OpreateConstantsResponse is the response struct for api OpreateConstants
type OpreateConstantsResponse struct {
	*responses.BaseResponse
	TotalCount      int64  `json:"TotalCount" xml:"TotalCount"`
	MasterData      string `json:"MasterData" xml:"MasterData"`
	PageSize        int64  `json:"PageSize" xml:"PageSize"`
	ZoneData        string `json:"ZoneData" xml:"ZoneData"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	PageNumber      int64  `json:"PageNumber" xml:"PageNumber"`
	AccessData      string `json:"AccessData" xml:"AccessData"`
	ProductCodeData string `json:"ProductCodeData" xml:"ProductCodeData"`
	OsversionData   string `json:"OsversionData" xml:"OsversionData"`
	Data            string `json:"Data" xml:"Data"`
	RegionData      string `json:"RegionData" xml:"RegionData"`
	EndpointData    string `json:"EndpointData" xml:"EndpointData"`
}

// CreateOpreateConstantsRequest creates a request to invoke OpreateConstants API
func CreateOpreateConstantsRequest() (request *OpreateConstantsRequest) {
	request = &OpreateConstantsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "OpreateConstants", "", "")
	request.Method = requests.POST
	return
}

// CreateOpreateConstantsResponse creates a response to parse from OpreateConstants response
func CreateOpreateConstantsResponse() (response *OpreateConstantsResponse) {
	response = &OpreateConstantsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
