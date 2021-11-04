package request

type UserCount struct {
	//是否包含未激活钉钉人数
	OnlyActive bool `json:"only_active" validate:"required"`
}

func NewUserCount(active bool) *UserCount {
	return &UserCount{active}
}
