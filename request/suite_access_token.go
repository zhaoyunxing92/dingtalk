package request

//SuiteAccessToken 获取suite套件token
type SuiteAccessToken struct {
	Key string `json:"suite_key" validate:"required"`

	Secret string `json:"suite_secret" validate:"required"`

	Ticket string `json:"suite_ticket" validate:"required"`
}

type suiteAccessTokenBuilder struct {
	sat *SuiteAccessToken
}

func NewSuiteAccessToken() *suiteAccessTokenBuilder {
	return &suiteAccessTokenBuilder{sat: &SuiteAccessToken{}}
}

func (s *suiteAccessTokenBuilder) SetKey(key string) *suiteAccessTokenBuilder {
	s.sat.Key = key
	return s
}

func (s *suiteAccessTokenBuilder) SetSecret(secret string) *suiteAccessTokenBuilder {
	s.sat.Secret = secret
	return s
}

func (s *suiteAccessTokenBuilder) SetTicket(ticket string) *suiteAccessTokenBuilder {
	s.sat.Ticket = ticket
	return s
}

func (s *suiteAccessTokenBuilder) Build() *SuiteAccessToken {
	return s.sat
}
