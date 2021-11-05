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
func (ding *dingTalk) CreateDept(res *request.CreateDept) (rsp response.CreateDept, err error) {

	return rsp, ding.Request(http.MethodPost, constant.CreateDeptKey, nil, res, &rsp)
}

//DeleteDept 删除部门
func (ding *dingTalk) DeleteDept(res *request.DeleteDept) (rsp response.Response, err error) {

	return rsp, ding.Request(http.MethodPost, constant.DeleteDeptKey, nil, res, &rsp)
}

//UpdateDept 更新部门
func (ding *dingTalk) UpdateDept(res *request.UpdateDept) (rsp response.Response, err error) {

	return rsp, ding.Request(http.MethodPost, constant.UpdateDeptKey, nil, res, &rsp)
}

//GetDeptDetail 获取部门详情
func (ding *dingTalk) GetDeptDetail(res *request.DeptDetail) (rsp response.DeptDetail, err error) {

	return rsp, ding.Request(http.MethodPost, constant.GetDeptDetailKey, nil, res, &rsp)
}

//GetDeptList 获取部门列表
func (ding *dingTalk) GetDeptList(res *request.DeptList) (rsp response.DeptList, err error) {

	return rsp, ding.Request(http.MethodPost, constant.GetDeptListKey, nil, res, &rsp)
}

//GetSubDeptList 获取子部门列表
func (ding *dingTalk) GetSubDeptList(deptId int, lang string, fetch bool) (rsp model.GetSubDeptResponse, err error) {

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
func (ding *dingTalk) GetDeptUserIds(du *request.DeptUserId) (req response.DeptUserId, err error) {

	return req, ding.Request(http.MethodPost, constant.GetDeptUserIdKey, nil, du, &req)
}

//GetDeptSimpleUserInfo 获取部门用户基础信息
func (ding *dingTalk) GetDeptSimpleUserInfo(res *request.DeptSimpleUserInfo) (req response.DeptSimpleUserInfo,
	err error) {

	return req, ding.Request(http.MethodPost, constant.GetDeptSimpleUserKey, nil, res, &req)
}

//GetDeptDetailUserInfo 获取部门用户详情
func (ding *dingTalk) GetDeptDetailUserInfo(res *request.DeptDetailUserInfo) (req response.DeptDetailUserInfo,
	err error) {

	return req, ding.Request(http.MethodPost, constant.GetDeptDetailUserKey, nil, res, &req)
}

//GetDeptUserDetail:获取子部门ID列表
func (ding *dingTalk) GetSubDeptIds(deptId int) (req model.GetSubDeptIdsResponse, err error) {

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))

	err = ding.Request(http.MethodGet, constant.GetSubDeptIdsKey, params, nil, &req)
	return req, err
}

//GetParentIdsByUserId:查询指定用户的所有上级父部门路径
func (ding *dingTalk) GetParentIdsByUserId(userId string) (req model.GetParentIdsByUserIdResponse, err error) {

	params := url.Values{}
	params.Set("userId", userId)

	err = ding.Request(http.MethodGet, constant.GetParentDeptsByUserKey, params, nil, &req)
	return req, err
}

//GetParentIdsByUserId:查询部门的所有上级父部门路径
func (ding *dingTalk) GetParentIdsByDeptId(deptId int) (req model.GetParentIdsByDeptIdResponse, err error) {

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))

	err = ding.Request(http.MethodGet, constant.GetParentDeptsByDeptKey, params, nil, &req)
	return req, err
}
