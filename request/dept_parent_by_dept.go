package request

type GetParentIdsByDeptId struct {
	DeptId int `json:"dept_id" validate:"required"`
}

func NewGetParentIdsByDeptId(id int) *GetParentIdsByDeptId {
	return &GetParentIdsByDeptId{id}
}
