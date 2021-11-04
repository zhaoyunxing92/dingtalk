package response

type DeptUserId struct {
	Response
	deptUserId `json:"result"`
}

type deptUserId struct {
	UserIds []string `json:"userid_list"`
}
