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

package model

//卡片按钮
type CardButton struct {
	Title string `json:"title"`
	Url   string `json:"action_url"` //跳转的url
}

//卡片
type Card struct {
	Title          string       `json:"title"`           //是透出到会话列表和通知的文案
	Content        string       `json:"markdown"`        //消息内容，支持markdown，语法参考标准markdown语法。
	SingleTitle    string       `json:"single_title"`    //使用整体跳转ActionCard样式时的标题，必须与single_url同时设置，最长20个字符
	SingleUrl      string       `json:"single_url"`      //消息点击链接地址，当发送消息为小程序时支持小程序跳转链接，最长500个字符。消息链接跳转，
	BtnOrientation string       `json:"btn_orientation"` //使用独立跳转ActionCard样式时的按钮排列方式,1：横向排列必须与btn_json_list同时设置。
	CardButtons    []CardButton `json:"btn_json_list"`
}

//卡片消息
type cardMessage struct {
	message
	Card `json:"action_card"`
}

func newCardMessage(card Card) cardMessage {
	return cardMessage{message: message{MsgType: "action_card"}, Card: card}
}
