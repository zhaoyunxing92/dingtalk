package request

type SnsUserInfo struct {
	//用户授权的临时授权码，只能使用一次。获取方法请参考：
	//
	//扫码登录第三方网站：开发流程
	//
	//免登第三方网站：开发流程
	//
	//使用钉钉账号登录第三方网站：开发流程
	AutoCode string `json:"tmp_auth_code,omitempty" validate:"required"`
}

func NewSnsUserInfo(code string) *SnsUserInfo {
	return &SnsUserInfo{code}
}
