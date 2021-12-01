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

//SuiteAccessToken 获取suite套件token
type SuiteAccessToken struct {
	Key string `json:"suite_key" validate:"required"`

	Secret string `json:"suite_secret" validate:"required"`

	Ticket string `json:"suite_ticket" validate:"required"`
}

func NewSuiteAccessToken(key, secret, ticket string) *SuiteAccessToken {
	return &SuiteAccessToken{key, secret, ticket}
}
