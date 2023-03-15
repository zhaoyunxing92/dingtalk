package request

const (
	ActionGetUserAuthToken = "authorization_code"
	ActionRefreshAuthToken = "refresh_token"
)

type UserAccessToken struct {
	ClientId     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	Code         string `json:"code,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	GrantType    string `json:"grantType,omitempty"`
}

func NewUserAuthToken(authCode string) *UserAccessToken {
	return &UserAccessToken{
		Code:      authCode,
		GrantType: ActionGetUserAuthToken,
	}
}

func NewRefreshUserAuthToken(refreshToken string) *UserAccessToken {
	return &UserAccessToken{
		RefreshToken: refreshToken,
		GrantType:    ActionRefreshAuthToken,
	}
}
