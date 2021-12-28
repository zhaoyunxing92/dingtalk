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

// LabelConfAdd 增加角色或者角色组
// 好像无法区分是角色还是角色组
//{
//    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
//    "EventType": "label_conf_add",
//    "LabelIdList": [
//        2393600321
//    ],
//    "scope": "1",
//    "TimeStamp": "1640672959393"
//}
type LabelConfAdd struct {
	Event

	// 时间
	TimeStamp int `json:"TimeStamp,string"`

	Scope string `json:"scope"`

	// 角色id或者角色组id
	LabelIds []int `json:"LabelIdList"`
}
