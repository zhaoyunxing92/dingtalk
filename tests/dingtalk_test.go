package tests

import (
	"github.com/zhaoyunxing92/dingtalk"
	"log"
	"os"
	"strings"
	"testing"
)

var dingTalk = dingtalk.NewDingTalk("dingdv9vbx9mcrj18g1n", "bDBJd67ct1Ik0GFqhNNWH4Lbo4aqZGglaE1wJ3mnbG6ANRjOruuGzs6Z0glNEU63")

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
	t.Logf("path:%s",strings.Replace(dir, "\\", "/", -1))
}