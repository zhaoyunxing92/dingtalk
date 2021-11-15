/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package domain

//
//type WorkNotify struct {
//	AgentId    int         `json:"agent_id"` //发送消息时使用的微应用的ID。
//	UserIds    []string    `json:"-" validate:"lte=100"`
//	StrUserIds string      `json:"userid_list,omitempty"`  //接收者的企业内部用户的userId列表，最大用户列表长度100。
//	StrDeptIds string      `json:"dept_id_list,omitempty"` //接收者的部门id列表，最大列表长度20。接收者是部门Id下包括子部门下的所有用户。
//	DeptIds    []int       `json:"-" validate:"lte=20"`
//	All        bool        `json:"to_all_user"` //是否发送给企业全部用户。当设置为false时必须指定userid_list或dept_id_list其中一个参数的值。
//	Msg        interface{} `json:"msg" validate:"required"`
//}
//
////NewOAWorkNotify:构建oa消息,工作通知title默认使用小程序名称
////url:消息点击链接地址，当发送消息为小程序时支持小程序跳转链接。
////bgColor:消息头部的背景颜色
////fs:消息体的表单，最多显示6个，超过会被隐藏。
//func (notify *WorkNotify) NewOAWorkNotify(oa message2.OA) {
//	notify.Msg = message2.NewOaMessage(oa)
//}
//
////NewTextWorkNotify: text文本消息通知
////content:发送的内容
//func (notify *WorkNotify) NewTextWorkNotify(content string) {
//	notify.Msg = message2.NewTextMessage(content)
//}
//
////NewMarkdownWorkNotify:markdown消息
////title:标题
////content:内容
//func (notify *WorkNotify) NewMarkdownWorkNotify(title, content string) {
//	notify.Msg = message2.newMarkDownMessage(title, content)
//}
//
////NewCardMessage：card消息
//func (notify *WorkNotify) NewCardMessage(card message2.Card) {
//	notify.Msg = message2.newCardMessage(card)
//}
//
////NewFileMessage：file消息
//func (notify *WorkNotify) NewFileMessage(mediaId string) {
//	notify.Msg = message2.NewFileMessage(mediaId)
//}
//
////组装部门，移除重复部门
//func (notify *WorkNotify) AssembleDept() {
//	var size = len(notify.DeptIds)
//	if size == 0 {
//		return
//	}
//	ds := make([]string, 0, size)
//	temp := map[int]int{}
//	for idx, id := range notify.DeptIds {
//		if _, ok := temp[id]; !ok {
//			temp[id] = idx
//			ds = append(ds, strconv.Itoa(id))
//		}
//	}
//	//部门id字符串拼接
//	notify.StrDeptIds = strings.Join(ds, ",")
//}
//
////组装用户,移除重复数据
//func (notify *WorkNotify) AssembleUser() {
//	var size = len(notify.UserIds)
//	if size == 0 {
//		return
//	}
//	users := make([]string, 0, size)
//	temp := map[string]int{}
//
//	for i, id := range notify.UserIds {
//		if _, ok := temp[id]; !ok {
//			temp[id] = i
//			users = append(users, id)
//		}
//	}
//	//用户id字符串拼接
//	notify.StrUserIds = strings.Join(users, ",")
//}
//
////设置agentId
//func (notify *WorkNotify) SetAgentId(agentId int) {
//	notify.AgentId = agentId
//}
//
////请求参数验证
//func (notify *WorkNotify) Validate(valid *validator.Validate, trans translator.Translator) error {
//	if err := valid.Struct(notify); err != nil {
//		errs := err.(validator.ValidationErrors)
//		var slice []string
//		for _, msg := range errs {
//			slice = append(slice, msg.Translate(trans))
//		}
//		return errors.New(strings.Join(slice, ","))
//	}
//	if notify.All == false && len(notify.UserIds) == 0 && len(notify.DeptIds) == 0 {
//		return errors.New("当字段All设置为false时必须指定UserIds或DeptIds其中一个参数的值")
//	}
//	return nil
//}
//
////发送工作通知请求
//type WorkNotifyRes interface {
//	Request
//	//拼接部门
//	AssembleDept()
//	//拼接用户
//	AssembleUser()
//	//设置agentId
//	SetAgentId(id int)
//}
