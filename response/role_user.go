package response

type RoleUser struct {
	Response

	roleUserResult `json:"result"`
}
type roleUserResult struct {
	HasMore    bool `json:"hasMore"`
	NextCursor int  `json:"nextCursor"`

	Users []roleUser `json:"list"`
}

type roleUser struct {
	UserId string `json:"userid"`

	Name string `json:"name"`

	ManageScopes []manageScope `json:"manageScopes"`
}

//管理范围。
type manageScope struct {
	//部门Id
	DeptId int `json:"dept_id"`
	//部门名称
	DeptName string `json:"name"`
}
