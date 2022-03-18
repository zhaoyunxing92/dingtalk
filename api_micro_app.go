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
	"net/url"

	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/domain"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

// GetMicroAppList 获取应用列表
func (ding *DingTalk) GetMicroAppList() (rsp response.MicroAppList, err error) {
	return rsp, ding.Request(http.MethodPost, constant.MicroAppListKey, nil, nil, &rsp)
}

// GetMicroAppByAgentId 根据id获取应用
func (ding *DingTalk) GetMicroAppByAgentId(agentId int) (domain.MicroApp, error) {
	var (
		app  domain.MicroApp
		apps response.MicroAppList
		err  error
	)
	if apps, err = ding.GetMicroAppList(); err != nil {
		return domain.MicroApp{}, err
	}
	for _, app = range apps.Apps {
		if app.AgentId == agentId {
			return app, nil
		}
	}
	return domain.MicroApp{}, errors.New(fmt.Sprintf("agentId:%d is not exist", agentId))
}

// GetMicroAppVisibleScopes 获取应用可见范围
func (ding *DingTalk) GetMicroAppVisibleScopes(agentId int) (scopes response.MicroAppVisibleScopes, err error) {
	return scopes, ding.Request(http.MethodPost, constant.MicroAppVisibleScopesKey, nil,
		request.NewMicroAppVisibleScopes(agentId), &scopes)
}

//GetUserMicroAppVisibleScopes 获取员工可见的应用列表
func (ding *DingTalk) GetUserMicroAppVisibleScopes(userId string) (scopes response.MicroAppList, err error) {
	query := url.Values{}
	query.Set("userid", userId)
	return scopes, ding.Request(http.MethodGet, constant.UserMicroAppVisibleScopesKey, query, nil, &scopes)
}
