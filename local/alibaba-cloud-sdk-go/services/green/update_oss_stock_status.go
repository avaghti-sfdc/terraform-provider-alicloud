package green

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

// UpdateOssStockStatus invokes the green.UpdateOssStockStatus API synchronously
func (client *Client) UpdateOssStockStatus(request *UpdateOssStockStatusRequest) (response *UpdateOssStockStatusResponse, err error) {
	response = CreateUpdateOssStockStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateOssStockStatusWithChan invokes the green.UpdateOssStockStatus API asynchronously
func (client *Client) UpdateOssStockStatusWithChan(request *UpdateOssStockStatusRequest) (<-chan *UpdateOssStockStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateOssStockStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateOssStockStatus(request)
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

// UpdateOssStockStatusWithCallback invokes the green.UpdateOssStockStatus API asynchronously
func (client *Client) UpdateOssStockStatusWithCallback(request *UpdateOssStockStatusRequest, callback func(response *UpdateOssStockStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateOssStockStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateOssStockStatus(request)
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

// UpdateOssStockStatusRequest is the request struct for api UpdateOssStockStatus
type UpdateOssStockStatusRequest struct {
	*requests.RpcRequest
	ResourceTypeList         string           `position:"Query" name:"ResourceTypeList"`
	VideoFrameInterval       requests.Integer `position:"Query" name:"VideoFrameInterval"`
	StartDate                string           `position:"Query" name:"StartDate"`
	SourceIp                 string           `position:"Query" name:"SourceIp"`
	VideoAutoFreezeSceneList string           `position:"Query" name:"VideoAutoFreezeSceneList"`
	AudioMaxSize             requests.Integer `position:"Query" name:"AudioMaxSize"`
	Lang                     string           `position:"Query" name:"Lang"`
	ImageAutoFreeze          string           `position:"Query" name:"ImageAutoFreeze"`
	AudioAutoFreezeSceneList string           `position:"Query" name:"AudioAutoFreezeSceneList"`
	VideoMaxSize             requests.Integer `position:"Query" name:"VideoMaxSize"`
	AutoFreezeType           string           `position:"Query" name:"AutoFreezeType"`
	EndDate                  string           `position:"Query" name:"EndDate"`
	BucketConfigList         string           `position:"Query" name:"BucketConfigList"`
	SceneList                string           `position:"Query" name:"SceneList"`
	VideoMaxFrames           requests.Integer `position:"Query" name:"VideoMaxFrames"`
	Operation                string           `position:"Query" name:"Operation"`
}

// UpdateOssStockStatusResponse is the response struct for api UpdateOssStockStatus
type UpdateOssStockStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateOssStockStatusRequest creates a request to invoke UpdateOssStockStatus API
func CreateUpdateOssStockStatusRequest() (request *UpdateOssStockStatusRequest) {
	request = &UpdateOssStockStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "UpdateOssStockStatus", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateOssStockStatusResponse creates a response to parse from UpdateOssStockStatus response
func CreateUpdateOssStockStatusResponse() (response *UpdateOssStockStatusResponse) {
	response = &UpdateOssStockStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
