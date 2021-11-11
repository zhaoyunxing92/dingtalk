/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
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

	query := url.Values{}
	query.Set("accessKey", ding.Key)
	query.Set("suiteTicket", ding.Ticket)
	query.Set("timestamp", timestamp)
	query.Set("signature", sign)

	res := &response.SuiteAccessToken{}
	if err = ding.Request(http.MethodPost, constant.CorpAccessToken, query,
		request.NewCorpAccessToken(ding.CorpId), res); err != nil {
		return "", err
	}
	return res.Token, err
}
