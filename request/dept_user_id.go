package request

type DeptUserId struct {
	//部门ID，根部门ID为1
	DeptId int `json:"dept_id" validate:"required,min=1"`
}

func NewDeptUserId(deptId int) *DeptUserId {

	return &DeptUserId{deptId}
}
