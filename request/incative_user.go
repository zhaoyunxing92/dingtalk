package request

import "sort"

//InactiveUser 获取未登录钉钉的员工列表
type InactiveUser struct {
	//是否活跃：
	//
	//false：未登录
	//
	//true：登录
	Active bool `json:"is_active" validate:"required"`

	//过滤部门ID列表，不传表示查询整个企业
	DeptIds []int `json:"dept_ids,omitempty"`

	//支持分页查询，与size参数同时设置时才生效，此参数代表偏移量，偏移量从0开始
	Offset int `json:"offset" validate:"required,min=0"`

	//支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，最大100
	Size int `json:"size" validate:"required,max=100"`

	//查询日期，日期格式为：yyyyMMdd
	Date string `json:"query_date" validate:"required,max=100"`
}

type inactiveUserBuilder struct {
	iu *InactiveUser
}

func NewInactiveUser(active bool, offset, size int, date string) *inactiveUserBuilder {
	return &inactiveUserBuilder{iu: &InactiveUser{Active: active, Offset: offset, Size: size, Date: date}}
}

func (receiver *inactiveUserBuilder) SetDeptIds(deptId int, ids ...int) *inactiveUserBuilder {
	deptIds := receiver.iu.DeptIds
	deptIds = append(deptIds, deptId)
	deptIds = append(deptIds, ids...)
	receiver.iu.DeptIds = deptIds
	return receiver
}

func (receiver inactiveUserBuilder) Build() *InactiveUser {
	receiver.iu.DeptIds = receiver.getDeptIds()
	return receiver.iu
}

func (receiver inactiveUserBuilder) getDeptIds() (ids []int) {
	deptIds := receiver.iu.DeptIds
	sort.Ints(deptIds)
	for i, id := range deptIds {
		if i >= 1 && id == deptIds[i-1] {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}
