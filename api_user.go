package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
)

type userResult struct {
	UserId string `json:"userid"`
}

//创建用户返回
type UserCreateResponse struct {
	model.Response
	userResult `json:"result"`
}

//创建用户
//name:姓名
//mobile:手机号
//deptId:部门
func (talk *DingTalk) CreateUser(name, mobile string, deptId int) (req UserCreateResponse, err error) {

	form := map[string]interface{}{
		"name":         name,
		"mobile":       mobile,
		"dept_id_list": deptId,
	}

	err = talk.request(http.MethodGet, global.CreateUserKey, nil, form, &req)
	return req, err
}

//创建详细的用户
func (talk *DingTalk) CreateDetailUser(res model.Request) (req UserCreateResponse, err error) {

	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return UserCreateResponse{}, err
	}

	err = talk.request(http.MethodGet, global.CreateUserKey, nil, res, &req)
	return req, err
}
