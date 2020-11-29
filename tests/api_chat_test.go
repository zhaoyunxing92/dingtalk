package tests

import (
	"encoding/json"
	"github.com/zhaoyunxing92/dingtalk/model"
	"testing"
)

//TestCreateChat:创建群
func TestCreateChat(t *testing.T) {
	//chat44231f20c2e5c88cf67651e1b82bfb86
	//cidQm8R6iRg2CcePBHqlUU3xQ==
	userIds := []string{"manager164", "011857506107842061", "manager164"}
	rsp, err := dingTalk.CreateChat("golang群", "manager164", userIds)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestCreateDetailChat:创建群
func TestCreateDetailChat(t *testing.T) {
	//chatcec1d860b41d42691cc83aa47b5b095c
	chat := new(model.CreateChat)
	chat.Name = "测试"
	chat.Owner = "manager164"
	chat.ChatBannedType = 1
	chat.UserIds = []string{"manager164", "011857506107842061", "manager164"}
	chat.Validation = 1
	chat.MentionAllAuthority = 1
	chat.ShowHistory = 1
	chat.ManagementType = 1
	chat.Searchable = 1

	rsp, err := dingTalk.CreateDetailChat(chat)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestGetChatInfo:获取群信息
func TestGetChatInfo(t *testing.T) {
	//chat44231f20c2e5c88cf67651e1b82bfb86
	//cidQm8R6iRg2CcePBHqlUU3xQ==
	rsp, err := dingTalk.GetChatInfo("chat44231f20c2e5c88cf67651e1b82bfb86")
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestChatFriendSwitch:设置禁止群成员私聊
func TestChatFriendSwitch(t *testing.T) {
	//chat44231f20c2e5c88cf67651e1b82bfb86
	//cidQm8R6iRg2CcePBHqlUU3xQ==
	rsp, err := dingTalk.ChatFriendSwitch("chat44231f20c2e5c88cf67651e1b82bfb86", true)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestChatSubAdmin:设置群管理员
func TestChatSubAdmin(t *testing.T) {
	//chat44231f20c2e5c88cf67651e1b82bfb86
	//cidQm8R6iRg2CcePBHqlUU3xQ==
	rsp, err := dingTalk.ChatSubAdmin("chat44231f20c2e5c88cf67651e1b82bfb86", "manager164", 2)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}
