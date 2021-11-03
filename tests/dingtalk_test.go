package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/api"
	"log"
	"os"
	"strings"
	"testing"
)

var dingTalk = api.NewDingTalk().
	SetId(1368002).
	SetKey("dingkjy4w80esdwgjuyo").
	SetSecret("bDKa_nfJg3zYRsFrj-wTohTuoJCtxTEHaGmybYF9vgaVAZJOz-mICsLGStB288nW").
	Build()
var robot = api.NewRobot("f00184626027aafcd6b4fe07b90ec11250d78b8735282185256a0c72186a5f49")

func TestDingTalkGetToken(t *testing.T) {

	token, err := dingTalk.GetAccessToken()
	assert.Nil(t, err)
	t.Log(token)
}

func TestCurrentPath(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("path:%s", strings.Replace(dir, "\\", "/", -1))
}
