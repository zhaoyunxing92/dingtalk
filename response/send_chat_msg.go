package response

type SendChatMessage struct {
	Response

	//指定员工的部门信息。
	MessageId string `json:"messageId"`
}
