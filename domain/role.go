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

package domain

// role:角色
type role struct {
	Id   int    `json:"id"`   // 角色id
	Name string `json:"name"` // 角色名称
}

// openRole:角色
type openRole struct {
	Id   int    `json:"role_id"`   // 角色id
	Name string `json:"role_name"` // 角色名称
}

// roleDetail:角色详情
type roleDetail struct {
	GroupId int    `json:"groupId"` // 角色id
	Name    string `json:"name"`    // 角色名称
}

// roleGroup:角色组
type roleGroup struct {
	GroupId int    `json:"groupId"` // 角色id
	Name    string `json:"name"`    // 角色组名称
	Roles   []role `json:"roles"`   // 角色列表
}

// roleGroup:角色组
type openRoleGroup struct {
	Name  string     `json:"group_name"` // 角色组名称
	Roles []openRole `json:"roles"`      // 角色列表
}

// roleListResult:角色列表返回
type roleListResult struct {
	HasMore bool        `json:"hasMore"` // 是否还有更多数据
	List    []roleGroup `json:"list"`    // 角色组列表。
}

type roleUserListResult struct {
	HasMore bool   `json:"hasMore"` // 是否还有更多数据
	List    []User `json:"list"`
}

// RoleListResponse:获取角色列表返回
type RoleListResponse struct {
	Response
	RoleGroup roleListResult `json:"result"`
}

// RoleUserListResponse:角色用户列表
type RoleUserListResponse struct {
	Response
	Result roleUserListResult `json:"result"`
}

// RoleGroupResponse:获取角色组
type RoleGroupResponse struct {
	Response
	Result openRoleGroup `json:"role_group"`
}

// RoleDetailResponse:获取角色详情
type RoleDetailResponse struct {
	Response
	Role roleDetail `json:"role"`
}

// CreateRoleGroupResponse:创建角色组
type CreateRoleGroupResponse struct {
	Response
	GroupId int `json:"groupId"` // 角色组id
}

// CreateRoleResponse:创建角色
type CreateRoleResponse struct {
	Response
	RoleId int `json:"roleId"` // 角色id
}
