package nlp_automl

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

// DeployModel invokes the nlp_automl.DeployModel API synchronously
func (client *Client) DeployModel(request *DeployModelRequest) (response *DeployModelResponse, err error) {
	response = CreateDeployModelResponse()
	err = client.DoAction(request, response)
	return
}

// DeployModelWithChan invokes the nlp_automl.DeployModel API asynchronously
func (client *Client) DeployModelWithChan(request *DeployModelRequest) (<-chan *DeployModelResponse, <-chan error) {
	responseChan := make(chan *DeployModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeployModel(request)
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

// DeployModelWithCallback invokes the nlp_automl.DeployModel API asynchronously
func (client *Client) DeployModelWithCallback(request *DeployModelRequest, callback func(response *DeployModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeployModelResponse
		var err error
		defer close(result)
		response, err = client.DeployModel(request)
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

// DeployModelRequest is the request struct for api DeployModel
type DeployModelRequest struct {
	*requests.RpcRequest
	Product      string           `position:"Body" name:"Product"`
	ModelId      requests.Integer `position:"Body" name:"ModelId"`
	OptType      string           `position:"Body" name:"OptType"`
	ProjectId    requests.Integer `position:"Body" name:"ProjectId"`
	ModelVersion string           `position:"Body" name:"ModelVersion"`
}

// DeployModelResponse is the response struct for api DeployModel
type DeployModelResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Data      map[string]interface{} `json:"Data" xml:"Data"`
	Code      int                    `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	Success   bool                   `json:"Success" xml:"Success"`
}

// CreateDeployModelRequest creates a request to invoke DeployModel API
func CreateDeployModelRequest() (request *DeployModelRequest) {
	request = &DeployModelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("nlp-automl", "2019-11-11", "DeployModel", "nlpautoml", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeployModelResponse creates a response to parse from DeployModel response
func CreateDeployModelResponse() (response *DeployModelResponse) {
	response = &DeployModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
