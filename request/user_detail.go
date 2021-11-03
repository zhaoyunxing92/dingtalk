package request

type UserDetail struct {
	*DeleteUser

	//通讯录语言，默认zh_CN。如果是英文，请传入en_US。
	Language string `json:"language,omitempty" validate:"omitempty,oneof=zh_CN en_US"`
}

type userDetailBuilder struct {
	user *UserDetail
}

func NewUserDetail(userId string) *userDetailBuilder {
	return &userDetailBuilder{user: &UserDetail{DeleteUser: NewDeleteUser(userId)}}
}

//SetLanguage 通讯录语言，默认zh_CN。如果是英文，请传入en_US。
func (ub *userDetailBuilder) SetLanguage(language string) *userDetailBuilder {
	ub.user.Language = language
	return ub
}

func (ub *userDetailBuilder) Build() *UserDetail {
	return ub.user
}
