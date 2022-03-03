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

// ChatRemoveMember 群删除人员
//{
//    "Operator": "manager164",
//    "OpenConversationId": "cidOuTaLz7/D7zkkM/5PlrQQA==",
//    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
//    "EventType": "chat_remove_member",
//    "UserId": [
//        "011505184066774889"
//    ],
//    "UnionId": [
//        "oiiRiSE5704kGgOZEHGGYmcQiEiE"
//    ],
//    "OperatorUnionId": "ABNiSWeAolg5OETyYT60wdQiEiE",
//    "TimeStamp": 1640601958259,
//    "ChatId": "chat691a41db53b100115bec3603472d78a9"
//}
type ChatRemoveMember struct {
	Event

	// 操作人
	Operator string `json:"Operator"`

	// 群回话id
	OpenConversationId string `json:"OpenConversationId"`

	OperatorUnionId string `json:"OperatorUnionId"`

	UserIds []string `json:"UserId"`

	UnionIds []string `json:"UnionId"`

	TimeStamp int `json:"TimeStamp"`

	ChatId string `json:"ChatId"`
}
