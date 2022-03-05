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

package event

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChatDisband(t *testing.T) {
	str := `{
    "Operator": "manager3060",
    "OpenConversationId": "cid740SMooADFxmWX3DvIyp4g==",
    "CorpId": "ding0761931a826dfde2ffe93478753d9884",
    "EventType": "chat_disband",
    "OperatorUnionId": "3bOBVxFv0J8VOu3J4jGhZQiEiE",
    "TimeStamp": 1640754514696,
    "ChatId": "chat6418e0c8c141114e18491ea2603c09be"
}`

	chat := &ChatDisband{}

	err := json.Unmarshal([]byte(str), chat)

	assert.Nil(t, err)
	assert.Equal(t, chat.Operator, "manager3060")
	assert.Equal(t, chat.EventType, "chat_disband")
	assert.Equal(t, chat.TimeStamp, 1640754514696)
	assert.Equal(t, chat.OperatorUnionId, "3bOBVxFv0J8VOu3J4jGhZQiEiE")
	assert.Equal(t, chat.CorpId, "ding0761931a826dfde2ffe93478753d9884")
	assert.Equal(t, chat.ChatId, "chat6418e0c8c141114e18491ea2603c09be")
	assert.Equal(t, chat.OpenConversationId, "cid740SMooADFxmWX3DvIyp4g==")
}
