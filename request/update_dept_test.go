package request

import "testing"

func TestNewEntryUpdateDept(t *testing.T) {
	str := NewUpdateDept(1).Build().String()
	t.Log(str)
}

func TestNewUpdateDept(t *testing.T) {
	str := NewUpdateDept(1).Build().String()
	t.Log(str)
}
