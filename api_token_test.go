package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	client := NewDingTalk().
		SetId(1368002).
		SetKey("dingkjy4w80esdwgjuyo").
		SetSecret("bDKa_nfJg3zYRsFrj-wTohTuoJCtxTEHaGmybYF9vgaVAZJOz-mICsLGStB288nW").
		Build()

	token, err := client.GetAccessToken()
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Log(token)
}

func TestGetSuiteAccessToken(t *testing.T) {
	client := NewDingTalk().
		SetId(1368002).
		SetKey("dingkjy4w80esdwgjuyo").
		SetSecret("bDKa_nfJg3zYRsFrj-wTohTuoJCtxTEHaGmybYF9vgaVAZJOz-mICsLGStB288nW").
		Build()

	token, err := client.GetSuiteAccessToken()
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Log(token)
}

func TestGetCorpAccessToken(t *testing.T) {
	client := NewDingTalk().
		SetId(1368002).
		SetKey("dingkjy4w80esdwgjuyo").
		SetSecret("bDKa_nfJg3zYRsFrj-wTohTuoJCtxTEHaGmybYF9vgaVAZJOz-mICsLGStB288nW").
		Build()

	token, err := client.GetCorpAccessToken()
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Log(token)
}
