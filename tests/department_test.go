package tests

import "testing"

func TestCreateDepartment(t *testing.T) {
	//427285169
	rsp, err := dingTalk.CreateDepartment("golang", 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp)
}

//TestDeleteDepartment:删除部门
func TestDeleteDepartment(t *testing.T) {

	rsp, err := dingTalk.DeleteDepartment(427285169)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp)
}
