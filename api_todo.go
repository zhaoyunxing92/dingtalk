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
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

//CreateTodo 新增钉钉待办任务
func (ding *dingTalk) CreateTodo(res *request.CreateTodo) (req response.CreateTodo, err error) {

	query := url.Values{}
	query.Set("operatorId", res.CreatorId)
	return req, ding.Request(http.MethodPost, fmt.Sprintf(constant.CreateTodoKey, res.CreatorId), nil, res, &req)
}

//GetTodoDetail 获取钉钉待办任务详情
func (ding *dingTalk) GetTodoDetail(unionId, taskId string) (req response.GetTodoDetail, err error) {

	return req, ding.Request(http.MethodPost, fmt.Sprintf(constant.GetTodoDetailKey, unionId, taskId), nil,
		nil, &req)
}
