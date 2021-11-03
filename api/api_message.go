package api

import (
	"github.com/zhaoyunxing92/dingtalk/constant"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
)

//SendToConversation:发送普通消息
func (ding *DingTalk) SendToConversation(senderId, chatId string, msg model.Request) (req model.SendToConversationResponse, err error) {

	form := make(map[string]interface{}, 3)
	form["sender"] = senderId
	form["cid"] = chatId
	form["msg"] = msg

	err = ding.request(http.MethodPost, constant.SendToConversationKey, nil, form, &req)
	return req, err
}
