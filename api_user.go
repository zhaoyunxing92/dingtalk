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
	"net/http"
	"net/url"
	"strconv"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

import (
	"github.com/pkg/errors"
)

//CreateUser 创建用户
func (ding *dingTalk) CreateUser(user *request.CreateUser) (req response.CreateUser, err error) {

	return req, ding.Request(http.MethodPost, constant.CreateUserKey, nil, user, &req)
}

//UpdateUser 更新用户信息
func (ding *dingTalk) UpdateUser(user *request.UpdateUser) (req response.Response, err error) {

	return req, ding.Request(http.MethodPost, constant.UpdateUserKey, nil, user, &req)
}

//DeleteUser 删除用户
func (ding *dingTalk) DeleteUser(userId string) (req response.Response, err error) {

	return req, ding.Request(http.MethodPost, constant.DeleteUserKey, nil, request.NewDeleteUser(userId), &req)
}

//GetUserDetail 根据userid获取用户详情
func (ding *dingTalk) GetUserDetail(user *request.UserDetail) (req response.UserDetail, err error) {

	return req, ding.Request(http.MethodPost, constant.GetUserDetailKey, nil, user, &req)
}

//GetUserIdByUnionId 根据unionid获取用户userid
func (ding *dingTalk) GetUserIdByUnionId(res *request.UnionIdGetUserId) (req response.UnionIdGetUserId, err error) {

	return req, ding.Request(http.MethodPost, constant.GetUserIdByUnionIdKey, nil, res, &req)
}

//GetUserIdByMobile 根据手机号获取userid
func (ding *dingTalk) GetUserIdByMobile(res *request.MobileGetUserId) (req response.MobileGetUserId, err error) {

	return req, ding.Request(http.MethodPost, constant.GetUserIdByMobileKey, nil, res, &req)
}

//GetOrgAdminUser 获取管理员列表
func (ding *dingTalk) GetOrgAdminUser() (req response.OrgAdminUser, err error) {

	return req, ding.Request(http.MethodPost, constant.GetOrgAdminUserKey, nil, nil, &req)
}

//GetOrgAdminScope 获取管理员通讯录权限范围
func (ding *dingTalk) GetOrgAdminScope(res *request.AdminUserScope) (req response.AdminUserScope, err error) {

	return req, ding.Request(http.MethodPost, constant.GetOrgAdminScopeKey, nil, res, &req)
}

//GetUserCanAccessApplet 获取管理员的应用管理权限
func (ding *dingTalk) GetUserCanAccessApplet(appId int, userId string) (req response.UserCanAccessApplet, err error) {

	if !ding.isv() {
		return response.UserCanAccessApplet{}, errors.New("应用必须是产品方案商所开发")
	}

	query := url.Values{}
	query.Set("appId", strconv.Itoa(appId))
	query.Set("userId", userId)

	return req, ding.Request(http.MethodGet, constant.GetUserCanAccessAppletKey, query, nil, &req)
}

//GetUserCount 获取员工人数
func (ding *dingTalk) GetUserCount(res *request.UserCount) (req response.UserCont, err error) {

	return req, ding.Request(http.MethodPost, constant.GetUserCountKey, nil, res, &req)
}

//GetInactiveUser 获取未登录钉钉的员工列表
func (ding *dingTalk) GetInactiveUser(res *request.InactiveUser) (req response.InactiveUser, err error) {

	return req, ding.Request(http.MethodPost, constant.GetInactiveUserKey, nil, res, &req)
}

//GetUserInfoByCode 通过免登码获取用户信息
func (ding *dingTalk) GetUserInfoByCode(code string) (req response.CodeGetUserInfo, err error) {

	return req, ding.Request(http.MethodPost, constant.GetUserInfoByCodeKey, nil,
		request.NewCodeGetUserInfo(code), &req)
}

//GetSSOUserInfo 获取应用管理员的身份信息
func (ding *dingTalk) GetSSOUserInfo(code string) (req response.SSOUserInfo, err error) {
	var (
		corpId = ding.CorpId
		secret = ding.SSOSecret
		token  string
	)

	if len(corpId) <= 0 || len(secret) <= 0 {
		return response.SSOUserInfo{}, err
	}

	if token, err = ding.GetSSOToken(corpId, secret); err != nil {
		return response.SSOUserInfo{}, err
	}

	query := url.Values{}
	query.Set("code", code)
	query.Set("access_token", token)

	return req, ding.Request(http.MethodGet, constant.GetSSOUserInfoKey, query, nil, &req)
}
