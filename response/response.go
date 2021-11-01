package response

import (
	"github.com/pkg/errors"
)

//Response 响应
//{"errcode":40035,"errmsg":"缺少参数 corpid or appkey"}
type Response struct {
	Code      int    `json:"errcode"`
	Msg       string `json:"errmsg,omitempty"`
	Success   bool   `json:"success,omitempty"`
	RequestId string `json:"request_id,omitempty"`
}

//Unmarshalled 统一检查返回异常异常
type Unmarshalled interface {
	CheckError() error
}

func (res *Response) CheckError() (err error) {
	if res.Code != 0 {
		err = errors.Errorf("%d: %s", res.Code, res.Msg)
	}
	return err
}
