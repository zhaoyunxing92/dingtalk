package request

import (
	"testing"
)
import (
	"github.com/stretchr/testify/assert"
)

func TestNewUpdateExtContact(t *testing.T) {

	uc := NewUpdateExtContact("123", "李四", "", 123, 456, 789, 123).
		SetTitle("资深开发").
		SetAddress("杭州").
		SetShareDept(123, 123, 678).
		SetRemark("技术人员").
		SetCompanyName("小番茄").
		SetShareUser("123", "123", "456").
		Build()

	assert.NotEmpty(t, uc)
	assert.Equal(t, len(uc.Contact.ShareUser), 2)
	assert.Equal(t, len(uc.Contact.Labels), 3)
	assert.Equal(t, len(uc.Contact.ShareDept), 2)

	t.Log(uc.String())
}
