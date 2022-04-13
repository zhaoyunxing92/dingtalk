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
	"net/http"

	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

// RegisterEvent 注册事件
func (ding *DingTalk) RegisterEvent(req *request.RegisterEvent) (res response.Response, err error) {

	return res, ding.Request(http.MethodPost, constant.RegisterCallBackKey, nil, req, &res)
}

// GetRegisterFailedEvent 获取推送失败的事件列表
func (ding *DingTalk) GetRegisterFailedEvent() (res response.FailedRegisterEvent, err error) {

	return res, ding.Request(http.MethodGet, constant.GetCallBackFailedDetailKey, nil, nil, &res)
}

// UpdateRegisterEvent 更新事件接口
func (ding *DingTalk) UpdateRegisterEvent(req *request.RegisterEvent) (res response.Response, err error) {

	return res, ding.Request(http.MethodPost, constant.UpdateCallBackKey, nil, req, &res)
}

// GetRegisterEvent 查询订阅事件
func (ding *DingTalk) GetRegisterEvent() (res response.RegisterEvent, err error) {

	return res, ding.Request(http.MethodGet, constant.GetCallBackKey, nil, nil, &res)
}

// DeleteRegisterEvent 删除事件回调接口
func (ding *DingTalk) DeleteRegisterEvent() (res response.Response, err error) {

	return res, ding.Request(http.MethodGet, constant.DeleteCallBackKey, nil, nil, &res)
}
