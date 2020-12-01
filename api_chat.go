package dingtalk

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/lihongchen/dingtalk/global"
	"github.com/lihongchen/dingtalk/model"
)

// CreateChat:创建群
// name:群名称
// owner:群主管
// userIds:群成员
func (talk *DingTalk) CreateChat(name, owner string, userIds []string) (rsp model.CreateChatResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["name"] = name
	form["owner"] = owner
	form["useridlist"] = userIds

	err = talk.request(http.MethodPost, global.CreateChatKey, nil, form, &rsp)

	return rsp, err
}

//CreateDetailChat:创建详细的群
func (talk *DingTalk) CreateDetailChat(res model.Request) (req model.CreateChatResponse, err error) {

	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return req, err
	}
	err = talk.request(http.MethodPost, global.CreateChatKey, nil, res, &req)

	return req, err
}

//GetChatInfo:获取群信息
func (talk *DingTalk) GetChatInfo(chatId string) (req model.GetChatInfoResponse, err error) {
	params := url.Values{}
	params.Set("chatid", chatId)

	err = talk.request(http.MethodGet, global.GetChatInfoKey, params, nil, &req)
	return req, err
}

//UpdateChat:更新群
func (talk *DingTalk) UpdateChat(res model.Request) (req model.Response, err error) {
	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return req, err
	}
	err = talk.request(http.MethodPost, global.UpdateChatKey, nil, res, &req)
	return req, err
}

//ChatFriendSwitch:设置禁止群成员私聊
func (talk *DingTalk) ChatFriendSwitch(chatId string, prohibit bool) (req model.ChatSetResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["chatid"] = chatId
	form["is_prohibit"] = prohibit

	err = talk.request(http.MethodPost, global.ChatFriendSwitchKey, nil, form, &req)
	return req, err
}

//ChatSubAdmin:设置群管理员,如果设置的用户id不存在则会返回code=400001的系统错误
//chatId:群id
//userId:用户id
//role:2：添加为管理员。 3：删除该管理员。
func (talk *DingTalk) ChatSubAdmin(chatId, userId string, role int) (req model.ChatSetResponse, err error) {

	if role > 3 || role < 2 {
		role = 3
	}
	form := make(map[string]interface{}, 3)
	form["chatid"] = chatId
	form["userids"] = userId
	form["role"] = role

	err = talk.request(http.MethodPost, global.ChatSubAdminKey, nil, form, &req)
	return req, err
}

//SendMsgToChat:发送消息到群
func (talk *DingTalk) SendMsgToChat(chatId string, msg model.Request) (req model.MessageResponse, err error) {

	if err = msg.Validate(talk.validate, talk.trans); err != nil {
		return req, err
	}

	form := make(map[string]interface{}, 2)
	form["chatid"] = chatId
	form["msg"] = msg

	err = talk.request(http.MethodPost, global.SendMsgToChatKey, nil, form, &req)
	return req, err
}

//GetChatMsgReadUser:查询群消息已读人员列表
func (talk *DingTalk) GetChatMsgReadUser(messageId string, cursor, size int) (req model.GetChatMsgReadResponse, err error) {
	if size > 100 || size < 0 {
		size = 100
	}
	params := url.Values{}
	params.Set("messageId", messageId)
	params.Set("cursor", strconv.Itoa(cursor))
	params.Set("size", strconv.Itoa(size))

	err = talk.request(http.MethodGet, global.GetChatReadUserKey, params, nil, &req)
	return req, err
}
