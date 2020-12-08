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
