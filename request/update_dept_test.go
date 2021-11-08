package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUpdateDept(t *testing.T) {

	str := NewUpdateDept(0).
		SetHideDept(true).
		SetOuterDept(true).
		SetUserPermits("123", "456", "").
		SetOrder(0).
		SetUserPermitsUserIds("123", "123", "", "456").
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

func TestNewEntryUserPermitsUpdateDept(t *testing.T) {

	str := NewUpdateDept(0).
		SetHideDept(true).
		SetOuterDept(true).
		//SetUserPermits("123", "456", "").
		SetOrder(0).
		SetUserPermitsUserIds("123", "123", "", "456").
		SetUserPermitsDeptIds(1, 0, 345, 567, 345).
		SetDeptPermits(1, 0, 345, 567, 345).
		Build()

	assert.NotNil(t, str)
	assert.Equal(t, len(str.UserPermitsUserIds), 2)
	assert.Equal(t, len(str.UserPermitsDeptIds), 3)
	//assert.Equal(t, len(str.UserPermit), 2)
	assert.Equal(t, len(str.DeptPermit), 3)
	t.Log(str.String())
}

func TestNewEntryUserPermitsUserIdsUpdateDept(t *testing.T) {

	str := NewUpdateDept(0).
		SetHideDept(true).
		SetOuterDept(true).
		//SetUserPermits("123", "456", "").
		SetOrder(0).
		//SetUserPermitsUserIds("123", "123", "", "456").
		SetUserPermitsDeptIds(1, 0, 345, 567, 345).
		SetDeptPermits(1, 0, 345, 567, 345).
		Build()

	assert.NotNil(t, str)
	//assert.Equal(t, len(str.UserPermitsUserIds), 2)
	assert.Equal(t, len(str.UserPermitsDeptIds), 3)
	//assert.Equal(t, len(str.UserPermit), 2)
	assert.Equal(t, len(str.DeptPermit), 3)
	t.Log(str.String())

}

func TestNewHideDeptUpdateDept(t *testing.T) {

	str := NewUpdateDept(0).
		SetHideDept(false).
		SetDeptPermits(1, 1, 2).
		SetOuterDeptOnlySelf(true).
		SetOrder(1).
		Build()

	assert.NotNil(t, str)
	assert.Equal(t, len(str.DeptPermit), 2)
	assert.Nil(t, str.OuterDeptOnlySelf)
	t.Log(str.String())

}
