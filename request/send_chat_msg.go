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

package request

import "github.com/zhaoyunxing92/dingtalk/v2/domain/message"

type SendChatMessage struct {
	//群会话的ID
	ChatId string `json:"chatid" validate:"required"`

	//消息内容，最长不超过2048个字节。
	Msg message.Message `json:"msg" validate:"required"`
}

func NewSendChatMessage(chatId string, msg message.Message) *SendChatMessage {
	return &SendChatMessage{chatId, msg}
}
