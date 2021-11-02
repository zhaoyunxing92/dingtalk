package request

//SuiteAccessToken 获取suite套件token
type CorpAccessToken struct {
	CorpId string `json:"auth_corpid" validate:"required"`
}

type corpAccessTokenBuilder struct {
	sat *CorpAccessToken
}

func NewCorpAccessToken() *corpAccessTokenBuilder {
	return &corpAccessTokenBuilder{sat: &CorpAccessToken{}}
}

func (s *corpAccessTokenBuilder) SetCorpId(corpId string) *corpAccessTokenBuilder {
	s.sat.CorpId = corpId
	return s
}

func (s *corpAccessTokenBuilder) Build() *CorpAccessToken {
	return s.sat
}
