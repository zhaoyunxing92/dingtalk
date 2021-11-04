package request

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCreateDept(t *testing.T) {

	cd := NewCreateDept("golang", 0).
		SetHideDept(true).
		SetUserPermits("123", "456", "").
		SetOrder(0).
		SetDeptPermits(-1, 0, 345, 567, 345).
		Build()

	d, err := json.Marshal(cd)
	assert.Nil(t, err)
	fmt.Println(string(d))
}
