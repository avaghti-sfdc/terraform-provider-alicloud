package privatelink

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

// ListVpcEndpointZones invokes the privatelink.ListVpcEndpointZones API synchronously
func (client *Client) ListVpcEndpointZones(request *ListVpcEndpointZonesRequest) (response *ListVpcEndpointZonesResponse, err error) {
	response = CreateListVpcEndpointZonesResponse()
	err = client.DoAction(request, response)
	return
}

// ListVpcEndpointZonesWithChan invokes the privatelink.ListVpcEndpointZones API asynchronously
func (client *Client) ListVpcEndpointZonesWithChan(request *ListVpcEndpointZonesRequest) (<-chan *ListVpcEndpointZonesResponse, <-chan error) {
	responseChan := make(chan *ListVpcEndpointZonesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVpcEndpointZones(request)
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

// ListVpcEndpointZonesWithCallback invokes the privatelink.ListVpcEndpointZones API asynchronously
func (client *Client) ListVpcEndpointZonesWithCallback(request *ListVpcEndpointZonesRequest, callback func(response *ListVpcEndpointZonesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVpcEndpointZonesResponse
		var err error
		defer close(result)
		response, err = client.ListVpcEndpointZones(request)
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

// ListVpcEndpointZonesRequest is the request struct for api ListVpcEndpointZones
type ListVpcEndpointZonesRequest struct {
	*requests.RpcRequest
	EndpointId string           `position:"Query" name:"EndpointId"`
	NextToken  string           `position:"Query" name:"NextToken"`
	MaxResults requests.Integer `position:"Query" name:"MaxResults"`
}

// ListVpcEndpointZonesResponse is the response struct for api ListVpcEndpointZones
type ListVpcEndpointZonesResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	NextToken  string `json:"NextToken" xml:"NextToken"`
	MaxResults string `json:"MaxResults" xml:"MaxResults"`
	Zones      []Zone `json:"Zones" xml:"Zones"`
}

// CreateListVpcEndpointZonesRequest creates a request to invoke ListVpcEndpointZones API
func CreateListVpcEndpointZonesRequest() (request *ListVpcEndpointZonesRequest) {
	request = &ListVpcEndpointZonesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Privatelink", "2020-04-15", "ListVpcEndpointZones", "privatelink", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListVpcEndpointZonesResponse creates a response to parse from ListVpcEndpointZones response
func CreateListVpcEndpointZonesResponse() (response *ListVpcEndpointZonesResponse) {
	response = &ListVpcEndpointZonesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
