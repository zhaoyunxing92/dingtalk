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

import "encoding/json"

//CardButton 卡片按钮
type CardButton struct {
	//使用独立跳转ActionCard样式时的按钮的标题，最长20个字符。
	Title string `json:"title,omitempty" validate:"omitempty,max=20"`

	//使用独立跳转ActionCard样式时的跳转链接。
	Url string `json:"action_url,omitempty" validate:"omitempty,url"`
}

//card 卡片
type card struct {
	//是透出到会话列表和通知的文案
	Title string `json:"title,omitempty"`

	//消息内容，支持markdown，语法参考标准markdown语法。
	Content string `json:"markdown,omitempty" validate:"required"`

	//使用整体跳转ActionCard样式时的标题，必须与single_url同时设置，最长20个字符
	SingleTitle string `json:"single_title,omitempty" validate:"omitempty,max=20"`

	//消息点击链接地址，当发送消息为小程序时支持小程序跳转链接，最长500个字符。消息链接跳转，
	SingleUrl string `json:"single_url,omitempty" validate:"omitempty,url,max=500"`

	//使用独立跳转ActionCard样式时的按钮排列方式,1：横向排列必须与btn_json_list同时设置。
	BtnOrientation string `json:"btn_orientation,omitempty" validate:"omitempty,oneof=0 1"`

	//使用独立跳转ActionCard样式时的按钮列表
	CardButtons []CardButton `json:"btn_json_list,omitempty"`
}

//卡片消息
type cardMessage struct {
	message
	card `json:"action_card,omitempty"`
}

func (cm *cardMessage) MessageType() string {
	return "action_card"
}

func (cm *cardMessage) String() string {
	str, _ := json.Marshal(cm)
	return string(str)
}

func NewCardMessage(content string) *cardMessage {
	msg := &cardMessage{}
	msg.MsgType = msg.MessageType()
	msg.Content = content
	return msg
}
