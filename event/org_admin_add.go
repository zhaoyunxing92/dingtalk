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

// OrgAdminAdd 通讯录用户被设为管理员
// 会一起触发label_user_change事件
//{
//    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
//    "EventType": "org_admin_add",
//    "UserId": [
//        "011505184066774889"
//    ],
//    "TimeStamp": "1640670965261"
//}
type OrgAdminAdd struct {
	Event

	// 时间戳
	TimeStamp int `json:"TimeStamp,string"`

	UserIds []string `json:"UserId"`
}
