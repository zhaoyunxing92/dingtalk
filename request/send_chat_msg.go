package request

import "github.com/zhaoyunxing92/dingtalk/v2/domain/message"

type SendChatMessage struct {
	//群会话的ID
	ChatId string `json:"chatid" validate:"required"`

	//消息内容，最长不超过2048个字节。
	Msg message.Message `json:"msg" validate:"required"`
}

func NewSendChatMessage(chatId string, msg message.Message) *SendChatMessage {
	return &SendChatMessage{chatId, msg}
}
