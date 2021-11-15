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

//AccessToken 钉钉token结构体
type AccessToken struct {
	Response
	Expires int16  `json:"expires_in"`   //过期时间
	Created int64  `json:"created"`      //创建时间
	Token   string `json:"access_token"` //token
}

//CreatedAt Expired.CreatedAt is when the access token is generated
func (token *AccessToken) CreatedAt() int64 {
	return token.Created
}

//ExpiresIn is how soon the access token is expired
func (token *AccessToken) ExpiresIn() int16 {
	return token.Expires
}
