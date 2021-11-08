package response

type ChatQRCode struct {
	Response
	Url string `json:"result"`
}
