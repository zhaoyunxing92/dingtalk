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
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/request"
)

func TestDingTalk_CreateRoleGroup(t *testing.T) {
	res, err := client.CreateRoleGroup("dubbo")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_CreateRole(t *testing.T) {
	res, err := client.CreateRole("java", 2308825440)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_UpdateRole(t *testing.T) {
	res, err := client.UpdateRole(2309075248, "go")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetRoleList(t *testing.T) {
	res, err := client.GetRoleList(0, 200)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.RoleGroups), 2)
}

func TestDingTalk_SetUserRoleManageScope(t *testing.T) {
	res, err := client.SetUserRoleManageScope(
		request.NewSetUserRoleManageScope("manager164", 2309075248).
			SetDeptIds(560900478).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Ok(), true)
}

func TestDingTalk_GetGroupRoles(t *testing.T) {
	res, err := client.GetGroupRoles(1299380989)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, len(res.Roles), 4)
}

func TestDingTalk_GetRoleDetail(t *testing.T) {
	res, err := client.GetRoleDetail(1299380990)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.RoleName, "主管理员")
	assert.Equal(t, res.GroupId, 1299380989)
}

func TestDingTalk_GetRoleUserList(t *testing.T) {
	res, err := client.GetRoleUserList(1299380990, 0, 10)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
