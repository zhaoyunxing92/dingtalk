package domain

import (
	"fmt"
)

//{"errcode":40035,"errmsg":"缺少参数 corpid or appkey"}
type Response struct {
	Code uint64 `json:"errcode"` //code
	Msg  string `json:"errmsg"`  //msg
}

//统一检查异常
type Unmarshallable interface {
	CheckError() error
}

func (data *Response) CheckError() (err error) {
	if data.Code != 0 {
		err = fmt.Errorf("%d: %s", data.Code, data.Msg)
	}
	return err
}
