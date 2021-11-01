package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/constant"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
	"net/url"
	"strconv"
)

// CreateChat:创建群
// name:群名称
// owner:群主管
// userIds:群成员
func (ding *DingTalk) CreateChat(name, owner string, userIds []string) (rsp model.CreateChatResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["name"] = name
	form["owner"] = owner
	form["useridlist"] = userIds

	err = ding.request(http.MethodPost, constant.CreateChatKey, nil, form, &rsp)

	return rsp, err
}

//CreateDetailChat:创建详细的群
func (ding *DingTalk) CreateDetailChat(res model.Request) (req model.CreateChatResponse, err error) {

	err = ding.request(http.MethodPost, constant.CreateChatKey, nil, res, &req)

	return req, err
}

//GetChatInfo:获取群信息
func (ding *DingTalk) GetChatInfo(chatId string) (req model.GetChatInfoResponse, err error) {
	params := url.Values{}
	params.Set("chatid", chatId)

	err = ding.request(http.MethodGet, constant.GetChatInfoKey, params, nil, &req)
	return req, err
}

//UpdateChat:更新群
func (ding *DingTalk) UpdateChat(res model.Request) (req model.Response, err error) {

	err = ding.request(http.MethodPost, constant.UpdateChatKey, nil, res, &req)
	return req, err
}

//ChatFriendSwitch:设置禁止群成员私聊
func (ding *DingTalk) ChatFriendSwitch(chatId string, prohibit bool) (req model.ChatSetResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["chatid"] = chatId
	form["is_prohibit"] = prohibit

	err = ding.request(http.MethodPost, constant.ChatFriendSwitchKey, nil, form, &req)
	return req, err
}

//ChatSubAdmin:设置群管理员,如果设置的用户id不存在则会返回code=400001的系统错误
//chatId:群id
//userId:用户id
//role:2：添加为管理员。 3：删除该管理员。
func (ding *DingTalk) ChatSubAdmin(chatId, userId string, role int) (req model.ChatSetResponse, err error) {

	if role > 3 || role < 2 {
		role = 3
	}
	form := make(map[string]interface{}, 3)
	form["chatid"] = chatId
	form["userids"] = userId
	form["role"] = role

	err = ding.request(http.MethodPost, constant.ChatSubAdminKey, nil, form, &req)
	return req, err
}

//SendMsgToChat:发送消息到群
func (ding *DingTalk) SendMsgToChat(chatId string, msg model.Request) (req model.MessageResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["chatid"] = chatId
	form["msg"] = msg

	err = ding.request(http.MethodPost, constant.SendMsgToChatKey, nil, form, &req)
	return req, err
}

//GetChatMsgReadUser:查询群消息已读人员列表
func (ding *DingTalk) GetChatMsgReadUser(messageId string, cursor, size int) (req model.GetChatMsgReadResponse, err error) {
	if size > 100 || size < 0 {
		size = 100
	}
	params := url.Values{}
	params.Set("messageId", messageId)
	params.Set("cursor", strconv.Itoa(cursor))
	params.Set("size", strconv.Itoa(size))

	err = ding.request(http.MethodGet, constant.GetChatReadUserKey, params, nil, &req)
	return req, err
}
