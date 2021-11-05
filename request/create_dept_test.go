package request

import (
	"testing"
)

func TestNewCreateDept(t *testing.T) {

	str := NewCreateDept("golang", 0).
		SetHideDept(true).
		SetOuterDept(true).
		SetUserPermits("123", "456", "").
		SetOrder(0).
		SetUserPermitsDeptIds(1, 0, 345, 567, 345).
		SetDeptPermits(1, 0, 345, 567, 345).
		Build().String()

	t.Log(str)
}

func TestNewEntryCreateDept(t *testing.T) {

	str := NewCreateDept("golang", 0).
		Build().String()

	t.Log(str)
}
