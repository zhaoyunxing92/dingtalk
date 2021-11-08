/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
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
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/model"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

//CreateRole 创建角色
func (ding *dingTalk) CreateRole(name string, groupId int) (apps response.CreateRole, err error) {

	return apps, ding.Request(http.MethodPost, constant.CreateRoleKey, nil,
		request.NewCreateRole(name, groupId), &apps)
}

//CreateRoleGroup 创建角色组
func (ding *dingTalk) CreateRoleGroup(name string) (apps response.CreateRoleGroup, err error) {

	return apps, ding.Request(http.MethodPost, constant.CreateRoleGroupKey, nil,
		request.NewCreateRoleGroup(name), &apps)
}

//UpdateRole 更新角色
func (ding *dingTalk) UpdateRole(id int, name string) (apps response.Response, err error) {

	return apps, ding.Request(http.MethodPost, constant.UpdateRoleKey, nil,
		request.NewUpdateRole(id, name), &apps)
}

//BatchAddUserRole  批量增加员工角色
func (ding *dingTalk) BatchAddUserRole(rs []int, us []string) (apps response.Response, err error) {

	return apps, ding.Request(http.MethodPost, constant.RoleBatchAddUserKey, nil,
		request.NewRoleAddUser(rs, us), &apps)
}

//GetRoleList 获取角色列表
func (ding *dingTalk) GetRoleList(offset, size int) (apps response.RoleList, err error) {

	return apps, ding.Request(http.MethodPost, constant.GetRoleListKey, nil,
		request.NewRoleList(offset, size), &apps)
}

//DeleteRole 删除角色
func (ding *dingTalk) DeleteRole(id int) (apps response.Response, err error) {

	return apps, ding.Request(http.MethodPost, constant.DeleteRoleKey, nil, request.NewDeleteRole(id), &apps)
}

//SetUserRoleManageScope 设定角色成员管理范围 todo:官方接口不通
func (ding *dingTalk) SetUserRoleManageScope(res *request.SetUserRoleManageScope) (apps response.Response, err error) {

	return apps, ding.Request(http.MethodPost, constant.RoleUpdateUserManageScopeKey, nil, res, &apps)
}

//GetGroupRoles 获取角色组列表
func (ding *dingTalk) GetGroupRoles(groupId int) (apps response.GroupRole, err error) {

	return apps, ding.Request(http.MethodPost, constant.GetRoleGroupKey, nil,
		request.NewGroupRole(groupId), &apps)
}

//GetRoleDetail 获取角色详情
func (ding *dingTalk) GetRoleDetail(roleId int) (apps response.RoleDetail, err error) {

	return apps, ding.Request(http.MethodPost, constant.GetRoleDetailKey, nil,
		request.NewRoleDetail(roleId), &apps)
}

//GetRoleUserList 获取指定角色的员工列表
func (ding *dingTalk) GetRoleUserList(roleId, offset, size int) (apps response.RoleUser, err error) {

	return apps, ding.Request(http.MethodPost, constant.GetRoleUserListKey, nil,
		request.NewRoleUser(roleId, offset, size), &apps)
}

//BatchRemoveUserRole 批量删除员工角色
func (ding *dingTalk) BatchRemoveUserRole(roleIds []int, userIds []string) (apps model.Response, err error) {
	return apps, ding.Request(http.MethodPost, constant.RoleBatchRemoveUserKey, nil,
		request.NewBatchRemoveUserRole(roleIds, userIds), &apps)
}
