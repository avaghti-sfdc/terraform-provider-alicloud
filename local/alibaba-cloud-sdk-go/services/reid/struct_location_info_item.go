package reid

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

// LocationInfoItem is a nested struct in reid response
type LocationInfoItem struct {
	ExternalId       string   `json:"ExternalId" xml:"ExternalId"`
	LocationType     string   `json:"LocationType" xml:"LocationType"`
	Status           int      `json:"Status" xml:"Status"`
	StoreId          int64    `json:"StoreId" xml:"StoreId"`
	GmtCreate        int64    `json:"GmtCreate" xml:"GmtCreate"`
	ParentLocationId int64    `json:"ParentLocationId" xml:"ParentLocationId"`
	GmtModified      int64    `json:"GmtModified" xml:"GmtModified"`
	LocationName     string   `json:"LocationName" xml:"LocationName"`
	LayerType        string   `json:"LayerType" xml:"LayerType"`
	LocationId       int64    `json:"LocationId" xml:"LocationId"`
	RectRois         RectRois `json:"RectRois" xml:"RectRois"`
}
