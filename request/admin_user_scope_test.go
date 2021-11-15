package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAdminUserScope(t *testing.T) {

	scope := NewAdminUserScope("123456")

	assert.NotNil(t, scope)
}

func TestAdminUserScope_String(t *testing.T) {
	scope := NewAdminUserScope("123456").String()

	assert.NotNil(t, scope)
}
