package model

import (
	"fmt"
)

//{"errcode":40035,"errmsg":"缺少参数 corpid or appkey"}
//响应
type Response struct {
	Code int    `json:"errcode"` //code
	Msg  string `json:"errmsg"`  //msg
}

//统一检查返回异常异常
type Unmarshallable interface {
	CheckError() error
}

func (res *Response) CheckError() (err error) {
	if res.Code != 0 {
		err = fmt.Errorf("%d: %s", res.Code, res.Msg)
	}
	return err
}
