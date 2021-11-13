package request

type ReauthCorp struct {

	//套件下的微应用ID
	AppId int `json:"app_id" validate:"required,min=1"`

	//授权未激活应用的CorpId列表
	CorpIds []string `json:"corpid_list" validate:"required,min=1"`
}

func NewReauthCorp(appId int, corpIds []string) *ReauthCorp {
	return &ReauthCorp{appId, removeStringDuplicates(corpIds)}
}
