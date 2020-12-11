package tests

import (
	"encoding/json"
	"testing"
)

//TestGetRoleList:获取角色列表
func TestGetRoleList(t *testing.T) {

	res, err := dingTalk.GetRoleList(0, 50)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestGetUserRoleList:获取角色列表
func TestGetUserRoleList(t *testing.T) {

	res, err := dingTalk.GetRoleUserList(1578723538, 0, 50)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestGetRoleGroup:获取角色组
func TestGetRoleGroup(t *testing.T) {

	res, err := dingTalk.GetRoleGroup(1299380989)

	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}
