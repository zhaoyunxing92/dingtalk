package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExtContact(t *testing.T) {

	ec := NewExtContact().
		SetTitle("资深开发").
		SetAddress("杭州").
		SetLabels(1234, 1234, 5678).
		SetShareDept(123, 123, 678).
		SetRemark("技术人员").
		SetStateCode("86").
		SetCompanyName("小番茄").
		SetShareUser("123", "123", "456").
		SetMobile("123456789").
		Build()

	assert.NotEmpty(t, ec)
	assert.Equal(t, len(ec.Contact.ShareUser), 2)
	assert.Equal(t, len(ec.Contact.Labels), 2)
	assert.Equal(t, len(ec.Contact.ShareDept), 2)

	t.Log(ec.String())
}
