package response

type GetParentIdsByUserId struct {
	Response
	parentIdsByUserId `json:"result"`
}

type parentIdsByUserId struct {
	Parent []parentIdList `json:"parent_list"`
}

type parentIdList struct {
	ParentIds []int `json:"parent_dept_id_list"`
}
