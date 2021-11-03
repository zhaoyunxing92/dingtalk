package response

import "github.com/zhaoyunxing92/dingtalk/v2/request"

type CreateUser struct {
	Response
	request.DeleteUser `json:"result"`
}
