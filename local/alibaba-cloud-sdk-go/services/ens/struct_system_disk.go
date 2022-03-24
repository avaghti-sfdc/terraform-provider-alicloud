package ens

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

// SystemDisk is a nested struct in ens response
type SystemDisk struct {
	DeviceType string `json:"device_type" xml:"device_type"`
	DiskType   string `json:"disk_type" xml:"disk_type"`
	Size       int    `json:"Size" xml:"Size"`
	DiskName   string `json:"DiskName" xml:"DiskName"`
	Uuid       string `json:"uuid" xml:"uuid"`
	Storage    int    `json:"storage" xml:"storage"`
	DiskId     string `json:"DiskId" xml:"DiskId"`
	Category   string `json:"Category" xml:"Category"`
	Name       string `json:"name" xml:"name"`
}
