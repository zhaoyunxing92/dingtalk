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

type GetHrmMeta struct {
	Response

	GroupMeta []HrmGroupMeta `json:"result"`
}

type HrmGroupMeta struct {

	// 分组名称。
	GroupName string `json:"group_name"`

	// 分组标识。
	GroupId string `json:"group_id"`

	// 分组是否支持明细。
	Detail bool `json:"detail"`

	FieldMeta []FieldMeta `json:"field_meta_info_list"`
}

type FieldMeta struct {
	// 字段名称
	FieldName string `json:"field_name"`

	// 字段标识
	FieldCode string `json:"field_code"`

	// 是否衍生字段，例如司龄、年龄等系统计算的字段
	Derived bool `json:"derived"`
}
