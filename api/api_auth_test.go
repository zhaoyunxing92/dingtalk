package api

import "testing"

func TestDingTalk_GetAuthInfo(t *testing.T) {
	client := NewDingTalk().
		SetId(10435002).
		SetKey("suitegqebx814dt8immqw").
		SetCorpId("ding3b1e912394011559bc961a6cb783455b").
		SetTicket("5nvE5VXD6BP43b9IBv8SLc4FBx3lgRZL0wdPPrXf48mi12yjPyftZwSA6r1NEgyF86sbR3Ksbe32TAsOHuDbGn").
		SetSecret("AXjjwYhZ7Bwh1e8vlkg7pPQHUACwl8rSJWFma1taYMDLUjmIAtl9d9yAdTBg4K3m").
		Build()

	client.GetAuthInfo("ding3b1e912394011559bc961a6cb783455b")
}
