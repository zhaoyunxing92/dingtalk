package dingtalk

import (
	"errors"
	"github.com/zhaoyunxing92/dingtalk/v2/cache"
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/crypto"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//GetAccessToken 获取token
func (ding *dingTalk) GetAccessToken() (token string, err error) {
	var (
		ch  = ding.Cache
		res = &response.AccessToken{}
	)
	//先缓存中获取
	if err = ch.Get(res); err == nil {
		return res.Token, nil
	}
	//读取本地文件
	args := url.Values{}
	args.Set("appkey", ding.Key)
	args.Set("appsecret", ding.Secret)

	if err = ding.Request(http.MethodGet, constant.GetTokenKey, args, nil, res); err != nil {
		return "", err
	}
	res.Create = time.Now().Unix()
	if err = ch.Set(res); err != nil {
		return "", err
	}
	return res.Token, nil
}

// GetSuiteAccessToken 获取第三方企业应用的suite_access_token
func (ding *dingTalk) GetSuiteAccessToken() (token string, err error) {
	var (
		ch  = cache.NewFileCache(strings.Join([]string{".token", "suite"}, "/"), ding.Key)
		res = &response.SuiteAccessToken{}
	)

	if err = ch.Get(res); err == nil {
		return res.Token, nil
	}
	req := request.NewSuiteAccessToken().
		SetKey(ding.Key).
		SetSecret(ding.Secret).
		SetTicket(ding.Ticket).
		Build()

	if err = ding.Request(http.MethodPost, constant.SuiteAccessToken, nil, req, res); err != nil {
		return "", err
	}
	res.Create = time.Now().Unix()
	if err = ch.Set(res); err != nil {
		return "", err
	}
	return res.Token, err
}

// GetCorpAccessToken 服务商获取第三方应用授权企业的access_token
func (ding *dingTalk) GetCorpAccessToken() (token string, err error) {
	// check ticket and corpId
	if len(ding.Ticket) <= 0 || len(ding.CorpId) <= 0 {
		return "", errors.New("ticket or corpId is null")
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	sign := crypto.GetSignature(timestamp, ding.Secret, ding.Ticket)

	args := url.Values{}
	args.Set("accessKey", ding.Key)
	args.Set("suiteTicket", ding.Ticket)
	args.Set("timestamp", timestamp)
	args.Set("signature", sign)

	req := request.NewCorpAccessToken(ding.CorpId)
	res := &response.SuiteAccessToken{}
	if err = ding.Request(http.MethodPost, constant.CorpAccessToken, args, req, res); err != nil {
		return "", err
	}
	return res.Token, err
}
