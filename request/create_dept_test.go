package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCreateDept(t *testing.T) {

	str := NewCreateDept("golang", 0).
		SetHideDept(true).
		SetOuterDept(true).
		SetUserPermits("123", "456", "").
		SetOrder(0).
		SetUserPermitsUsers("123", "123", "", "456").
		SetUserPermitsDeptIds(1, 0, 345, 567, 345).
		SetDeptPermits(1, 0, 345, 567, 345).
		Build()

	assert.NotNil(t, str)
	assert.Equal(t, len(str.UserPermitsUserIds), 2)
	assert.Equal(t, len(str.UserPermitsDeptIds), 3)
	assert.Equal(t, len(str.UserPermit), 2)
	assert.Equal(t, len(str.DeptPermit), 3)
	t.Log(str.String())
}
