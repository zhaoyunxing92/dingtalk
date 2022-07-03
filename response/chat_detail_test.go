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

package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChat_Detail(t *testing.T) {
	str := `{
        "errcode":0,
        "errmsg":"ok",
        "chat_info":{
                "owner":"manager4220",
                "mentionAllAuthority":0,
                "showHistoryType":0,
                "chatid":"chate39f540dxxxx",
                "conversationTag":2,
                "validationType":0,
                "useridlist":[
                        "user123",
                        "manager4220"
                ],
                "icon":"@lADPDfJ6Q7LLejDNAljNAlg",
                "name":"全员群",
                "searchable":0,
                "chatBannedType":0,
                "managementType":0
        }
}`
	chat := &GetChatInfo{}

	err := json.Unmarshal([]byte(str), chat)

	assert.Nil(t, err)
	assert.Equal(t, chat.Code, 0)
	assert.Equal(t, chat.Msg, "ok")
	assert.Equal(t, chat.ChatInfo.Owner, "manager4220")
	assert.Equal(t, chat.ChatInfo.MentionAllAuthority, 0)
	assert.Equal(t, chat.ChatInfo.ShowHistoryType, 0)
	assert.Equal(t, chat.ChatInfo.ChatId, "chate39f540dxxxx")
	assert.Equal(t, chat.ChatInfo.ConversationTag, 2)
	assert.Equal(t, chat.ChatInfo.ValidationType, 0)
	assert.Equal(t, chat.ChatInfo.UserIds[0], "user123")
	assert.Equal(t, chat.ChatInfo.UserIds[1], "manager4220")
	assert.Equal(t, chat.ChatInfo.Icon, "@lADPDfJ6Q7LLejDNAljNAlg")
	assert.Equal(t, chat.ChatInfo.Name, "全员群")
	assert.Equal(t, chat.ChatInfo.Searchable, 0)
	assert.Equal(t, chat.ChatInfo.ChatBannedType, 0)
	assert.Equal(t, chat.ChatInfo.ManagementType, 0)
}
