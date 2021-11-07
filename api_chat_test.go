package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"testing"
)

func TestDingTalk_CreateChat(t *testing.T) {

	res, err := client.CreateChat(
		request.NewCreatChat("golang", "manager164",
			"manager164", "manager164", "manager164").
			Build())

	// chat79e713edce46ea72ea0dd60fb4f9d6f6
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetChatInfo(t *testing.T) {

	res, err := client.GetChatInfo("chat79e713edce46ea72ea0dd60fb4f9d6f6")

	// chat79e713edce46ea72ea0dd60fb4f9d6f6
	assert.Nil(t, err)
	assert.NotNil(t, res)
}