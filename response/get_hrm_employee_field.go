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

type GetHrmEmployeeField struct {
	Response

	//应用AgentID，可在钉钉开发者后台的应用详情页获取。
	EmpFields []empRosterField `json:"result"`
}

type empRosterField struct {

	//企业的corpid
	CorpId string `json:"corp_id"`

	//员工的userid
	UserId string `json:"userid"`

	//返回的字段信息列表
	Fields []empField `json:"field_data_list"`
}

type empField struct {

	//字段描述
	FieldName string `json:"field_name"`

	//字段code
	FieldCode string `json:"field_code"`

	//分组标识
	GroupId string `json:"group_id"`

	//字段值列表。
	//
	//明细分组字段包含多条。
	//
	//非明细分组仅一条记录。
	Values []fieldValue `json:"field_value_list"`
}

type fieldValue struct {

	//第几条的明细标识，下标从0开始。
	Index int `json:"item_index"`

	//字段展示值，选项类型字段对应选项的value
	Label string `json:"label"`

	//字段取值，选项类型字段对应选项的key。
	Value string `json:"value"`
}
