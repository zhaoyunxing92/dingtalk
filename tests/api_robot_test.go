package tests

import (
	"encoding/json"
	"github.com/zhaoyunxing92/dingtalk/model"
	"testing"
)

//TestTextRobotMsg：机器消息
func TestTextRobotMsg(t *testing.T) {
	msg := model.NewTextRobotMsg("测试 robot\n换行")
	msg.AtMobiles=[]string{"18357154439"}
	res, err := robot.SendRobotMsg(msg)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}
