package edas

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

// GetK8sStorageInfo invokes the edas.GetK8sStorageInfo API synchronously
func (client *Client) GetK8sStorageInfo(request *GetK8sStorageInfoRequest) (response *GetK8sStorageInfoResponse, err error) {
	response = CreateGetK8sStorageInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetK8sStorageInfoWithChan invokes the edas.GetK8sStorageInfo API asynchronously
func (client *Client) GetK8sStorageInfoWithChan(request *GetK8sStorageInfoRequest) (<-chan *GetK8sStorageInfoResponse, <-chan error) {
	responseChan := make(chan *GetK8sStorageInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetK8sStorageInfo(request)
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

// GetK8sStorageInfoWithCallback invokes the edas.GetK8sStorageInfo API asynchronously
func (client *Client) GetK8sStorageInfoWithCallback(request *GetK8sStorageInfoRequest, callback func(response *GetK8sStorageInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetK8sStorageInfoResponse
		var err error
		defer close(result)
		response, err = client.GetK8sStorageInfo(request)
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

// GetK8sStorageInfoRequest is the request struct for api GetK8sStorageInfo
type GetK8sStorageInfoRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// GetK8sStorageInfoResponse is the response struct for api GetK8sStorageInfo
type GetK8sStorageInfoResponse struct {
	*responses.BaseResponse
	Code        int         `json:"Code" xml:"Code"`
	Message     string      `json:"Message" xml:"Message"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	StorageInfo StorageInfo `json:"StorageInfo" xml:"StorageInfo"`
}

// CreateGetK8sStorageInfoRequest creates a request to invoke GetK8sStorageInfo API
func CreateGetK8sStorageInfoRequest() (request *GetK8sStorageInfoRequest) {
	request = &GetK8sStorageInfoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetK8sStorageInfo", "/pop/v5/k8s/acs/k8s_storage", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetK8sStorageInfoResponse creates a response to parse from GetK8sStorageInfo response
func CreateGetK8sStorageInfoResponse() (response *GetK8sStorageInfoResponse) {
	response = &GetK8sStorageInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
