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
	"encoding/json"
	"strings"
	"time"

	"github.com/zhaoyunxing92/dingtalk/v2/constant/employee"
)

type HrmCreateEmployee struct {
	// 参数
	Param *hrmCreateEmployee `json:"param" validate:"required"`
}

type hrmCreateEmployee struct {

	// 待入职员工姓名
	Name string `json:"name" validate:"required"`

	// 待入职员工手机号
	Mobile string `json:"mobile" validate:"required"`

	// 预期入职时间
	EntryTime string `json:"pre_entry_time,omitempty"`

	// 操作人userid
	Operator string `json:"op_userid,omitempty"`

	// 扩展信息，json串格式，按要求传入有效信息，无效信息不会保存
	Ext *hrmCreateEmployeeExt `json:"extend_info,omitempty"`
}

// hrmCreateEmployeeExt 扩展字段
type hrmCreateEmployeeExt struct {

	// 部门ID列表，多个部门用"|"分隔
	Depts string `json:"depts,omitempty"`

	// 部门列表
	DeptList []int `json:"-"`

	// 主部门ID
	MainDeptId *int `json:"mainDeptId,omitempty"`

	// 主部门名称
	MainDeptName string `json:"mainDeptName,omitempty"`

	// 职位
	Position string `json:"position,omitempty"`

	// 工作地点
	WorkPlace string `json:"workPlace,omitempty"`

	// 工号
	JobNumber string `json:"jobNumber,omitempty"`

	//0：无类型
	//
	//1：全职
	//
	//2：兼职
	//
	//3：实习
	//
	//4：劳务派遣
	//
	//5：退休返聘
	//
	//6：劳务外包
	EmployeeType string `json:"employeeType,omitempty"`
}

type hrmCreateEmployeeBuilder struct {
	emp *HrmCreateEmployee
}

func (h *HrmCreateEmployee) String() string {
	if str, err := json.Marshal(h); err != nil {
		return ""
	} else {
		return string(str)
	}
}

func NewHrmCreateEmployee(name, mobile string) *hrmCreateEmployeeBuilder {
	return &hrmCreateEmployeeBuilder{emp: &HrmCreateEmployee{Param: &hrmCreateEmployee{
		Name: name, Mobile: mobile,
		Ext: &hrmCreateEmployeeExt{},
	}}}
}

func (h *hrmCreateEmployeeBuilder) SetEntryTime(time time.Time) *hrmCreateEmployeeBuilder {
	h.emp.Param.EntryTime = time.Format("2006-01-02 15:04:05")
	return h
}

func (h *hrmCreateEmployeeBuilder) SetOperator(operator string) *hrmCreateEmployeeBuilder {
	h.emp.Param.Operator = operator
	return h
}

func (h *hrmCreateEmployeeBuilder) SetDeptList(ids []int) *hrmCreateEmployeeBuilder {
	h.emp.Param.Ext.DeptList = ids
	return h
}

func (h *hrmCreateEmployeeBuilder) SetMainDeptId(id int) *hrmCreateEmployeeBuilder {
	h.emp.Param.Ext.MainDeptId = &id
	return h
}

func (h *hrmCreateEmployeeBuilder) SetMainDeptName(deptName string) *hrmCreateEmployeeBuilder {
	h.emp.Param.Ext.MainDeptName = deptName
	return h
}

func (h *hrmCreateEmployeeBuilder) SetPosition(position string) *hrmCreateEmployeeBuilder {
	h.emp.Param.Ext.Position = position
	return h
}

func (h *hrmCreateEmployeeBuilder) SetWorkPlace(workPlace string) *hrmCreateEmployeeBuilder {
	h.emp.Param.Ext.WorkPlace = workPlace
	return h
}

func (h *hrmCreateEmployeeBuilder) SetJobNumber(jobNumber string) *hrmCreateEmployeeBuilder {
	h.emp.Param.Ext.JobNumber = jobNumber
	return h
}

func (h *hrmCreateEmployeeBuilder) SetEmployeeType(mold employee.Mold) *hrmCreateEmployeeBuilder {
	h.emp.Param.Ext.EmployeeType = string(mold)
	return h
}

func (h *hrmCreateEmployeeBuilder) Build() *HrmCreateEmployee {
	ids := h.emp.Param.Ext.DeptList
	if len(ids) >= 0 {
		h.emp.Param.Ext.Depts = strings.Join(removeIntDuplicatesToString(ids), "|")
	}
	return h.emp
}
