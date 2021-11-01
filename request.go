package dingtalk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/zhaoyunxing92/dingtalk/constant"
	"github.com/zhaoyunxing92/dingtalk/model"
	"github.com/zhaoyunxing92/dingtalk/response"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

//统一请求
func (ding *DingTalk) request(method, path string, params url.Values, form interface{},
	data response.Unmarshalled) (err error) {

	if form != nil {
		if err = validate(form); err != nil {
			return err
		}
	}

	if params == nil {
		params = url.Values{}
	}
	if path != constant.GetTokenKey {
		var token string
		if token, err = ding.GetAccessToken(); err != nil {
			return err
		}
		//set token
		params.Set("access_token", token)
	}
	return ding.httpRequest(method, path, params, form, data)
}

func (robot *Robot) send(form interface{}, data response.Unmarshalled) (err error) {
	args := url.Values{}
	args.Set("access_token", robot.Token)

	return robot.httpRequest(http.MethodPost, constant.SendRobotMsgKey, args, form, data)
}

func (robot *Robot) httpRequest(method, path string, args url.Values, form interface{},
	data response.Unmarshalled) error {

	client := robot.client

	uri, _ := url.Parse(constant.Host + path)
	uri.RawQuery = args.Encode()

	var (
		req     *http.Request
		res     *http.Response
		err     error
		content []byte
	)

	//表单不为空
	d, _ := json.Marshal(form)
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

func (ding *DingTalk) httpRequest(method, path string, args url.Values, form interface{},
	data response.Unmarshalled) error {

	client := ding.Client
	uri, _ := url.Parse(constant.Host + path)
	uri.RawQuery = args.Encode()

	var (
		req     *http.Request
		res     *http.Response
		err     error
		content []byte
	)

	if form != nil {
		//检查提交表单类型
		switch form.(type) {
		case model.UploadFile:
			var b bytes.Buffer
			w := multipart.NewWriter(&b)

			file := form.(model.UploadFile)
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
			d, _ := json.Marshal(form)
			req, _ = http.NewRequest(method, uri.String(), bytes.NewReader(d))
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
	fmt.Println(string(content))
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
