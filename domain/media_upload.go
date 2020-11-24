package domain

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

//媒体文件类型：
//image：图片，图片最大1MB，支持JPG格式。
//voice：语音，语音文件最大2MB，支持AMR格式。
//file：普通文件，最大10MB。支持上传以下文件：doc、docx、xls、xlsx、ppt、pptx、zip、pdf、rar。
type MediaUpload struct {
	Response
	Type      string `json:"type"`       // 媒体文件类型
	MediaId   []byte `json:"media_id"`   // 媒体文件类型
	CreatedAt int    `json:"created_at"` // 媒体文件类型
}

//请求
type UploadFile struct {
	Type  string `json:"type" validate:"required,oneof=image voice file"` // 媒体文件类型
	Media []byte `json:"media" validate:"required"`                       // 媒体文件类型
}

func NewUploadFile(tp string, media []byte) UploadFile {
	return UploadFile{tp, media}
}

//请求参数验证
func (req *UploadFile) Validate(valid *validator.Validate, trans translator.Translator) error {
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
