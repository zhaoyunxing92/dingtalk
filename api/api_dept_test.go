package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"testing"
)

func TestDingTalk_GetDeptSimpleUserInfo(t *testing.T) {

	info, err := client.GetDeptSimpleUserInfo(
		request.NewDeptSimpleUserInfo(1, 0, 1).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, info.DeptUsers)
}

func TestDingTalk_GetDeptUserIds(t *testing.T) {

	userId, err := client.GetDeptUserIds(request.NewDeptUserId(1))

	assert.Nil(t, err)
	assert.NotNil(t, userId.UserIds)
}

func TestDingTalk_GetDeptDetailUserInfo(t *testing.T) {

	res, err := client.GetDeptDetailUserInfo(
		request.NewDeptDetailUserInfo(1, 0, 10).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.DeptDetailUsers)
}

func TestDingTalk_CreateDept(t *testing.T) {

	res, err := client.CreateDept(
		request.NewCreateDept("golang", 1).
			SetHideDept(false).
			SetOrder(1).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
