package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
)

//SendToConversation:发送普通消息
func (talk *DingTalk) SendToConversation(senderId, chatId string, msg model.Request) (req model.SendToConversationResponse, err error) {

	if err = msg.Validate(talk.validate, talk.trans); err != nil {
		return req, err
	}

	form := make(map[string]interface{}, 3)
	form["sender"] = senderId
	form["cid"] = chatId
	form["msg"] = msg

	err = talk.request(http.MethodPost, global.SendToConversationKey, nil, form, &req)
	return req, err
}
