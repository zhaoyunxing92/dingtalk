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
func (ding dingTalk) GetHrmEmployee(offset, size int, status []employee.Status) (res response.GetHrmEmployee,
	err error) {
	req := request.NewGetHrmEmployee(offset, size, status)
	return res, ding.Request(http.MethodPost, constant.GetHrmEmployeeKey, nil, req, &res)
}

// GetHrmToBeHiredEmployee 获取待入职员工列表
func (ding dingTalk) GetHrmToBeHiredEmployee(offset, size int) (res response.GetHrmEmployee,
	err error) {
	req := request.NewGetHrmToBeHiredEmployee(offset, size)
	return res, ding.Request(http.MethodPost, constant.GetHrmToBeHiredEmployeeKey, nil, req, &res)
}

// GetHrmResignEmployeeIds 获取待入职员工列表
func (ding dingTalk) GetHrmResignEmployeeIds(offset, size int) (res response.GetHrmEmployee, err error) {
	req := request.NewGetHrmToBeHiredEmployee(offset, size)
	return res, ding.Request(http.MethodPost, constant.GetHrmResignEmployeeKey, nil, req, &res)
}

// GetHrmResignEmployee 获取员工离职信息
func (ding dingTalk) GetHrmResignEmployee(userIds []string) (res response.HrmResignEmployee, err error) {
	req := request.NewGetHrmResignEmployee(userIds)
	return res, ding.Request(http.MethodPost, constant.GetHrmResignEmployeeInfoKey, nil, req, &res)
}

// HrmCreateEmployee 添加企业待入职员工
func (ding dingTalk) HrmCreateEmployee(req *request.HrmCreateEmployee) (res response.HrmCreateEmployee, err error) {
	return res, ding.Request(http.MethodPost, constant.HrmCreateEmployeeKey, nil, req, &res)
}
