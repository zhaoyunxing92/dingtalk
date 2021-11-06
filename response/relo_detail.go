package response

type RoleDetail struct {
	Response

	roleDetail `json:"role"`
}

type roleDetail struct {
	//角色名称
	RoleName string `json:"name"`

	//所属的角色组Id
	GroupId int `json:"groupId"`
}
