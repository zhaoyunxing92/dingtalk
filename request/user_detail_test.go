package request

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUserDetail(t *testing.T) {

	build := NewUserDetail("").SetLanguage("ch").Build()

	js, err := json.Marshal(build)
	assert.Nil(t, err)
	t.Log(string(js))
}
