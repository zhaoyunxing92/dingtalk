package request

type SubDeptList struct {
	DeptId int `json:"dept_id" validate:"required"`
}

func NewSubDept(id int) *SubDeptList {
	return &SubDeptList{id}
}
