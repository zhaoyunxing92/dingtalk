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

type DeptDetail struct {
	Response

	deptDetail `json:"result"`
}

type deptDetail struct {
	//部门id
	Id int `json:"dept_id"`

	Name string `json:"name"`

	ParentId int `json:"parent_id"`

	//部门标识字段
	SourceIdentifier string `json:"source_identifier"`

	//是否同步创建一个关联此部门的企业群：
	//
	//true：创建
	//
	//false：不创建
	CreateDeptGroup bool `json:"create_dept_group"`

	//当部门群已经创建后，是否有新人加入部门会自动加入该群：
	//
	//true：自动加入群
	//
	//false：不会自动加入群
	AutoAddUser bool `json:"auto_add_user"`

	//是否默认同意加入该部门的申请：
	//
	//true：表示加入该部门的申请将默认同意
	//
	//false：表示加入该部门的申请需要有权限的管理员同意
	AutoApproveApply bool `json:"auto_approve_apply"`

	//部门是否来自关联组织：
	//
	//true：是
	//
	//false：不是
	FromUnionOrg bool `json:"from_union_org"`

	//教育部门标签：
	//
	//campus：校区
	//
	//period：学段
	//
	//grade：年级
	//
	//class：班级
	Tags string `json:"tags"`

	//在父部门中的排序值，order值小的排序靠前
	Order int `json:"order"`

	//部门群ID
	DeptGroupChatId bool `json:"dept_group_chat_id"`

	//
	//true：包含
	//
	//false：不包含
	//
	//不传值，则保持不变
	GroupContainSubDept bool `json:"group_contain_sub_dept"`

	//企业群群主的userid
	OrgDeptOwner string `json:"org_dept_owner"`

	//部门的主管userid列表，多个userid之间使用英文逗号分隔
	DeptManagerUseridList []string `json:"dept_manager_userid_list"`

	//是否限制本部门成员查看通讯录：
	//
	//true：开启限制。开启后本部门成员只能看到限定范围内的通讯录
	//
	//false（默认值）：不限制
	OuterDept bool `json:"outer_dept"`

	//指定本部门成员可查看的通讯录部门ID列表，总数不能超过200
	//
	//当outer_dept为true时，此参数生效
	UserPermitsDeptIds []int `json:"outer_permit_depts"`

	//指定本部门成员可查看的通讯录用户userid列表，总数不能超过200
	//
	//当outer_dept为true时，此参数生效。
	UserPermitsUsers []string `json:"outer_permit_users"`

	//是否隐藏本部门：
	//
	//true：隐藏部门，隐藏后本部门将不会显示在公司通讯录中
	//
	//false（默认值）：显示部门
	HideDept *bool `json:"hide_dept"`

	//指定可以查看本部门的人员userid列表，总数不能超过200
	//
	//当hide_dept为true时，则此值生效
	UserPermits []string `json:"user_permits"`

	//指定可以查看本部门的其他部门列表，总数不能超过200
	//
	//当hide_dept为true时，则此值生效
	DeptPermits []int `json:"dept_permits"`
}
