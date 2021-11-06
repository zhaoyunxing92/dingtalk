package dingtalk

import (
	"github.com/stretchr/testify/assert"
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

func TestDingTalk_GetRoleUserList(t *testing.T) {

	res, err := client.GetRoleList(0, 200)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.RoleGroups), 2)
}
