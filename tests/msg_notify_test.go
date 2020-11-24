package tests

import (
	"github.com/zhaoyunxing92/dingtalk/domain"
	"testing"
)

//发送工作通知
func TestSendWorkNotify(t *testing.T) {

	res := domain.NewTextWorkNotify("hello dingtalk 0")
	res.UserIds = []string{"manager164"}
	notify, err := dingTalk.SendWorkNotify(res)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(notify)
}
