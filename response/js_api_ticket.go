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

package response

type JsApiTicket struct {
	Response
	// Expires 过期时间
	Expires int16 `json:"expires_in"`
	// Create 创建时间
	Create int64 `json:"create"`
	// Ticket 生成的临时jsapi_ticket。
	Ticket string `json:"ticket"`
}

// CreatedAt is when the access token is generated
func (token *JsApiTicket) CreatedAt() int64 {
	return token.Create
}

// ExpiresIn is how soon the access token is expired
func (token *JsApiTicket) ExpiresIn() int16 {
	return token.Expires
}
