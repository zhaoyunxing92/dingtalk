package response

type CreatChat struct {
	Response

	//群会话的ID,后续版本中chatid将不再使用，请将openConversationId作为群会话唯一标识
	ChatId string `json:"chatid"`

	//群会话的ID
	OpenConversationId string `json:"openConversationId"`

	//会话类型：
	//
	//2：企业群。
	ConversationTag int `json:"conversationTag"`
}
