package request

type CorpPermanentCode struct {
	Code string `json:"tmp_auth_code" validate:"required"`
}

func NewCorpPermanentCode(code string) *CorpPermanentCode {
	return &CorpPermanentCode{code}
}
