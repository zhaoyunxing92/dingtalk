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

package request

import (
	"encoding/json"
	"strings"
)

type SetUserRoleManageScope struct {
	//员工在企业中的userid
	UserId string `json:"userid" validate:"required"`

	//角色Id
	RoleId int `json:"role_id"`

	DeptIds string `json:"dept_ids,omitempty"`

	Ds []string `json:"-" validate:"max=50"`
}

func (s *SetUserRoleManageScope) String() string {
	str, _ := json.Marshal(s)
	return string(str)
}

type setUserRoleManageScopeBuilder struct {
	rs *SetUserRoleManageScope
}

func NewSetUserRoleManageScope(userId string, roleId int) *setUserRoleManageScopeBuilder {

	return &setUserRoleManageScopeBuilder{rs: &SetUserRoleManageScope{UserId: userId, RoleId: roleId}}
}

func (b *setUserRoleManageScopeBuilder) SetDeptIds(deptId int, deptIds ...int) *setUserRoleManageScopeBuilder {
	ids := []int{deptId}
	ids = append(ids, deptIds...)
	b.rs.Ds = removeIntDuplicatesToString(ids)
	return b
}

func (b *setUserRoleManageScopeBuilder) Build() *SetUserRoleManageScope {
	b.rs.DeptIds = strings.Join(b.rs.Ds, ",")
	return b.rs
}
