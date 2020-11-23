package domain

// 钉钉消息结构体
type Text struct {
	Content string `json:"content" validate:"required"`
}

//文本消息
type TextMessage struct {
	MsgType string `json:"msgtype" validate:"required"`
	Text    Text   `json:"text" validate:"required"`
}

// 文本对象
func NewTextMessages(context string) TextMessage {
	return TextMessage{MsgType: "text", Text: Text{Content: context}}
}
