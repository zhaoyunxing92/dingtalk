package request

import "encoding/json"

type AdminUserScope struct {
	UserId string `json:"userid" validate:"required"`
}

func (a *AdminUserScope) String() string {
	str, _ := json.Marshal(a)
	return string(str)
}
func NewAdminUserScope(userId string) *AdminUserScope {
	return &AdminUserScope{userId}
}
