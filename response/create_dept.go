package response

type CreateDept struct {
	Response
	createDept `json:"result"`
}

type createDept struct {
	DeptId int `json:"dept_id"`
}
