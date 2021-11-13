package response

type UnactiveCorp struct {
	Response

	//微应用ID
	AppId int `json:"app_id"`

	//授权未激活应用的CorpId列表
	CorpIds []string `json:"corp_list"`

	//是否还有更多数据
	HasMore bool `json:"has_more"`
}
