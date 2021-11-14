package response

type SNSUserInfo struct {
	Response
	snsUserInfo `json:"user_info"`
}
type snsUserInfo struct {
}
