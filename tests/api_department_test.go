package tests

import (
	"encoding/json"
	"fmt"
	"github.com/zhaoyunxing92/dingtalk/model"
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

	dept := model.Department{}
	dept.Name = "测试"
	dept.ParentId = 1 //根部门
	dept.Hide = true
	dept.DeptPermits = []int{411415721, 412362849}
	dept.OuterDept = true
	dept.OuterDeptOnlySelf = false
	dept.OuterPermitDepts = []int{411415721, 412362849}
	dept.Order = 10

	js, err := json.Marshal(dept)
	fmt.Println(string(js))
	rsp, err := dingTalk.CreateV2DetailDepartment(dept)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp)

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

	rsp, err := dingTalk.GetDepartmentDetail(411415721)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp)
}
