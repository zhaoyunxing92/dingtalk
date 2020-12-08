package model

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type markdown struct {
	Title string `json:"title" validate:"required"`
	Text  string `json:"text" validate:"required"`
}

//markdown消息
type markdownMessage struct {
	message
	markdown `json:"markdown" validate:"required"`
}

func newMarkDownMessage(title, content string) markdownMessage {
	return markdownMessage{message{MsgType: "markdown"}, markdown{title, content}}
}

//请求参数验证
func (req *markdownMessage) Validate(valid *validator.Validate, trans translator.Translator) error {
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
