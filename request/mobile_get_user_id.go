package request

type MobileGetUserId struct {
	Mobile string `json:"mobile" validate:"required"`
}

func NewMobileGetUserId(mobile string) *MobileGetUserId {

	return &MobileGetUserId{mobile}
}
