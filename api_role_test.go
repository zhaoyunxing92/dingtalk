package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"testing"
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
		request.NewSetUserRoleManageScope("1948546245774889", 2309075248).
			SetDeptIds(560900478).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Success(), true)
}

func TestDingTalk_GetGroupRoles(t *testing.T) {

	res, err := client.GetGroupRoles(1299380989)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, len(res.Roles), 4)
}

func TestDingTalk_GetRoleDetail(t *testing.T) {

	res, err := client.GetRoleDetail(1299380992)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.RoleName, "负责人")
	assert.Equal(t, res.GroupId, 1299380989)
}

func TestDingTalk_GetRoleUserList(t *testing.T) {

	res, err := client.GetRoleUserList(1299380992, 0, 10)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}