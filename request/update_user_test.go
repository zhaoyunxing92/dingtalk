package request

import (
	"testing"
)

func TestNewUpdateUser(t *testing.T) {

	str := NewUpdateUser("134567").
		SetForceUpdateFields("").
		SetHideMobile(false).
		SetLoginId("").
		SetForceUpdateFields("").
		Build().
		String()

	t.Log(str)
}
