package response

type SubDeptList struct {
	Response
	DeptIds subDeptList `json:"result"`
}

type subDeptList struct {
	Ids []int `json:"dept_id_list"`
}
