/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package request

import (
	"encoding/json"
	"strings"
)

type UpdateDept struct {
	//部门id
	Id uint `json:"dept_id" validate:"required"`

	//部门名称。长度限制为1~64个字符，不允许包含字符"-"","以及","
	Name string `json:"name,omitempty" validate:"omitempty,min=1,max=64"`

	//父部门ID，根部门ID为1。
	ParentId *uint `json:"parent_id,omitempty" validate:"omitempty,min=1"`

	//是否隐藏本部门：
	//
	//true：隐藏部门，隐藏后本部门将不会显示在公司通讯录中
	//
	//false（默认值）：显示部门
	HideDept *bool `json:"hide_dept,omitempty"`

	//指定可以查看本部门的其他部门列表，总数不能超过200
	//
	//当hide_dept为true时，则此值生效
	DeptPermits string `json:"dept_permits,omitempty"`

	DeptPermit []int `json:"-" validate:"max=200"`

	//指定可以查看本部门的人员userid列表，总数不能超过200
	//
	//当hide_dept为true时，则此值生效
	UserPermits string `json:"user_permits,omitempty"`

	UserPermit []string `json:"-" validate:"max=200"`

	//是否限制本部门成员查看通讯录：
	//
	//true：开启限制。开启后本部门成员只能看到限定范围内的通讯录
	//
	//false（默认值）：不限制
	OuterDept *bool `json:"outer_dept,omitempty"`

	//本部门成员是否只能看到所在部门及下级部门通讯录：
	//
	//true：只能看到所在部门及下级部门通讯录
	//
	//false：不能查看所有通讯录，在通讯录中仅能看到自己
	//
	//当outer_dept为true时，此参数生效。
	OuterDeptOnlySelf *bool `json:"outer_dept_only_self,omitempty"`

	//指定本部门成员可查看的通讯录用户userid列表，总数不能超过200
	//
	//当outer_dept为true时，此参数生效。
	UserPermitsUsers string `json:"outer_permit_users,omitempty"`

	UserPermitsUserIds []string `json:"-" validate:"max=200"`

	//指定本部门成员可查看的通讯录部门ID列表，总数不能超过200
	//
	//当outer_dept为true时，此参数生效
	UserPermitsDepts string `json:"outer_permit_depts,omitempty"`

	UserPermitsDeptIds []int `json:"-" validate:"max=200"`

	//是否创建一个关联此部门的企业群，默认为false即不创建
	CreateDeptGroup *bool `json:"create_dept_group,omitempty"`

	//是否默认同意加入该部门的申请：
	//
	//true：表示加入该部门的申请将默认同意
	//
	//false：表示加入该部门的申请需要有权限的管理员同意
	AutoApproveApply *bool `json:"auto_approve_apply,omitempty"`

	//在父部门中的排序值，order值小的排序靠前
	Order uint `json:"order,omitempty"`

	//部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射
	SourceIdentifier string `json:"source_identifier,omitempty"`

	//当部门群已经创建后，有新人加入部门时是否会自动加入该群：
	//
	//true：自动加入群
	//
	//false：不会自动加入群
	//
	//不传值，则保持不变
	AutoAddUser *bool `json:"auto_add_user,omitempty"`

	//部门的主管userid列表，多个userid之间使用英文逗号分隔
	DeptManagerUseridList string `json:"dept_manager_userid_list,omitempty"`

	deptManagerUseridList []string

	//部门群是否包含子部门：
	//
	//true：包含
	//
	//false：不包含
	//
	//不传值，则保持不变
	GroupContainSubDept *bool `json:"group_contain_sub_dept,omitempty"`

	//部门群是否包含外包部门：
	//
	//true：包含
	//
	//false：不包含
	//
	//不传值，则保持不变
	GroupContainOuterDept *bool `json:"group_contain_outer_dept,omitempty"`

	//部门群是否包含隐藏部门：
	//
	//true：包含
	//
	//false：不包含
	//
	//不传值，则保持不变
	GroupContainHiddenDept *bool `json:"group_contain_hidden_dept,omitempty"`

	//企业群群主的userid
	OrgDeptOwner string `json:"org_dept_owner,omitempty"`

	//强制更新的字段，支持清空指定的字段，多个字段之间使用英文逗号分隔。目前支持字段: dept_manager_userid_list
	ForceUpdateFields string `json:"force_update_fields,omitempty"`

	forceUpdateFields []string
}

func (u *UpdateDept) String() string {
	str, _ := json.Marshal(u)
	return string(str)
}

type updateDeptBuilder struct {
	cd *UpdateDept
}

