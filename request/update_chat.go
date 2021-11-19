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
	"encoding/json"
)

type UpdateChat struct {
	ChatId string `json:"chatid" validate:"required"`

	//群名称，长度限制为1~20个字符
	Name string `json:"name,omitempty" validate:"omitempty,min=1,max=20"`

	//群主的userid,该员工必须为会话useridlist的成员之一。
	Owner string `json:"owner,omitempty"`

	//群主类型：
	//
	//emp：企业员工
	//
	//ext：外部联系人
	OwnerType string `json:"ownerType,omitempty" validate:"omitempty,oneof=emp ext"`

	//群成员列表，每次最多支持40人，群人数上限为1000
	AddUsers []string `json:"add_useridlist,omitempty" validate:"omitempty,max=40"`

	//删除的成员列表
	DelUsers []string `json:"del_useridlist,omitempty"`

	//添加的外部联系人成员列表。
	ExtUsers []string `json:"add_extidlist,omitempty"`

	//删除的外部联系人成员列表。
	DelExtUsers []string `json:"del_extidlist,omitempty"`

	//群头像的mediaId
	Icon string `json:"icon,omitempty"`

	//新成员是否可查看100条历史消息：
	//
	//1：可查看
	//
	//0：不可查看
	//
	//如果不传值，代表不可查看。
	ShowHistory *int `json:"showHistoryType,omitempty" validate:"omitempty,max=1,min=0"`

	//群是否可以被搜索：
	//
	//0（默认）：不可搜索
	//
	//1：可搜索
	Searchable *int `json:"searchable,omitempty" validate:"omitempty,max=1,min=0"`

	//入群是否需要验证
	//
	//0（默认）：不验证
	//
	//1：入群验证
	Validation *int `json:"validationType,omitempty" validate:"omitempty,max=1,min=0"`

	//@all 使用范围：
	//
	//0（默认）：所有人可使用
	//
	//1：仅群主可@all
	MentionAllAuthority *int `json:"mentionAllAuthority,omitempty" validate:"omitempty,max=1,min=0"`

	//群管理类型：
	//
	//0（默认）：所有人可管理
	//1：仅群主可管理
	ManagementType *int `json:"managementType,omitempty" validate:"omitempty,max=1,min=0"`

	//是否开启群禁言：
	//0（默认）：不禁言
	//1：全员禁言
	ChatBannedType *int `json:"chatBannedType,omitempty" validate:"omitempty,max=1,min=0"`
}

func (uc *UpdateChat) String() string {
	str, _ := json.Marshal(uc)
	return string(str)
}

type updateChatBuilder struct {
	uc *UpdateChat
}

func NewUpdateChat(chatId string) *updateChatBuilder {

	return &updateChatBuilder{uc: &UpdateChat{ChatId: chatId}}
}

func (cb *updateChatBuilder) SetName(name string) *updateChatBuilder {
	cb.uc.Name = name
	return cb
}

func (cb *updateChatBuilder) SetOwner(userId string) *updateChatBuilder {
	cb.uc.Owner = userId
	return cb
}

func (cb *updateChatBuilder) SetAddUsers(userId string, userIds ...string) *updateChatBuilder {
	cb.uc.AddUsers = append(userIds, userId)
	return cb
}

func (cb *updateChatBuilder) SetDelUsers(userId string, userIds ...string) *updateChatBuilder {
	cb.uc.DelUsers = append(userIds, userId)
	return cb
}

func (cb *updateChatBuilder) SetExtUsers(userId string, userIds ...string) *updateChatBuilder {
	cb.uc.ExtUsers = append(userIds, userId)
	return cb
}

func (cb *updateChatBuilder) SetDelExtUsers(userId string, userIds ...string) *updateChatBuilder {
	cb.uc.DelExtUsers = append(userIds, userId)
	return cb
}

func (cb *updateChatBuilder) SetIcon(mediaId string) *updateChatBuilder {
	cb.uc.Icon = mediaId
	return cb
}

//SetShowHistory 新成员是否可查看100条历史消息：
//1：可查看
//0：不可查看
//
//如果不传值，代表不可查看。
func (cb *updateChatBuilder) SetShowHistory(history int) *updateChatBuilder {
	cb.uc.ShowHistory = &history
	return cb
}

//SetSearchable 群是否可以被搜索：
//
//0（默认）：不可搜索
//
//1：可搜索
func (cb *updateChatBuilder) SetSearchable(search int) *updateChatBuilder {
	cb.uc.Searchable = &search
	return cb
}

//SetValidation 入群是否需要验证
//
//0（默认）：不验证
//
//1：入群验证
func (cb *updateChatBuilder) SetValidation(valida int) *updateChatBuilder {
	cb.uc.Validation = &valida
	return cb
}

//SetMentionAllAuthority @all 使用范围：
//
//0（默认）：所有人可使用
//
//1：仅群主可@all
func (cb *updateChatBuilder) SetMentionAllAuthority(mentionAllAuthority int) *updateChatBuilder {
	cb.uc.MentionAllAuthority = &mentionAllAuthority
	return cb
}

//SetManagementType 群管理类型：
//
//0（默认）：所有人可管理
//1：仅群主可管理
func (cb *updateChatBuilder) SetManagementType(managementType int) *updateChatBuilder {
	cb.uc.ManagementType = &managementType
	return cb
}

//SetChatBannedType 是否开启群禁言：
//0（默认）：不禁言
//1：全员禁言
func (cb *updateChatBuilder) SetChatBannedType(chatBannedType int) *updateChatBuilder {
	cb.uc.ChatBannedType = &chatBannedType
	return cb
}

func (cb *updateChatBuilder) Build() *UpdateChat {
	chat := cb.uc
	chat.AddUsers = removeStringDuplicates(chat.AddUsers)
	chat.DelUsers = removeStringDuplicates(chat.DelUsers)
	chat.ExtUsers = removeStringDuplicates(chat.ExtUsers)
	chat.DelExtUsers = removeStringDuplicates(chat.DelExtUsers)
	return chat
}
