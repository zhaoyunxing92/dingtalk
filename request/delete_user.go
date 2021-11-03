package request

type DeleteUser struct {
	UserId string `json:"userid" validate:"required"`
}

func NewDeleteUser(userId string) *DeleteUser {
	return &DeleteUser{userId}
}
