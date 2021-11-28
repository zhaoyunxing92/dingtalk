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

type UpdateTodoDone struct {
	//当前访问资源所归属用户的unionId
	UnionId string `json:"unionId,omitempty" validate:"required"`

	//待办ID
	TaskId string `json:"taskId,omitempty" validate:"required"`

	//执行者状态列表，id需传用户的unionId
	TodoDone []*todoDone `json:"executorStatusList,omitempty"`
}

type todoDone struct {
	UnionId string `json:"id,omitempty"`
	Done    *bool  `json:"isDone"`
}

type updateTodoDoneBuilder struct {
	todo *UpdateTodoDone
}

func newTodoDone(unionId string, done bool) *todoDone {
	return &todoDone{unionId, &done}
}

func NewUpdateTodoDone(unionId, taskId string) *updateTodoDoneBuilder {
	return &updateTodoDoneBuilder{todo: &UpdateTodoDone{UnionId: unionId, TaskId: taskId}}
}

func (b *updateTodoDoneBuilder) SetTodoDone(unionId string, done bool) *updateTodoDoneBuilder {
	b.todo.TodoDone = append(b.todo.TodoDone, newTodoDone(unionId, done))
	return b
}

func (b *updateTodoDoneBuilder) Build() *UpdateTodoDone {
	return b.todo
}
