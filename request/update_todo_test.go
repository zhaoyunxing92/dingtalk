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

	"github.com/stretchr/testify/assert"
)

func TestNewUpdateTodo(t *testing.T) {
	todo := NewUpdateTodo("qwweer", "eretrtr").
		SetSubject("subject").
		SetDesc("desc").
		SetDueTime(12345).
		SetDone(false).
		SetExecutors("123", "456").
		SetParticipants("123", "456").
		Build()

	assert.NotNil(t, todo)
	assert.Equal(t, todo.UnionId, "qwweer")
	assert.Equal(t, todo.TaskId, "eretrtr")
}
