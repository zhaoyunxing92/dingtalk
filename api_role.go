package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
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
