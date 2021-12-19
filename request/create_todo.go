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
	"encoding/json"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant/priority"
)

// CreateTodo 新增钉钉待办任务
type CreateTodo struct {
	// 业务系统侧的唯一标识ID，即业务ID。
	SourceId string `json:"sourceId,omitempty"`

	// 待办标题
	Subject string `json:"subject" validate:"required"`

	// 创建者的unionId
	CreatorId string `json:"creatorId" validate:"required"`

	// 待办备注描述
	Desc string `json:"description,omitempty"`

	// 截止时间，Unix时间戳，单位毫秒。
	DueTime *int `json:"dueTime,omitempty"`

	// 执行者的unionId
	Executors []string `json:"executorIds,omitempty"`

	// 参与者的unionId
	Participants []string `json:"participantIds,omitempty"`

	// 详情页url跳转地址。
	Urls *detailUrl `json:"detailUrl,omitempty"`

	// 生成的待办是否仅展示在执行者的待办列表中
	OnlyShowExecutor *bool `json:"isOnlyShowExecutor,omitempty"`

	// 优先级，取值：
	// 10：较低
	// 20：普通
	// 30：紧急
	// 40：非常紧急
	Priority *int `json:"priority,omitempty"`

	// 待办通知配置
	NotifyConfigs *notifyConfigs `json:"notifyConfigs,omitempty"`
}

type detailUrl struct {
	// APP端详情页url跳转地址
	App string `json:"appUrl,omitempty"`

	// PC端详情页url跳转地址
	Pc string `json:"pcUrl,omitempty"`
}

// notifyConfigs 待办通知配置
type notifyConfigs struct {
	Ding string `json:"dingNotify,omitempty"`
}

type createTodoBuilder struct {
	todo *CreateTodo
}

func (todo *CreateTodo) String() string {
	str, _ := json.Marshal(todo)
	return string(str)
}

func NewCreateTodo(unionId, subject string) *createTodoBuilder {
	return &createTodoBuilder{todo: &CreateTodo{Subject: subject, CreatorId: unionId}}
}

func (c *createTodoBuilder) SetSourceId(sourceId string) *createTodoBuilder {
	c.todo.SourceId = sourceId
	return c
}

func (c *createTodoBuilder) SetDesc(desc string) *createTodoBuilder {
	c.todo.Desc = desc
	return c
}

// SetDueTime 截止时间，Unix时间戳，单位毫秒。
func (c *createTodoBuilder) SetDueTime(time int) *createTodoBuilder {
	c.todo.DueTime = &time
	return c
}

func (c *createTodoBuilder) SetExecutors(unionId string, unionIds ...string) *createTodoBuilder {
	c.todo.Executors = append(unionIds, unionId)
	return c
}

func (c *createTodoBuilder) SetParticipants(unionId string, unionIds ...string) *createTodoBuilder {
	c.todo.Participants = append(unionIds, unionId)
	return c
}

func (c *createTodoBuilder) SetAppUrl(url string) *createTodoBuilder {
	if c.todo.Urls == nil {
		c.todo.Urls = &detailUrl{}
	}
	c.todo.Urls.App = url
	return c
}

func (c *createTodoBuilder) SetPcUrl(url string) *createTodoBuilder {
	if c.todo.Urls == nil {
		c.todo.Urls = &detailUrl{}
	}
	c.todo.Urls.Pc = url
	return c
}

// SetPriority 优先级，取值：
// 10：较低
// 20：普通
// 30：紧急
// 40：非常紧急
func (c *createTodoBuilder) SetPriority(level priority.Level) *createTodoBuilder {
	l := new(priority.Level)
	*l = level
	c.todo.Priority = (*int)(l)
	return c
}

// SetDingNotify DING通知配置，目前仅支持取值为1，表示应用内DING。
func (c *createTodoBuilder) SetDingNotify() *createTodoBuilder {
	c.todo.NotifyConfigs = &notifyConfigs{Ding: "1"}
	return c
}

func (c *createTodoBuilder) Build() *CreateTodo {
	todo := c.todo
	todo.Executors = removeStringDuplicates(todo.Executors)
	todo.Participants = removeStringDuplicates(todo.Participants)
	return todo
}
