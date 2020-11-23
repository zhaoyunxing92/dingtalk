package dingtalk_go

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/zhaoyunxing92/dingtalk/domain"
	"github.com/zhaoyunxing92/dingtalk/global"
	"io/ioutil"
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
		request *http.Request
		resp    *http.Response
		err     error
		content []byte
	)

	switch method {
	case "POST":
		//表单不为空
		d, _ := json.Marshal(form)
		request, _ = http.NewRequest("POST", uri.String(), bytes.NewReader(d))
		break
	default:
		request, _ = http.NewRequest("GET", uri.String(), nil)
	}

	if resp, err = client.Do(request); err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("dingtalk server error: " + resp.Status)
	}

	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return err
	}
	if err = json.Unmarshal(content, data); err != nil {
		return err
	}
	return data.CheckError()
}
