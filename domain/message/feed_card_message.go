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
	"errors"
	"strings"
)

import (
	translator "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
)

//"title": "时代的火车向前开1",
//"messageURL": "https://www.dingtalk.com/",
//"picURL": "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png"
//
type FeedCardLink struct {
	Title  string `json:"title" validate:"required"`      //单条信息文本。
	Url    string `json:"messageURL" validate:"required"` //点击单条信息到跳转链接。
	BkgUrl string `json:"picURL"  validate:"required"`    //单条信息后面图片的URL。
}

type feedCard struct {
	Links []FeedCardLink `json:"links"`
}

type feedCardMessage struct {
	message
	feedCard `json:"feedCard"`
}

func NewFeedCardMessage(links []FeedCardLink) feedCardMessage {
	return feedCardMessage{message{MsgType: "feedCard"}, feedCard{Links: links}}
}

//请求参数验证
func (req feedCardMessage) Validate(valid *validator.Validate, trans translator.Translator) error {
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
