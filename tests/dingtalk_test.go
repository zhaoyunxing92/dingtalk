package tests

import (
	"github.com/zhaoyunxing92/dingtalk"
	"log"
	"os"
	"strings"
	"testing"
)

var dingTalk = dingtalk.NewDingTalk(1020345059, "dinggqaj3nmlgosailas", "qzNjaNQCVhA73jLKe2OwT1z4jNATcAz07KJ6pczeRw4Xn1XfP93brx5DiD_MMsQ7")

func TestDingTalkGetToken(t *testing.T) {

	token, err := dingTalk.GetToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestCurrentPath(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("path:%s", strings.Replace(dir, "\\", "/", -1))
}
