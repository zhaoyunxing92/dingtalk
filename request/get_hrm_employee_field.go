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

import "strings"

type GetHrmEmployeeField struct {

	// 应用AgentID，可在钉钉开发者后台的应用详情页获取。
	AgentId int `json:"agentid"`

	// 员工的userid列表，多个userid之间使用逗号分隔，一次最多支持传100个值
	UserIds string `json:"userid_list" validate:"required,max=100"`

	UserIdList []string `json:"-"`

	// 需要获取的花名册字段信息。查询字段越少，RT越低，建议按需查询，多个字段之间使用逗号分隔，一次最多支持传100个值
	Fields string `json:"field_filter_list,omitempty"`

	FieldList []string `json:"-" validate:"omitempty,max=100"`
}

func NewGetHrmEmployeeField(agentId int, userIds []string, fields []string) *GetHrmEmployeeField {
	us := removeStringDuplicates(userIds)
	fs := removeStringDuplicates(fields)
	return &GetHrmEmployeeField{
		AgentId: agentId, UserIdList: us, UserIds: strings.Join(us, ","),
		FieldList: fs, Fields: strings.Join(fs, ","),
	}
}
