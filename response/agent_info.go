package response

type AgentInfo struct {
	Response
	//企业应用名称
	Name string `json:"name"`

	//企业应用ID。
	AgentId int `json:"agentid"`

	//授权方应用头像。
	Logo string `json:"logo_url"`

	//企业应用描述
	Description string `json:"description"`

	//授权方企业应用是否被禁用：
	//
	//0：禁用
	//
	//1：正常
	//
	//2：待激活
	Close int `json:"close"`
}
