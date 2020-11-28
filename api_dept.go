package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
	"net/url"
	"strconv"
)

type deptResult struct {
	DeptId      int    `json:"dept_id,omitempty"` //部门id
	Name        string `json:"name,omitempty"`    //部门名称
	Nick        string `json:"nick,omitempty"`    //部门别名。
	Chain       []int  `json:"chain,omitempty"`   //节点链。从顶层节点到当前节点的中间节点，其内容不包含当前节点。
	Feature     string `json:"feature,omitempty"` //部门节点特有属性。
	ContactType string `json:"contact_type"`      //通讯录类型，classic:传统经典4层结构。校区/学段/年级/班级，custom：自定义结构
	DeptType    string `json:"dept_type"`         //节点类型
}

//DeptCreateResponse 创建部门返回
type DeptCreateResponse struct {
	model.Response
	deptResult `json:"result"`
}

type DeptUserIdsResponse struct {
	model.Response
	UserIds []string `json:"userIds,omitempty"` //用户列表
}

type GetDeptUserDetailResponse struct {
	model.Response
	More     bool         `json:"hasMore"`  //在分页查询时返回，代表是否还有下一页更多数据。
	UserList []UserDetail `json:"userlist"` //成员列表
}

//GetDepartmentDetail:获取部门详情
func (talk *DingTalk) GetDepartmentDetail(deptId int, lang string) (rsp model.Department, err error) {

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
func (talk *DingTalk) GetDeptUserIds(deptId int) (req DeptUserIdsResponse, err error) {

	params := url.Values{}
	params.Set("deptId", strconv.Itoa(deptId))

	err = talk.request(http.MethodGet, global.GetDeptUserIdKey, params, nil, &req)
	return req, err
}

//GetDeptUserDetail:获取部门用户详情
func (talk *DingTalk) GetDeptUserDetail(deptId, offset, size int, lang string) (req GetDeptUserDetailResponse, err error) {
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

//CreateV2Department 创建部门
//name:部门名称
//parentId:父部门id
func (talk *DingTalk) CreateV2Department(name string, pId int) (rsp DeptCreateResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["name"] = name
	form["parent_id"] = pId

	err = talk.request(http.MethodPost, global.DepartmentCreateV2Key, nil, form, &rsp)

	return rsp, err
}

//CreateV2DetailDepartment:创建详细的部门
func (talk *DingTalk) CreateV2DetailDepartment(res model.Request) (rsp DeptCreateResponse, err error) {

	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return rsp, err
	}

	err = talk.request(http.MethodPost, global.DepartmentCreateV2Key, nil, res, &rsp)

	return rsp, err
}

//UpdateV2Department:更新部门
func (talk *DingTalk) UpdateV2Department(name string, parentId int) (rsp DeptCreateResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["name"] = name
	form["parent_id"] = parentId

	err = talk.request(http.MethodPost, global.DepartmentUpdateV2Key, nil, form, &rsp)

	return rsp, err
}

//DeleteV2Department:删除部门
func (talk *DingTalk) DeleteV2Department(deptId int) (rsp DeptCreateResponse, err error) {

	form := make(map[string]int, 1)
	form["dept_id"] = deptId

	err = talk.request(http.MethodPost, global.DepartmentDeleteV2Key, nil, form, &rsp)

	return rsp, err
}
