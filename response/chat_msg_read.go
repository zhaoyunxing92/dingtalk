package response

type ChatMsgReadUser struct {
	Response

	//下次分页获取的起始游标
	NextCursor int `json:"next_cursor"`

	//已读人员的userid列表。已读人员为空时不返回该参数。
	UserIds []string `json:"readUserIdList"`
}
