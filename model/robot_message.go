package model

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

//消息类型及数据格式
type textRobotMsg struct {
	message
	text
	AtAll     bool     `json:"isAtAll"`
	AtMobiles []string `json:"atMobiles"`
}

func NewTextRobotMsg(content string) *textRobotMsg {
	return &textRobotMsg{message: message{MsgType: "text"}, text: text{Content: content}}
}

//请求参数验证
func (t textRobotMsg) Validate(valid *validator.Validate, trans translator.Translator) error {
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
