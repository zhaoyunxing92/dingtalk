package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUpdateChat(t *testing.T) {

	chat := NewUpdateChat("123ewer").
		SetAddUsers("123456", "123456", "123").
		SetDelUsers("123456", "123456", "123").
		SetExtUsers("123456", "123456", "123").
		SetDelExtUsers("123456", "123456", "123").
		SetShowHistory(0).
		Build()

	err := validate(chat)
	assert.Nil(t, err)

	assert.Equal(t, len(chat.AddUsers), 2)
	assert.Equal(t, len(chat.DelUsers), 2)
	assert.Equal(t, len(chat.ExtUsers), 2)
	assert.Equal(t, len(chat.DelExtUsers), 2)
}
