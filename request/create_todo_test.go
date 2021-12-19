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
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant/priority"
)

func TestCreateTodoBuilder_Build(t *testing.T) {
	todo := NewCreateTodo("ABNiSWeAolg5OETyYT60wdQiEiE", "测试代办").
		SetDesc("使用api接口创建代办").
		SetAppUrl("https://developers.dingtalk.com").
		SetPcUrl("https://baidu.com").
		SetDingNotify().
		SetDueTime(12345).
		SetExecutors("123", "456").
		SetParticipants("123", "456").
		SetSourceId("123").
		SetPriority(priority.Lower).
		Build()

	assert.Equal(t, todo.Urls.Pc, "https://baidu.com")
	assert.Equal(t, todo.Urls.App, "https://developers.dingtalk.com")

	t.Log(todo.String())
}
