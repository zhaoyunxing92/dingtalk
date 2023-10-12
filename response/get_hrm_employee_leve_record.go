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

type GetHrmEmployeeLevelRecord struct {
	Response

	// 离职记录列表。
	Records []employeeField `json:"records"`

	// 分页游标，首次查询，该参数传0；非首次查询，该参数传上次调用本接口返回的nextToken。
	NextToken string `json:"nextToken"`
}

type employeeField struct {

	// 员工的userId
	UserId string `json:"userId"`

	// 员工名字。

	Name string `json:"name"`

	// 员工的userId
	StateCode string `json:"stateCode"`

	// 手机号码
	Mobile string `json:"mobile"`

	// 离职时间 YYYY-MM-DDTHH:mm:ssZ（ISO 8601/RFC 3339）
	LeaveTime string `json:"leaveTime"`

	// 退出企业方式，取值：oapi：调用接口删除；cancel：注销；leave：主动离职；unknown：未知原因；delete：管理员删除
	LeaveReason string `json:"leaveReason"`
}
