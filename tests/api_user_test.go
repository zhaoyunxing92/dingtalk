package tests

import (
	"encoding/json"
	"github.com/zhaoyunxing92/dingtalk/model"
	"testing"
)

//TestCreatUser:用户管理-创建用户
func TestCreatUser(t *testing.T) {
	//userId:011755000243774889
	deptIds := []int{1, 427805278}
	user, err := dingTalk.CreateUser("张三", "18513027676", deptIds)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", user)
}

//TestDeleteUser:删除用户
func TestDeleteUser(t *testing.T) {
	//userId:011755000243774889
	rsp, err := dingTalk.DeleteUser("011755000243774889")
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestDeleteUser:删除用户
func TestUpdateUser(t *testing.T) {

	user := new(model.User)
	user.Name = "李四"
	user.Id = "011755000243774889"
	user.Remark = "测试用户"
	user.Position = "技术支持"
	user.PositionInDepts = map[int]string{1: "技术支持"}
	user.JobNumber = "001"
	user.OrderInDepts = map[int]string{1: "1", 427805278: "1"}
	user.Department = []int{1, 427805278}
	user.Hide = true
	user.OrgEmail = "1@dingtalk.com"
	user.Email = "1@dingtalk.com"
	user.Lang = "en_US"

	rsp, err := dingTalk.UpdateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

// TestGetUser:获取用户
func TestGetUser(t *testing.T) {

	rsp, err := dingTalk.GetUser("011755000243774889", "zh_CN")
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestCreatUserV2:用户管理2-创建用户
func TestCreatUserV2(t *testing.T) {
	//011755000243774889
	user, err := dingTalk.CreateUserV2("张三", "18513027676", 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", user)
}
