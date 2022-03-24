package ens

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

// DescribeLoadBalancerAttribute invokes the ens.DescribeLoadBalancerAttribute API synchronously
func (client *Client) DescribeLoadBalancerAttribute(request *DescribeLoadBalancerAttributeRequest) (response *DescribeLoadBalancerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerAttributeWithChan invokes the ens.DescribeLoadBalancerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerAttributeWithChan(request *DescribeLoadBalancerAttributeRequest) (<-chan *DescribeLoadBalancerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerAttribute(request)
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

// DescribeLoadBalancerAttributeWithCallback invokes the ens.DescribeLoadBalancerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerAttributeWithCallback(request *DescribeLoadBalancerAttributeRequest, callback func(response *DescribeLoadBalancerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerAttribute(request)
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

// DescribeLoadBalancerAttributeRequest is the request struct for api DescribeLoadBalancerAttribute
type DescribeLoadBalancerAttributeRequest struct {
	*requests.RpcRequest
	LoadBalancerId string `position:"Query" name:"LoadBalancerId"`
}

// DescribeLoadBalancerAttributeResponse is the response struct for api DescribeLoadBalancerAttribute
type DescribeLoadBalancerAttributeResponse struct {
	*responses.BaseResponse
	RequestId                 string     `json:"RequestId" xml:"RequestId"`
	LoadBalancerId            string     `json:"LoadBalancerId" xml:"LoadBalancerId"`
	LoadBalancerName          string     `json:"LoadBalancerName" xml:"LoadBalancerName"`
	LoadBalancerStatus        string     `json:"LoadBalancerStatus" xml:"LoadBalancerStatus"`
	EnsRegionId               string     `json:"EnsRegionId" xml:"EnsRegionId"`
	Address                   string     `json:"Address" xml:"Address"`
	NetworkId                 string     `json:"NetworkId" xml:"NetworkId"`
	VSwitchId                 string     `json:"VSwitchId" xml:"VSwitchId"`
	Bandwidth                 int        `json:"Bandwidth" xml:"Bandwidth"`
	LoadBalancerSpec          string     `json:"LoadBalancerSpec" xml:"LoadBalancerSpec"`
	CreateTime                string     `json:"CreateTime" xml:"CreateTime"`
	EndTime                   string     `json:"EndTime" xml:"EndTime"`
	AddressIPVersion          string     `json:"AddressIPVersion" xml:"AddressIPVersion"`
	ListenerPorts             []string   `json:"ListenerPorts" xml:"ListenerPorts"`
	BackendServers            []Rs       `json:"BackendServers" xml:"BackendServers"`
	ListenerPortsAndProtocols []Listener `json:"ListenerPortsAndProtocols" xml:"ListenerPortsAndProtocols"`
}

// CreateDescribeLoadBalancerAttributeRequest creates a request to invoke DescribeLoadBalancerAttribute API
func CreateDescribeLoadBalancerAttributeRequest() (request *DescribeLoadBalancerAttributeRequest) {
	request = &DescribeLoadBalancerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeLoadBalancerAttribute", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeLoadBalancerAttributeResponse creates a response to parse from DescribeLoadBalancerAttribute response
func CreateDescribeLoadBalancerAttributeResponse() (response *DescribeLoadBalancerAttributeResponse) {
	response = &DescribeLoadBalancerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
