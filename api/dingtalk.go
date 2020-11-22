package api

import (
	"errors"
	"github.com/zhaoyunxing92/dingtalk/domain"
	"github.com/zhaoyunxing92/dingtalk/global"
	"net/http"
	"net/url"
	"time"
)

type DingTalk struct {
	AppKey    string //应用key
	AppSecret string //应用秘钥
	client    *http.Client
	Cache     global.Cache
}

// 创建钉钉客户端
func NewDingTalk(appKey, appSecret string) (*DingTalk, error) {
	if appKey == "" {
		return nil, errors.New("appKey不能为空")
	}
	if appSecret == "" {
		return nil, errors.New("appSecret不能为空")
	}
	return &DingTalk{appKey, appSecret, &http.Client{
		Timeout: 10 * time.Second,
	}, global.NewFileCache(".token/" + appKey)}, nil
}

//获取token
func (talk *DingTalk) GetToken() (token string, err error) {
	cache := talk.Cache
	var accessToken domain.AccessToken
	//先缓存中获取
	if err = cache.Get(&accessToken); err == nil {
		return accessToken.Token, nil
	}
	//读取本地文件
	args := url.Values{}
	args.Set("appkey", talk.AppKey)
	args.Set("appsecret", talk.AppSecret)

	if err = talk.request(http.MethodGet, global.GetToken, args, nil, &accessToken); err != nil {
		return "", err
	}
	if err = cache.Set(&accessToken); err != nil {
		return "", err
	}
	return accessToken.Token, nil
}
