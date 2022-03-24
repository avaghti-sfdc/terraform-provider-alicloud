package unimkt

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

// BrandMessageMigrate invokes the unimkt.BrandMessageMigrate API synchronously
func (client *Client) BrandMessageMigrate(request *BrandMessageMigrateRequest) (response *BrandMessageMigrateResponse, err error) {
	response = CreateBrandMessageMigrateResponse()
	err = client.DoAction(request, response)
	return
}

// BrandMessageMigrateWithChan invokes the unimkt.BrandMessageMigrate API asynchronously
func (client *Client) BrandMessageMigrateWithChan(request *BrandMessageMigrateRequest) (<-chan *BrandMessageMigrateResponse, <-chan error) {
	responseChan := make(chan *BrandMessageMigrateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BrandMessageMigrate(request)
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

// BrandMessageMigrateWithCallback invokes the unimkt.BrandMessageMigrate API asynchronously
func (client *Client) BrandMessageMigrateWithCallback(request *BrandMessageMigrateRequest, callback func(response *BrandMessageMigrateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BrandMessageMigrateResponse
		var err error
		defer close(result)
		response, err = client.BrandMessageMigrate(request)
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

// BrandMessageMigrateRequest is the request struct for api BrandMessageMigrate
type BrandMessageMigrateRequest struct {
	*requests.RpcRequest
	BrandUserId   string `position:"Body" name:"BrandUserId"`
	BrandUserNick string `position:"Body" name:"BrandUserNick"`
	Industry      string `position:"Body" name:"Industry"`
	ProxyUserId   string `position:"Body" name:"ProxyUserId"`
	UserId        string `position:"Body" name:"UserId"`
	ContactName   string `position:"Body" name:"ContactName"`
	AccountNo     string `position:"Body" name:"AccountNo"`
	Company       string `position:"Body" name:"Company"`
	ProxyUserNick string `position:"Body" name:"ProxyUserNick"`
	ContactPhone  string `position:"Body" name:"ContactPhone"`
}

// BrandMessageMigrateResponse is the response struct for api BrandMessageMigrate
type BrandMessageMigrateResponse struct {
	*responses.BaseResponse
	Status    int    `json:"Status" xml:"Status"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBrandMessageMigrateRequest creates a request to invoke BrandMessageMigrate API
func CreateBrandMessageMigrateRequest() (request *BrandMessageMigrateRequest) {
	request = &BrandMessageMigrateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-07", "BrandMessageMigrate", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBrandMessageMigrateResponse creates a response to parse from BrandMessageMigrate response
func CreateBrandMessageMigrateResponse() (response *BrandMessageMigrateResponse) {
	response = &BrandMessageMigrateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
