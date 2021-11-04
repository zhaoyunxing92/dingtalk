package api

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/model"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
	"net/http"
	"net/url"
	"strconv"
)

//CreateUser 创建用户
func (ding *DingTalk) CreateUser(user *request.CreateUser) (req response.CreateUser, err error) {

	return req, ding.Request(http.MethodPost, constant.CreateUserKey, nil, user, &req)
}

//UpdateUser 更新用户信息
func (ding *DingTalk) UpdateUser(user *request.UpdateUser) (req response.Response, err error) {

	return req, ding.Request(http.MethodPost, constant.UpdateUserKey, nil, user, &req)
}

//DeleteUser 删除用户
func (ding *DingTalk) DeleteUser(userId string) (req response.Response, err error) {

	return req, ding.Request(http.MethodPost, constant.DeleteUserKey, nil, request.NewDeleteUser(userId), &req)
}

//GetUserDetail 根据userid获取用户详情
func (ding *DingTalk) GetUserDetail(user *request.UserDetail) (req response.UserDetail, err error) {

	return req, ding.Request(http.MethodPost, constant.GetUserDetailKey, nil, user, &req)
}

//GetUserIdByUnionId 根据unionid获取用户userid
func (ding *DingTalk) GetUserIdByUnionId(res *request.UnionIdGetUserId) (req response.UnionIdGetUserId, err error) {

	return req, ding.Request(http.MethodPost, constant.GetUserIdByUnionIdKey, nil, res, &req)
}

//GetUserIdByMobile 根据手机号获取userid
func (ding *DingTalk) GetUserIdByMobile(res *request.MobileGetUserId) (req response.MobileGetUserId, err error) {

	return req, ding.Request(http.MethodPost, constant.GetUserIdByMobileKey, nil, res, &req)
}

//GetOrgUserCount:获取企业员工人数
//active:0:包含未激活钉钉的人员数量,1:不包含未激活钉钉的人员数量
func (ding *DingTalk) GetOrgUserCount(active int) (req model.OrgUserCountResponse, err error) {

	if active < 0 || active > 2 {
		active = 0
	}

	params := url.Values{}
	params.Set("onlyActive", strconv.Itoa(active))

	err = ding.Request(http.MethodGet, constant.GetOrgUserCountKey, params, nil, &req)
	return req, err
}

//GetOrgAdminUser:获取未登录钉钉的员工列表
func (ding *DingTalk) GetOrgAdminUser() (req model.OrgAdminUserResponse, err error) {
	err = ding.Request(http.MethodGet, constant.GetOrgAdminUserKey, nil, nil, &req)
	return req, err
}

//GetOrgAdminScope:获取管理员通讯录权限范围
func (ding *DingTalk) GetOrgAdminScope(userId string) (req model.OrgAdminScopeResponse, err error) {
	form := make(map[string]string, 1)
	form["userid"] = userId

	err = ding.Request(http.MethodPost, constant.GetOrgAdminScopeKey, nil, form, &req)
	return req, err
}

//GetUserByAuthCode:通过免登码获取用户信息(v2)
func (ding *DingTalk) GetUserByAuthCode(code string) (req model.UserGetByCodeResponse, err error) {
	form := make(map[string]string, 1)
	form["code"] = code

	err = ding.Request(http.MethodPost, constant.GetUserByAuthCodeKey, nil, form, &req)
	return req, err
}

//GetUserCount 获取员工人数
func (ding *DingTalk) GetUserCount(res *request.UserCount) (req response.UserCont, err error) {

	return req, ding.Request(http.MethodPost, constant.GetUserCountKey, nil, res, &req)
}

//GetInactiveUser 获取未登录钉钉的员工列表
func (ding *DingTalk) GetInactiveUser(res *request.InactiveUser) (req response.InactiveUser, err error) {

	return req, ding.Request(http.MethodPost, constant.GetInactiveUserKey, nil, res, &req)
}
