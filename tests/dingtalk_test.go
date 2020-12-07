package tests

import (
	"github.com/zhaoyunxing92/dingtalk"
	"log"
	"os"
	"strings"
	"testing"
)

var dingTalk = dingtalk.NewDingTalk(1020345059, "dinggqaj3nmlgosailas", "qzNjaNQCVhA73jLKe2OwT1z4jNATcAz07KJ6pczeRw4Xn1XfP93brx5DiD_MMsQ7")
var robot = dingtalk.NewRobot("f00184626027aafcd6b4fe07b90ec11250d78b8735282185256a0c72186a5f49")

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
