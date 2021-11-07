package request

import (
	"encoding/json"
	"strings"
)

type SetUserRoleManageScope struct {
	//员工在企业中的userid
	UserId string `json:"userid" validate:"required"`

	//角色Id
	RoleId int `json:"role_id"`

	DeptIds string `json:"dept_ids,omitempty"`

	Ds []string `json:"-" validate:"max=50"`
}

func (s *SetUserRoleManageScope) String() string {
	str, _ := json.Marshal(s)
	return string(str)
}

type setUserRoleManageScopeBuilder struct {
	rs *SetUserRoleManageScope
}

func NewSetUserRoleManageScope(userId string, roleId int) *setUserRoleManageScopeBuilder {

	return &setUserRoleManageScopeBuilder{rs: &SetUserRoleManageScope{UserId: userId, RoleId: roleId}}
}

func (b *setUserRoleManageScopeBuilder) SetDeptIds(deptId int, deptIds ...int) *setUserRoleManageScopeBuilder {
	ids := []int{deptId}
	ids = append(ids, deptIds...)
	b.rs.Ds = removeIntDuplicatesToString(ids)
	return b
}

func (b *setUserRoleManageScopeBuilder) Build() *SetUserRoleManageScope {
	b.rs.DeptIds = strings.Join(b.rs.Ds, ",")
	return b.rs
}
