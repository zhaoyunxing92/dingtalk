package response

type JsApiTicket struct {
	Response
	// Expires 过期时间
	Expires int16 `json:"expires_in"`
	// Create 创建时间
	Create int64 `json:"create"`
	// Ticket 生成的临时jsapi_ticket。
	Ticket string `json:"ticket"`
}


//CreatedAt is when the access token is generated
func (token *JsApiTicket) CreatedAt() int64 {
	return token.Create
}

//ExpiresIn is how soon the access token is expired
func (token *JsApiTicket) ExpiresIn() int16 {
	return token.Expires
}
