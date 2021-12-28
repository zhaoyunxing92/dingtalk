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

// LabelUserChange 员工角色信息发生变更
//{
//    "UserIdList": [
//        "011505184066774889"
//    ],
//    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
//    "EventType": "label_user_change",
//    "LabelIdList": [
//        1299380990
//    ],
//    "action": "add",
//    "TimeStamp": "1640670965490"
//}
type LabelUserChange struct {
	Event

	// 动作
	Action string `json:"action"`

	// 时间戳
	TimeStamp int `json:"TimeStamp,string"`

	// 角色组id
	LabelIds []int `json:"LabelIdList"`

	// 用户id
	UserIds []string `json:"UserIdList"`
}
