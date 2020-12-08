package model

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type RobotMsg struct {
	At `json:"at"`
}
type At struct {
	AtAll     bool     `json:"isAtAll"`
	AtMobiles []string `json:"atMobiles"`
}

//TextRobotMsg:消息类型及数据格式
type TextRobotMsg struct {
	textMessage
	RobotMsg
}

//MarkdownRobotMsg:markdown机器人消息
type MarkdownRobotMsg struct {
	markdownMessage
	RobotMsg
}

func NewTextRobotMsg(content string) *TextRobotMsg {
	return &TextRobotMsg{textMessage: NewTextMessages(content)}
}

//markdown消息
func NewMarkDownRobotMessage(title, content string) *MarkdownRobotMsg {
	return &MarkdownRobotMsg{markdownMessage: NewMarkDownMessage(title, content)}
}

//请求参数验证
func (t TextRobotMsg) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(t); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}

//请求参数验证
func (t MarkdownRobotMsg) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(t); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
