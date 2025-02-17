package vod

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

// SubmitDetectionJob invokes the vod.SubmitDetectionJob API synchronously
func (client *Client) SubmitDetectionJob(request *SubmitDetectionJobRequest) (response *SubmitDetectionJobResponse, err error) {
	response = CreateSubmitDetectionJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitDetectionJobWithChan invokes the vod.SubmitDetectionJob API asynchronously
func (client *Client) SubmitDetectionJobWithChan(request *SubmitDetectionJobRequest) (<-chan *SubmitDetectionJobResponse, <-chan error) {
	responseChan := make(chan *SubmitDetectionJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitDetectionJob(request)
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

// SubmitDetectionJobWithCallback invokes the vod.SubmitDetectionJob API asynchronously
func (client *Client) SubmitDetectionJobWithCallback(request *SubmitDetectionJobRequest, callback func(response *SubmitDetectionJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitDetectionJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitDetectionJob(request)
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

// SubmitDetectionJobRequest is the request struct for api SubmitDetectionJob
type SubmitDetectionJobRequest struct {
	*requests.RpcRequest
	WhiteListUrls      string           `position:"Query" name:"WhiteListUrls"`
	CopyrightEndTime   string           `position:"Query" name:"CopyrightEndTime"`
	CopyrightStatus    string           `position:"Query" name:"CopyrightStatus"`
	CopyrightBeginTime string           `position:"Query" name:"CopyrightBeginTime"`
	Duration           requests.Integer `position:"Query" name:"Duration"`
	EndTime            string           `position:"Query" name:"EndTime"`
	VideoId            string           `position:"Query" name:"VideoId"`
	BeginTime          string           `position:"Query" name:"BeginTime"`
	ShortVideo         requests.Boolean `position:"Query" name:"ShortVideo"`
	TemplateId         string           `position:"Query" name:"TemplateId"`
	CopyrightFile      string           `position:"Query" name:"CopyrightFile"`
}

// SubmitDetectionJobResponse is the response struct for api SubmitDetectionJob
type SubmitDetectionJobResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	DetectionJob DetectionJob `json:"DetectionJob" xml:"DetectionJob"`
}

// CreateSubmitDetectionJobRequest creates a request to invoke SubmitDetectionJob API
func CreateSubmitDetectionJobRequest() (request *SubmitDetectionJobRequest) {
	request = &SubmitDetectionJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SubmitDetectionJob", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitDetectionJobResponse creates a response to parse from SubmitDetectionJob response
func CreateSubmitDetectionJobResponse() (response *SubmitDetectionJobResponse) {
	response = &SubmitDetectionJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
