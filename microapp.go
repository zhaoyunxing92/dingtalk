package dingtalk

import (
	"errors"
	"fmt"
	"github.com/zhaoyunxing92/dingtalk/domain"
	"github.com/zhaoyunxing92/dingtalk/global"
	"net/http"
)

//获取应用列表
//https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-applications
func (talk *DingTalk) GetMicroAppList() (apps domain.MicroAppList, err error) {

	err = talk.request(http.MethodPost, global.MicroAppListKey, nil, nil, &apps)

	return apps, err
}

//根据id获取应用
func (talk *DingTalk) GetMicroAppByAgentId(agentId uint64) (app domain.MicroApp, err error) {
	var apps domain.MicroAppList
	if apps, err = talk.GetMicroAppList(); err != nil {
		return domain.MicroApp{}, err
	}

	for _, item := range apps.AppList {
		if item.AgentId == agentId {
			return item, nil
		}
	}

	return domain.MicroApp{},errors.New(fmt.Sprintf("agentId:%d is not exist",agentId))
}

//获取应用可见范围
//https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-the-application-visible-range
func (talk *DingTalk) GetMicroAppVisibleScopes(agentId uint64) (scopes domain.MicroAppVisibleScopes, err error) {
	form := map[string]interface{}{
		"agentId": agentId,
	}
	err = talk.request(http.MethodPost, global.MicroAppVisibleScopesKey, nil, form, &scopes)
	return scopes, err
}
