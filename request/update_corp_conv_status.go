package request

type UpdateCorpConvMsgStatus struct {
	//发送消息时使用的微应用的ID。
	AgentId int `json:"agent_id" validate:"required"`

	TaskId int `json:"task_id" validate:"required"`

	//状态栏值
	StatusValue string `json:"status_value,omitempty" validate:"required"`

	//状态栏背景色，推荐0xFF加六位颜色值
	StatusBgColor string `json:"status_bg,omitempty"`
}

type updateCorpConvMsgStatusBuilder struct {
	conv *UpdateCorpConvMsgStatus
}

func NewUpdateCorpConvMsgStatus(taskId int, value string) *updateCorpConvMsgStatusBuilder {
	return &updateCorpConvMsgStatusBuilder{conv: &UpdateCorpConvMsgStatus{TaskId: taskId, StatusValue: value}}
}

func (b *updateCorpConvMsgStatusBuilder) SetAgentId(agentId int) *updateCorpConvMsgStatusBuilder {
	b.conv.AgentId = agentId
	return b
}

func (b *updateCorpConvMsgStatusBuilder) SetStatusBgColor(color string) *updateCorpConvMsgStatusBuilder {
	b.conv.StatusBgColor = color
	return b
}

func (b *updateCorpConvMsgStatusBuilder) Build() *UpdateCorpConvMsgStatus {
	return b.conv
}
