package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/model"
	"net/http"
)

//SendToConversation:发送普通消息
func (ding *dingTalk) SendToConversation(senderId, chatId string, msg model.Request) (req model.SendToConversationResponse, err error) {

	form := make(map[string]interface{}, 3)
	form["sender"] = senderId
	form["cid"] = chatId
	form["msg"] = msg

	err = ding.Request(http.MethodPost, constant.SendToConversationKey, nil, form, &req)
	return req, err
}
