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
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

// CreateDept 创建部门
func (ding *dingTalk) CreateDept(res *request.CreateDept) (rsp response.CreateDept, err error) {
	return rsp, ding.Request(http.MethodPost, constant.CreateDeptKey, nil, res, &rsp)
}

// DeleteDept 删除部门
func (ding *dingTalk) DeleteDept(deptId int) (rsp response.Response, err error) {
	return rsp, ding.Request(http.MethodPost, constant.DeleteDeptKey, nil, request.NewDeleteDept(deptId), &rsp)
}

// UpdateDept 更新部门
func (ding *dingTalk) UpdateDept(res *request.UpdateDept) (rsp response.Response, err error) {
	return rsp, ding.Request(http.MethodPost, constant.UpdateDeptKey, nil, res, &rsp)
}

// GetDeptDetail 获取部门详情
func (ding *dingTalk) GetDeptDetail(res *request.DeptDetail) (rsp response.DeptDetail, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetDeptDetailKey, nil, res, &rsp)
}

// GetDeptList 获取部门列表
func (ding *dingTalk) GetDeptList(res *request.DeptList) (rsp response.DeptList, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetDeptListKey, nil, res, &rsp)
}

// GetSubDeptList 获取子部门列表
func (ding *dingTalk) GetSubDeptList(deptId int) (rsp response.SubDeptList, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetSubDeptListKey, nil, request.NewSubDept(deptId), &rsp)
}

// GetDeptUserIds 获取部门用户userid列表
func (ding *dingTalk) GetDeptUserIds(du *request.DeptUserId) (req response.DeptUserId, err error) {
	return req, ding.Request(http.MethodPost, constant.GetDeptUserIdKey, nil, du, &req)
}

// GetDeptSimpleUserInfo 获取部门用户基础信息
func (ding *dingTalk) GetDeptSimpleUserInfo(res *request.DeptSimpleUserInfo) (req response.DeptSimpleUserInfo,
	err error,
) {
	return req, ding.Request(http.MethodPost, constant.GetDeptSimpleUserKey, nil, res, &req)
}

// GetDeptDetailUserInfo 获取部门用户详情
func (ding *dingTalk) GetDeptDetailUserInfo(res *request.DeptDetailUserInfo) (req response.DeptDetailUserInfo,
	err error,
) {
	return req, ding.Request(http.MethodPost, constant.GetDeptDetailUserKey, nil, res, &req)
}

// GetParentIdsByUserId 获取指定用户的所有父部门列表
func (ding *dingTalk) GetParentIdsByUserId(userId string) (req response.GetParentIdsByUserId, err error) {
	return req, ding.Request(http.MethodPost, constant.GetParentDeptsByUserKey, nil,
		request.NewGetParentIdsByUserId(userId), &req)
}

// GetParentIdsByDeptId 获取指定部门的所有父部门列表
func (ding *dingTalk) GetParentIdsByDeptId(deptId int) (req response.GetParentIdsByDeptId, err error) {
	return req, ding.Request(http.MethodPost, constant.GetParentDeptsByDeptKey, nil,
		request.NewGetParentIdsByDeptId(deptId), &req)
}
