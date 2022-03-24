package cloudcallcenter

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

// ModifyMedia invokes the cloudcallcenter.ModifyMedia API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifymedia.html
func (client *Client) ModifyMedia(request *ModifyMediaRequest) (response *ModifyMediaResponse, err error) {
	response = CreateModifyMediaResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMediaWithChan invokes the cloudcallcenter.ModifyMedia API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifymedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMediaWithChan(request *ModifyMediaRequest) (<-chan *ModifyMediaResponse, <-chan error) {
	responseChan := make(chan *ModifyMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMedia(request)
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

// ModifyMediaWithCallback invokes the cloudcallcenter.ModifyMedia API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifymedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMediaWithCallback(request *ModifyMediaRequest, callback func(response *ModifyMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMediaResponse
		var err error
		defer close(result)
		response, err = client.ModifyMedia(request)
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

// ModifyMediaRequest is the request struct for api ModifyMedia
type ModifyMediaRequest struct {
	*requests.RpcRequest
	Description  string `position:"Query" name:"Description"`
	OssFilePath  string `position:"Query" name:"OssFilePath"`
	UploadResult string `position:"Query" name:"UploadResult"`
	Type         string `position:"Query" name:"Type"`
	Content      string `position:"Query" name:"Content"`
	OssFileName  string `position:"Query" name:"OssFileName"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	FileName     string `position:"Query" name:"FileName"`
	Name         string `position:"Query" name:"Name"`
}

// ModifyMediaResponse is the response struct for api ModifyMedia
type ModifyMediaResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Success          bool             `json:"Success" xml:"Success"`
	Code             string           `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	HttpStatusCode   int              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	MediaUploadParam MediaUploadParam `json:"MediaUploadParam" xml:"MediaUploadParam"`
}

// CreateModifyMediaRequest creates a request to invoke ModifyMedia API
func CreateModifyMediaRequest() (request *ModifyMediaRequest) {
	request = &ModifyMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyMedia", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyMediaResponse creates a response to parse from ModifyMedia response
func CreateModifyMediaResponse() (response *ModifyMediaResponse) {
	response = &ModifyMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
