package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var client = NewClient(1244553273, "dingkjy4w80esdwgjuyo",
	"bDKa_nfJg3zYRsFrj-wTohTuoJCtxTEHaGmybYF9vgaVAZJOz-mICsLGStB288nW").
	Build()

//var client = NewClient(1354379668, "ding4xqt9h4tb5hmovlq",
//	"rNzB2gWYLG5ZjVtnSCxk19nmJo1FQytEG5UsvzizIGLJYkwdfITsnjBs31AKQJdM").
//	Build()

func TestDingTalk_GetAccessToken(t *testing.T) {

	token, err := client.GetAccessToken()
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Log(token)
}

func TestDingTalk_GetSuiteAccessToken(t *testing.T) {

	token, err := client.GetSuiteAccessToken()
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Log(token)
}

func TestDingTalk_GetCorpAccessToken(t *testing.T) {

	token, err := client.GetCorpAccessToken()
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Log(token)
}
