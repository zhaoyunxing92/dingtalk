package request

type ChatQRCode struct {
	//群会话的chatid
	ChatId string `json:"chatid,omitempty" validate:"required"`

	//分享二维码用户的userid
	UserId string `json:"userid,omitempty" validate:"required"`
}

func NewChatQRCode(chatId, userId string) *ChatQRCode {
	return &ChatQRCode{chatId, userId}
}
