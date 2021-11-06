package request

type UpdateRole struct {
	//角色id
	Id int `json:"roleId" validate:"required"`
	//角色名称
	Name string `json:"roleName" validate:"required"`
}

func NewUpdateRole(id int, name string) *UpdateRole {
	return &UpdateRole{id, name}
}
