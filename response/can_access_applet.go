package response

type UserCanAccessApplet struct {
	Response

	Access bool `json:"canAccess"`
}
