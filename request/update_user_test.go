package request

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUpdateUser(t *testing.T) {
	build := NewUpdateUser("").SetForceUpdateFields("").Build()
	js, err := json.Marshal(build)

	assert.Nil(t, err)
	t.Log(string(js))
}
