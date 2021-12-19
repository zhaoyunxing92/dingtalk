package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/employee"
)

func TestNewGetHrmEmployee(t *testing.T) {
	e := NewGetHrmEmployee(1, 0,
		[]employee.Status{employee.Formal, employee.Formal, employee.ToBeResigned, employee.Unknown})

	assert.NotNil(t, e)
	assert.Equal(t, e.Offset, 1)
	assert.Equal(t, e.Size, 0)
	assert.Equal(t, e.Status, "3,5,-1")
}
