/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package event

// CrmCustomerTrack CRM客户动态变更
//{
//    "CorpId": "dingf5d3534a855bdef9a39a90f97fcb1e09",
//    "corpId": "dingf5d3534a855bdef9a39a90f97fcb1e09",
//    "EventType": "crm_customer_track",
//    "EventTime": 1640749233707,
//    "BizId": "21057D5806944EF9C93192B36DC5F7B4",
//    "tracks": [
//        {
//            "creator": "manager4268",
//            "corpId": "dingf5d3534a855bdef9a39a90f97fcb1e09",
//            "customerId": "ad6b2b05-91f1-4393-8ed5-96e9e558b604",
//            "subType": 1,
//            "id": "qoZCjgT7of7XAtUjo/4JgFy+RR432dJXcXqRbNxwTVM=",
//            "gmtCreate": 1640749233707,
//            "type": 1
//        }
//    ]
//}
type CrmCustomerTrack struct {
	Event

	TimeStamp int `json:"EventTime"`

	BizId string `json:"BizId"`

	Tracks []Track `json:"tracks"`
}

type Track struct {
	Creator string `json:"creator"`

	CorpId string `json:"corpId"`

	CustomerId string `json:"customerId"`

	SubType int `json:"subType"`

	Id string `json:"id"`

	GmtCreate int `json:"gmtCreate"`

	Type int `json:"type"`
}
