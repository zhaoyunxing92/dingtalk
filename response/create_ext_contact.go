package response

type CreateExtContact struct {
	Response
	UserId string `json:"userid"`
}
