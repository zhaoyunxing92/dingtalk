package request

import (
	"testing"
)

func TestNewCreateUser(t *testing.T) {

	str := NewCreateUser("zhangsan", "18357154439", 123, 123, 7890, 123456, 123456).
		SetDeptOrder(123, 1).
		SetDeptOrder(456, 2).
		Build().
		String()

	t.Log(str)
}
