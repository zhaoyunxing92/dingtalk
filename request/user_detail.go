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

import "encoding/json"

type UserDetail struct {
	UserId string `json:"userid" validate:"required"`

	//通讯录语言，默认zh_CN。如果是英文，请传入en_US。
	Language string `json:"language,omitempty" validate:"omitempty,oneof=zh_CN en_US"`
}

func (u *UserDetail) String() string {
	str, _ := json.Marshal(u)
	return string(str)
}

type userDetailBuilder struct {
	user *UserDetail
}

func NewUserDetail(userId string) *userDetailBuilder {
	return &userDetailBuilder{user: &UserDetail{UserId: userId}}
}

//SetLanguage 通讯录语言，默认zh_CN。如果是英文，请传入en_US。
func (ub *userDetailBuilder) SetLanguage(language string) *userDetailBuilder {
	ub.user.Language = language
	return ub
}

func (ub *userDetailBuilder) Build() *UserDetail {
	return ub.user
}
