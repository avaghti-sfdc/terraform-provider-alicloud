package ocr

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

// RecognizeCharacter invokes the ocr.RecognizeCharacter API synchronously
func (client *Client) RecognizeCharacter(request *RecognizeCharacterRequest) (response *RecognizeCharacterResponse, err error) {
	response = CreateRecognizeCharacterResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeCharacterWithChan invokes the ocr.RecognizeCharacter API asynchronously
func (client *Client) RecognizeCharacterWithChan(request *RecognizeCharacterRequest) (<-chan *RecognizeCharacterResponse, <-chan error) {
	responseChan := make(chan *RecognizeCharacterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeCharacter(request)
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

// RecognizeCharacterWithCallback invokes the ocr.RecognizeCharacter API asynchronously
func (client *Client) RecognizeCharacterWithCallback(request *RecognizeCharacterRequest, callback func(response *RecognizeCharacterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeCharacterResponse
		var err error
		defer close(result)
		response, err = client.RecognizeCharacter(request)
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

// RecognizeCharacterRequest is the request struct for api RecognizeCharacter
type RecognizeCharacterRequest struct {
	*requests.RpcRequest
	ImageType         requests.Integer `position:"Body" name:"ImageType"`
	OutputProbability requests.Boolean `position:"Body" name:"OutputProbability"`
	ImageURL          string           `position:"Body" name:"ImageURL"`
	MinHeight         requests.Integer `position:"Body" name:"MinHeight"`
}

// RecognizeCharacterResponse is the response struct for api RecognizeCharacter
type RecognizeCharacterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeCharacterRequest creates a request to invoke RecognizeCharacter API
func CreateRecognizeCharacterRequest() (request *RecognizeCharacterRequest) {
	request = &RecognizeCharacterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "RecognizeCharacter", "ocr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeCharacterResponse creates a response to parse from RecognizeCharacter response
func CreateRecognizeCharacterResponse() (response *RecognizeCharacterResponse) {
	response = &RecognizeCharacterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
