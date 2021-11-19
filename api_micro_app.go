/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

import (
	"errors"
	"fmt"
	"net/http"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/domain"
)

//获取应用列表
//https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-applications
func (ding *dingTalk) GetMicroAppList() (apps domain.MicroAppList, err error) {

	err = ding.Request(http.MethodPost, constant.MicroAppListKey, nil, nil, &apps)

	return apps, err
}

//根据id获取应用
func (ding *dingTalk) GetMicroAppByAgentId(agentId uint64) (app domain.MicroApp, err error) {
	var apps domain.MicroAppList
	if apps, err = ding.GetMicroAppList(); err != nil {
		return domain.MicroApp{}, err
	}

	for _, item := range apps.AppList {
		if item.AgentId == agentId {
			return item, nil
		}
	}

	return domain.MicroApp{}, errors.New(fmt.Sprintf("agentId:%d is not exist", agentId))
}

//获取应用可见范围
//https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-the-application-visible-range
func (ding *dingTalk) GetMicroAppVisibleScopes(agentId uint64) (scopes domain.MicroAppVisibleScopes, err error) {
	form := map[string]interface{}{
		"agentId": agentId,
	}
	err = ding.Request(http.MethodPost, constant.MicroAppVisibleScopesKey, nil, form, &scopes)
	return scopes, err
}
