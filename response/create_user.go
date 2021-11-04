package response

type CreateUser struct {
	Response
	createUserResponse `json:"result"`
}

type createUserResponse struct {
	UserId string `json:"userid"`
}
