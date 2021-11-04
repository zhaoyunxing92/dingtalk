package tests

import (
	"encoding/json"
	"github.com/zhaoyunxing92/dingtalk/v2/model"
	"testing"
)

func TestCreateDepartment(t *testing.T) {
	//427772502
	rsp, err := dingTalk.CreateDept("golang", 1)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestCreateDetailDept:创建详细的部门
func TestCreateDetailDept(t *testing.T) {

	dept := new(model.Dept)
	dept.Name = "测试部门"
	dept.ParentId = 1 //根部门
	dept.SourceIdentifier = "123"
	dept.Hide = true
	dept.ShareBalance = true
	dept.ParentBalanceFirst = true
	dept.DeptPermits = []int{411415721, 411415721, 412362849}
	dept.OuterDept = true
	dept.OuterDeptOnlySelf = false
	dept.UserPermits = []string{"manager164", "manager164"}
	dept.DeptPermits = []int{411415721, 411415721, 412362849}
	dept.OuterPermitUsers = []string{"manager164", "manager164"}
	dept.OuterPermitDepts = []int{411415721, 411415721, 412362849}
	dept.Order = 1
	dept.Extension = `{"职能":"人事"}`

	rsp, err := dingTalk.CreateDetailDept(dept)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestGetDeptDetail：获取部门详情
func TestGetDeptDetail(t *testing.T) {

	rsp, err := dingTalk.GetDeptDetail(428051259, "")
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


//TestDeleteDept:删除部门
func TestDeleteDept(t *testing.T) {

	rsp, err := dingTalk.DeleteDept(428051259)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestGetSubDeptList:获取部门列表
func TestGetSubDeptList(t *testing.T) {

	rsp, err := dingTalk.GetSubDeptList(1, "", true)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestGetSubDeptIds:获取部门列表
func TestGetSubDeptIds(t *testing.T) {

	rsp, err := dingTalk.GetSubDeptIds(1)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestGetUserParentDeptIds:查询指定用户的所有上级父部门路径
func TestGetUserParentDeptIds(t *testing.T) {

	rsp, err := dingTalk.GetParentIdsByUserId("manager164")
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestGetUserParentDeptIds:查询指定用户的所有上级父部门路径
func TestGetParentIdsByDeptId(t *testing.T) {

	rsp, err := dingTalk.GetParentIdsByDeptId(427922598)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestBatchDelDept 批量删除部门
func TestBatchDelDept(t *testing.T) {
	rsp, _ := dingTalk.GetSubDeptIds(1)
	for _, id := range rsp.SubDeptIds {
		//需要保留的部门id
		if id != 492171596 {
			_, _ = dingTalk.DeleteDept(id)
		}
	}
}
