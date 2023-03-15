package response

import "github.com/pkg/errors"

type UserAccessToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	ExpireIn     int    `json:"expireIn,omitempty"`
	CorpId       string `json:"corpId,omitempty"`

	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	RequestId string `json:"requestid,omitempty"`
}

func (u UserAccessToken) CheckError() (err error) {
	if u.Code != "" {
		err = errors.Errorf("code:'%s', msg:%s, requestId: %s", u.Code, u.Message, u.RequestId)
	}
	return
}
