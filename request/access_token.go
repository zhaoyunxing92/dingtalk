package request

//AccessToken 获取企业内部应用的access_token
// https://developers.dingtalk.com/document/app/obtain-orgapp-token
type AccessToken struct {
	//应用的唯一标识key
	AppKey string

	//应用的密钥
	AppSecret string
}

//NewAccessTokenRequest  获取AccessToken对象
func NewAccessTokenRequest(AppKey, AppSecret string) *AccessToken {

	return &AccessToken{AppKey, AppSecret}
}
