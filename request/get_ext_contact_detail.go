package request

//GetExtContactDetail 获取外部联系人详情
type GetExtContactDetail struct {
	UserId string `json:"user_id" validate:"required"`
}

func NewGetExtContactDetail(userId string) *GetExtContactDetail {
	return &GetExtContactDetail{userId}
}
