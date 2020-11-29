package model

import (
	"fmt"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

//{"errcode":40035,"errmsg":"缺少参数 corpid or appkey"}
//Response:响应
type Response struct {
	Code      int    `json:"errcode"` //code
	Msg       string `json:"errmsg"`  //msg
	Success   bool   `json:"success"`
	RequestId string `json:"request_id"`
}

//Unmarshallable:统一检查返回异常异常
type Unmarshallable interface {
	CheckError() error
}

//Request:请求
type Request interface {
	Validate(valid *validator.Validate, trans translator.Translator) error
}

func (res *Response) CheckError() (err error) {
	if res.Code != 0 {
		err = fmt.Errorf("%d: %s", res.Code, res.Msg)
	}
	return err
}
