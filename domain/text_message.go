package domain

// 钉钉消息结构体
type text struct {
	Content string `json:"content" validate:"required"`
}

//文本消息
type textMessage struct {
	MsgType string `json:"msgtype" validate:"required"`
	text    `json:"text" validate:"required"`
}

// 文本对象
func newTextMessages(context string) textMessage {
	return textMessage{"text", text{Content: context}}
}
