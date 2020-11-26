package model

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
func newTextMessages(context string) textMessage {
	return textMessage{message: message{MsgType: "text"}, text: text{Content: context}}
}
