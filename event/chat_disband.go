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

// ChatDisband 群会话解散群
//{
//    "Operator": "manager3060",
//    "OpenConversationId": "cid740SMooADFxmWX3DvIyp4g==",
//    "CorpId": "ding0761931a826dfde2ffe93478753d9884",
//    "EventType": "chat_disband",
//    "OperatorUnionId": "3bOBVxFv0J8VOu3J4jGhZQiEiE",
//    "TimeStamp": 1640754514696,
//    "ChatId": "chat6418e0c8c141114e18491ea2603c09be"
//}
type ChatDisband struct {
	Event

	// 操作人
	Operator string `json:"Operator"`

	// 群回话id
	OpenConversationId string `json:"OpenConversationId"`

	OperatorUnionId string `json:"OperatorUnionId"`

	TimeStamp int `json:"TimeStamp"`

	ChatId string `json:"ChatId"`
}
