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

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant/language"
)

type DeptDetail struct {
	DeptId int `json:"dept_id" validate:"required"`

	//通讯录语言，默认zh_CN。如果是英文，请传入en_US。
	Language string `json:"language,omitempty" validate:"omitempty,oneof=zh_CN en_US"`
}

type deptDetailBuilder struct {
	d *DeptDetail
}

func NewDeptDetail(deptId int) *deptDetailBuilder {
	return &deptDetailBuilder{d: &DeptDetail{DeptId: deptId}}
}

func (db *deptDetailBuilder) SetLanguage(language language.Language) *deptDetailBuilder {
	db.d.Language = string(language)
	return db
}

func (db *deptDetailBuilder) Build() *DeptDetail {
	return db.d
}
