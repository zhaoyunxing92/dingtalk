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
	assert.NotNil(t, info)
}

func TestDingTalk_GetDeptUserIds(t *testing.T) {

	userId, err := client.GetDeptUserIds(request.NewDeptUserId(1))

	assert.Nil(t, err)
	assert.NotNil(t, userId.UserIds)
}
