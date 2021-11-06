package request

type DeleteRole struct {
	Id int `json:"role_id"`
}

func NewDeleteRole(id int) *DeleteRole {
	return &DeleteRole{id}
}
