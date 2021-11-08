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

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

//oa消息的头
type Header struct {
	BgColor string `json:"bgcolor" validate:"required"`     //消息头部的背景颜色。长度限制为8个英文字符，其中前2为表示透明度，后6位表示颜色值。不要添加0x。
	Text    string `json:"text" validate:"required,max=10"` //消息的头部标题 (向普通会话发送时有效，向企业会话发送时会被替换为微应用的名字)。长度限制为最多10个字符。
}

type Form struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//单行富文本信息。
type Rich struct {
	Num  string `json:"num"`  //单行富文本信息的数目。
	Unit string `json:"unit"` //单行富文本信息的单位。
}

//消息体
type Body struct {
	Title   string `json:"title" validate:"omitempty,max=50"`       //消息体的标题，建议50个字符以内。
	Forms   []Form `json:"form" validate:"omitempty,lte=6"`         //消息体的表单，最多显示6个，超过会被隐藏。
	Content string `json:"content"`                                 //消息体的内容，最多显示3行。
	ImageId string `json:"image" validate:"omitempty,startswith=@"` //图片id,以@开头的
	Author  string `json:"author"`                                  //自定义的作者名字。
}

//消息状态栏。
type StatusBar struct {
	Value     string `json:"status_value"` //状态栏文案。
	BackColor string `json:"status_bg"`    //状态栏背景色，默认为黑色，推荐0xFF加六位颜色值。
}

// oa
type OA struct {
	MessageUrl   string                            `json:"message_url" validate:"required"`         //消息点击链接地址，当发送消息为小程序时支持小程序跳转链接。
	PcMessageUrl string                            `json:"pc_message_url" validate:"omitempty,url"` //PC端点击消息时跳转到的地址。
	StatusBar    `json:"status_bar"`               //消息状态栏
	Header       `json:"head" validate:"required"` //消息头部内容。
	Body         `json:"body" validate:"required"` //消息体。
}

//OA
type OAMessage struct {
	message
	OA `json:"oa" validate:"required"`
}

//构建oa消息
func NewOaMessage(oa OA) *OAMessage {
	return &OAMessage{message{MsgType: "oa"}, oa}
}

//请求参数验证
func (o OAMessage) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(o); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
