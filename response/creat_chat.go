/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package response

type CreatChat struct {
	Response

	//群会话的ID,后续版本中chatid将不再使用，请将openConversationId作为群会话唯一标识
	ChatId string `json:"chatid"`

	//群会话的ID
	OpenConversationId string `json:"openConversationId"`

	//会话类型：
	//
	//2：企业群。
	ConversationTag int `json:"conversationTag"`
}
