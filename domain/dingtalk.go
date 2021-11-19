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

package domain

import (
	"fmt"
)

import (
	translator "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
)

//Response 响应
//{"errcode":40035,"errmsg":"缺少参数 corpid or appkey"}
type Response struct {
	Code      int    `json:"errcode"`          //code
	Msg       string `json:"errmsg,omitempty"` //msg
	Success   bool   `json:"success,omitempty"`
	RequestId string `json:"request_id,omitempty"`
}

//Request 请求
type Request interface {
	Validate(valid *validator.Validate, trans translator.Translator) error
}

func (res *Response) CheckError() (err error) {
	if res.Code != 0 {
		err = fmt.Errorf("%d: %s", res.Code, res.Msg)
	}
	return err
}
