package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCorpPermanentCode(t *testing.T) {
	pc := NewCorpPermanentCode("123456")
	assert.NotNil(t, pc)
}
