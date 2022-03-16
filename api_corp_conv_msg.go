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

	"github.com/pkg/errors"

	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

// SendTemplateMessage 使用模板发送工作通知消息
func (ding *DingTalk) SendTemplateMessage(req *request.SendTemplateMessage) (rsp response.SendTemplateMessage,
	err error) {
	if !ding.isv() {
		return response.SendTemplateMessage{}, errors.New("只支持isv调用")
	}
	return rsp, ding.Request(http.MethodPost, constant.SendTemplateMessageKey, nil, req, &rsp)
}

// SendCorpConvMessage 发送工作通知
func (ding *DingTalk) SendCorpConvMessage(req *request.CorpConvMessage) (rsp response.CorpConvMessage,
	err error) {
	return rsp, ding.Request(http.MethodPost, constant.SendCorpConversationMessageKey, nil, req, &rsp)
}

// UpdateCorpConvMessageStatus 更新工作通知状态栏
func (ding *DingTalk) UpdateCorpConvMessageStatus(req *request.UpdateCorpConvMsgStatus) (rsp response.Response,
	err error) {
	return rsp, ding.Request(http.MethodPost, constant.UpdateCorpConvMessageStatusKey, nil, req, &rsp)
}

// GetCorpConvMsgProgress 获取工作通知消息的发送进度
func (ding *DingTalk) GetCorpConvMsgProgress(agentId, taskId int) (rsp response.MessageProgress, err error) {
	return rsp, ding.Request(http.MethodPost, constant.MessageProgressKey, nil,
		request.NewMessageProgress(agentId, taskId), &rsp)
}

// GetMessageSendResult 获取工作通知消息的发送结果
func (ding *DingTalk) GetMessageSendResult(agentId, taskId int) (rsp response.MessageSendResult, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetMessageSendResultKey, nil,
		request.NewMessageProgress(agentId, taskId), &rsp)
}

// RecallCorpConvMessage 撤回工作通知消息
func (ding *DingTalk) RecallCorpConvMessage(agentId, taskId int) (rsp response.Response, err error) {
	return rsp, ding.Request(http.MethodPost, constant.RecallCorpConvMessageKey, nil,
		request.NewRecallCorpConvMessage(agentId, taskId), &rsp)
}
