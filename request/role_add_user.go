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

type RoleAddUser struct {
	Rs []string `validate:"required,max=20" json:"-"`

	//角色roleId列表 多个roleId用英文逗号（,）分隔，最多可传20个
	RoleIds string `validate:"required" json:"roleIds"`

	Us []string `validate:"required,max=20" json:"-"`

	//员工的userId 多个userId用英文逗号（,）分隔，最多可传20个
	UserIds string `validate:"required" json:"userIds"`
}

func (r *RoleAddUser) String() string {
	str, _ := json.Marshal(r)
	return string(str)
}

func NewRoleAddUser(roleIds []int, userIds []string) *RoleAddUser {
	rs := removeIntDuplicatesToString(roleIds)
	us := removeStringDuplicates(userIds)
	return &RoleAddUser{rs, strings.Join(rs, ","), us, strings.Join(us, ",")}
}
