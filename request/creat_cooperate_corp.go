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

//CreateCooperateCorp 创建合作空间
type CreateCooperateCorp struct {
	//合作空间组织名称
	Name string `json:"orgName" validate:"required"`

	//合作空间的logo
	LogoMediaId string `json:"logoMediaId,omitempty"`

	//行业code
	IndustryCode int `json:"industryCode,omitempty"`
}

type createCooperateCorpBuilder struct {
	c *CreateCooperateCorp
}

func NewCreateCooperateCorp(name string) *createCooperateCorpBuilder {
	return &createCooperateCorpBuilder{c: &CreateCooperateCorp{Name: name}}
}

func (cb *createCooperateCorpBuilder) SetLogoMediaId(mediaId string) *createCooperateCorpBuilder {
	cb.c.LogoMediaId = mediaId
	return cb
}

func (cb *createCooperateCorpBuilder) SetIndustryCode(code int) *createCooperateCorpBuilder {
	cb.c.IndustryCode = code
	return cb
}

func (cb *createCooperateCorpBuilder) Build() *CreateCooperateCorp {
	return cb.c
}
