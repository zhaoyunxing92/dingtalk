package request

type UnactiveCorp struct {
	//微应用ID
	AppId int `json:"app_id"`
}

func NewUnactiveCorp(appId int) *UnactiveCorp {
	return &UnactiveCorp{appId}
}
