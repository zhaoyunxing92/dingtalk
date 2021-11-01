package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/constant"
	"github.com/zhaoyunxing92/dingtalk/crypto"
	"github.com/zhaoyunxing92/dingtalk/request"
	"github.com/zhaoyunxing92/dingtalk/response"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//GetAccessToken 获取token
func (ding *DingTalk) GetAccessToken() (token string, err error) {
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

	if err = ding.request(http.MethodGet, constant.GetTokenKey, args, nil, res); err != nil {
		return "", err
	}
	res.Create = time.Now().Unix()
	if err = ch.Set(res); err != nil {
		return "", err
	}
	return res.Token, nil
}

// GetSuiteAccessToken 获取第三方企业应用的suite_access_token
func (ding *DingTalk) GetSuiteAccessToken() (token string, err error) {

	req := request.NewSuiteAccessToken().
		SetKey(ding.Key).
		SetSecret(ding.Secret).
		SetTicket(ding.Ticket).
		Build()

	res := &response.SuiteAccessToken{}

	if err = ding.request(http.MethodPost, constant.SuiteAccessToken, nil, req, res); err != nil {
		return "", err
	}
	return res.Token, err
}

// GetCorpAccessToken 服务商获取第三方应用授权企业的access_token
func (ding *DingTalk) GetCorpAccessToken() (token string, err error) {

	req := request.NewCorpAccessToken().
		SetKey(ding.Key).
		SetSecret(ding.Secret).
		SetTicket(ding.Ticket).
		SetCorpId(ding.CorpId).
		Build()

	if err = validate(req); err != nil {
		return "", err
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	sign := crypto.GetSignature(timestamp, req.Secret, req.Ticket)

	args := url.Values{}
	args.Set("accessKey", req.Key)
	args.Set("suiteTicket", req.Ticket)
	args.Set("timestamp", timestamp)
	args.Set("signature", sign)

	res := &response.SuiteAccessToken{}
	if err = ding.request(http.MethodPost, constant.CorpAccessToken, args, req, res); err != nil {
		return "", err
	}
	return res.Token, err
}
