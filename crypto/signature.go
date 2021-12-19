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

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// GetSignature 第三方访问接口的签名计算方法
// https://developers.dingtalk.com/document/app/signature-calculation-method-for-third-party-access-interfaces-1
func GetSignature(timestamp, secret string, ticket string) string {
	str := timestamp + "\n" + ticket
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(str))
	// return url.QueryEscape(sign)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// GetAvoidLoginSignature 个人免登场景的签名计算方法
// https://developers.dingtalk.com/document/app/signature-calculation-for-logon-free-scenarios-1
func GetAvoidLoginSignature(timestamp, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(timestamp))
	// return url.QueryEscape(sign)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
