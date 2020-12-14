package dingtalk

import (
	"errors"
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
	"strings"
)

//GetRoleList:获取角色列表
//offset:支持分页查询，与size参数同时设置时才生效，此参数代表偏移量，偏移量从0开始
//size:支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，默认值20，最大值200。
func (talk *DingTalk) GetRoleList(offset, size int) (apps model.RoleListResponse, err error) {
	if size > 200 || size < 0 {
		size = 200
	}
	form := map[string]int{
		"offset": offset,
		"size":   size,
	}
	err = talk.request(http.MethodPost, global.GetRoleListKey, nil, form, &apps)
	return apps, err
}

//GetRoleUserList:获取指定角色的员工列表
//roleId:角色Id
//offset:支持分页查询，与size参数同时设置时才生效，此参数代表偏移量，偏移量从0开始
//size:支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，默认值20，最大值200。
//
//{
//    "errcode": 0,
//    "result": {
//        "hasMore": false,
//        "list": [
//            {
//                "name": "xx",
//                "userid": "xx"
//            },
//            {
//                "name": "xx",
//                "userid": "xx"
//            }
//        ]
//    },
//    "request_id": "y5y017a37yhd"
//}
func (talk *DingTalk) GetRoleUserList(roleId, offset, size int) (apps model.RoleUserListResponse, err error) {
	if size > 200 || size < 0 {
		size = 200
	}
	form := map[string]int{
		"role_id": roleId,
		"offset":  offset,
		"size":    size,
	}
	err = talk.request(http.MethodPost, global.GetRoleUserListKey, nil, form, &apps)
	return apps, err
}

//GetRoleGroup:获取角色组
//groupId:组id
func (talk *DingTalk) GetRoleGroup(groupId int) (apps model.RoleGroupResponse, err error) {

	form := map[string]int{
		"group_id": groupId,
	}
	err = talk.request(http.MethodPost, global.GetRoleGroupKey, nil, form, &apps)
	return apps, err
}

//GetRoleDetail:获取角色详情
//todo：如果你传的是角色组id，那么会返回角色组的信息
func (talk *DingTalk) GetRoleDetail(roleId int) (apps model.RoleDetailResponse, err error) {

	form := map[string]int{
		"roleId": roleId,
	}
	err = talk.request(http.MethodPost, global.GetRoleDetailKey, nil, form, &apps)
	return apps, err
}

//CreateRoleGroup:创建角色组
func (talk *DingTalk) CreateRoleGroup(name string) (apps model.CreateRoleGroupResponse, err error) {

	form := map[string]string{
		"name": name,
	}
	err = talk.request(http.MethodPost, global.CreateRoleGroupKey, nil, form, &apps)
	return apps, err
}

//CreateRole:创建角色
//name:角色名称
//groupId:角色组id
func (talk *DingTalk) CreateRole(name string, groupId int) (apps model.CreateRoleResponse, err error) {

	form := map[string]interface{}{
		"roleName": name,
		"groupId":  groupId,
	}
	err = talk.request(http.MethodPost, global.CreateRoleKey, nil, form, &apps)
	return apps, err
}

//UpdateRole:更新角色
// todo:如果传入的是角色组id则会修改角色组信息
//name:角色名称
//roleId:角色id
func (talk *DingTalk) UpdateRole(name string, roleId int) (apps model.Response, err error) {

	form := map[string]interface{}{
		"roleName": name,
		"roleId":   roleId,
	}
	err = talk.request(http.MethodPost, global.UpdateRoleKey, nil, form, &apps)
	return apps, err
}

//DeleteRole:删除角色
// todo:如果传入的是角色组id则会删除角色组
//roleId:角色id
func (talk *DingTalk) DeleteRole(roleId int) (apps model.Response, err error) {

	form := map[string]int{
		"role_id": roleId,
	}
	err = talk.request(http.MethodPost, global.DeleteRoleKey, nil, form, &apps)
	return apps, err
}

//RoleBatchAddUser:批量增加员工角色
//roleIds:角色roleId列表，最多可传20个。
//userIds:员工的userId,，最多可传20个。
func (talk *DingTalk) RoleBatchAddUser(roleIds []string, userIds []string) (apps model.Response, err error) {

	if len(roleIds) > 20 {
		err = errors.New("一次最多20个角色")
	}

	if len(userIds) > 20 {
		err = errors.New("一次最多20个用户")
	}

	var form = map[string]string{
		"userIds": strings.Join(userIds, ","),
		"roleIds": strings.Join(roleIds, ","),
	}
	err = talk.request(http.MethodPost, global.RoleBatchAddUserKey, nil, form, &apps)
	return apps, err
}

//RoleBatchRemoveUser:批量删除员工角色
//roleIds:角色roleId列表，最多可传20个。
//userIds:员工的userId,，最多可传20个。
func (talk *DingTalk) RoleBatchRemoveUser(roleIds []string, userIds []string) (apps model.Response, err error) {

	if len(roleIds) > 20 {
		err = errors.New("一次最多20个角色")
	}

	if len(userIds) > 20 {
		err = errors.New("一次最多20个用户")
	}

	var form = map[string]string{
		"userIds": strings.Join(userIds, ","),
		"roleIds": strings.Join(roleIds, ","),
	}
	err = talk.request(http.MethodPost, global.RoleBatchRemoveUserKey, nil, form, &apps)
	return apps, err
}

//RoleUpdateUserManageScope:设定角色成员管理范围
//userId:用户id
//roleId:角色id
//deptIds:部门ID列表数。最多50个，不传则设置范围为所有人
func (talk *DingTalk) RoleUpdateUserManageScope(userId string, roleId int, deptIds []int) (apps model.Response, err error) {

	if deptIds != nil && len(deptIds) > 50 {
		err = errors.New("最多50个部门")
	}

	var form = map[string]interface{}{
		"userid":  userId,
		"role_id": roleId,
	}

	if deptIds != nil {
		form["dept_ids"] = deptIds
	}
	err = talk.request(http.MethodPost, global.RoleUpdateUserManageScopeKey, nil, form, &apps)
	return apps, err
}