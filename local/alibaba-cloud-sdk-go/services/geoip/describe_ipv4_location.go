package geoip

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

// DescribeIpv4Location invokes the geoip.DescribeIpv4Location API synchronously
func (client *Client) DescribeIpv4Location(request *DescribeIpv4LocationRequest) (response *DescribeIpv4LocationResponse, err error) {
	response = CreateDescribeIpv4LocationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIpv4LocationWithChan invokes the geoip.DescribeIpv4Location API asynchronously
func (client *Client) DescribeIpv4LocationWithChan(request *DescribeIpv4LocationRequest) (<-chan *DescribeIpv4LocationResponse, <-chan error) {
	responseChan := make(chan *DescribeIpv4LocationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIpv4Location(request)
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

// DescribeIpv4LocationWithCallback invokes the geoip.DescribeIpv4Location API asynchronously
func (client *Client) DescribeIpv4LocationWithCallback(request *DescribeIpv4LocationRequest, callback func(response *DescribeIpv4LocationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIpv4LocationResponse
		var err error
		defer close(result)
		response, err = client.DescribeIpv4Location(request)
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

// DescribeIpv4LocationRequest is the request struct for api DescribeIpv4Location
type DescribeIpv4LocationRequest struct {
	*requests.RpcRequest
	Ip           string `position:"Query" name:"Ip"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeIpv4LocationResponse is the response struct for api DescribeIpv4Location
type DescribeIpv4LocationResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Ip          string `json:"Ip" xml:"Ip"`
	Country     string `json:"Country" xml:"Country"`
	Province    string `json:"Province" xml:"Province"`
	City        string `json:"City" xml:"City"`
	County      string `json:"County" xml:"County"`
	Isp         string `json:"Isp" xml:"Isp"`
	CountryCode string `json:"CountryCode" xml:"CountryCode"`
	CountryEn   string `json:"CountryEn" xml:"CountryEn"`
	ProvinceEn  string `json:"ProvinceEn" xml:"ProvinceEn"`
	CityEn      string `json:"CityEn" xml:"CityEn"`
	Longitude   string `json:"Longitude" xml:"Longitude"`
	Latitude    string `json:"Latitude" xml:"Latitude"`
}

// CreateDescribeIpv4LocationRequest creates a request to invoke DescribeIpv4Location API
func CreateDescribeIpv4LocationRequest() (request *DescribeIpv4LocationRequest) {
	request = &DescribeIpv4LocationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("geoip", "2020-01-01", "DescribeIpv4Location", "geoip", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeIpv4LocationResponse creates a response to parse from DescribeIpv4Location response
func CreateDescribeIpv4LocationResponse() (response *DescribeIpv4LocationResponse) {
	response = &DescribeIpv4LocationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
