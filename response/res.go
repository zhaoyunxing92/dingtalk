package response

//Response 响应
//{"errcode":40035,"errmsg":"缺少参数 corpid or appkey"}
type Response struct {
	Code      int    `json:"errcode"`
	Msg       string `json:"errmsg,omitempty"`
	Success   bool   `json:"success,omitempty"`
	RequestId string `json:"request_id,omitempty"`
}
