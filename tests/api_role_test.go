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

	res, err := dingTalk.GetRoleGroup(1765793487)

	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestGetRoleDetail:获取角色详情
func TestGetRoleDetail(t *testing.T) {

	res, err := dingTalk.GetRoleDetail(1765793487)

	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestCreateRoleGroup:创建角色组
func TestCreateRoleGroup(t *testing.T) {

	res, err := dingTalk.CreateRoleGroup("制造")

	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestCreateRole:创建角色
func TestCreateRole(t *testing.T) {

	res, err := dingTalk.CreateRole("服装制造", 1765944351)

	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestCreateRole:更新角色
func TestUpdateRole(t *testing.T) {

	res, err := dingTalk.UpdateRole("服装制造", 1765883402)

	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestRoleBatchAddUser:批量增加员工角色
func TestRoleBatchAddUser(t *testing.T) {
	roleIds := []string{"1765883402"}
	userIds := []string{"manager164", "035267695923139924"}
	res, err := dingTalk.RoleBatchAddUser(roleIds, userIds)

	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestRoleBatchRemoveUser:批量增加员工角色
func TestRoleBatchRemoveUser(t *testing.T) {
	roleIds := []string{"1765883402"}
	userIds := []string{"manager164", "035267695923139924"}
	res, err := dingTalk.RoleBatchRemoveUser(roleIds, userIds)

	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestRoleUpdateUserManageScope:设定角色成员管理范围
func TestRoleUpdateUserManageScope(t *testing.T) {
	deptIds := []int{411415721}

	res, err := dingTalk.RoleUpdateUserManageScope("manager164", 1765883402, deptIds)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}
