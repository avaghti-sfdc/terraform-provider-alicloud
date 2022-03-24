package airec

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

// DescribeSyncReportOutliers invokes the airec.DescribeSyncReportOutliers API synchronously
func (client *Client) DescribeSyncReportOutliers(request *DescribeSyncReportOutliersRequest) (response *DescribeSyncReportOutliersResponse, err error) {
	response = CreateDescribeSyncReportOutliersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSyncReportOutliersWithChan invokes the airec.DescribeSyncReportOutliers API asynchronously
func (client *Client) DescribeSyncReportOutliersWithChan(request *DescribeSyncReportOutliersRequest) (<-chan *DescribeSyncReportOutliersResponse, <-chan error) {
	responseChan := make(chan *DescribeSyncReportOutliersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSyncReportOutliers(request)
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

// DescribeSyncReportOutliersWithCallback invokes the airec.DescribeSyncReportOutliers API asynchronously
func (client *Client) DescribeSyncReportOutliersWithCallback(request *DescribeSyncReportOutliersRequest, callback func(response *DescribeSyncReportOutliersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSyncReportOutliersResponse
		var err error
		defer close(result)
		response, err = client.DescribeSyncReportOutliers(request)
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

// DescribeSyncReportOutliersRequest is the request struct for api DescribeSyncReportOutliers
type DescribeSyncReportOutliersRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"instanceId"`
	LevelType  string           `position:"Query" name:"levelType"`
	EndTime    requests.Integer `position:"Query" name:"endTime"`
	StartTime  requests.Integer `position:"Query" name:"startTime"`
	Type       string           `position:"Query" name:"type"`
	Key        string           `position:"Query" name:"key"`
}

// DescribeSyncReportOutliersResponse is the response struct for api DescribeSyncReportOutliers
type DescribeSyncReportOutliersResponse struct {
	*responses.BaseResponse
	Code      string                 `json:"code" xml:"code"`
	Message   string                 `json:"message" xml:"message"`
	RequestId string                 `json:"requestId" xml:"requestId"`
	Result    map[string]interface{} `json:"result" xml:"result"`
}

// CreateDescribeSyncReportOutliersRequest creates a request to invoke DescribeSyncReportOutliers API
func CreateDescribeSyncReportOutliersRequest() (request *DescribeSyncReportOutliersRequest) {
	request = &DescribeSyncReportOutliersRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "DescribeSyncReportOutliers", "/v2/openapi/instances/[instanceId]/sync-reports/outliers", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeSyncReportOutliersResponse creates a response to parse from DescribeSyncReportOutliers response
func CreateDescribeSyncReportOutliersResponse() (response *DescribeSyncReportOutliersResponse) {
	response = &DescribeSyncReportOutliersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
