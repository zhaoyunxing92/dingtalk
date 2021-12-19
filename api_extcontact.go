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

// CreateExtContact 添加外部联系人
func (ding *dingTalk) CreateExtContact(res *request.CreateExtContact) (rsp response.CreateExtContact, err error) {
	return rsp, ding.Request(http.MethodPost, constant.CreateExtContactKey, nil, res, &rsp)
}

// DeleteExtContact 删除外部联系人
func (ding *dingTalk) DeleteExtContact(userId string) (rsp response.Response, err error) {
	return rsp, ding.Request(http.MethodPost, constant.DeleteExtContactKey, nil,
		request.NewDeleteExtContact(userId), &rsp)
}

// UpdateExtContact 更新外部联系人
func (ding *dingTalk) UpdateExtContact(res *request.UpdateExtContact) (rsp response.CreateExtContact, err error) {
	return rsp, ding.Request(http.MethodPost, constant.UpdateExtContactKey, nil, res, &rsp)
}

// GetExtContact 获取外部联系人列表
func (ding *dingTalk) GetExtContact(offset, size int) (rsp response.GetExtContact, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetExtContactKey, nil,
		request.NewGetExtContact(offset, size), &rsp)
}

// GetExtContactLabel 获取外部联系人标签列表
func (ding *dingTalk) GetExtContactLabel(offset, size int) (rsp response.GetExtContactLabel, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetExtContactLabelKey, nil,
		request.NewGetExtContact(offset, size), &rsp)
}

// GetExtContactDetail 获取外部联系人详情
func (ding *dingTalk) GetExtContactDetail(userId string) (rsp response.GetExtContactDetail, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetExtContactDetailKey, nil,
		request.NewGetExtContactDetail(userId), &rsp)
}
