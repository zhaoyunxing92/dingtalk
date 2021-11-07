package request

type GetIndustryDept struct {
	DeptId int `json:"dept_id" validate:"required"`

	//支持分页查询，与size参数同时设置时才生效，此参数代表偏移量，偏移量从0开始
	Offset int `json:"cursor"`

	//支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，最大100
	Size int `json:"size" validate:"max=1000"`
}

func NewIndustryDept(deptId, offset, size int) *GetIndustryDept {
	return &GetIndustryDept{deptId, offset, size}
}
