package response

type CorpPermanentCode struct {
	Response
	corpAuthInfo  `json:"auth_corp_info"`
	PermanentCode string `json:"permanent_code"`
}

type corpAuthInfo struct {
	//授权方企业CorpId
	CorpId string `json:"corpid"`

	//授权方企业名称。
	CorpName string `json:"corp_name"`
}
