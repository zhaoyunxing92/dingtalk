package dingtalk

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zhaoyunxing92/dingtalk/model"
	"github.com/zhaoyunxing92/dingtalk/global"
	"net/http"
	"net/url"
	"time"
)

type DingTalk struct {
	AgentId   int    //应用id
	AppKey    string //应用key
	AppSecret string //应用秘钥
	client    *http.Client
	cache     global.Cache
	validate  *validator.Validate //参数校验
	trans     translator.Translator
}

// 创建钉钉客户端
func NewDingTalk(agentId int, appKey, appSecret string) *DingTalk {
	validate := validator.New()
	uni := translator.New(en.New(), zh.New())
	trans, _ := uni.GetTranslator("zh")
	_ = zh_trans.RegisterDefaultTranslations(validate, trans)

	return &DingTalk{agentId, appKey, appSecret, &http.Client{
		Timeout: 10 * time.Second,
	}, global.NewFileCache(".token", appKey), validate, trans}
}

//获取token
func (talk *DingTalk) GetToken() (token string, err error) {
	cache := talk.cache
	var accessToken model.AccessToken
	//先缓存中获取
	if err = cache.Get(&accessToken); err == nil {
		return accessToken.Token, nil
	}
	//读取本地文件
	args := url.Values{}
	args.Set("appkey", talk.AppKey)
	args.Set("appsecret", talk.AppSecret)

	if err = talk.request(http.MethodGet, global.GetTokenKey, args, nil, &accessToken); err != nil {
		return "", err
	}
	accessToken.Created = time.Now().Unix()
	if err = cache.Set(&accessToken); err != nil {
		return "", err
	}
	return accessToken.Token, nil
}
