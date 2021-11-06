package response

type GroupRole struct {
	Response

	groupRole `json:"role_group"`
}

type groupRole struct {
	//角色组名
	GroupName string `json:"group_name"`

	Roles []gr `json:"roles"`
}

type gr struct {
	//角色Id
	Id int `json:"role_id"`

	//角色名
	Name string `json:"role_name"`
}
