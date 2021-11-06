package request

import "testing"

func TestNewSetUserRoleManageScope(t *testing.T) {

	str := NewSetUserRoleManageScope("1948546245774889", 2309075248).
		//SetDeptIds(560900478).
		Build().String()

	t.Log(str)
}
