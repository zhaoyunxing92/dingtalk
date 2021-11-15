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

package domain

//
//import (
//	"errors"
//	translator "github.com/go-playground/universal-translator"
//	"github.com/go-playground/validator/v10"
//	"github.com/zhaoyunxing92/dingtalk/v2/domain/message"
//	"strings"
//)
//
//type RobotMsg struct {
//	At `json:"at"`
//}
//type At struct {
//	AtAll     bool     `json:"isAtAll"`
//	AtMobiles []string `json:"atMobiles"`
//}
//
////TextRobotMsg:消息类型及数据格式
//type TextRobotMsg struct {
//	*message.textMessage
//	RobotMsg
//}
//
////MarkdownRobotMsg:markdown机器人消息
//type MarkdownRobotMsg struct {
//	markdownMessage
//	RobotMsg
//}
//
////NewTextRobotMsg:
//func NewTextRobotMsg(content string) *TextRobotMsg {
//	return &TextRobotMsg{textMessage: message.NewTextMessage(content)}
//}
//
////NewMarkDownRobotMessage:markdown消息
//func NewMarkDownRobotMessage(title, content string) *MarkdownRobotMsg {
//	return &MarkdownRobotMsg{markdownMessage: newMarkDownMessage(title, content)}
//}
//
////FeedCard类型
//
////请求参数验证
//func (t TextRobotMsg) Validate(valid *validator.Validate, trans translator.Translator) error {
//	if err := valid.Struct(t); err != nil {
//		errs := err.(validator.ValidationErrors)
//		var slice []string
//		for _, msg := range errs {
//			slice = append(slice, msg.Translate(trans))
//		}
//		return errors.New(strings.Join(slice, ","))
//	}
//	return nil
//}
//
////请求参数验证
//func (t MarkdownRobotMsg) Validate(valid *validator.Validate, trans translator.Translator) error {
//	if err := valid.Struct(t); err != nil {
//		errs := err.(validator.ValidationErrors)
//		var slice []string
//		for _, msg := range errs {
//			slice = append(slice, msg.Translate(trans))
//		}
//		return errors.New(strings.Join(slice, ","))
//	}
//	return nil
//}
