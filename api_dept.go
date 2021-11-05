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

//CreateDept 创建部门
func (ding *DingTalk) CreateDept(res *request.CreateDept) (rsp response.CreateDept, err error) {

	return rsp, ding.Request(http.MethodPost, constant.CreateDeptKey, nil, res, &rsp)
}

//GetDeptDetail:获取部门详情
func (ding *DingTalk) GetDeptDetail(deptId int, lang string) (rsp model.DeptDetail, err error) {

	if lang != "en_US" {
		lang = "zh_CN"
	}

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))
	params.Set("lang", lang)

	err = ding.Request(http.MethodGet, constant.GetDeptDetailKey, params, nil, &rsp)

	return rsp, err
}

//DeleteDept:删除部门
func (ding *DingTalk) DeleteDept(deptId int) (rsp model.Response, err error) {

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))

	err = ding.Request(http.MethodGet, constant.DeleteDeptKey, params, nil, &rsp)

	return rsp, err
}

//UpdateDept:更新部门
func (ding *DingTalk) UpdateDept(res model.CreateDetailDeptRequest) (rsp model.DeptCreateResponse, err error) {

	//组装参数
	res.JoinOuterPermitUsers()
	res.JoinOuterPermitDepts()
	res.JoinDeptPermits()
	res.JoinUserPermits()
	res.JoinDeptManagerUserIds()

	err = ding.Request(http.MethodPost, constant.UpdateDeptKey, nil, res, &rsp)

	return rsp, err
}

//GetSubDeptList:获取子部门列表
func (ding *DingTalk) GetSubDeptList(deptId int, lang string, fetch bool) (rsp model.GetSubDeptResponse, err error) {

	if lang != "en_US" {
		lang = "zh_CN"
	}

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))
	params.Set("lang", lang)
	params.Set("fetch_child", strconv.FormatBool(fetch))

	err = ding.Request(http.MethodGet, constant.GetSubDeptListKey, params, nil, &rsp)

	return rsp, err
}

//GetDeptUserIds 获取部门用户userid列表
func (ding *DingTalk) GetDeptUserIds(du *request.DeptUserId) (req response.DeptUserId, err error) {

	return req, ding.Request(http.MethodPost, constant.GetDeptUserIdKey, nil, du, &req)
}

//GetDeptSimpleUserInfo 获取部门用户基础信息
func (ding *DingTalk) GetDeptSimpleUserInfo(res *request.DeptSimpleUserInfo) (req response.DeptSimpleUserInfo,
	err error) {

	return req, ding.Request(http.MethodPost, constant.GetDeptSimpleUserKey, nil, res, &req)
}

//GetDeptDetailUserInfo 获取部门用户详情
func (ding *DingTalk) GetDeptDetailUserInfo(res *request.DeptDetailUserInfo) (req response.DeptDetailUserInfo,
	err error) {

	return req, ding.Request(http.MethodPost, constant.GetDeptDetailUserKey, nil, res, &req)
}

//GetDeptUserDetail:获取子部门ID列表
func (ding *DingTalk) GetSubDeptIds(deptId int) (req model.GetSubDeptIdsResponse, err error) {

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))

	err = ding.Request(http.MethodGet, constant.GetSubDeptIdsKey, params, nil, &req)
	return req, err
}

//GetParentIdsByUserId:查询指定用户的所有上级父部门路径
func (ding *DingTalk) GetParentIdsByUserId(userId string) (req model.GetParentIdsByUserIdResponse, err error) {

	params := url.Values{}
	params.Set("userId", userId)

	err = ding.Request(http.MethodGet, constant.GetParentDeptsByUserKey, params, nil, &req)
	return req, err
}

//GetParentIdsByUserId:查询部门的所有上级父部门路径
func (ding *DingTalk) GetParentIdsByDeptId(deptId int) (req model.GetParentIdsByDeptIdResponse, err error) {

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))

	err = ding.Request(http.MethodGet, constant.GetParentDeptsByDeptKey, params, nil, &req)
	return req, err
}
