package request

import "encoding/json"

type DeleteDept struct {
	DeptId int `json:"dept_id" validate:"required"`
}

func NewDeleteDept(deptId int) *DeleteDept {
	return &DeleteDept{deptId}
}

func (d *DeleteDept) String() string {
	str, _ := json.Marshal(d)
	return string(str)
}
