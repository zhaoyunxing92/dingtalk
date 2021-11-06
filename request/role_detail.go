package request

type RoleDetail struct {
	Id int `json:"roleId"`
}

func NewRoleDetail(id int) *RoleDetail {
	return &RoleDetail{id}
}
