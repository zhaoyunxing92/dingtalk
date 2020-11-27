package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
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
func (talk *DingTalk) CreateDetailDepartment(res model.Request) (rsp DeptCreateResponse, err error) {

	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return rsp, err
	}

	err = talk.request(http.MethodPost, global.DepartmentCreateKey, nil, res, &rsp)

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

//DeleteDepartment:删除部门
func (talk *DingTalk) GetDepartmentDetail(deptId int) (rsp DeptCreateResponse, err error) {

	form := make(map[string]int, 1)
	form["dept_id"] = deptId

	err = talk.request(http.MethodPost, global.DepartmentEduGetKey, nil, form, &rsp)

	return rsp, err
}
