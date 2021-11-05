package request

type DeptList struct {
	//父部门ID。
	//
	//如果不传，默认部门为根部门，根部门ID为1。只支持查询下一级子部门，不支持查询多级子部门。
	DeptId int `json:"dept_id,omitempty"`

	//通讯录语言，默认zh_CN。如果是英文，请传入en_US。
	Language string `json:"language,omitempty" validate:"omitempty,oneof=zh_CN en_US"`
}

type deptListBuilder struct {
	dl *DeptList
}

func NewDeptList() *deptListBuilder {
	return &deptListBuilder{dl: &DeptList{}}
}

func (db *deptListBuilder) SetDeptId(deptId int) *deptListBuilder {
	db.dl.DeptId = deptId
	return db
}

func (db *deptListBuilder) SetLanguage(language string) *deptListBuilder {
	db.dl.Language = language
	return db
}

func (db *deptListBuilder) Build() *DeptList {
	return db.dl
}
