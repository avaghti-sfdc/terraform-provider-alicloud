package dataworks_public

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

// TopTenElapsedTimeInstance invokes the dataworks_public.TopTenElapsedTimeInstance API synchronously
func (client *Client) TopTenElapsedTimeInstance(request *TopTenElapsedTimeInstanceRequest) (response *TopTenElapsedTimeInstanceResponse, err error) {
	response = CreateTopTenElapsedTimeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// TopTenElapsedTimeInstanceWithChan invokes the dataworks_public.TopTenElapsedTimeInstance API asynchronously
func (client *Client) TopTenElapsedTimeInstanceWithChan(request *TopTenElapsedTimeInstanceRequest) (<-chan *TopTenElapsedTimeInstanceResponse, <-chan error) {
	responseChan := make(chan *TopTenElapsedTimeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TopTenElapsedTimeInstance(request)
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

// TopTenElapsedTimeInstanceWithCallback invokes the dataworks_public.TopTenElapsedTimeInstance API asynchronously
func (client *Client) TopTenElapsedTimeInstanceWithCallback(request *TopTenElapsedTimeInstanceRequest, callback func(response *TopTenElapsedTimeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TopTenElapsedTimeInstanceResponse
		var err error
		defer close(result)
		response, err = client.TopTenElapsedTimeInstance(request)
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

// TopTenElapsedTimeInstanceRequest is the request struct for api TopTenElapsedTimeInstance
type TopTenElapsedTimeInstanceRequest struct {
	*requests.RpcRequest
	BusinessDate string           `position:"Body" name:"BusinessDate"`
	ProjectId    requests.Integer `position:"Body" name:"ProjectId"`
}

// TopTenElapsedTimeInstanceResponse is the response struct for api TopTenElapsedTimeInstance
type TopTenElapsedTimeInstanceResponse struct {
	*responses.BaseResponse
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	InstanceConsumeTimeRank InstanceConsumeTimeRank `json:"InstanceConsumeTimeRank" xml:"InstanceConsumeTimeRank"`
}

// CreateTopTenElapsedTimeInstanceRequest creates a request to invoke TopTenElapsedTimeInstance API
func CreateTopTenElapsedTimeInstanceRequest() (request *TopTenElapsedTimeInstanceRequest) {
	request = &TopTenElapsedTimeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "TopTenElapsedTimeInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateTopTenElapsedTimeInstanceResponse creates a response to parse from TopTenElapsedTimeInstance response
func CreateTopTenElapsedTimeInstanceResponse() (response *TopTenElapsedTimeInstanceResponse) {
	response = &TopTenElapsedTimeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
