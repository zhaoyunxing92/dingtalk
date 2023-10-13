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

import "github.com/pkg/errors"

type UserAccessToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	ExpireIn     int    `json:"expireIn,omitempty"`
	CorpId       string `json:"corpId,omitempty"`

	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	RequestId string `json:"requestid,omitempty"`
}

func (u UserAccessToken) CheckError() (err error) {
	if u.Code != "" {
		err = errors.Errorf("code:'%s', msg:%s, requestId: %s", u.Code, u.Message, u.RequestId)
	}
	return
}
