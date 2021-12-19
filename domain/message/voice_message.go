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
type voice struct {
	// 媒体文件ID
	MediaId string `json:"media_id" validate:"required"`

	// 正整数，小于60，表示音频时长。
	Duration int `json:"duration,string" validate:"required,max=60"`
}

// 文本消息
type voiceMessage struct {
	message
	voice `json:"voice" validate:"required"`
}

func (v *voiceMessage) String() string {
	str, _ := json.Marshal(v)
	return string(str)
}

func (v *voiceMessage) MessageType() string {
	return "voice"
}

// NewVoiceMessage 语音消息
func NewVoiceMessage(mediaId string, duration int) *voiceMessage {
	msg := &voiceMessage{}
	msg.MsgType = msg.MessageType()
	msg.voice = voice{mediaId, duration}
	return msg
}
