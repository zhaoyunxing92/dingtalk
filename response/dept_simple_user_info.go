package response

type DeptSimpleUserInfo struct {
	Response
	DeptUserInfoPage `json:"result"`
}

type DeptUserInfoPage struct {
	//是否还有更多的数据
	HasMore bool `json:"has_more"`

	//下一次分页的游标，如果has_more为false，表示没有更多的分页数据。
	NextCursor int `json:"next_cursor"`

	DeptUsers []simpleUserInfo `json:"list"`
}

//simpleUserInfo 简单用户信息
type simpleUserInfo struct {
	UserId string `json:"userid"`

	Name string `json:"name"`
}
