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

func TestNewUpdateChat(t *testing.T) {

	chat := NewUpdateChat("123ewer").
		SetName("name").
		SetOwner("name").
		SetIcon("name").
		SetSearchable(0).
		SetValidation(0).
		SetShowHistory(0).
		SetChatBannedType(0).
		SetManagementType(0).
		SetMentionAllAuthority(0).
		SetAddUsers("123456", "123456", "123").
		SetDelUsers("123456", "123456", "123").
		SetExtUsers("123456", "123456", "123").
		SetDelExtUsers("123456", "123456", "123").
		Build()

	err := validate(chat)
	assert.Nil(t, err)

	assert.Equal(t, len(chat.AddUsers), 2)
	assert.Equal(t, len(chat.DelUsers), 2)
	assert.Equal(t, len(chat.ExtUsers), 2)
	assert.Equal(t, len(chat.DelExtUsers), 2)
	t.Log(chat.String())
}
