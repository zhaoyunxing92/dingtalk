package model

//role:角色
type role struct {
	Id   int    `json:"id"`   //角色id
	Name string `json:"name"` //角色名称
}

//openRole:角色
type openRole struct {
	Id   int    `json:"role_id"`   //角色id
	Name string `json:"role_name"` //角色名称
}

//roleGroup:角色组
type roleGroup struct {
	GroupId int        `json:"groupId"` //角色id
	Name    string     `json:"name"`    //角色组名称
	Roles   []role `json:"roles"`   //角色列表
}

//roleGroup:角色组
type openRoleGroup struct {
	Name  string `json:"group_name"` //角色组名称
	Roles []openRole `json:"roles"`      //角色列表
}

//roleListResult:角色列表返回
type roleListResult struct {
	HasMore bool        `json:"hasMore"` //是否还有更多数据
	List    []roleGroup `json:"list"`    //角色组列表。
}

type roleUserListResult struct {
	HasMore bool   `json:"hasMore"` //是否还有更多数据
	List    []User `json:"list"`
}

//RoleListResponse:获取角色列表返回
type RoleListResponse struct {
	Response
	RoleGroup roleListResult `json:"result"`
}

//RoleUserListResponse:角色用户列表
type RoleUserListResponse struct {
	Response
	Result roleUserListResult `json:"result"`
}

//RoleGroupResponse:获取角色组
type RoleGroupResponse struct {
	Response
	Result openRoleGroup `json:"role_group"`
}
