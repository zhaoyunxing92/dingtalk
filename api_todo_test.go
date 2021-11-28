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
	"testing"
	"time"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/request"
)

func TestDingTalk_CreateTodo(t *testing.T) {

	todo := request.NewCreateTodo("ABNiSWeAolg5OETyYT60wdQiEiE", "使用api接口创建代办").
		SetDesc("使用api接口创建代办").
		SetAppUrl("https://developers.dingtalk.com").
		SetPcUrl("https://developers.dingtalk.com").
		SetDueTime(int(time.Now().UnixMilli())).
		SetPriority(40).
		SetExecutors("ABNiSWeAolg5OETyYT60wdQiEiE").
		SetParticipants("ABNiSWeAolg5OETyYT60wdQiEiE").
		//SetDingNotify().
		Build()

	res, err := client.CreateTodo(todo)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.CreatorId, "ABNiSWeAolg5OETyYT60wdQiEiE")
}

func TestDingTalk_GetTodoDetail(t *testing.T) {

	res, err := client.GetTodoDetail("ABNiSWeAolg5OETyYT60wdQiEiE", "task5c7c849d9127b64f360870d5ae086a7c")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
