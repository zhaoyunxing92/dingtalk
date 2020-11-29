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
	rsp, err := dingTalk.ChatSubAdmin("chat44231f20c2e5c88cf67651e1b82bfb86", "123", 2)
	js, err := json.Marshal(rsp)
	t.Log(string(js))
	if err != nil {
		t.Fatal(err)
	}
}

//TestSendTextMsgToChat:发送消息到群
func TestSendTextMsgToChat(t *testing.T) {
	//chat44231f20c2e5c88cf67651e1b82bfb86

	msg := model.NewTextMessages("hello")

	//cidQm8R6iRg2CcePBHqlUU3xQ==
	rsp, err := dingTalk.SendMsgToChat("chat44231f20c2e5c88cf67651e1b82bfb86", msg)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestSendOAMsgToChat:发送oa消息
func TestSendOAMsgToChat(t *testing.T) {
	//chat44231f20c2e5c88cf67651e1b82bfb86

	oa := model.OA{}

	f := model.Form{}
	f.Key = "姓名:"
	f.Value = "赵云兴"

	f2 := model.Form{}
	f2.Key = "年龄:"
	f2.Value = "20"

	f3 := model.Form{}
	f3.Key = "身高:"
	f3.Value = "192"

	f4 := model.Form{}
	f4.Key = "体重:"
	f4.Value = "80KG"

	f5 := model.Form{}
	f5.Key = "爱好:"
	f5.Value = "go、java、docker、vue"

	f6 := model.Form{}
	f6.Key = "测试:"
	f6.Value = "go、java、docker、vue"

	//设置体最多只有6个
	oa.Body.Forms = append(oa.Body.Forms, f, f2, f3, f4, f5, f6)
	oa.Body.Content = "validator用于对数据进行校验。在 Web 开发中，对用户传过来的数据我们都需要进行严格校验，防止用户的恶意请求。例如日期格式，用户年龄，性别等必须是正常的值，不能随意设置。"
	oa.Body.Title = "头部标题"
	oa.Body.Author = "赵云兴"
	oa.Body.ImageId = "@lALPDe7sx7z5xEJgzQJA"

	//设置头
	oa.Header.BgColor = "FFBBBBBB"
	oa.Header.Text = "被替换为应用名称"

	oa.MessageUrl = "https://ding-doc.dingtalk.com/document#/org-dev-guide/message-types-and-data-format"
	oa.PcMessageUrl = "https://ding-doc.dingtalk.com"

	res := new(model.WorkNotify)
	res.NewOAWorkNotify(oa)
	res.UserIds = []string{"manager164"}

	msg := model.NewOaMessage(oa)
	//cidQm8R6iRg2CcePBHqlUU3xQ==
	rsp, err := dingTalk.SendMsgToChat("chat44231f20c2e5c88cf67651e1b82bfb86", msg)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestSendImageMsgToChat:发送消息到群
func TestSendImageMsgToChat(t *testing.T) {
	//msgRnwvnhx7sqGgL5LEebu0cg==

	msg := model.NewImageMessages("@lALPDe7sx7z5xEJgzQJA")

	//cidQm8R6iRg2CcePBHqlUU3xQ==
	rsp, err := dingTalk.SendMsgToChat("chat44231f20c2e5c88cf67651e1b82bfb86", msg)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestGetChatMsgReadUser:查询群消息已读人员列表
func TestGetChatMsgReadUser(t *testing.T) {
	//msgRnwvnhx7sqGgL5LEebu0cg==
	rsp, err := dingTalk.GetChatMsgReadUser("msgRnwvnhx7sqGgL5LEebu0cg==", 1, 100)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}
