package tests

import (
	"encoding/json"
	"github.com/zhaoyunxing92/dingtalk/model"
	"testing"
)

//TestTextRobotMsg：机器消息
func TestTextRobotMsg(t *testing.T) {
	msg := model.NewTextRobotMsg("测试 robot\n换行")
	msg.AtMobiles = []string{"18357154439"}
	res, err := robot.SendRobotMsg(msg)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestMarkDownRobotMessage：机器消息
func TestMarkDownRobotMessage(t *testing.T) {
	md := `
>  这个是markdown，目前不知道要写啥，但是我明确的知道他支持渲染表格
# 这个是一级标题
## 二级标题测试图片

![钉钉开放平台](https://img.alicdn.com/tfs/TB1bB9QKpzqK1RjSZFoXXbfcXXa-576-96.png)

### 三级标题测试列表

* 吃饭
* 写bug

### 测试任务列表

- [x] 支持oa消息
- [x] 支持text消息
- [ ] 图片消息

### 下划线
钉钉为企业和组织提供了很多基础办公应用例如审批、日志、视频会议等。企业可基于钉钉开放平台的能力，根据实际需要定制开发企业应用。钉钉开放平台提供了丰富的接口能力，例如通讯录管理、群会话管理、消息通知、智能工作流、待办任务、考勤等以满足企业多样的业务需求，同时基于开放平台的授权机制，降低开发者的对接成本和安全风险。`
	//todo:markdown不支持at
	msg := model.NewMarkDownRobotMessage("markdown", md)
	msg.AtMobiles = []string{"18357154439"}
	res, err := robot.SendRobotMsg(msg)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestLinkMessage:link消息
func TestLinkMessage(t *testing.T) {
	msg := model.NewLinkMessage()
	msg.Title = "链接消息测试"
	msg.MediaId = "@lALPDe7sx7z5xEJgzQJA"
	msg.Url = "https://ding-doc.dingtalk.com/document#/org-dev-guide/message-types-and-data-format/title-48c-gj4-sbp"
	msg.Describe = "是以企业工作通知会话中某个微应用的名义推送到员工的通知消息，例如生日祝福、入职提醒等"

	res, err := robot.SendRobotMsg(msg)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

//TestRobotFileMessage:机器发送文件
func TestRobotFileMessage(t *testing.T) {
	//msg := model.NewFileMessage("@lALPDe7sx7z5xEJgzQJA")
	//
	//res, err := robot.SendRobotMsg(msg)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//js, err := json.Marshal(res)
	//t.Log(string(js))
}
