package response

type AdminUserScope struct {
	Response
	//可管理的部门ID列表
	DeptIds []int `json:"dept_ids"`
}
