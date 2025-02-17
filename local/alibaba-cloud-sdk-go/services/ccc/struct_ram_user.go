package ccc

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

// RamUser is a nested struct in ccc response
type RamUser struct {
	AliyunUid   int64  `json:"AliyunUid" xml:"AliyunUid"`
	DisplayName string `json:"DisplayName" xml:"DisplayName"`
	Email       string `json:"Email" xml:"Email"`
	LoginName   string `json:"LoginName" xml:"LoginName"`
	Mobile      string `json:"Mobile" xml:"Mobile"`
	Primary     bool   `json:"Primary" xml:"Primary"`
	RamId       string `json:"RamId" xml:"RamId"`
}
