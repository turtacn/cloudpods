// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type Voucher struct {

    /*  (Optional) */
    CashingFee float64 `json:"cashingFee"`

    /*  (Optional) */
    CreateTime string `json:"createTime"`

    /*  (Optional) */
    EndTime string `json:"endTime"`

    /*  (Optional) */
    ErpOrderId string `json:"erpOrderId"`

    /*  (Optional) */
    Fee float64 `json:"fee"`

    /*  (Optional) */
    Id int `json:"id"`

    /*  (Optional) */
    PaymentChannel int `json:"paymentChannel"`

    /*  (Optional) */
    PaymentNumber string `json:"paymentNumber"`

    /*  (Optional) */
    Pin string `json:"pin"`

    /*  (Optional) */
    RechargeType int `json:"rechargeType"`

    /*  (Optional) */
    Status int `json:"status"`

    /*  (Optional) */
    UpdateTime string `json:"updateTime"`

    /*  (Optional) */
    UsableFee float64 `json:"usableFee"`

    /*  (Optional) */
    UserAllFee float64 `json:"userAllFee"`
}