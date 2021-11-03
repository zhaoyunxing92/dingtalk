package response

type CreateUser struct {
	Response
	userCreateResponse `json:"result"`
}

type userCreateResponse struct {
	UserId string `json:"userid"`
}
