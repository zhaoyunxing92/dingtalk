package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDingTalk_GetIndustryDeptDetail(t *testing.T) {

	res, err := client.GetIndustryDeptDetail(561570467)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetIndustryDept(t *testing.T) {

	res, err := client.GetIndustryDept(1, 0, 50)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
