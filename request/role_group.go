package request

type GroupRole struct {
	//员工在企业中的userid
	GroupId int `json:"group_id"`
}

func NewGroupRole(id int) *GroupRole {
	return &GroupRole{id}
}
