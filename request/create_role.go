package request

type CreateRole struct {
	//角色名称
	Name string `json:"roleName" validate:"required"`
	//角色组ID
	GroupId int `json:"groupId" validate:"required"`
}

func NewCreateRole(name string, groupId int) *CreateRole {
	return &CreateRole{name, groupId}
}
