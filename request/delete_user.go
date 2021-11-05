package request

import "encoding/json"

type DeleteUser struct {
	UserId string `json:"userid" validate:"required"`
}

func (d *DeleteUser) String() string {
	str, _ := json.Marshal(d)
	return string(str)
}
func NewDeleteUser(userId string) *DeleteUser {
	return &DeleteUser{userId}
}
