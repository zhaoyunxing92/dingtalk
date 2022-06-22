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

package dingtalk

import (
	"net/http"

	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/employee"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

// GetHrmEmployee 获取在职员工列表
func (ding DingTalk) GetHrmEmployee(offset, size int, status []employee.Status) (res response.GetHrmEmployee,
	err error) {
	req := request.NewGetHrmEmployee(offset, size, status)
	return res, ding.Request(http.MethodPost, constant.GetHrmEmployeeKey, nil, req, &res)
}

// GetHrmToBeHiredEmployee 获取待入职员工列表
func (ding DingTalk) GetHrmToBeHiredEmployee(offset, size int) (res response.GetHrmEmployee,
	err error) {
	req := request.NewGetHrmToBeHiredEmployee(offset, size)
	return res, ding.Request(http.MethodPost, constant.GetHrmToBeHiredEmployeeKey, nil, req, &res)
}

// GetHrmResignEmployeeIds 获取离职员工列表
func (ding DingTalk) GetHrmResignEmployeeIds(offset, size int) (res response.GetHrmEmployee, err error) {
	req := request.NewGetHrmToBeHiredEmployee(offset, size)
	return res, ding.Request(http.MethodPost, constant.GetHrmResignEmployeeKey, nil, req, &res)
}

// GetHrmResignEmployee 获取员工离职信息
func (ding DingTalk) GetHrmResignEmployee(userIds []string) (res response.HrmResignEmployee, err error) {
	req := request.NewGetHrmResignEmployee(userIds)
	return res, ding.Request(http.MethodPost, constant.GetHrmResignEmployeeInfoKey, nil, req, &res)
}

// HrmCreateEmployee 添加企业待入职员工
func (ding DingTalk) HrmCreateEmployee(req *request.HrmCreateEmployee) (res response.HrmCreateEmployee, err error) {
	return res, ding.Request(http.MethodPost, constant.HrmCreateEmployeeKey, nil, req, &res)
}

// GetHrmField 获取花名册字段组详情
func (ding DingTalk) GetHrmField(agentId int) (res response.GetHrmField, err error) {
	return res, ding.Request(http.MethodPost, constant.GetHrmFieldKey, nil, request.NewGetHrmField(agentId), &res)
}

// GetHrmEmployeeField 获取员工花名册字段信息
func (ding DingTalk) GetHrmEmployeeField(agentId int, userIds []string, fields []string) (res response.GetHrmEmployeeField, err error) {
	return res, ding.Request(http.MethodPost, constant.GetHrmEmployeeFieldKey, nil,
		request.NewGetHrmEmployeeField(agentId, userIds, fields), &res)
}

// UpdateHrmEmployeeField 更新员工花名册信息
func (ding DingTalk) UpdateHrmEmployeeField(req *request.UpdateHrmEmpField) (res response.Response, err error) {
	return res, ding.Request(http.MethodPost, constant.UpdateHrmEmployeeFieldKey, nil, req, &res)
}

// GetHrmMeta 获取花名册元数据
func (ding DingTalk) GetHrmMeta(agentId int) (res response.GetHrmMeta, err error) {
	return res, ding.Request(http.MethodPost, constant.GetHrmMetaKey, nil,
		request.NewGetHrmField(agentId), &res)
}
