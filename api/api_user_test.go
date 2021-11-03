package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/request"
	"testing"
)

func TestDingTalk_CreateUser(t *testing.T) {
	client := NewDingTalk().
		SetId(1244553273).
		SetKey("dingkjy4w80esdwgjuyo").
		SetSecret("bDKa_nfJg3zYRsFrj-wTohTuoJCtxTEHaGmybYF9vgaVAZJOz-mICsLGStB288nW").
		Build()

	user, err := client.CreateUser(request.NewCreateUser("张三", "15669019215", 554656655).Build())
	assert.Nil(t, err)
	assert.NotNil(t, user.UserId)
	t.Log(user.UserId)
}
