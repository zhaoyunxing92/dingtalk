package response

type UnionIdGetUserId struct {
	Response

	unionIdGetUserId `json:"result"`
}

type unionIdGetUserId struct {
	UserId string `json:"userid"`

	//联系类型：
	//
	//0：企业内部员工
	//
	//1：企业外部联系人
	ContactType int `json:"contact_type"`
}
