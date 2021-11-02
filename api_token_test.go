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
		SetId(10435002).
		SetKey("suitegqebx814dt8immqw").
		SetCorpId("ding3b1e912394011559bc961a6cb783455b").
		SetTicket("iR4i4yOsNrSZCYLpfbrE6jigqU0hJuWVNmFQFuJ93IDQOir9bSQ3OP93gDUttfAVsljIXyXfBfVB0nr63dwOAT").
		SetSecret("AXjjwYhZ7Bwh1e8vlkg7pPQHUACwl8rSJWFma1taYMDLUjmIAtl9d9yAdTBg4K3m").
		Build()

	token, err := client.GetCorpAccessToken()
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Log(token)
}
