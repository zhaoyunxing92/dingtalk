package model

//文件消息
type file struct {
	MediaId string `json:"media_id" validate:"required"` //媒体文件ID。引用的媒体文件最大10MB。可以通过上传媒体文件接口获取。
}

type fileMessage struct {
	message
	file `json:"file" validate:"required"`
}

func NewFileMessage(mediaId string) fileMessage {
	return fileMessage{message: message{MsgType: "file"}, file: file{MediaId: mediaId}}
}

////请求参数验证
//func (req fileMessage) Validate(valid *validator.Validate, trans translator.Translator) error {
//	if err := valid.Struct(req); err != nil {
//		errs := err.(validator.ValidationErrors)
//		var slice []string
//		for _, msg := range errs {
//			slice = append(slice, msg.Translate(trans))
//		}
//		return errors.New(strings.Join(slice, ","))
//	}
//	return nil
//}
