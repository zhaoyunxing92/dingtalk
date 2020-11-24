package domain

//消息
type Message interface {
}

// 钉钉消息结构体
type Text struct {
	Content string `json:"content" validate:"required"`
}

//文本消息
type TextMessage struct {
	MsgType string `json:"msgtype" validate:"required"`
	Text    `json:"text" validate:"required"`
}

type OA struct {
}

//OA
type OAMessage struct {
	MsgType string `json:"msgtype" validate:"required"`
	OA      `json:"oa" validate:"required"`
}

// 文本对象
func newTextMessages(context string) TextMessage {
	return TextMessage{"text", Text{Content: context}}
}

//构建oa消息
func newOaMessage(oa OA) OAMessage {
	return OAMessage{"oa", oa}
}
