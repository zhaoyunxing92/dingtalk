package request

import (
	"encoding/json"
	"strings"
)

type RoleAddUser struct {
	Rs []string `validate:"required,max=20" json:"-"`

	//角色roleId列表 多个roleId用英文逗号（,）分隔，最多可传20个
	RoleIds string `validate:"required" json:"roleIds"`

	Us []string `validate:"required,max=20" json:"-"`

	//员工的userId 多个userId用英文逗号（,）分隔，最多可传20个
	UserIds string `validate:"required" json:"userIds"`
}

func (r *RoleAddUser) String() string {
	str, _ := json.Marshal(r)
	return string(str)
}

func NewRoleAddUser(roleIds []int, userIds []string) *RoleAddUser {
	rs := removeIntDuplicates(roleIds)
	us := removeStringDuplicates(userIds)
	return &RoleAddUser{rs, strings.Join(rs, ","), us, strings.Join(us, ",")}
}
