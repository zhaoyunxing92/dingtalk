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

package domain

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

//链接消息
type link struct {
	Title    string `json:"title" validate:"required,max=100,min=1"`  //消息标题，建议100字符以内。
	Describe string `json:"text" validate:"required,max=500,min=1"`   //消息描述，建议500字符以内。
	MediaId  string `json:"picUrl" validate:"required,max=500,min=1"` //图片地址，可以通过上传媒体文件接口获取。
	Url      string `json:"messageUrl" validate:"required"`           //消息点击链接地址，当发送消息为小程序时支持小程序跳转链接
}

type linkMessage struct {
	message
	link `json:"link" validate:"required"`
}

func NewLinkMessage() *linkMessage {
	return &linkMessage{message: message{MsgType: "link"}}
}

//请求参数验证
func (req linkMessage) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(req); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
