package tests

import (
	"encoding/json"
	"github.com/zhaoyunxing92/dingtalk/model"
	"testing"
)

//TestSendImageToConversation:发送消息
func TestSendImageToConversation(t *testing.T) {
	//msgRnwvnhx7sqGgL5LEebu0cg==
	msg := model.NewImageMessages("@lALPDe7sx7z5xEJgzQJA")
	//cidQm8R6iRg2CcePBHqlUU3xQ==
	rsp, err := dingTalk.SendToConversation("manager164", "chat44231f20c2e5c88cf67651e1b82bfb86", msg)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}
