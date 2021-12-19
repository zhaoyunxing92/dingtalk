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
	"fmt"
	"net/http"
	"net/url"

	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

// CreateTodo 新增钉钉待办任务
func (ding *dingTalk) CreateTodo(res *request.CreateTodo) (req response.CreateTodo, err error) {
	query := url.Values{}
	query.Set("operatorId", res.CreatorId)
	return req, ding.Request(http.MethodPost, fmt.Sprintf(constant.CreateTodoKey, res.CreatorId), nil, res, &req)
}

// GetTodoDetail 获取钉钉待办任务详情
func (ding *dingTalk) GetTodoDetail(unionId, taskId string) (req response.GetTodoDetail, err error) {
	return req, ding.Request(http.MethodPost, fmt.Sprintf(constant.GetTodoDetailKey, unionId, taskId), nil,
		nil, &req)
}

// DeleteTodo 删除钉钉待办任务
func (ding *dingTalk) DeleteTodo(unionId, taskId string) (req response.DeleteTodo, err error) {
	query := url.Values{}
	query.Set("operatorId", unionId)

	return req, ding.Request(http.MethodDelete, fmt.Sprintf(constant.DeleteTodoKey, unionId, taskId), query,
		nil, &req)
}

// UpdateTodo 更新钉钉待办任务
func (ding *dingTalk) UpdateTodo(res *request.UpdateTodo) (req response.UpdateTodo, err error) {
	return req, ding.Request(http.MethodPut, fmt.Sprintf(constant.UpdateTodoKey, res.UnionId, res.TaskId),
		nil, res, &req)
}

// UpdateTodoDone 更新钉钉待办执行者状态
func (ding *dingTalk) UpdateTodoDone(res *request.UpdateTodoDone) (req response.UpdateTodo, err error) {
	return req, ding.Request(http.MethodPut, fmt.Sprintf(constant.UpdateTodoDoneKey, res.UnionId, res.TaskId),
		nil, res, &req)
}

// GetTodoListBySourceId 根据sourceId获取钉钉待办任务详情
func (ding *dingTalk) GetTodoListBySourceId(unionId, sourceId string) (req response.GetTodoDetail, err error) {
	return req, ding.Request(http.MethodGet, fmt.Sprintf(constant.GetTodoListBySourceIdKey, unionId, sourceId),
		nil, nil, &req)
}

// GetTodoList 查询企业下用户待办列表
func (ding *dingTalk) GetTodoList(unionId, token string, done bool) (req response.GetTodoList, err error) {
	return req, ding.Request(http.MethodPost, fmt.Sprintf(constant.GetTodoListKey, unionId), nil,
		request.NewGetTodoList(token, done), &req)
}
