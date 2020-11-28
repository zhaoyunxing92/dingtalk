package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
	"net/url"
)

type userResult struct {
	UserId string `json:"userid"`
}

//创建用户返回
type UserCreateResponse struct {
	model.Response
	userResult `json:"result,omitempty"`
}

//type getUserResult struct {
//	model.Response
//}


//CreateUser:创建用户
//name:姓名
//mobile:手机号
//deptId:部门
func (talk *DingTalk) CreateUser(name, mobile string, deptIds []int) (req UserCreateResponse, err error) {

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

//GetUser:获取用户
//userId:用户id
//lang:语言，默认是zh_CN
//TODO:PositionInDepts:只用写一个，因为写多了也是无用的
func (talk *DingTalk) GetUser(userId, lang string) (req model.Response, err error) {

	if lang != "en_US" {
		lang = "zh_CN"
	}

	params := url.Values{}
	params.Set("userid", userId)
	params.Set("lang", lang)

	err = talk.request(http.MethodGet, global.GetUserKey, params, nil, &req)
	return req, err
}

//CreateUserV2:创建用户
//name:姓名
//mobile:手机号
//deptId:部门
func (talk *DingTalk) CreateUserV2(name, mobile string, deptId int) (req UserCreateResponse, err error) {

	form := map[string]interface{}{
		"name":         name,
		"mobile":       mobile,
		"dept_id_list": deptId,
	}

	err = talk.request(http.MethodGet, global.CreateUserKey, nil, form, &req)
	return req, err
}

//创建详细的用户
func (talk *DingTalk) CreateDetailUserV2(res model.Request) (req UserCreateResponse, err error) {

	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return req, err
	}

	err = talk.request(http.MethodGet, global.CreateUserKey, nil, res, &req)
	return req, err
}
