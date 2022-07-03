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

type GetChatInfo struct {
	Response

	ChatInfo struct {
		ChatId string `json:"chatid"`

		// 群名称
		Name string `json:"name"` // 群名称，长度限制为1~20个字符。

		//状态
		//
		//1：正常
		//
		//2：已解散
		Status int `json:"status"`

		// 群头像的mediaID
		Icon string `json:"icon"`

		// 会话类型 2:企业群
		ConversationTag int `json:"conversationTag"`

		// 群主的userid，
		Owner string `json:"owner"`

		// 群成员列表，每次最多支持40人，群人数上限为1000
		UserIds []string `json:"useridlist"`

		// 新成员是否可查看100条历史消息：1:可看
		ShowHistoryType int `json:"showHistoryType"`

		// 群是否可以被搜索.1:可搜索
		Searchable int `json:"searchable"`

		// 入群是否需要验证.1:要验证
		ValidationType int `json:"validationType"`

		//@all 使用范围：.1:仅群主可用
		MentionAllAuthority int `json:"mentionAllAuthority"`

		// 群管理类型：.1:仅群主可管理
		ManagementType int `json:"managementType" `

		// 是否开启群禁言：.1:全员禁言
		ChatBannedType int `json:"chatBannedType" `
	} `json:"chat_info"`
}
