/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/model"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
	"net/http"
	"net/url"
	"strconv"
)

//CreateChat 创建群
func (ding *dingTalk) CreateChat(res *request.CreatChat) (rsp response.CreatChat, err error) {

	return rsp, ding.Request(http.MethodPost, constant.CreateChatKey, nil, res, &rsp)
}

//GetChatInfo 获取群信息
func (ding *dingTalk) GetChatInfo(chatId string) (req response.GetChatInfo, err error) {

	query := url.Values{}
	query.Set("chatid", chatId)

	return req, ding.Request(http.MethodGet, constant.GetChatInfoKey, query, nil, &req)
}

//UpdateChat:更新群
func (ding *dingTalk) UpdateChat(res model.Request) (req model.Response, err error) {

	err = ding.Request(http.MethodPost, constant.UpdateChatKey, nil, res, &req)
	return req, err
}

//ChatFriendSwitch:设置禁止群成员私聊
func (ding *dingTalk) ChatFriendSwitch(chatId string, prohibit bool) (req model.ChatSetResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["chatid"] = chatId
	form["is_prohibit"] = prohibit

	err = ding.Request(http.MethodPost, constant.ChatFriendSwitchKey, nil, form, &req)
	return req, err
}

//ChatSubAdmin:设置群管理员,如果设置的用户id不存在则会返回code=400001的系统错误
//chatId:群id
//userId:用户id
//role:2：添加为管理员。 3：删除该管理员。
func (ding *dingTalk) ChatSubAdmin(chatId, userId string, role int) (req model.ChatSetResponse, err error) {

	if role > 3 || role < 2 {
		role = 3
	}
	form := make(map[string]interface{}, 3)
	form["chatid"] = chatId
	form["userids"] = userId
	form["role"] = role

	err = ding.Request(http.MethodPost, constant.ChatSubAdminKey, nil, form, &req)
	return req, err
}

//SendMsgToChat:发送消息到群
func (ding *dingTalk) SendMsgToChat(chatId string, msg model.Request) (req model.MessageResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["chatid"] = chatId
	form["msg"] = msg

	err = ding.Request(http.MethodPost, constant.SendMsgToChatKey, nil, form, &req)
	return req, err
}

//GetChatMsgReadUser:查询群消息已读人员列表
func (ding *dingTalk) GetChatMsgReadUser(messageId string, cursor, size int) (req model.GetChatMsgReadResponse, err error) {
	if size > 100 || size < 0 {
		size = 100
	}
	params := url.Values{}
	params.Set("messageId", messageId)
	params.Set("cursor", strconv.Itoa(cursor))
	params.Set("size", strconv.Itoa(size))

	err = ding.Request(http.MethodGet, constant.GetChatReadUserKey, params, nil, &req)
	return req, err
}

//GetChatQRCode 获取入群二维码链接
func (ding *dingTalk) GetChatQRCode(chatId, userId string) (req response.ChatQRCode, err error) {

	return req, ding.Request(http.MethodPost, constant.GetChatQRCodeKey, nil,
		request.NewChatQRCode(chatId, userId), &req)
}
