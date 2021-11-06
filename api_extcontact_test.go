package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDingTalk_GetExtContact(t *testing.T) {

	res, err := client.GetExtContact(0, 10)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetExtContactLabel(t *testing.T) {

	res, err := client.GetExtContactLabel(0, 10)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Results), 4)
}
