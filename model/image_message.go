package model

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

//图片消息
// 钉钉消息结构体
type image struct {
	MediaId string `json:"media_id" validate:"required"`
}

//文本消息
type imageMessage struct {
	message
	image `json:"image" validate:"required"`
}

// 文本对象
func NewImageMessages(mediaId string) imageMessage {
	return imageMessage{message{MsgType: "image"}, image{MediaId: mediaId}}
}

//请求参数验证
func (i imageMessage) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
