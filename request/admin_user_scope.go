package request

type AdminUserScope struct {
	UserId string `json:"userid" validate:"required"`
}

func NewAdminUserScope(userId string) *AdminUserScope {
	return &AdminUserScope{userId}
}
