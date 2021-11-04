package response

type MobileGetUserId struct {
	Response

	mobileGetUserId `json:"result"`
}

type mobileGetUserId struct {
	UserId string `json:"userid"`
}
