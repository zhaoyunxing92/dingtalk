package model

type message struct {
	MsgType string `json:"msgtype" validate:"required,oneof=text image voice file link oa markdown action_card feedCard"`
}

//MessageResponse:发送消息返回
type MessageResponse struct {
	Response
	MessageId string `json:"messageId"` //指定员工的部门信息。
}

//SendToConversationResponse:发送普通消息返回
type SendToConversationResponse struct {
	Response
	Receiver string `json:"receiver"`
}
