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
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/zhaoyunxing92/dingtalk/v2/crypto"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/zhaoyunxing92/dingtalk/v2/cache"
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/logger"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

type DingTalk struct {
	// 企业内部应用对应:AgentId，如果是应套件:SuiteId
	//id int

	// 企业内部应用对应:AppKey，套件对应:SuiteKey
	key string

	// 企业内部对应:AppSecret，套件对应:SuiteSecret
	secret string

	// isv 钉钉开放平台会向应用的回调URL推送的suite_ticket（约5个小时推送一次）
	ticket string

	// 授权企业的id
	corpId string

	// 在开发者后台的基本信息 > 开发信息（旧版）页面获取微应用管理后台SSOSecret
	ssoSecret string

	// 日志级别
	Level zapcore.Level

	log *zap.SugaredLogger

	client *http.Client

	cache cache.Cache
}

type OptionFunc func(*DingTalk)

func WithTicket(ticket string) OptionFunc {
	return func(dt *DingTalk) {
		dt.ticket = ticket
	}
}

func WithCorpId(corpId string) OptionFunc {
	return func(dt *DingTalk) {
		dt.corpId = corpId
	}
}

func WithSSOSecret(secret string) OptionFunc {
	return func(dt *DingTalk) {
		dt.ssoSecret = secret
	}
}

func WithLevel(level zapcore.Level) OptionFunc {
	return func(dt *DingTalk) {
		dt.Level = level
	}
}

func WithCache(cache cache.Cache) OptionFunc {
	return func(dt *DingTalk) {
		dt.cache = cache
	}
}

// isv 是否isv
func (ding *DingTalk) isv() bool {
	return len(ding.ticket) > 0 && len(ding.corpId) > 0
}

func (ding *DingTalk) validate() error {
	if ding.key == "" {
		return errors.New("key is empty")
	}
	if ding.secret == "" {
		return errors.New("secret is empty")
	}
	return nil
}

// NewClient new client
func NewClient(key, secret string, opts ...OptionFunc) (ding *DingTalk, err error) {
	ding = &DingTalk{key: key, secret: secret, Level: zapcore.InfoLevel}

	for _, opt := range opts {
		opt(ding)
	}

	if ding.cache == nil {
		if ding.isv() {
			ding.cache = cache.NewFileCache(strings.Join([]string{".token", "isv", key}, "/"), ding.corpId)
		} else {
			ding.cache = cache.NewFileCache(strings.Join([]string{".token", "corp"}, "/"), key)
		}
	}

	ding.client = &http.Client{Timeout: 10 * time.Second}
	ding.log = logger.GetLogger(ding.Level).Sugar()

	if err = ding.validate(); err != nil {
		return nil, err
	}
	return
}

// Request 统一请求
func (ding *DingTalk) Request(method, path string, query url.Values, body interface{},
	data response.Unmarshalled) (err error) {
	if body != nil {
		if err = validate(body); err != nil {
			return err
		}
	}

	if query == nil {
		query = url.Values{}
	}

	if query.Get("access_token") == "" && path != constant.GetTokenKey && path != constant.CorpAccessToken &&
		path != constant.SuiteAccessToken && path != constant.GetAuthInfo && path != constant.GetAgentKey &&
		path != constant.ActivateSuiteKey && path != constant.GetSSOTokenKey && path != constant.GetUnactiveCorpKey &&
		path != constant.ReauthCorpKey && path != constant.GetCorpPermanentCodeKey && path != constant.GetUserAccessToken {

		var token string

		if ding.isv() {
			if token, err = ding.GetCorpAccessToken(); err != nil {
				return err
			}
		} else {
			if token, err = ding.GetAccessToken(); err != nil {
				return err
			}
		}
		// set token
		query.Set("access_token", token)
	}
	return ding.httpRequest(method, path, query, body, data)
}

func (robot *Robot) send(form interface{}, data response.Unmarshalled) (err error) {
	args := url.Values{}
	args.Set("access_token", robot.Token)

	return robot.httpRequest(http.MethodPost, constant.SendRobotMsgKey, args, form, data)
}

