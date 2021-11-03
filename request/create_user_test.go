package request

import (
	"fmt"
	"testing"
)

func TestNewCreateUser(t *testing.T) {

	user := NewCreateUser("zhangsan", "18357154439", 123, 5678, 7890, 123456, 123456).
		SetDeptOrder(123, 1).
		SetDeptOrder(456, 2).
		Build()

	fmt.Println(user)
}
