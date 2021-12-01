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

func TestNewCreatChat(t *testing.T) {
	c := NewCreatChat("golang", "manager164",
		"1948546245774889", "16347147862675011", "1948546245774889").
		SetSearchable(1).
		SetValidation(1).
		SetShowHistory(1).
		SetManagementType(1).
		SetChatBannedType(1).
		SetMentionAllAuthority(1).
		Build()

	assert.NotNil(t, c)
	assert.Equal(t, len(c.Users), 3)
}
