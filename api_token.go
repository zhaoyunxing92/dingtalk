/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/cache"
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/crypto"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

//GetAccessToken 获取token
func (ding *dingTalk) GetAccessToken() (token string, err error) {
	var (
		ch  = ding.cache
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
		return res.Token, err
	}
	return res.Token, nil
}

// GetSuiteAccessToken 获取第三方企业应用的suite_access_token
func (ding *dingTalk) GetSuiteAccessToken() (token string, err error) {
	if !ding.isv() {
		return "", errors.New("ticket or corpId is null")
	}
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
		SetTicket(ding.ticket).
		Build()

	if err = ding.Request(http.MethodPost, constant.SuiteAccessToken, nil, req, res); err != nil {
		return "", err
	}
	res.Create = time.Now().Unix()
	if err = ch.Set(res); err != nil {
		return res.Token, err
	}
	return res.Token, err
}

// GetCorpAccessToken 服务商获取第三方应用授权企业的access_token
func (ding *dingTalk) GetCorpAccessToken() (token string, err error) {
	var (
		ch  = ding.cache
		res = &response.AccessToken{}
	)
	// check ticket and corpId
	if !ding.isv() {
		return "", errors.New("ticket or corpId is null")
	}
	if err = ch.Get(res); err == nil {
		return res.Token, nil
	}
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	sign := crypto.GetSignature(timestamp, ding.Secret, ding.ticket)

	query := url.Values{}
	query.Set("accessKey", ding.Key)
	query.Set("timestamp", timestamp)
	query.Set("suiteTicket", ding.ticket)
	query.Set("signature", sign)

	if err = ding.Request(http.MethodPost, constant.CorpAccessToken, query,
		request.NewCorpAccessToken(ding.corpId), res); err != nil {
		return "", err
	}
	res.Create = time.Now().Unix()
	if err = ch.Set(res); err != nil {
		return res.Token, err
	}
	return res.Token, err
}

//GetSSOToken 获取微应用后台免登的access_token
func (ding *dingTalk) GetSSOToken(corpId, secret string) (token string, err error) {
	var (
		ch  = cache.NewFileCache(strings.Join([]string{".token", "sso"}, "/"), ding.Key)
		res = &response.AccessToken{}
	)
	if err = ch.Get(res); err == nil {
		return res.Token, nil
	}
	query := url.Values{}
	query.Set("corpid", corpId)
	query.Set("corpsecret", secret)

	if err = ding.Request(http.MethodGet, constant.GetSSOTokenKey, query, nil, res); err != nil {
		return "", err
	}
	res.Create = time.Now().Unix()
	if err = ch.Set(res); err != nil {
		return res.Token, err
	}
	return res.Token, nil
}

//GetJsApiTicket 获取jsapi_ticket
func (ding *dingTalk) GetJsApiTicket() (ticket string, err error) {
	var (
		ch  = cache.NewFileCache(strings.Join([]string{".token", "ticket"}, "/"), ding.Key)
		res = &response.JsApiTicket{}
	)
	if err = ch.Get(res); err == nil {
		return res.Ticket, nil
	}

	if err = ding.Request(http.MethodGet, constant.GetJsApiTicketKey, nil, nil, res); err != nil {
		return "", err
	}
	res.Create = time.Now().Unix()
	if err = ch.Set(res); err != nil {
		return res.Ticket, err
	}
	return res.Ticket, nil
}
