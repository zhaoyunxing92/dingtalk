package request

type CreatChat struct {
	//群名称，长度限制为1~20个字符
	Name string `json:"name" validate:"required,min=1,max=20"`

	//群主的userid,该员工必须为会话useridlist的成员之一。
	Owner string `json:"owner" validate:"required"`

	//群成员列表，每次最多支持40人，群人数上限为1000
	Users []string `json:"useridlist" validate:"required,max=40"`

	//新成员是否可查看100条历史消息：
	//
	//1：可查看
	//
	//0：不可查看
	//
	//如果不传值，代表不可查看。
	ShowHistory *int `json:"showHistoryType,omitempty"`

	//群是否可以被搜索：
	//
	//0（默认）：不可搜索
	//
	//1：可搜索
	Searchable *int `json:"searchable,omitempty"`

	//入群是否需要验证
	//
	//0（默认）：不验证
	//
	//1：入群验证
	Validation *int `json:"validationType,omitempty"`

	//@all 使用范围：
	//
	//0（默认）：所有人可使用
	//
	//1：仅群主可@all
	MentionAllAuthority *int `json:"mentionAllAuthority,omitempty"`

	//群管理类型：
	//
	//0（默认）：所有人可管理
	//1：仅群主可管理
	ManagementType *int `json:"managementType,omitempty"`

	//是否开启群禁言：
	//0（默认）：不禁言
	//1：全员禁言
	ChatBannedType *int `json:"chatBannedType,omitempty"`
}

type creatChatBuilder struct {
	c *CreatChat
}

func NewCreatChat(name, owner string, userIds ...string) *creatChatBuilder {
	us := append(userIds, owner)
	return &creatChatBuilder{c: &CreatChat{Name: name, Owner: owner, Users: removeStringDuplicates(us)}}
}

//SetShowHistory 新成员是否可查看100条历史消息：
//1：可查看
//0：不可查看
//
//如果不传值，代表不可查看。
func (cb *creatChatBuilder) SetShowHistory(history int) *creatChatBuilder {
	cb.c.ShowHistory = &history
	return cb
}

//SetSearchable 群是否可以被搜索：
//
//0（默认）：不可搜索
//
//1：可搜索
func (cb *creatChatBuilder) SetSearchable(search int) *creatChatBuilder {
	cb.c.Searchable = &search
	return cb
}

//SetValidation 入群是否需要验证
//
//0（默认）：不验证
//
//1：入群验证
func (cb *creatChatBuilder) SetValidation(valida int) *creatChatBuilder {
	cb.c.Validation = &valida
	return cb
}

//SetMentionAllAuthority @all 使用范围：
//
//0（默认）：所有人可使用
//
//1：仅群主可@all
func (cb *creatChatBuilder) SetMentionAllAuthority(mentionAllAuthority int) *creatChatBuilder {
	cb.c.MentionAllAuthority = &mentionAllAuthority
	return cb
}

//SetManagementType 群管理类型：
//
//0（默认）：所有人可管理
//1：仅群主可管理
func (cb *creatChatBuilder) SetManagementType(managementType int) *creatChatBuilder {
	cb.c.ManagementType = &managementType
	return cb
}

//SetChatBannedType 是否开启群禁言：
//0（默认）：不禁言
//1：全员禁言
func (cb *creatChatBuilder) SetChatBannedType(chatBannedType int) *creatChatBuilder {
	cb.c.ChatBannedType = &chatBannedType
	return cb
}

func (cb *creatChatBuilder) Build() *CreatChat {
	return cb.c
}
