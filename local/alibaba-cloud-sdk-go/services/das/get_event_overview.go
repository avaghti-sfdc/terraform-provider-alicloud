package das

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

// GetEventOverview invokes the das.GetEventOverview API synchronously
func (client *Client) GetEventOverview(request *GetEventOverviewRequest) (response *GetEventOverviewResponse, err error) {
	response = CreateGetEventOverviewResponse()
	err = client.DoAction(request, response)
	return
}

// GetEventOverviewWithChan invokes the das.GetEventOverview API asynchronously
func (client *Client) GetEventOverviewWithChan(request *GetEventOverviewRequest) (<-chan *GetEventOverviewResponse, <-chan error) {
	responseChan := make(chan *GetEventOverviewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEventOverview(request)
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

// GetEventOverviewWithCallback invokes the das.GetEventOverview API asynchronously
func (client *Client) GetEventOverviewWithCallback(request *GetEventOverviewRequest, callback func(response *GetEventOverviewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEventOverviewResponse
		var err error
		defer close(result)
		response, err = client.GetEventOverview(request)
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

// GetEventOverviewRequest is the request struct for api GetEventOverview
type GetEventOverviewRequest struct {
	*requests.RpcRequest
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	Tags       string `position:"Query" name:"Tags"`
	InstanceId string `position:"Query" name:"InstanceId"`
	MinLevel   string `position:"Query" name:"MinLevel"`
	TicketId   string `position:"Query" name:"TicketId"`
}

// GetEventOverviewResponse is the response struct for api GetEventOverview
type GetEventOverviewResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateGetEventOverviewRequest creates a request to invoke GetEventOverview API
func CreateGetEventOverviewRequest() (request *GetEventOverviewRequest) {
	request = &GetEventOverviewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetEventOverview", "das", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetEventOverviewResponse creates a response to parse from GetEventOverview response
func CreateGetEventOverviewResponse() (response *GetEventOverviewResponse) {
	response = &GetEventOverviewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
