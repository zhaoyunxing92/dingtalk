package response

type GetHrmEmployee struct {
	Response

	getHrmEmployee `json:"result"`
}

type getHrmEmployee struct {

	// 查询到的员工userid
	UserIds []string

	// 下一次分页调用的offset值，当返回结果里没有next_cursor时，表示分页结束
	NextCursor int `json:"next_cursor"`
}
