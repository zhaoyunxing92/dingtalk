package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"testing"
)

var client = NewDingTalk().
	SetId(1244553273).
	SetKey("dingkjy4w80esdwgjuyo").
	SetSecret("bDKa_nfJg3zYRsFrj-wTohTuoJCtxTEHaGmybYF9vgaVAZJOz-mICsLGStB288nW").
	Build()

func TestDingTalk_CreateUser(t *testing.T) {

	user, err := client.CreateUser(request.NewCreateUser("张三", "15669019215", 554656655).Build())
	assert.Nil(t, err)
	assert.NotNil(t, user.UserId)
	t.Log(user.UserId)
}

func TestDingTalk_UpdateUser(t *testing.T) {

	res, err := client.UpdateUser(
		request.NewUpdateUser("4649322615774889").
			SetName("李四").
			SetDept(1, 554656655).
			SetTelephone("").
			SetForceUpdateFields("title,userid,title,userid").
			SetDeptTitle(1, "普通员工").
			SetDeptTitle(554656655, "设计人员").
			SetJobNumber("").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res.Code)
	t.Log(res.Msg)
}
