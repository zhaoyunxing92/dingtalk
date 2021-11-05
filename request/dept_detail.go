package request

type DeptDetail struct {
	DeptId int `json:"dept_id" validate:"required"`

	//通讯录语言，默认zh_CN。如果是英文，请传入en_US。
	Language string `json:"language,omitempty" validate:"omitempty,oneof=zh_CN en_US"`
}

type deptDetailBuilder struct {
	d *DeptDetail
}

func NewDeptDetail(deptId int) *deptDetailBuilder {
	return &deptDetailBuilder{d: &DeptDetail{DeptId: deptId}}
}

func (db *deptDetailBuilder) SetLanguage(language string) *deptDetailBuilder {
	db.d.Language = language
	return db
}

func (db *deptDetailBuilder) Build() *DeptDetail {
	return db.d
}
