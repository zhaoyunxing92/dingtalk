package request

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInactiveUser(t *testing.T) {
	iu := NewInactiveUser(false, 0, 10, "20211104").
		SetDeptIds(123, 123, 456, 789, 123).
		Build()

	assert.Equal(t, len(iu.DeptIds), 3)
	js, err := json.Marshal(iu)
	assert.Nil(t, err)
	t.Log(string(js))

}
