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

// ChatAddMember 群添加人员
//{
//    "Operator": "manager164",
//    "OpenConversationId": "cidOuTaLz7/D7zkkM/5PlrQQA==",
//    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
//    "EventType": "chat_add_member",
//    "UserId": [
//        "16399780922588973",
//        "011505184066774889"
//    ],
//    "UnionId": [
//        "iSjYiSU5816QIIG0p3h4iP2UQiEiE",
//        "oiiRiSE5704kGgOZEHGGYmcQiEiE"
//    ],
//    "OperatorUnionId": "ABNiSWeAolg5OETyYT60wdQiEiE",
//    "TimeStamp": 1640601921836,
//    "ChatId": "chat691a41db53b100115bec3603472d78a9"
//}
type ChatAddMember struct {
	ChatRemoveMember
}
