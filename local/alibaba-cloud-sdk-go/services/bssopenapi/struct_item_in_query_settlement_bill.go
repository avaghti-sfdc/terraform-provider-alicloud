package bssopenapi

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

// ItemInQuerySettlementBill is a nested struct in bssopenapi response
type ItemInQuerySettlementBill struct {
	RecordID                    string  `json:"RecordID" xml:"RecordID"`
	Item                        string  `json:"Item" xml:"Item"`
	PayerAccount                string  `json:"PayerAccount" xml:"PayerAccount"`
	OwnerID                     string  `json:"OwnerID" xml:"OwnerID"`
	CreateTime                  string  `json:"CreateTime" xml:"CreateTime"`
	UsageStartTime              string  `json:"UsageStartTime" xml:"UsageStartTime"`
	UsageEndTime                string  `json:"UsageEndTime" xml:"UsageEndTime"`
	SuborderID                  string  `json:"SuborderID" xml:"SuborderID"`
	OrderID                     string  `json:"OrderID" xml:"OrderID"`
	OrderType                   string  `json:"OrderType" xml:"OrderType"`
	LinkedCustomerOrderID       string  `json:"LinkedCustomerOrderID" xml:"LinkedCustomerOrderID"`
	OriginalOrderID             string  `json:"OriginalOrderID" xml:"OriginalOrderID"`
	PaymentTime                 string  `json:"PaymentTime" xml:"PaymentTime"`
	SolutionID                  string  `json:"SolutionID" xml:"SolutionID"`
	SolutionName                string  `json:"SolutionName" xml:"SolutionName"`
	BillID                      string  `json:"BillID" xml:"BillID"`
	ProductCode                 string  `json:"ProductCode" xml:"ProductCode"`
	ProductType                 string  `json:"ProductType" xml:"ProductType"`
	SubscriptionType            string  `json:"SubscriptionType" xml:"SubscriptionType"`
	Region                      string  `json:"Region" xml:"Region"`
	Config                      string  `json:"Config" xml:"Config"`
	Quantity                    string  `json:"Quantity" xml:"Quantity"`
	PretaxGrossAmount           float64 `json:"PretaxGrossAmount" xml:"PretaxGrossAmount"`
	ChargeDiscount              float64 `json:"ChargeDiscount" xml:"ChargeDiscount"`
	DeductedByCoupons           float64 `json:"DeductedByCoupons" xml:"DeductedByCoupons"`
	AccountDiscount             float64 `json:"AccountDiscount" xml:"AccountDiscount"`
	Promotion                   string  `json:"Promotion" xml:"Promotion"`
	PretaxAmount                float64 `json:"PretaxAmount" xml:"PretaxAmount"`
	Currency                    string  `json:"Currency" xml:"Currency"`
	PretaxAmountLocal           float64 `json:"PretaxAmountLocal" xml:"PretaxAmountLocal"`
	PreviousBillingCycleBalance float64 `json:"PreviousBillingCycleBalance" xml:"PreviousBillingCycleBalance"`
	Tax                         float64 `json:"Tax" xml:"Tax"`
	AfterTaxAmount              float64 `json:"AfterTaxAmount" xml:"AfterTaxAmount"`
	Status                      string  `json:"Status" xml:"Status"`
	ClearedTime                 string  `json:"ClearedTime" xml:"ClearedTime"`
	OutstandingAmount           float64 `json:"OutstandingAmount" xml:"OutstandingAmount"`
	DeductedByCashCoupons       float64 `json:"DeductedByCashCoupons" xml:"DeductedByCashCoupons"`
	DeductedByPrepaidCard       float64 `json:"DeductedByPrepaidCard" xml:"DeductedByPrepaidCard"`
	MybankPaymentAmount         float64 `json:"MybankPaymentAmount" xml:"MybankPaymentAmount"`
	PaymentAmount               float64 `json:"PaymentAmount" xml:"PaymentAmount"`
	PaymentCurrency             string  `json:"PaymentCurrency" xml:"PaymentCurrency"`
	Seller                      string  `json:"Seller" xml:"Seller"`
	InvoiceNo                   string  `json:"InvoiceNo" xml:"InvoiceNo"`
}
