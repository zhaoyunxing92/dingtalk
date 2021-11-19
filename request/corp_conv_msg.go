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

package request

import (
	"strings"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/domain/message"
)

// CorpConvMessage 工作通知
type CorpConvMessage struct {
	//发送消息时使用的微应用的ID。
	AgentId int `json:"agent_id" validate:"required"`

	//接收者的企业内部用户的userId列表，最大用户列表长度100。
	UserIds []string `json:"-" validate:"lte=100"`

	UserIdList string `json:"userid_list,omitempty"`

	//接收者的部门id列表，最大列表长度20。接收者是部门Id下包括子部门下的所有用户。
	DeptIdList string `json:"dept_id_list,omitempty"`

	DeptIds []int `json:"-" validate:"lte=20"`

	//是否发送给企业全部用户。当设置为false时必须指定userid_list或dept_id_list其中一个参数的值。
	All *bool `json:"to_all_user,omitempty"`

	Msg message.Message `json:"msg,omitempty" validate:"required"`
}

type corpConvMessageBuilder struct {
	cm *CorpConvMessage
}

func NewCorpConvMessage(msg message.Message) *corpConvMessageBuilder {
	return &corpConvMessageBuilder{cm: &CorpConvMessage{Msg: msg}}
}

func (sb *corpConvMessageBuilder) SetAgentId(agentId int) *corpConvMessageBuilder {
	sb.cm.AgentId = agentId
	return sb
}

func (sb *corpConvMessageBuilder) SetUserIds(userId string, userIds ...string) *corpConvMessageBuilder {
	sb.cm.UserIds = append(userIds, userId)
	return sb
}

func (sb *corpConvMessageBuilder) SetDeptIds(deptId int, deptIds ...int) *corpConvMessageBuilder {
	sb.cm.DeptIds = append(deptIds, deptId)
	return sb
}

func (sb *corpConvMessageBuilder) SetAllUser(all bool) *corpConvMessageBuilder {
	sb.cm.All = &all
	return sb
}

func (sb *corpConvMessageBuilder) Build() *CorpConvMessage {
	cm := sb.cm
	cm.DeptIds = removeIntDuplicates(cm.DeptIds)
	cm.DeptIdList = strings.Join(removeIntDuplicatesToString(cm.DeptIds), ",")

	cm.UserIds = removeStringDuplicates(cm.UserIds)
	cm.UserIdList = strings.Join(cm.UserIds, ",")
	return cm
}
