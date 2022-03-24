package imageprocess

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

// Element is a nested struct in imageprocess response
type Element struct {
	Category       string  `json:"Category" xml:"Category"`
	ImageY         float64 `json:"ImageY" xml:"ImageY"`
	Confidence     float64 `json:"Confidence" xml:"Confidence"`
	SOPInstanceUID string  `json:"SOPInstanceUID" xml:"SOPInstanceUID"`
	Z              float64 `json:"Z" xml:"Z"`
	Lobe           string  `json:"Lobe" xml:"Lobe"`
	ImageX         float64 `json:"ImageX" xml:"ImageX"`
	Y              float64 `json:"Y" xml:"Y"`
	Volume         float64 `json:"Volume" xml:"Volume"`
	X              float64 `json:"X" xml:"X"`
	Diameter       float64 `json:"Diameter" xml:"Diameter"`
	MeanValue      float64 `json:"MeanValue" xml:"MeanValue"`
	ImageZ         float64 `json:"ImageZ" xml:"ImageZ"`
	Lung           string  `json:"Lung" xml:"Lung"`
}
