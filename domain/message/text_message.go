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

package message

import (
	"encoding/json"
)

// 钉钉消息结构体
type text struct {
	Content string `json:"content" validate:"required"`
}

//文本消息
type textMessage struct {
	message
	text `json:"text" validate:"required"`
}

func (t *textMessage) String() string {
	str, _ := json.Marshal(t)
	return string(str)
}

func (t *textMessage) MessageType() string {
	return "text"
}

//NewTextMessage 文本对象
func NewTextMessage(context string) *textMessage {
	msg := &textMessage{}
	msg.MsgType = msg.MessageType()
	msg.text = text{Content: context}
	return msg
}
