package dingtalk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/employee"
)

func TestDingTalk_GetHrmEmployee(t *testing.T) {
	user, err := client.GetHrmEmployee(1, 10, employee.Formal, employee.Unknown, employee.ProbationPeriod)

	assert.Nil(t, err)
	assert.NotNil(t, user)
}
