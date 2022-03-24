package outboundbot

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

// QueryScriptWaveforms invokes the outboundbot.QueryScriptWaveforms API synchronously
func (client *Client) QueryScriptWaveforms(request *QueryScriptWaveformsRequest) (response *QueryScriptWaveformsResponse, err error) {
	response = CreateQueryScriptWaveformsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryScriptWaveformsWithChan invokes the outboundbot.QueryScriptWaveforms API asynchronously
func (client *Client) QueryScriptWaveformsWithChan(request *QueryScriptWaveformsRequest) (<-chan *QueryScriptWaveformsResponse, <-chan error) {
	responseChan := make(chan *QueryScriptWaveformsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryScriptWaveforms(request)
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

// QueryScriptWaveformsWithCallback invokes the outboundbot.QueryScriptWaveforms API asynchronously
func (client *Client) QueryScriptWaveformsWithCallback(request *QueryScriptWaveformsRequest, callback func(response *QueryScriptWaveformsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryScriptWaveformsResponse
		var err error
		defer close(result)
		response, err = client.QueryScriptWaveforms(request)
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

// QueryScriptWaveformsRequest is the request struct for api QueryScriptWaveforms
type QueryScriptWaveformsRequest struct {
	*requests.RpcRequest
	ScriptId      string `position:"Query" name:"ScriptId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	ScriptContent string `position:"Query" name:"ScriptContent"`
}

// QueryScriptWaveformsResponse is the response struct for api QueryScriptWaveforms
type QueryScriptWaveformsResponse struct {
	*responses.BaseResponse
	Code            string           `json:"Code" xml:"Code"`
	HttpStatusCode  int              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message         string           `json:"Message" xml:"Message"`
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	Success         bool             `json:"Success" xml:"Success"`
	ScriptWaveforms []ScriptWaveform `json:"ScriptWaveforms" xml:"ScriptWaveforms"`
}

// CreateQueryScriptWaveformsRequest creates a request to invoke QueryScriptWaveforms API
func CreateQueryScriptWaveformsRequest() (request *QueryScriptWaveformsRequest) {
	request = &QueryScriptWaveformsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "QueryScriptWaveforms", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryScriptWaveformsResponse creates a response to parse from QueryScriptWaveforms response
func CreateQueryScriptWaveformsResponse() (response *QueryScriptWaveformsResponse) {
	response = &QueryScriptWaveformsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