func (robot *Robot) httpRequest(method, path string, query url.Values, body interface{},
	data response.Unmarshalled) error {
	client := robot.client

	uri, _ := url.Parse(constant.Api + path)
	uri.RawQuery = query.Encode()

	var (
		req     *http.Request
		res     *http.Response
		err     error
		content []byte
	)

	// 表单不为空
	d, _ := json.Marshal(body)
	fmt.Println(string(d))

	req, _ = http.NewRequest(method, uri.String(), bytes.NewReader(d))
	req.Header.Set("Content-Type", "application/json")

	if res, err = client.Do(req); err != nil {
		return err
	}

	defer func(b io.ReadCloser) { _ = b.Close() }(res.Body)

	if res.StatusCode != 200 {
		return errors.New("dingtalk server error: " + res.Status)
	}

	if content, err = ioutil.ReadAll(res.Body); err != nil {
		return err
	}
	fmt.Println(string(content))
	if err = json.Unmarshal(content, data); err != nil {
		return err
	}
	return data.CheckError()
}

func (ding *DingTalk) httpRequest(method, path string, query url.Values, body interface{},
	response response.Unmarshalled) error {
	var (
		req    *http.Request
		res    *http.Response
		err    error
		form   []byte // body 数据
		data   []byte // 返回数据
		client = ding.client
		uri    *url.URL
		token  string
		newApi = isNewApi(path)
	)

	if newApi {
		token = query.Get("access_token")
		query.Del("access_token")
		uri, _ = url.Parse(constant.NewApi + path)
	} else {
		uri, _ = url.Parse(constant.Api + path)
	}
	uri.RawQuery = query.Encode()

	if body != nil {
		// 检查提交表单类型
		switch body.(type) {
		case request.UploadFile:
			var b bytes.Buffer
			var fw io.Writer

			w := multipart.NewWriter(&b)
			file := body.(request.UploadFile)
			fw, err = w.CreateFormFile(file.FieldName, file.FileName)
			if err != nil {
				return err
			}
			if _, err = io.Copy(fw, file.Reader); err != nil {
				return err
			}
			if err = w.Close(); err != nil {
				return err
			}
			req, _ = http.NewRequest(method, uri.String(), &b)
			req.Header.Set("Content-Type", w.FormDataContentType())
		default:
			// 表单不为空
			form, _ = json.Marshal(body)
			req, _ = http.NewRequest(method, uri.String(), bytes.NewReader(form))
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
		}
	} else {
		req, _ = http.NewRequest(method, uri.String(), nil)
	}

	if newApi {
		req.Header.Set("x-acs-dingtalk-access-token", token)
	}

	if res, err = client.Do(req); err != nil {
		return err
	}

	defer func(b io.ReadCloser) { _ = b.Close() }(res.Body)

	if data, err = ioutil.ReadAll(res.Body); err != nil {
		return err
	}

	log := ding.log.With(zap.Bool("new", newApi), zap.String("method", method), zap.String("path", path),
		zap.String("body", string(form)), zap.String("token", token),
		zap.ByteString("res", data))

	switch res.StatusCode {
	case 200:
		if err = json.Unmarshal(data, response); err != nil {
			log.Error(err)
			return err
		}
		log.Debug("request succeed")
		return response.CheckError()
	case 400:
		log.Errorf("ding fail status code %d,res:%s", res.StatusCode, data)
		return errors.Errorf("dingtalk server error,res:%s", data)
	case 404:
		log.Errorf("ding fail status code %d,res:%s", res.StatusCode, data)
		return errors.Errorf("dingtalk server error,res:%s", data)
	case 500:
		log.Errorf("ding fail status code %d,res:%s", res.StatusCode, data)
		return errors.Errorf("dingtalk server error,res:%s", data)
	}
	log.Error("ding fail")
	return errors.New("ding fail")
}

// validate 参数验证
func validate(s interface{}) error {
	v := validator.New()
	if err := v.Struct(s); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Error())
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}

func isNewApi(path string) bool {
	return strings.HasPrefix(path, "/v1.0/")
}

// GetDingTalkCrypto 钉钉事件解密类
func (ding *DingTalk) GetDingTalkCrypto(token, aesKey string) (*crypto.DingTalkCrypto, error) {
	var (
		block cipher.Block
		err   error
		key   []byte
	)

	if len(aesKey) != crypto.AesEncodeKeyLength {
		return nil, errors.New("不合法的aes key")
	}

	if key, err = base64.StdEncoding.DecodeString(aesKey + "="); err != nil {
		return nil, err
	}

	if block, err = aes.NewCipher(key); err != nil {
		return nil, err
	}
	return &crypto.DingTalkCrypto{
		Token:    token,
		SuiteKey: ding.key,
		AESKey:   key,
		Block:    block,
	}, nil
}
