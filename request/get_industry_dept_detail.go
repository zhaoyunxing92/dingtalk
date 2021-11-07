package request

type GetIndustryDeptDetail struct {
	DeptId int `json:"dept_id" validate:"required"`
}

func NewIndustryDeptDetail(deptId int) *GetIndustryDeptDetail {
	return &GetIndustryDeptDetail{deptId}
}
