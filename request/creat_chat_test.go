package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCreatChat(t *testing.T) {
	c := NewCreatChat("golang", "manager164",
		"1948546245774889", "16347147862675011", "1948546245774889").
		Build()

	assert.NotNil(t, c)
	assert.Equal(t, len(c.Users), 3)
}
