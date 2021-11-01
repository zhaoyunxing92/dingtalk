package model

//AccessToken 钉钉token结构体
type AccessToken struct {
	Response
	Expires int16  `json:"expires_in"`   //过期时间
	Created int64  `json:"created"`      //创建时间
	Token   string `json:"access_token"` //token
}

//CreatedAt Expired.CreatedAt is when the access token is generated
func (token *AccessToken) CreatedAt() int64 {
	return token.Created
}

//ExpiresIn is how soon the access token is expired
func (token *AccessToken) ExpiresIn() int16 {
	return token.Expires
}
