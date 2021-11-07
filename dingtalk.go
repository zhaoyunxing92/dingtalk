package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	"github.com/zhaoyunxing92/dingtalk/v2/model"
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
	Secret string `json:"secret,omitempty" validate:"required"`

	// isv 钉钉开放平台会向应用的回调URL推送的suite_ticket（约5个小时推送一次）
	Ticket string `json:"ticket,omitempty"`

	//授权企业的id
	CorpId string `json:"corpId"`

	Client *http.Client

	Cache cache.Cache
}

//isv 是否isv
func (ding *dingTalk) isv() bool {
	return len(ding.Ticket) > 0 && len(ding.CorpId) > 0
}

// builder builder for dingTalk
type builder struct {
	ding *dingTalk
}

// NewClient new DingTalkBuilder
func NewClient(id int, key, secret string) *builder {
	return &builder{ding: &dingTalk{Id: id, Key: key, Secret: secret}}
}

//SetTicket set app ticket
func (b *builder) SetTicket(ticket string) *builder {
	if len(ticket) <= 0 {
		panic(errors.New("ticket is invalid"))
	}
	b.ding.Ticket = ticket
	return b
}

//SetCorpId set app ticket
func (b *builder) SetCorpId(corpId string) *builder {
	if len(corpId) <= 0 {
		panic(errors.New("corpId is invalid"))
	}
	b.ding.CorpId = corpId
	return b
}

//Build return dingtalk
func (b *builder) Build() *dingTalk {

	if err := validate(b); err != nil {
		panic(err)
	}
	ding := b.ding
	key := ding.Key
	//判断是否isv
	if ding.isv() {
		ding.Cache = cache.NewFileCache(strings.Join([]string{".token", "isv", key}, "/"), ding.CorpId)
	} else {
		ding.Cache = cache.NewFileCache(strings.Join([]string{".token", "corp"}, "/"), key)
	}
	ding.Client = &http.Client{Timeout: 10 * time.Second}
	return b.ding
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

	if path != constant.GetTokenKey && path != constant.CorpAccessToken &&
		path != constant.SuiteAccessToken {
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

	uri, _ := url.Parse(constant.Host + path)
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

	defer res.Body.Close()
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
	data response.Unmarshalled) error {

	client := ding.Client
	uri, _ := url.Parse(constant.Host + path)
	uri.RawQuery = query.Encode()

	var (
		req     *http.Request
		res     *http.Response
		err     error
		content []byte
	)

	if method != http.MethodGet {
		//检查提交表单类型
		switch body.(type) {
		case model.UploadFile:
			var b bytes.Buffer
			w := multipart.NewWriter(&b)

			file := body.(model.UploadFile)
			fw, err := w.CreateFormFile(file.FieldName, file.FileName)
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
			fmt.Println(string(d))
			req, _ = http.NewRequest(method, uri.String(), bytes.NewReader(d))
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
		}
	} else {
		req, _ = http.NewRequest(method, uri.String(), nil)
	}

	if res, err = client.Do(req); err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("dingtalk server error: " + res.Status)
	}

	if content, err = ioutil.ReadAll(res.Body); err != nil {
		return err
	}
	fmt.Println("content=", string(content))
	if err = json.Unmarshal(content, data); err != nil {
		return err
	}

	return data.CheckError()
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