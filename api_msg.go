/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
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
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

import "github.com/pkg/errors"

// SendTemplateMessage 使用模板发送工作通知消息
func (ding *dingTalk) SendTemplateMessage(req *request.SendTemplateMessage) (rsp response.SendTemplateMessage,
	err error) {
	if !ding.isv() {
		return response.SendTemplateMessage{}, errors.New("只支持isv调用")
	}

	return rsp, ding.Request(http.MethodPost, constant.SendTemplateMessageKey, nil, req, &rsp)
}

// SendCorpConversationMessage 发送工作通知
func (ding *dingTalk) SendCorpConversationMessage(req *request.CorpConversationMessage) (rsp response.CorpConversationMessage,
	err error) {
	req.AgentId = ding.Id
	return rsp, ding.Request(http.MethodPost, constant.SendCorpConversationMessageKey, nil, req, &rsp)
}

// GetMessageProgress 获取工作通知消息的发送进度
func (ding *dingTalk) GetMessageProgress(agentId, taskId int) (rsp response.MessageProgress, err error) {

	return rsp, ding.Request(http.MethodPost, constant.MessageProgressKey, nil,
		request.NewMessageProgress(agentId, taskId), &rsp)
}
