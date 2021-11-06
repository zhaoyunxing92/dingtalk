package response

type RoleList struct {
	Response

	rs `json:"result"`
}

type rs struct {
	//是否还有更多数据。
	HasMore bool `json:"hasMore"`

	RoleGroups []roleGroup `json:"list"`
}

type roleGroup struct {
	//角色组Id
	GroupId int `json:"groupId"`

	//角色组名称
	Name string `json:"name"`

	Roles []role `json:"roles"`
}

type role struct {
	//角色Id
	Id int `json:"id"`

	//角色名称
	Name string `json:"name"`
}
