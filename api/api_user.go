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

//DeleteUser:删除用户
//userId:员工唯一标识userid
func (ding *DingTalk) DeleteUser(userId string) (req model.Response, err error) {

	params := url.Values{}
	params.Set("userid", userId)

	err = ding.Request(http.MethodGet, constant.DeleteUserKey, params, nil, &req)
	return req, err
}

//GetUserDetail:获取用户
//userId:用户id
//lang:语言，默认是zh_CN
func (ding *DingTalk) GetUserDetail(userId, lang string) (req model.UserDetail, err error) {

	if lang != "en_US" {
		lang = "zh_CN"
	}

	params := url.Values{}
	params.Set("userid", userId)
	params.Set("lang", lang)

	err = ding.Request(http.MethodGet, constant.GetUserKey, params, nil, &req)
	return req, err
}

//GetUserIdByUnionId:根据unionid获取userid
//unionId:员工在当前开发者企业账号范围内的唯一标识
func (ding *DingTalk) GetUserIdByUnionId(unionId string) (req model.UserIdResponse, err error) {

	params := url.Values{}
	params.Set("unionid", unionId)

	err = ding.Request(http.MethodGet, constant.GetUserIdByUnionIdKey, params, nil, &req)
	return req, err
}

//GetUserIdByMobile:根据手机号获取userid
//unionId:员工在当前开发者企业账号范围内的唯一标识
func (ding *DingTalk) GetUserIdByMobile(mobile string) (req model.UserIdResponse, err error) {

	params := url.Values{}
	params.Set("mobile", mobile)

	err = ding.Request(http.MethodGet, constant.GetUserIdByMobileKey, params, nil, &req)
	return req, err
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

//GetOrgInactiveUser:获取未登录钉钉的员工列表
//active:0:包含未激活钉钉的人员数量,1:不包含未激活钉钉的人员数量
func (ding *DingTalk) GetOrgInactiveUser(date string, offset, size int) (req model.InactiveUserResponse, err error) {
	if size < 0 || size > 100 {
		size = 100
	}

	form := make(map[string]interface{}, 3)
	form["query_date"] = date
	form["offset"] = offset
	form["size"] = size

	err = ding.Request(http.MethodPost, constant.GetOrgInactiveUserKey, nil, form, &req)
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
