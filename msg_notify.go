package dingtalk

import (
	"errors"
	"github.com/zhaoyunxing92/dingtalk/domain"
	"github.com/zhaoyunxing92/dingtalk/global"
	"net/http"
)

//发送普通文本工作通知
//部门id或用户id重复会剔除
//https://ding-doc.dingtalk.com/document#/org-dev-guide/send-work-notifications
func (talk *DingTalk) SendTextCorpConversation(res *domain.TextWorkNotifyRes) (resp domain.WorkNotifyResp, err error) {

	res.DeptIds = res.RemoveRepeatDept()
	res.UserIds = res.RemoveRepeatUser()

	if res.All == false && len(res.UserIds) == 0 && len(res.DeptIds) == 0 {
		return resp, errors.New("当字段All设置为false时必须指定UserIds或DeptIds其中一个参数的值")
	}

	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return resp, err
	}

	if res.All == false && len(res.UserIds) == 0 && len(res.DeptIds) == 0 {
		return resp, errors.New("当字段All设置为false时必须指定UserIds或DeptIds其中一个参数的值")
	}
	if res.All {
		res.UserIds = []string{}
		res.DeptIds = []uint64{}
	}

	form := map[string]interface{}{
		"agent_id":     res.AgentId,
		"userid_list":  res.All,
		"dept_id_list": res.All,
		"msg":          res.Msg,
	}

	err = talk.request(http.MethodPost, global.MicroAppListKey, nil, form, &resp)

	return resp, err
}
