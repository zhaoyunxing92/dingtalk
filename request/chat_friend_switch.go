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

//ChatFriendSwitch 设置禁止群成员私聊
type ChatFriendSwitch struct {
	//群会话ID，可通过创建群会话接口获取
	ChatId string `json:"chatid" validate:"required"`

	//true：开启禁止开关
	//false：关闭禁止开关
	Prohibit bool `json:"is_prohibit" validate:"required"`
}

func NewChatFriendSwitch(chatId string, prohibit bool) *ChatFriendSwitch {
	return &ChatFriendSwitch{chatId, prohibit}
}
