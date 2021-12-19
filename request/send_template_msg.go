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

	"github.com/pkg/errors"
)

type SendTemplateMessage struct {
	// 应用的agentId
	AgentId int `json:"agent_id" validate:"required"`

	// 消息模板ID
	TemplateId string `json:"template_id" validate:"required"`

	UserIdList string `json:"userid_list,omitempty"`

	// 接收者的用户userid列表。最大列表长度为5000
	UserIds []string `json:"-" validate:"omitempty,max=5000"`

	DeptIdList string `json:"dept_id_list,omitempty"`

	// 接收者的部门id列表。最大列表长度为500
	DeptIds []int `json:"-" validate:"omitempty,max=500"`

	// 消息模板动态参数赋值数据
	Data map[string]string `json:"data"`
}

type sendTemplateMessageBuilder struct {
	msg *SendTemplateMessage
}

func NewSendTemplateMessage(agentId int, templateId string) *sendTemplateMessageBuilder {
	return &sendTemplateMessageBuilder{msg: &SendTemplateMessage{AgentId: agentId, TemplateId: templateId}}
}

func (sb *sendTemplateMessageBuilder) SetUserIds(userId string, userIds ...string) *sendTemplateMessageBuilder {
	sb.msg.UserIds = append(userIds, userId)
	return sb
}

func (sb *sendTemplateMessageBuilder) SetDeptIds(deptId int, deptIds ...int) *sendTemplateMessageBuilder {
	sb.msg.DeptIds = append(deptIds, deptId)
	return sb
}

func (sb *sendTemplateMessageBuilder) SetData(key, value string) *sendTemplateMessageBuilder {
	data := sb.msg.Data
	if data == nil {
		data = make(map[string]string, 5)
	}
	data[key] = value
	sb.msg.Data = data
	return sb
}

func (sb *sendTemplateMessageBuilder) SetMessage(msg map[string]string) *sendTemplateMessageBuilder {
	data := sb.msg.Data
	if data != nil {
		for k, v := range msg {
			data[k] = v
		}
	} else {
		data = msg
	}
	sb.msg.Data = data
	return sb
}

func (sb *sendTemplateMessageBuilder) Build() *SendTemplateMessage {
	msg := sb.msg
	deptIds := msg.DeptIds
	userIds := msg.UserIds
	if len(deptIds) <= 0 && len(userIds) <= 0 {
		panic(errors.New("DeptIds和UserIds不能同时为空"))
	}
	deptIds = removeIntDuplicates(deptIds)
	userIds = removeStringDuplicates(userIds)

	msg.DeptIds = deptIds
	msg.DeptIdList = strings.Join(removeIntDuplicatesToString(deptIds), ",")

	msg.UserIds = userIds
	msg.UserIdList = strings.Join(userIds, ",")
	return msg
}
