package model

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type CreateChat struct {
	Name                string   `json:"name,omitempty" validate:"required,max=20,min=1"`              //群名称，长度限制为1~20个字符。
	Owner               string   `json:"owner,omitempty" validate:"required"`                          //群主的userid，
	UserIds             []string `json:"useridlist,omitempty" validate:"required,max=40"`              //群成员列表，每次最多支持40人，群人数上限为1000
	ShowHistory         int      `json:"showHistoryType,omitempty" validate:"omitempty,oneof=0 1"`     //新成员是否可查看100条历史消息：1:可看
	Searchable          int      `json:"searchable,omitempty" validate:"omitempty,oneof=0 1"`          //群是否可以被搜索.1:可搜索
	Validation          int      `json:"validationType,omitempty" validate:"omitempty,oneof=0 1"`      //入群是否需要验证.1:要验证
	MentionAllAuthority int      `json:"mentionAllAuthority,omitempty" validate:"omitempty,oneof=0 1"` //@all 使用范围：.1:仅群主可用
	ManagementType      int      `json:"managementType,omitempty" validate:"omitempty,oneof=0 1"`      //群管理类型：.1:仅群主可管理
	ChatBannedType      int      `json:"chatBannedType,omitempty" validate:"omitempty,oneof=0 1"`      //是否开启群禁言：.1:全员禁言
}

//UpdateChat:更新群
type UpdateChat struct {
	Id                  string   `json:"id" validate:"required"`                                       //群会话Id
	Name                string   `json:"name,omitempty" validate:"omitempty,max=20,min=1"`             //群名称，长度限制为1~20个字符。
	Owner               string   `json:"owner,omitempty"`                                              //群主的userid，
	OwnerType           string   `json:"ownerType,omitempty" validate:"omitempty,oneof=emp ext"`       //群主类型：emp：企业员工,ext：外部联系人
	AddUserIds          []string `json:"add_useridlist,omitempty" validate:"omitempty,max=40"`         //添加的群成员列表，每次最多支持40人，群人数上限为1000
	DelUserIds          []string `json:"del_useridlist,omitempty" validate:"omitempty,max=40"`         //删除的成员列表
	AddExtIds           []string `json:"add_extidlist,omitempty" validate:"omitempty,max=40"`          //添加的外部联系人成员列表
	DelExtIds           []string `json:"del_extidlist,omitempty" validate:"omitempty,max=40"`          //删除的外部联系人成员列表
	Icon                string   `json:"icon,omitempty"`                                               //群头像的mediaId
	Searchable          int      `json:"searchable,omitempty" validate:"omitempty,oneof=0 1"`          //群是否可以被搜索.1:可搜索
	ShowHistory         int      `json:"showHistoryType,omitempty" validate:"omitempty,oneof=0 1"`     //新成员是否可查看100条历史消息：1:可看
	Validation          int      `json:"validationType,omitempty" validate:"omitempty,oneof=0 1"`      //入群是否需要验证.1:要验证
	MentionAllAuthority int      `json:"mentionAllAuthority,omitempty" validate:"omitempty,oneof=0 1"` //@all 使用范围：.1:仅群主可用
	ManagementType      int      `json:"managementType,omitempty" validate:"omitempty,oneof=0 1"`      //群管理类型：.1:仅群主可管理
	ChatBannedType      int      `json:"chatBannedType,omitempty" validate:"omitempty,oneof=0 1"`      //是否开启群禁言：.1:全员禁言
}

//ChatDetail:群详情
type ChatDetail struct {
	Id                  string   `json:"chatid"`              //群id。
	Name                string   `json:"name"`                //群名称，长度限制为1~20个字符。
	Status              int      `json:"status"`              //状态
	Icon                string   `json:"icon"`                //群头像的mediaID
	ConversationTag     int      `json:"conversationTag"`     //会话类型 2:企业群
	Owner               string   `json:"owner"`               //群主的userid，
	UserIds             []string `json:"useridlist"`          //群成员列表，每次最多支持40人，群人数上限为1000
	ShowHistory         int      `json:"showHistoryType"`     //新成员是否可查看100条历史消息：1:可看
	Searchable          int      `json:"searchable"`          //群是否可以被搜索.1:可搜索
	Validation          int      `json:"validationType"`      //入群是否需要验证.1:要验证
	MentionAllAuthority int      `json:"mentionAllAuthority"` //@all 使用范围：.1:仅群主可用
	ManagementType      int      `json:"managementType" `     //群管理类型：.1:仅群主可管理
	ChatBannedType      int      `json:"chatBannedType" `     //是否开启群禁言：.1:全员禁言
}

//CreateChatResponse:创建群返回
type CreateChatResponse struct {
	Response
	ConversationTag    int    `json:"conversationTag"`    //会话类型 2:企业群
	OpenConversationId string `json:"openConversationId"` //公开群会话的Id
	ChatId             string `json:"chatid"`             //群会话的Id
}

//GetChatInfoResponse:获取群信息返回
type GetChatInfoResponse struct {
	Response
	ChatInfo ChatDetail `json:"chat_info"`
}

//ChatSetResponse:群设置返回
type ChatSetResponse struct {
	Response
	Success   bool   `json:"success"`
	RequestId string `json:"request_id"`
}

//ChatSetResponse:群设置返回
type GetChatMsgReadResponse struct {
	Response
	NextCursor int      `json:"next_cursor"`    //下次分页获取的起始游标。
	UserIds    []string `json:"readUserIdList"` //已读人员的userid列表。已读人员为空时不返回该参数。
}

//请求参数验证
func (u CreateChat) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(u); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}

//请求参数验证
func (u UpdateChat) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(u); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
