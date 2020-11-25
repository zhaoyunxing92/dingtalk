package dingtalk

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/zhaoyunxing92/dingtalk/domain"
	"github.com/zhaoyunxing92/dingtalk/global"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

//统一请求
func (talk *DingTalk) request(method, path string, params url.Values, form interface{}, data domain.Unmarshallable) (err error) {
	if params == nil {
		params = url.Values{}
	}
	if path != global.GetTokenKey {
		var token string
		if token, err = talk.GetToken(); err != nil {
			return err
		}
		//set token
		params.Set("access_token", token)
	}
	return talk.httpRequest(method, path, params, form, data)
}

func (talk *DingTalk) httpRequest(method, path string, args url.Values, form interface{}, data domain.Unmarshallable) error {
	client := talk.client

	uri, _ := url.Parse(global.Host + path)
	uri.RawQuery = args.Encode()

	var (
		res     *http.Request
		rep     *http.Response
		err     error
		content []byte
	)

	if form != nil {
		//检查提交表单类型
		switch form.(type) {
		case domain.UploadFile:
			var b bytes.Buffer
			w := multipart.NewWriter(&b)

			file := form.(domain.UploadFile)
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
			res, _ = http.NewRequest("POST", uri.String(), &b)
			res.Header.Set("Content-Type", w.FormDataContentType())
		default:
			//表单不为空
			d, _ := json.Marshal(form)
			res, _ = http.NewRequest("POST", uri.String(), bytes.NewReader(d))
		}
	} else {
		if method == "GET" {
			res, _ = http.NewRequest("GET", uri.String(), nil)
		} else {
			res, _ = http.NewRequest("POST", uri.String(), nil)
		}
	}

	if rep, err = client.Do(res); err != nil {
		return err
	}

	defer rep.Body.Close()
	if rep.StatusCode != 200 {
		return errors.New("dingtalk server error: " + rep.Status)
	}

	if content, err = ioutil.ReadAll(rep.Body); err != nil {
		return err
	}
	if err = json.Unmarshal(content, data); err != nil {
		return err
	}
	return data.CheckError()
}
