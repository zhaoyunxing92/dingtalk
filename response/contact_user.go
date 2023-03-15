package response

type ContactUser struct {
	Response

	Nick      string `json:"nick,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	UnionId   string `json:"unionid,omitempty"`
	OpenId    string `json:"openid,omitempty"`
	Email     string `json:"email,omitempty"`
	StateCode string `json:"stateCode,omitempty"`
}
