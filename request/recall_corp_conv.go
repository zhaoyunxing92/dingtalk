package request

type RecallCorpConvMessage struct {
	//发送消息时使用的微应用的ID
	AgentId int `json:"agent_id" validate:"required"`

	//发送消息时钉钉返回的任务ID。
	TaskId int `json:"msg_task_id" validate:"required"`
}

type recallCorpConvMessageBuilder struct {
	mp *RecallCorpConvMessage
}

func NewRecallCorpConvMessage(taskId int) *recallCorpConvMessageBuilder {
	return &recallCorpConvMessageBuilder{mp: &RecallCorpConvMessage{TaskId: taskId}}
}

func (b *recallCorpConvMessageBuilder) SetAgentId(agentId int) *recallCorpConvMessageBuilder {
	b.mp.AgentId = agentId
	return b
}

func (b *recallCorpConvMessageBuilder) Build() *RecallCorpConvMessage {
	return b.mp
}
