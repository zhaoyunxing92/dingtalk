package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
)

type result struct {
	DeptId int `json:"dept_id"`
}

//DeptCreateResponse 创建部门返回
type DeptCreateResponse struct {
	model.Response
	result `json:"result"`
}

//CreateDepartment 创建部门
//name:部门名称
//parentId:父部门id
func (talk *DingTalk) CreateDepartment(name string, parentId int) (rsp DeptCreateResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["name"] = name
	form["parent_id"] = parentId

	err = talk.request(http.MethodPost, global.DepartmentCreateKey, nil, form, &rsp)

	return rsp, err
}

//CreateDetailDepartment:创建详细的部门
func (talk *DingTalk) CreateDetailDepartment(name string, parentId int) (rsp DeptCreateResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["name"] = name
	form["parent_id"] = parentId

	err = talk.request(http.MethodPost, global.DepartmentCreateKey, nil, form, &rsp)

	return rsp, err
}

//UpdateDepartment:更新部门
func (talk *DingTalk) UpdateDepartment(name string, parentId int) (rsp DeptCreateResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["name"] = name
	form["parent_id"] = parentId

	err = talk.request(http.MethodPost, global.DepartmentUpdateKey, nil, form, &rsp)

	return rsp, err
}

//DeleteDepartment:删除部门
func (talk *DingTalk) DeleteDepartment(deptId int) (rsp DeptCreateResponse, err error) {

	form := make(map[string]int, 1)
	form["dept_id"] = deptId

	err = talk.request(http.MethodPost, global.DepartmentDeleteKey, nil, form, &rsp)

	return rsp, err
}
