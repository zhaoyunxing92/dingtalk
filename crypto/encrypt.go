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

package crypto

import "encoding/json"

// DingTalkEncrypt 钉钉事件回调加密
type DingTalkEncrypt struct {

	// 加密体
	Encrypt string `json:"encrypt,omitempty"`

	// 签名
	Sign string `json:"msg_signature,omitempty"`

	// 时间戳
	Timestamp string `json:"timeStamp,omitempty"`

	// 随机字符串
	Nonce string `json:"nonce,omitempty"`
}

// NewDingTalkEncrypt 钉钉返回体
func NewDingTalkEncrypt(encrypt, sign, timestamp, nonce string) *DingTalkEncrypt {
	return &DingTalkEncrypt{encrypt, sign, timestamp, nonce}
}

// String 转换成json格式
func (d *DingTalkEncrypt) String() string {
	if str, err := json.Marshal(d); err != nil {
		return ""
	} else {
		return string(str)
	}
}
