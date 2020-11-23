package tests

import (
	"github.com/zhaoyunxing92/dingtalk/domain"
	"testing"
)

//发送工作通知
func TestWorkNotify(t *testing.T) {
	res:=new(domain.TextWorkNotifyRes)
	res.AgentId=970282753
	res.Msg=domain.NewTextMessages("你好")
	res.UserIds=[]string{"manager164"}
	_, err := dingTalk.SendTextCorpConversation(res)
	if err != nil {
		t.Fatal(err)
	}
}
