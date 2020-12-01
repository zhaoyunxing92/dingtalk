package dingtalk

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/lihongchen/dingtalk/global"
	"github.com/lihongchen/dingtalk/model"
)

//获取应用列表
//https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-applications
func (talk *DingTalk) GetMicroAppList() (apps model.MicroAppList, err error) {

	err = talk.request(http.MethodPost, global.MicroAppListKey, nil, nil, &apps)

	return apps, err
}

//根据id获取应用
func (talk *DingTalk) GetMicroAppByAgentId(agentId uint64) (app model.MicroApp, err error) {
	var apps model.MicroAppList
	if apps, err = talk.GetMicroAppList(); err != nil {
		return model.MicroApp{}, err
	}

	for _, item := range apps.AppList {
		if item.AgentId == agentId {
			return item, nil
		}
	}

	return model.MicroApp{}, errors.New(fmt.Sprintf("agentId:%d is not exist", agentId))
}

//获取应用可见范围
//https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-the-application-visible-range
func (talk *DingTalk) GetMicroAppVisibleScopes(agentId uint64) (scopes model.MicroAppVisibleScopes, err error) {
	form := map[string]interface{}{
		"agentId": agentId,
	}
	err = talk.request(http.MethodPost, global.MicroAppVisibleScopesKey, nil, form, &scopes)
	return scopes, err
}
