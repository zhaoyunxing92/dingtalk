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

// 链接消息
type link struct {

	// 消息标题，建议100字符以内。
	Title string `json:"title,omitempty" validate:"required,max=100,min=1"`

	// 消息描述，建议500字符以内。
	Describe string `json:"text,omitempty" validate:"required,max=500,min=1"`

	// 图片地址，可以通过上传媒体文件接口获取。
	MediaId string `json:"picUrl,omitempty" validate:"required,max=500,min=1"`

	// 消息点击链接地址，当发送消息为小程序时支持小程序跳转链接
	Url string `json:"messageUrl,omitempty" validate:"required"`
}

type linkMessage struct {
	message
	link `json:"link" validate:"required"`
}

func (l *linkMessage) String() string {
	str, _ := json.Marshal(l)
	return string(str)
}

func (l *linkMessage) MessageType() string {
	return "link"
}

func NewLinkMessage(title, desc, mediaId, url string) *linkMessage {
	msg := &linkMessage{}
	msg.MsgType = msg.MessageType()
	msg.link = link{title, desc, mediaId, url}
	return msg
}
