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

type RecallCorpConvMessage struct {
	//发送消息时使用的微应用的ID
	AgentId int `json:"agent_id" validate:"required"`

	//发送消息时钉钉返回的任务ID。
	TaskId int `json:"msg_task_id" validate:"required"`
}

type recallCorpConvMessageBuilder struct {
	mp *RecallCorpConvMessage
}

func NewRecallCorpConvMessage(taskId int) *recallCorpConvMessageBuilder {
	return &recallCorpConvMessageBuilder{mp: &RecallCorpConvMessage{TaskId: taskId}}
}

func (b *recallCorpConvMessageBuilder) SetAgentId(agentId int) *recallCorpConvMessageBuilder {
	b.mp.AgentId = agentId
	return b
}

func (b *recallCorpConvMessageBuilder) Build() *RecallCorpConvMessage {
	return b.mp
}
