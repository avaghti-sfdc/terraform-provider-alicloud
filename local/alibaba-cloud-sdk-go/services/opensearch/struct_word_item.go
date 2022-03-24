package opensearch

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

// WordItem is a nested struct in opensearch response
type WordItem struct {
	Cmd       string                 `json:"cmd" xml:"cmd"`
	Word      string                 `json:"word" xml:"word"`
	Created   int64                  `json:"created" xml:"created"`
	Updated   int64                  `json:"updated" xml:"updated"`
	Status    string                 `json:"status" xml:"status"`
	Relevance map[string]interface{} `json:"relevance" xml:"relevance"`
	Tokens    []Token                `json:"tokens" xml:"tokens"`
}
