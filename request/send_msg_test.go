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
	"github.com/zhaoyunxing92/dingtalk/v2/domain/message"
)

func TestNewSendMessage(t *testing.T) {
	msg := NewSendMessage("sender", "chat", message.NewTextMessage("hello dingtalk"))

	assert.NotNil(t, msg)
	assert.Equal(t, msg.ChatId, "chat")
	assert.Equal(t, msg.Sender, "sender")
	assert.Equal(t, msg.Msg.MessageType(), "text")
}
