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

// OrgUserModify 通讯录用户更改
//{
//    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
//    "EventType": "user_modify_org",
//    "UserId": [
//        "manager164"
//    ],
//    "OptStaffId": "manager164",
//    "TimeStamp": "1640657432022"
//}
type OrgUserModify struct {
	Event

	// 时间戳
	TimeStamp int `json:"TimeStamp,string"`

	// 操作人，不一定会有，智能人事入职的人员没有该字段
	OptStaffId string `json:"OptStaffId"`

	UserIds []string `json:"UserId"`
}
