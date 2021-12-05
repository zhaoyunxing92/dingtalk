/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package response

type DriveSpacesFilesPermissions struct {
	Response
	//企业内成员权限列表
	Members []driveSpacesFilesPermissions `json:"members"`

	//企业外成员权限列表
	OutMembers []driveSpacesFilesPermissions `json:"outMembers"`
}

type driveSpacesFilesPermissions struct {
	//权限角色。
	//
	//owner：所有者
	//admin：管理员
	//editor：可编辑
	//viewer：可查看/下载
	//only_viewer：只读
	Role string `json:"role"`

	//是否是继承的权限
	Extend bool `json:"extend"`

	permissionMember `json:"member"`
}

//成员信息
type permissionMember struct {
	//企业的CorpId
	CorpId string `json:"corpId"`

	//成员ID。
	//
	//当memberType为org时，取值为corpId
	//当memberType为department时，取值为deptId
	//当memberType为conversation时，取值为chatId
	//当memberType为user时，取值为staffId
	MemberId string `json:"memberId"`

	//成员类型。
	//
	//org：企业
	//department：部门
	//conversation：群
	//user：用户
	MemberType string `json:"memberType"`

	//成员名称
	MemberName string `json:"memberName"`
}
