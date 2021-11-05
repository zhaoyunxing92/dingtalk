package request

type GetParentIdsByUserId struct {
	UserId string `json:"userid" validate:"required"`
}

func NewGetParentIdsByUserId(userId string) *GetParentIdsByUserId {
	return &GetParentIdsByUserId{userId}
}
