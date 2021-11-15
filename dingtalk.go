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
	"encoding/json"
	"fmt"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)
import (
	"github.com/zhaoyunxing92/dingtalk/v2/cache"
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)
import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type dingTalk struct {
	// 企业内部应用对应:AgentId，如果是应套件:SuiteId
	Id int `json:"id,omitempty" validate:"required"`

	// 企业内部应用对应:AppKey，套件对应:SuiteKey
	Key string `json:"key,omitempty" validate:"required"`

	//企业内部对应:AppSecret，套件对应:SuiteSecret
	Secret string `json:"Secret,omitempty" validate:"required"`

	// isv 钉钉开放平台会向应用的回调URL推送的suite_ticket（约5个小时推送一次）
	ticket string

	//授权企业的id
	corpId string

	//在开发者后台的基本信息 > 开发信息（旧版）页面获取微应用管理后台SSOSecret
	SSOSecret string `json:"SSOSecret"`

	client *http.Client

	cache cache.Cache
}

type OptionFunc func(*dingTalk)

func WithTicket(ticket string) OptionFunc {
	return func(dt *dingTalk) {
		dt.ticket = ticket
	}
}

func WithCorpId(corpId string) OptionFunc {
	return func(dt *dingTalk) {
		dt.corpId = corpId
	}
}

func WithSSOSecret(secret string) OptionFunc {
	return func(dt *dingTalk) {
		dt.SSOSecret = secret
	}
}

//isv 是否isv
func (ding *dingTalk) isv() bool {
	return len(ding.ticket) > 0 && len(ding.corpId) > 0
}

// NewClient new DingTalkBuilder
func NewClient(id int, key, secret string, opts ...OptionFunc) (ding *dingTalk) {
	ding = &dingTalk{Id: id, Key: key, Secret: secret}
	for _, opt := range opts {
		opt(ding)
	}

	if ding.isv() {
		ding.cache = cache.NewFileCache(strings.Join([]string{".token", "isv", key}, "/"), ding.corpId)
	} else {
		ding.cache = cache.NewFileCache(strings.Join([]string{".token", "corp"}, "/"), key)
	}
	ding.client = &http.Client{Timeout: 10 * time.Second}

	if err := validate(ding); err != nil {
		panic(err)
		return nil
	}
	return
}

//Request 统一请求
func (ding *dingTalk) Request(method, path string, query url.Values, body interface{},
	data response.Unmarshalled) (err error) {

	if body != nil {
		if err = validate(body); err != nil {
			return err
		}
	}

	if query == nil {
		query = url.Values{}
	}

	if !query.Has("access_token") && path != constant.GetTokenKey && path != constant.CorpAccessToken &&
		path != constant.SuiteAccessToken && path != constant.GetAuthInfo && path != constant.GetAgentKey &&
		path != constant.ActivateSuiteKey && path != constant.GetSSOTokenKey && path != constant.GetUnactiveCorpKey &&
		path != constant.ReauthCorpKey && path != constant.GetCorpPermanentCodeKey {

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
		//set token
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

	//表单不为空
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

func (ding *dingTalk) httpRequest(method, path string, query url.Values, body interface{},
	response response.Unmarshalled) error {

	var (
		req    *http.Request
		res    *http.Response
		err    error
		data   []byte
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
		//检查提交表单类型
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
			//表单不为空
			d, _ := json.Marshal(body)
			req, _ = http.NewRequest(method, uri.String(), bytes.NewReader(d))
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

	switch res.StatusCode {
	case 400:
		return errors.Errorf("dingtalk server error: status=%d", res.StatusCode)
	case 500:
		return errors.Errorf("dingtalk server error,status=%d", res.StatusCode)
	}

	if err = json.Unmarshal(data, response); err != nil {
		return err
	}
	fmt.Println("res=", string(data))
	return response.CheckError()
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
