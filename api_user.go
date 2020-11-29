package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
	"net/url"
	"strconv"
)

//CreateUser:创建用户
//name:姓名
//mobile:手机号
//deptId:部门
func (talk *DingTalk) CreateUser(name, mobile string, deptIds []int) (req model.UserIdResponse, err error) {

	form := map[string]interface{}{
		"name":       name,
		"mobile":     mobile,
		"department": deptIds,
	}

	err = talk.request(http.MethodPost, global.CreateUserKey, nil, form, &req)
	return req, err
}

//DeleteUser:删除用户
//userId:员工唯一标识userid
func (talk *DingTalk) DeleteUser(userId string) (req model.Response, err error) {

	params := url.Values{}
	params.Set("userid", userId)

	err = talk.request(http.MethodGet, global.DeleteUserKey, params, nil, &req)
	return req, err
}

//UpdateUser:更新用户
//res:员工信息
//TODO:PositionInDepts:只用写一个，因为写多了也是无用的
func (talk *DingTalk) UpdateUser(res model.Request) (req model.Response, err error) {

	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return req, err
	}

	err = talk.request(http.MethodPost, global.UpdateUserKey, nil, res, &req)
	return req, err
}

//GetUserDetail:获取用户
//userId:用户id
//lang:语言，默认是zh_CN
func (talk *DingTalk) GetUserDetail(userId, lang string) (req model.UserDetail, err error) {

	if lang != "en_US" {
		lang = "zh_CN"
	}

	params := url.Values{}
	params.Set("userid", userId)
	params.Set("lang", lang)

	err = talk.request(http.MethodGet, global.GetUserKey, params, nil, &req)
	return req, err
}

//GetUserIdByUnionId:根据unionid获取userid
//unionId:员工在当前开发者企业账号范围内的唯一标识
func (talk *DingTalk) GetUserIdByUnionId(unionId string) (req model.UserIdResponse, err error) {

	params := url.Values{}
	params.Set("unionid", unionId)

	err = talk.request(http.MethodGet, global.GetUserIdByUnionIdKey, params, nil, &req)
	return req, err
}

//GetUserIdByMobile:根据手机号获取userid
//unionId:员工在当前开发者企业账号范围内的唯一标识
func (talk *DingTalk) GetUserIdByMobile(mobile string) (req model.UserIdResponse, err error) {

	params := url.Values{}
	params.Set("mobile", mobile)

	err = talk.request(http.MethodGet, global.GetUserIdByMobileKey, params, nil, &req)
	return req, err
}

//GetOrgUserCount:获取企业员工人数
//active:0:包含未激活钉钉的人员数量,1:不包含未激活钉钉的人员数量
func (talk *DingTalk) GetOrgUserCount(active int) (req model.OrgUserCountResponse, err error) {

	if active < 0 || active > 2 {
		active = 0
	}

	params := url.Values{}
	params.Set("onlyActive", strconv.Itoa(active))

	err = talk.request(http.MethodGet, global.GetOrgUserCountKey, params, nil, &req)
	return req, err
}

//GetOrgInactiveUser:获取未登录钉钉的员工列表
//active:0:包含未激活钉钉的人员数量,1:不包含未激活钉钉的人员数量
func (talk *DingTalk) GetOrgInactiveUser(date string, offset, size int) (req model.InactiveUserResponse, err error) {
	if size < 0 || size > 100 {
		size = 100
	}

	form := make(map[string]interface{}, 3)
	form["query_date"] = date
	form["offset"] = offset
	form["size"] = size

	err = talk.request(http.MethodPost, global.GetOrgInactiveUserKey, nil, form, &req)
	return req, err
}

//GetOrgAdminUser:获取未登录钉钉的员工列表
func (talk *DingTalk) GetOrgAdminUser() (req model.OrgAdminUserResponse, err error) {
	err = talk.request(http.MethodGet, global.GetOrgAdminUserKey, nil, nil, &req)
	return req, err
}

//GetOrgAdminScope:获取管理员通讯录权限范围
func (talk *DingTalk) GetOrgAdminScope(userId string) (req model.OrgAdminScopeResponse, err error) {
	form := make(map[string]string, 1)
	form["userid"] = userId

	err = talk.request(http.MethodPost, global.GetOrgAdminScopeKey, nil, form, &req)
	return req, err
}
