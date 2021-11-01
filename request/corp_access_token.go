package request

//SuiteAccessToken 获取suite套件token
type CorpAccessToken struct {
	Key string `json:"accessKey" validate:"required"`

	Secret string `json:"suite_secret" validate:"required"`

	Ticket string `json:"suiteTicket" validate:"required"`

	CorpId string `json:"auth_corpid" validate:"required"`
}

type corpAccessTokenBuilder struct {
	sat *CorpAccessToken
}

func NewCorpAccessToken() *corpAccessTokenBuilder {
	return &corpAccessTokenBuilder{sat: &CorpAccessToken{}}
}

func (s *corpAccessTokenBuilder) SetKey(key string) *corpAccessTokenBuilder {
	s.sat.Key = key
	return s
}

func (s *corpAccessTokenBuilder) SetSecret(secret string) *corpAccessTokenBuilder {
	s.sat.Secret = secret
	return s
}

func (s *corpAccessTokenBuilder) SetTicket(ticket string) *corpAccessTokenBuilder {
	s.sat.Ticket = ticket
	return s
}

func (s *corpAccessTokenBuilder) SetCorpId(corpId string) *corpAccessTokenBuilder {
	s.sat.CorpId = corpId
	return s
}

func (s *corpAccessTokenBuilder) Build() *CorpAccessToken {
	return s.sat
}
