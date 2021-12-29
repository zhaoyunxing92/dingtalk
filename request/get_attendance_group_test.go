package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGetAttendanceGroup(t *testing.T) {
	group := NewGetAttendanceGroup(0, 10)

	assert.NotNil(t, group)
	assert.Equal(t, group.Offset, 0)
	assert.Equal(t, group.Size, 10)
}
