package request

import "testing"

func TestNewDeleteDept(t *testing.T) {
	str := NewDeleteDept(12).String()
	t.Log(str)
}
