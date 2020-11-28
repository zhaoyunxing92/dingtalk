package tests

import (
	"encoding/json"
	"testing"
)

func TestCreateDepartment(t *testing.T) {
	//427285169
	rsp, err := dingTalk.CreateV2Department("golang", 1)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestCreateDetailDepartment:创建详细的部门
func TestCreateDetailDepartment(t *testing.T) {

	//dept := new(model.Department)
	//dept.Name = "测试"
	//dept.ParentId = 1 //根部门
	//dept.Hide = true
	//dept.DeptPermits = []int{411415721, 412362849}
	//dept.OuterDept = true
	//dept.OuterDeptOnlySelf = false
	////dept.OuterPermitDepts = []int{411415721, 412362849}
	//dept.Order = 10
	//
	//js, err := json.Marshal(dept)
	//fmt.Println(string(js))
	//rsp, err := dingTalk.CreateV2DetailDepartment(dept)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Logf("%v", rsp)

}

//TestDeleteDepartment:删除部门
func TestDeleteDepartment(t *testing.T) {

	rsp, err := dingTalk.DeleteV2Department(427497463)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp)
}

//TestGetDepartmentDetail：获取部门详情
func TestGetDepartmentDetail(t *testing.T) {

	rsp, err := dingTalk.GetDepartmentDetail(427805278, "")
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestGetDepartmentUserIds：获取部门用户userid列表
func TestGetDepartmentUserIds(t *testing.T) {

	rsp, err := dingTalk.GetDeptUserIds(1)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//GetDeptUserDetail：获取部门用户userid列表
func TestGetDeptUserDetail(t *testing.T) {

	rsp, err := dingTalk.GetDeptUserDetail(1, 0, 3, "")
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}
