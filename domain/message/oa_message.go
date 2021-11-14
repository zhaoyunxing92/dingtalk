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

import json "encoding/json"

//oaMessageHeader 消息的头
type oaMessageHeader struct {
	//消息头部的背景颜色。长度限制为8个英文字符，其中前2为表示透明度，后6位表示颜色值。不要添加0x。
	BgColor string `json:"bgcolor" validate:"required"`

	//消息的头部标题 (向普通会话发送时有效，向企业会话发送时会被替换为微应用的名字)。长度限制为最多10个字符。
	Text string `json:"text" validate:"required,max=10"`
}

type Form struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//Rich 单行富文本信息。
type Rich struct {
	//单行富文本信息的数目。
	Num string `json:"num"`

	//单行富文本信息的单位。
	Unit string `json:"unit"`
}

//Body 消息体
type Body struct {
	//消息体的标题，建议50个字符以内。
	Title string `json:"title,omitempty" validate:"omitempty,max=50"`

	//消息体的表单，最多显示6个，超过会被隐藏。
	Forms []Form `json:"form,omitempty" validate:"omitempty,lte=6"`

	//消息体的内容，最多显示3行。
	Content string `json:"content,omitempty"`

	//图片id,以@开头的
	MediaId string `json:"image,omitempty" validate:"omitempty,startswith=@"`

	//单行富文本信息
	Rich `json:"rich,omitempty"`

	//自定义的作者名字。
	Author string `json:"author,omitempty"`
}

//StatusBar 消息状态栏。
type StatusBar struct {
	//状态栏文案。
	Value string `json:"status_value,omitempty"`

	//状态栏背景色，默认为黑色，推荐0xFF加六位颜色值。
	BackColor string `json:"status_bg,omitempty"`
}

//OA oa
type oa struct {
	//消息点击链接地址，当发送消息为小程序时支持小程序跳转链接。
	MessageUrl string `json:"message_url" validate:"required"`

	//PC端点击消息时跳转到的地址。
	PcMessageUrl string `json:"pc_message_url" validate:"omitempty,url"`

	//消息状态栏
	StatusBar `json:"status_bar,omitempty"`

	//消息头部内容。
	oaMessageHeader `json:"head,omitempty" validate:"required"`

	//消息体。
	Body `json:"body,omitempty" validate:"required"`
}

//oaMessage oa
type oaMessage struct {
	message
	oa `json:"oa" validate:"required"`
}

func (oa *oaMessage) MessageType() string {
	return "oa"
}

func (oa *oaMessage) String() string {
	str, _ := json.Marshal(oa)
	return string(str)
}

type oaMessageBuilder struct {
	msg *oaMessage
}

//NewOaMessage 构建oa消息
func NewOaMessage(title, bgColor, url, pcUrl string) *oaMessage {
	msg := &oaMessage{}
	msg.MsgType = msg.MessageType()
	msg.MessageUrl = url
	msg.PcMessageUrl = pcUrl
	msg.oaMessageHeader = oaMessageHeader{bgColor, title}
	return msg
}
