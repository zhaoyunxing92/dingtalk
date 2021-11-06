package request

type CreateRoleGroup struct {
	//角色名称
	Name string `json:"name" validate:"required"`
}

func NewCreateRoleGroup(name string) *CreateRoleGroup {
	return &CreateRoleGroup{name}
}
