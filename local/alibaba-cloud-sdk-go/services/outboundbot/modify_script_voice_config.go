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

// ModifyScriptVoiceConfig invokes the outboundbot.ModifyScriptVoiceConfig API synchronously
func (client *Client) ModifyScriptVoiceConfig(request *ModifyScriptVoiceConfigRequest) (response *ModifyScriptVoiceConfigResponse, err error) {
	response = CreateModifyScriptVoiceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyScriptVoiceConfigWithChan invokes the outboundbot.ModifyScriptVoiceConfig API asynchronously
func (client *Client) ModifyScriptVoiceConfigWithChan(request *ModifyScriptVoiceConfigRequest) (<-chan *ModifyScriptVoiceConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyScriptVoiceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScriptVoiceConfig(request)
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

// ModifyScriptVoiceConfigWithCallback invokes the outboundbot.ModifyScriptVoiceConfig API asynchronously
func (client *Client) ModifyScriptVoiceConfigWithCallback(request *ModifyScriptVoiceConfigRequest, callback func(response *ModifyScriptVoiceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScriptVoiceConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyScriptVoiceConfig(request)
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

// ModifyScriptVoiceConfigRequest is the request struct for api ModifyScriptVoiceConfig
type ModifyScriptVoiceConfigRequest struct {
	*requests.RpcRequest
	Type                   string `position:"Query" name:"Type"`
	ScriptId               string `position:"Query" name:"ScriptId"`
	ScriptVoiceConfigId    string `position:"Query" name:"ScriptVoiceConfigId"`
	InstanceId             string `position:"Query" name:"InstanceId"`
	ScriptWaveformRelation string `position:"Query" name:"ScriptWaveformRelation"`
}

// ModifyScriptVoiceConfigResponse is the response struct for api ModifyScriptVoiceConfig
type ModifyScriptVoiceConfigResponse struct {
	*responses.BaseResponse
	Code              string            `json:"Code" xml:"Code"`
	HttpStatusCode    int               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message           string            `json:"Message" xml:"Message"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Success           bool              `json:"Success" xml:"Success"`
	ScriptVoiceConfig ScriptVoiceConfig `json:"ScriptVoiceConfig" xml:"ScriptVoiceConfig"`
}

// CreateModifyScriptVoiceConfigRequest creates a request to invoke ModifyScriptVoiceConfig API
func CreateModifyScriptVoiceConfigRequest() (request *ModifyScriptVoiceConfigRequest) {
	request = &ModifyScriptVoiceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ModifyScriptVoiceConfig", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyScriptVoiceConfigResponse creates a response to parse from ModifyScriptVoiceConfig response
func CreateModifyScriptVoiceConfigResponse() (response *ModifyScriptVoiceConfigResponse) {
	response = &ModifyScriptVoiceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
