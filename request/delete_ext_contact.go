package request

type DeleteExtContact struct {
	UserId string `json:"user_id" validate:"required"`
}

func NewDeleteExtContact(userId string) *DeleteExtContact {
	return &DeleteExtContact{userId}
}
