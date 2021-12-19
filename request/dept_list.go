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

type DeptList struct {
	//父部门ID。
	//
	//如果不传，默认部门为根部门，根部门ID为1。只支持查询下一级子部门，不支持查询多级子部门。
	DeptId int `json:"dept_id,omitempty"`

	// 通讯录语言，默认zh_CN。如果是英文，请传入en_US。
	Language string `json:"language,omitempty" validate:"omitempty,oneof=zh_CN en_US"`
}

type deptListBuilder struct {
	dl *DeptList
}

func NewDeptList() *deptListBuilder {
	return &deptListBuilder{dl: &DeptList{}}
}

func (db *deptListBuilder) SetDeptId(deptId int) *deptListBuilder {
	db.dl.DeptId = deptId
	return db
}

func (db *deptListBuilder) SetLanguage(language language.Language) *deptListBuilder {
	db.dl.Language = string(language)
	return db
}

func (db *deptListBuilder) Build() *DeptList {
	return db.dl
}
