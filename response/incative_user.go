package response

//InactiveUser 获取未登录钉钉的员工列表
type InactiveUser struct {
	Response
	inactiveUser `json:"result"`
}

type inactiveUser struct {
	//是否还有更多的数据
	HasMore bool `json:"has_more"`

	//下一次分页的游标，如果has_more为false，表示没有更多的分页数据。
	NextCursor int `json:"next_cursor"`

	UserIds []string `json:"list"`
}
