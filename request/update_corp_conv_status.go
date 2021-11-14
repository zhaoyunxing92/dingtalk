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

package request

type UpdateCorpConvMsgStatus struct {
	//发送消息时使用的微应用的ID。
	AgentId int `json:"agent_id" validate:"required"`

	TaskId int `json:"task_id" validate:"required"`

	//状态栏值
	StatusValue string `json:"status_value,omitempty" validate:"required"`

	//状态栏背景色，推荐0xFF加六位颜色值
	StatusBgColor string `json:"status_bg,omitempty"`
}

type updateCorpConvMsgStatusBuilder struct {
	conv *UpdateCorpConvMsgStatus
}

func NewUpdateCorpConvMsgStatus(taskId int, value string) *updateCorpConvMsgStatusBuilder {
	return &updateCorpConvMsgStatusBuilder{conv: &UpdateCorpConvMsgStatus{TaskId: taskId, StatusValue: value}}
}

func (b *updateCorpConvMsgStatusBuilder) SetAgentId(agentId int) *updateCorpConvMsgStatusBuilder {
	b.conv.AgentId = agentId
	return b
}

func (b *updateCorpConvMsgStatusBuilder) SetStatusBgColor(color string) *updateCorpConvMsgStatusBuilder {
	b.conv.StatusBgColor = color
	return b
}

func (b *updateCorpConvMsgStatusBuilder) Build() *UpdateCorpConvMsgStatus {
	return b.conv
}
