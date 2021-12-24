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

type GetHrmField struct {
	Response

	//应用AgentID，可在钉钉开发者后台的应用详情页获取。
	FieldGroups []fieldGroupMeta `json:"result"`
}

type fieldGroupMeta struct {

	//字段组ID
	GroupId string `json:"group_id"`

	//是否支持明细
	HasDetail bool `json:"has_detail"`

	Fields []employeeField `json:"field_list"`
}

type employeeField struct {

	//字段类型：
	//
	//DDSelectField
	//
	//TextField
	//
	//DDDateField
	//
	//DDDateWithLongField
	//
	//DDPhotoField
	FieldType string `json:"field_type"`

	//字段描述
	FieldName string `json:"field_name"`

	//字段code
	FieldCode string `json:"field_code"`
}
