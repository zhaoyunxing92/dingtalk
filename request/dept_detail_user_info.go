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

type DeptDetailUserInfo struct {
	//部门ID，根部门ID为1
	DeptId int `json:"dept_id" validate:"required,min=1"`

	//分页查询的游标，最开始传0，后续传返回参数中的next_cursor值
	Cursor int `json:"cursor" validate:"omitempty,min=0"`

	//分页长度，最大值100。
	Size int `json:"size" validate:"required,max=100"`

	//部门成员的排序规则：
	//
	//entry_asc：代表按照进入部门的时间升序。
	//
	//entry_desc：代表按照进入部门的时间降序。
	//
	//modify_asc：代表按照部门信息修改时间升序。
	//
	//modify_desc：代表按照部门信息修改时间降序。
	//
	//custom：代表用户定义(未定义时按照拼音)排序。
	//
	//默认值：custom。
	OrderField string `json:"order_field,omitempty" validate:"omitempty,oneof=entry_asc entry_desc modify_asc modify_desc"`

	//是否返回访问受限的员工
	ContainAccessLimit bool `json:"contain_access_limit,omitempty"`

	//通讯录语言，默认zh_CN。如果是英文，请传入en_US。
	Language string `json:"language,omitempty" validate:"omitempty,oneof=zh_CN en_US"`
}

type deptDetailUserInfoBuilder struct {
	du *DeptDetailUserInfo
}

func NewDeptDetailUserInfo(deptId, cursor, size int) *deptDetailUserInfoBuilder {
	return &deptDetailUserInfoBuilder{du: &DeptDetailUserInfo{DeptId: deptId, Cursor: cursor, Size: size}}
}

//SetOrderField 部门成员的排序规则：
//
//entry_asc：代表按照进入部门的时间升序。
//
//entry_desc：代表按照进入部门的时间降序。
//
//modify_asc：代表按照部门信息修改时间升序。
//
//modify_desc：代表按照部门信息修改时间降序。
//
//custom：代表用户定义(未定义时按照拼音)排序。
//
//默认值：custom。
func (dub *deptDetailUserInfoBuilder) SetOrderField(orderField string) *deptDetailUserInfoBuilder {
	dub.du.OrderField = orderField
	return dub
}

//SetContainAccessLimit 是否返回访问受限的员工
func (dub *deptDetailUserInfoBuilder) SetContainAccessLimit(containAccessLimit bool) *deptDetailUserInfoBuilder {
	dub.du.ContainAccessLimit = containAccessLimit
	return dub
}

//SetLanguage 通讯录语言，默认zh_CN。如果是英文，请传入en_US。
func (dub *deptDetailUserInfoBuilder) SetLanguage(language string) *deptDetailUserInfoBuilder {
	dub.du.Language = language
	return dub
}

func (dub *deptDetailUserInfoBuilder) Build() *DeptDetailUserInfo {
	return dub.du
}
