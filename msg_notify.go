package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/domain"
	"github.com/zhaoyunxing92/dingtalk/global"
	"net/http"
)

//发送普通文本工作通知
//部门id或用户id重复会剔除
//https://ding-doc.dingtalk.com/document#/org-dev-guide/send-work-notifications
//发送工作通知
func (talk *DingTalk) SendWorkNotify(res domain.WorkNotifyRes) (resp domain.WorkNotifyResp, err error) {
	//组装部门、用户
	res.AssembleDept()
	res.AssembleUser()
	res.SetAgentId(talk.AgentId)
	//字段验证
	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return resp, err
	}

	err = talk.request(http.MethodPost, global.CorpConversationKey, nil, res, &resp)

	return resp, err
}
