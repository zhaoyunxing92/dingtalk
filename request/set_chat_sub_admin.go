/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package request

type ChatSetSubAdmin struct {
	//群会话ID，可通过创建群会话接口获取
	ChatId string `json:"chatid" validate:"required"`

	//群成员userid，可通过根据手机号获取userid接口获取userid。
	UserId string `json:"userids" validate:"required"`

	//2：添加为管理员
	//3：删除该管理员
	Role int `json:"role" validate:"required,min=2,max=3"`
}

func NewChatSetSubAdmin(chatId, userId string, role int) *ChatSetSubAdmin {

	return &ChatSetSubAdmin{chatId, userId, role}
}
