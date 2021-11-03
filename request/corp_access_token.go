package request

//CorpAccessToken 获取suite套件token
type CorpAccessToken struct {
	CorpId string `json:"auth_corpid" validate:"required"`
}

func NewCorpAccessToken(corpId string) *CorpAccessToken {
	return &CorpAccessToken{corpId}
}
