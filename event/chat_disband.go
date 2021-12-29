package event

// ChatDisband 群会话解散群
//{
//    "Operator": "manager3060",
//    "OpenConversationId": "cid740SMooADFxmWX3DvIyp4g==",
//    "CorpId": "ding0761931a826dfde2ffe93478753d9884",
//    "EventType": "chat_disband",
//    "OperatorUnionId": "3bOBVxFv0J8VOu3J4jGhZQiEiE",
//    "TimeStamp": 1640754514696,
//    "ChatId": "chat6418e0c8c141114e18491ea2603c09be"
//}
type ChatDisband struct {
	Event

	// 操作人
	Operator string `json:"Operator"`

	// 群回话id
	OpenConversationId string `json:"OpenConversationId"`

	OperatorUnionId string `json:"OperatorUnionId"`

	TimeStamp int `json:"TimeStamp"`

	ChatId string `json:"ChatId"`
}
