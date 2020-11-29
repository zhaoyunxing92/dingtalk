package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
	"net/url"
	"strconv"
)

// CreateDept:创建部门
// name:部门名称
// parentId:父部门id
func (talk *DingTalk) CreateDept(name string, parentId int) (rsp model.DeptCreateResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["name"] = name
	form["parentid"] = parentId

	err = talk.request(http.MethodPost, global.CreateDeptKey, nil, form, &rsp)

	return rsp, err
}

//CreateDetailDept:创建详细的部门
func (talk *DingTalk) CreateDetailDept(res model.CreateDetailDeptRequest) (req model.DeptCreateResponse, err error) {

	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return req, err
	}
	//组装参数
	res.JoinOuterPermitUsers()
	res.JoinOuterPermitDepts()
	res.JoinDeptPermits()
	res.JoinUserPermits()

	err = talk.request(http.MethodPost, global.CreateDeptKey, nil, res, &req)

	return req, err
}

//GetDeptDetail:获取部门详情
func (talk *DingTalk) GetDeptDetail(deptId int, lang string) (rsp model.DeptDetail, err error) {

	if lang != "en_US" {
		lang = "zh_CN"
	}

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))
	params.Set("lang", lang)

	err = talk.request(http.MethodGet, global.DepartmentGetKey, params, nil, &rsp)

	return rsp, err
}

//GetDeptUserIds:获取部门用户userid列表
//deptId:部门id
func (talk *DingTalk) GetDeptUserIds(deptId int) (req model.DeptUserIdsResponse, err error) {

	params := url.Values{}
	params.Set("deptId", strconv.Itoa(deptId))

	err = talk.request(http.MethodGet, global.GetDeptUserIdKey, params, nil, &req)
	return req, err
}

//GetDeptUserDetail:获取部门用户详情
func (talk *DingTalk) GetDeptUserDetail(deptId, offset, size int, lang string) (req model.GetDeptUserDetailResponse, err error) {
	if lang != "en_US" {
		lang = "zh_CN"
	}

	if size < 0 || size > 100 {
		size = 100
	}
	params := url.Values{}
	params.Set("lang", lang)
	params.Set("department_id", strconv.Itoa(deptId))
	params.Set("offset", strconv.Itoa(offset))
	params.Set("size", strconv.Itoa(size))
	params.Set("order", "entry_desc")

	err = talk.request(http.MethodGet, global.GetDeptUserDetailKey, params, nil, &req)
	return req, err
}
