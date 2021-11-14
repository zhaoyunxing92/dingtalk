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

package message

//图片消息
type image struct {
	MediaId string `json:"media_id" validate:"required"`
}

//文本消息
type imageMessage struct {
	message
	image `json:"image" validate:"required"`
}

func (im *imageMessage) MessageType() string {
	return "image"
}

//NewImageMessages 文本对象
func NewImageMessages(mediaId string) *imageMessage {
	msg := &imageMessage{}
	msg.MsgType = msg.MessageType()
	msg.image = image{MediaId: mediaId}
	return msg
}
