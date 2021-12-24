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

type HrmResignEmployee struct {
	Response

	// 离职人员
	ResignEmployee []resignEmployee `json:"result"`
}

type resignEmployee struct {

	// 离职员工的userid。
	UserId string `json:"userid"`

	// 最后工作日
	LastWorkDay int `json:"last_work_day"`

	// 离职部门列表
	DeptList []EmpDept `json:"dept_list"`

	// 离职原因备注
	ReasonMemo string `json:"reason_memo"`

	//离职原因类型：
	//
	//1：家庭原因
	//
	//2：个人原因
	//
	//3：发展原因
	//
	//4：合同到期不续签
	//
	//5：协议解除
	//
	//6：无法胜任工作
	//
	//7：经济性裁员
	//
	//8：严重违法违纪
	//
	//9：其他
	ReasonType int `json:"reason_type"`

	//离职状态：
	//
	//1：待离职
	//
	//2：已离职
	//
	//3：未离职
	//
	//4：发起离职审批但还未通过
	//
	//5：失效（离职流程被其他流程强制终止后的状态）
	PreStatus int `json:"pre_status"`

	// 离职前主部门名称
	MainDeptName string `json:"main_dept_name"`

	// 离职前主部门ID
	MainDeptId int `json:"main_dept_id"`
}

type EmpDept struct {
	// 部门名称
	Name string `json:"dept_path"`

	// 部门id
	Id int `json:"dept_id"`
}
