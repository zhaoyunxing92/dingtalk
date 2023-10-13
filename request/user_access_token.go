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

package request

const (
	ActionGetUserAuthToken = "authorization_code"
	ActionRefreshAuthToken = "refresh_token"
)

type UserAccessToken struct {
	ClientId     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	Code         string `json:"code,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	GrantType    string `json:"grantType,omitempty"`
}

func NewUserAuthToken(authCode string) *UserAccessToken {
	return &UserAccessToken{
		Code:      authCode,
		GrantType: ActionGetUserAuthToken,
	}
}

func NewRefreshUserAuthToken(refreshToken string) *UserAccessToken {
	return &UserAccessToken{
		RefreshToken: refreshToken,
		GrantType:    ActionRefreshAuthToken,
	}
}
