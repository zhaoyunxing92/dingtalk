package domain

//oa消息的头
type Header struct {
	BgColor string `json:"bgcolor" validate:"required"`     //消息头部的背景颜色。长度限制为8个英文字符，其中前2为表示透明度，后6位表示颜色值。不要添加0x。
	Text    string `json:"text" validate:"required,max=10"` //消息的头部标题 (向普通会话发送时有效，向企业会话发送时会被替换为微应用的名字)。长度限制为最多10个字符。
}

type Form struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Body struct {
	Title   string `json:"title" validate:"omitempty,max=50"` //消息体的标题，建议50个字符以内。
	Forms   []Form `json:"form" validate:"omitempty,lte=6"`   //消息体的表单，最多显示6个，超过会被隐藏。
	Content string `json:"content"`                           //消息体的内容，最多显示3行。
	Author  string `json:"author"`                            //自定义的作者名字。
}

// oa
type OA struct {
	MessageUrl   string `json:"message_url" validate:"required"`         //消息点击链接地址，当发送消息为小程序时支持小程序跳转链接。
	PcMessageUrl string `json:"pc_message_url" validate:"omitempty,url"` //PC端点击消息时跳转到的地址。
	Header       `json:"head" validate:"required"`                       //消息头部内容。
	Body         `json:"body" validate:"required"`                       //消息体。

}

//OA
type OAMessage struct {
	MsgType string `json:"msgtype" validate:"required"`
	OA      `json:"oa" validate:"required"`
}

//构建oa消息
func newOaMessage(oa OA) *OAMessage {
	return &OAMessage{"oa", oa}
}
