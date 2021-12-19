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

type UpdateTodo struct {
	// 当前访问资源所归属用户的unionId
	UnionId string `json:"unionId,omitempty" validate:"required"`

	// 待办ID
	TaskId string `json:"taskId,omitempty" validate:"required"`

	// 待办标题
	Subject string `json:"subject,omitempty"`

	// 待办描述
	Desc string `json:"description,omitempty"`

	// 截止时间，Unix时间戳，单位毫秒。
	DueTime *int `json:"dueTime,omitempty"`

	// 完成状态
	Done *bool `json:"done,omitempty"`

	// 执行者的unionId
	Executors []string `json:"executorIds,omitempty"`

	// 参与者的unionId
	Participants []string `json:"participantIds,omitempty"`
}

type updateTodoBuilder struct {
	todo *UpdateTodo
}

func NewUpdateTodo(unionId, taskId string) *updateTodoBuilder {
	return &updateTodoBuilder{todo: &UpdateTodo{UnionId: unionId, TaskId: taskId}}
}

func (b *updateTodoBuilder) SetSubject(subject string) *updateTodoBuilder {
	b.todo.Subject = subject
	return b
}

func (b *updateTodoBuilder) SetDesc(desc string) *updateTodoBuilder {
	b.todo.Desc = desc
	return b
}

func (b *updateTodoBuilder) SetDueTime(time int) *updateTodoBuilder {
	b.todo.DueTime = &time
	return b
}

func (b *updateTodoBuilder) SetDone(done bool) *updateTodoBuilder {
	b.todo.Done = &done
	return b
}

func (b *updateTodoBuilder) SetExecutors(unionId string, unionIds ...string) *updateTodoBuilder {
	b.todo.Executors = append(unionIds, unionId)
	return b
}

func (b *updateTodoBuilder) SetParticipants(unionId string, unionIds ...string) *updateTodoBuilder {
	b.todo.Participants = append(unionIds, unionId)
	return b
}

func (b *updateTodoBuilder) Build() *UpdateTodo {
	todo := b.todo
	todo.Executors = removeStringDuplicates(todo.Executors)
	todo.Participants = removeStringDuplicates(todo.Participants)
	return todo
}
