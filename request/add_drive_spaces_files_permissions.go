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

package request

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant/member"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/role"
)

type AddDriveSpacesFilesPermissions struct {
	// 钉盘空间ID
	SpaceId string `json:"-" validate:"required"`

	// 文件ID
	FileId string `json:"-" validate:"required"`

	// 用户unionId
	UnionId string `json:"unionId" validate:"required"`

	//权限角色。
	//
	//owner：所有者
	//admin：管理员
	//editor：可编辑
	//viewer：可查看/下载
	//only_viewer：只读
	Role string `json:"role"`

	// 成员权限列表
	Members []driveSpacesFileMember `json:"members"`
}

type driveSpacesFileMember struct {
	// 企业的CorpId
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
}

type addDriveSpacesFilesPermissionsBuilder struct {
	fp *AddDriveSpacesFilesPermissions
}

func newDriveSpacesFileMember(corpId, memberId string, mt member.Type) driveSpacesFileMember {
	return driveSpacesFileMember{corpId, memberId, string(mt)}
}

func NewAddDriveSpacesFilesPermissions() *addDriveSpacesFilesPermissionsBuilder {
	return &addDriveSpacesFilesPermissionsBuilder{fp: &AddDriveSpacesFilesPermissions{}}
}

func (b *addDriveSpacesFilesPermissionsBuilder) SetSpaceId(spaceId string) *addDriveSpacesFilesPermissionsBuilder {
	b.fp.SpaceId = spaceId
	return b
}

func (b *addDriveSpacesFilesPermissionsBuilder) SetFileId(fileId string) *addDriveSpacesFilesPermissionsBuilder {
	b.fp.FileId = fileId
	return b
}

func (b *addDriveSpacesFilesPermissionsBuilder) SetUnionId(unionId string) *addDriveSpacesFilesPermissionsBuilder {
	b.fp.UnionId = unionId
	return b
}

func (b *addDriveSpacesFilesPermissionsBuilder) SetRole(role role.SpacesRole) *addDriveSpacesFilesPermissionsBuilder {
	b.fp.Role = string(role)
	return b
}

func (b *addDriveSpacesFilesPermissionsBuilder) SetSpacesFileMember(corpId, memberId string,
	mt member.Type) *addDriveSpacesFilesPermissionsBuilder {
	members := b.fp.Members
	b.fp.Members = append(members, newDriveSpacesFileMember(corpId, memberId, mt))
	return b
}

func (b *addDriveSpacesFilesPermissionsBuilder) Build() *AddDriveSpacesFilesPermissions {
	return b.fp
}
