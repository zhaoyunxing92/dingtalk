package model

type message struct {
	MsgType string `json:"msgtype" validate:"required,oneof=text image voice file link oa markdown action_card"`
}