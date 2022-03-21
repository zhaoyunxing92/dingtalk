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

import "encoding/json"

// 文件消息
type file struct {
	// 媒体文件ID。引用的媒体文件最大10MB。可以通过上传媒体文件接口获取。
	MediaId string `json:"media_id" validate:"required"`
}

type fileMessage struct {
	message
	file `json:"file" validate:"required"`
}

func (f *fileMessage) MessageType() string {
	return "file"
}

func (f *fileMessage) String() string {
	str, _ := json.Marshal(f)
	return string(str)
}

func NewFileMessage(mediaId string) *fileMessage {
	msg := &fileMessage{}
	msg.MsgType = msg.MessageType()
	msg.file = file{MediaId: mediaId}
	return msg
}
