package request

type AuthInfo struct {
	SuiteKey string `json:"suite_key" validate:"required"`
	CorpId   string `json:"auth_corpid" validate:"required"`
}

func NewAuthInfo(suiteKey, corpId string) *AuthInfo {
	return &AuthInfo{suiteKey, corpId}
}
