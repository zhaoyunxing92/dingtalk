package response

type UserCont struct {
	Response
	userCont `json:"result"`
}

type userCont struct {
	Count int `json:"count"`
}
