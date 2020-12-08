package model

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

// 钉钉消息结构体
type text struct {
	Content string `json:"content" validate:"required"`
}

//文本消息
type textMessage struct {
	message
	text `json:"text" validate:"required"`
}

// 文本对象
func NewTextMessage(context string) textMessage {
	return textMessage{message: message{MsgType: "text"}, text: text{Content: context}}
}

//请求参数验证
func (t textMessage) Validate(valid *validator.Validate, trans translator.Translator) error {
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
