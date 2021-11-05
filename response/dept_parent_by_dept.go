package response

type GetParentIdsByDeptId struct {
	Response
	parentIdsByDeptId `json:"result"`
}

type parentIdsByDeptId struct {
	ParentIds []int `json:"parent_id_list"`
}