func NewUpdateDept(Id uint) *updateDeptBuilder {
	return &updateDeptBuilder{cd: &UpdateDept{Id: Id}}
}

func (cdb *updateDeptBuilder) SetName(name string) *updateDeptBuilder {
	cdb.cd.Name = name
	return cdb
}

func (cdb *updateDeptBuilder) SetParentId(pId uint) *updateDeptBuilder {
	cdb.cd.ParentId = &pId
	return cdb
}

func (cdb *updateDeptBuilder) SetHideDept(hide bool) *updateDeptBuilder {
	cdb.cd.HideDept = &hide
	return cdb
}

//SetDeptPermits 指定可以查看本部门的其他部门列表，总数不能超过200,当hide_dept为true时，则此值生效
func (cdb *updateDeptBuilder) SetDeptPermits(deptId int, deptIds ...int) *updateDeptBuilder {
	//if cdb.cd.HideDept != nil && *cdb.cd.HideDept == true {
	cdb.cd.DeptPermit = append(deptIds, deptId)
	//}
	return cdb
}

//SetUserPermits 指定可以查看本部门的人员userid列表，总数不能超过200,当hide_dept为true时，则此值生效
func (cdb *updateDeptBuilder) SetUserPermits(userId string, userIds ...string) *updateDeptBuilder {
	//if cdb.cd.HideDept != nil && *cdb.cd.HideDept == true {
	cdb.cd.UserPermit = append(userIds, userId)
	//}
	return cdb
}

func (cdb *updateDeptBuilder) SetOuterDept(outer bool) *updateDeptBuilder {
	cdb.cd.OuterDept = &outer
	return cdb
}

func (cdb *updateDeptBuilder) SetOuterDeptOnlySelf(self bool) *updateDeptBuilder {
	if cdb.cd.OuterDept != nil && *cdb.cd.OuterDept == true {
		cdb.cd.OuterDeptOnlySelf = &self
	}
	return cdb
}

func (cdb *updateDeptBuilder) SetUserPermitsUserIds(userId string, userIds ...string) *updateDeptBuilder {
	//if cdb.cd.OuterDept != nil && *cdb.cd.OuterDept == true {
	cdb.cd.UserPermitsUserIds = append(userIds, userId)
	//}
	return cdb
}

func (cdb *updateDeptBuilder) SetUserPermitsDeptIds(deptId int, deptIds ...int) *updateDeptBuilder {
	//if cdb.cd.OuterDept != nil && *cdb.cd.OuterDept == true {
	cdb.cd.UserPermitsDeptIds = append(deptIds, deptId)
	//}
	return cdb
}

func (cdb *updateDeptBuilder) SetCreateDeptGroup(group bool) *updateDeptBuilder {
	cdb.cd.CreateDeptGroup = &group
	return cdb
}
func (cdb *updateDeptBuilder) SetAutoApproveApply(approve bool) *updateDeptBuilder {
	cdb.cd.AutoApproveApply = &approve
	return cdb
}

func (cdb *updateDeptBuilder) SetOrder(order uint) *updateDeptBuilder {
	cdb.cd.Order = order
	return cdb
}

func (cdb *updateDeptBuilder) SetSourceIdentifier(id string) *updateDeptBuilder {
	cdb.cd.SourceIdentifier = id
	return cdb
}

func (cdb *updateDeptBuilder) SetAutoAddUser(auto bool) *updateDeptBuilder {
	cdb.cd.AutoAddUser = &auto
	return cdb
}

func (cdb *updateDeptBuilder) SetDeptManagerUseridList(userId string, userIds ...string) *updateDeptBuilder {
	cdb.cd.deptManagerUseridList = append(userIds, userId)
	return cdb
}

func (cdb *updateDeptBuilder) Build() *UpdateDept {
	cd := cdb.cd

	ds := removeIntDuplicates(cd.DeptPermit)
	us := removeStringDuplicates(cd.UserPermit)

	deptIds := removeIntDuplicates(cd.UserPermitsDeptIds)
	userIds := removeStringDuplicates(cd.UserPermitsUserIds)

	cd.DeptPermit = ds
	cd.UserPermit = us

	cd.UserPermitsDeptIds = deptIds
	cd.UserPermitsUserIds = userIds

	if cd.HideDept != nil && *cd.HideDept == true {
		cd.DeptPermits = strings.Join(removeIntDuplicatesToString(ds), ",")
		cd.UserPermits = strings.Join(us, ",")
	}
	if cd.OuterDept != nil && *cd.OuterDept == true {
		cd.UserPermitsDepts = strings.Join(removeIntDuplicatesToString(deptIds), ",")
		cd.UserPermitsUsers = strings.Join(userIds, ",")
	}
	return cd
}
