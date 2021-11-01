package response

//SuiteAccessToken 获取suite套件token
type SuiteAccessToken struct {
	Response
	//Token 第三方企业应用的凭证。
	Token string `json:"suite_access_token"`
	//Expires 第三方企业应用的凭证过期时间，单位秒。
	Expires int16 `json:"expires_in"`
	// Create 创建时间
	Create int64 `json:"create"`
}

//CreatedAt Expired.CreatedAt is when the access token is generated
func (token *SuiteAccessToken) CreatedAt() int64 {
	return token.Create
}

//ExpiresIn is how soon the access token is expired
func (token *SuiteAccessToken) ExpiresIn() int16 {
	return token.Expires
}
