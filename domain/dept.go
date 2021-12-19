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

package domain

import (
	"errors"
	"strconv"
	"strings"
)

import (
	translator "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
)

// Dept:部门实体
type Dept struct {
	Id                    int      `json:"id,omitempty"`                                     // 部门id
	Name                  string   `json:"name,omitempty" validate:"omitempty,max=64,min=1"` // 部门名称。长度限制为1~64个字符，不允许包含字符"-"","以及","。
	ParentId              int      `json:"parentid,omitempty"`                               // 父部门ID，根部门ID为1。
	ParentBalanceFirst    bool     `json:"parentBalanceFirst,omitempty"`                     // 是否优先使用父部门的预算
	ShareBalance          bool     `json:"shareBalance,omitempty"`                           // 是否共享预算
	OuterPermitUsers      []string `json:"-" validate:"omitempty,lte=200"`                   // 指定本部门成员可查看的通讯录用户userid列表，总数不能超过200。当outer_dept为true时，此参数生效。
	StrOuterPermitUsers   string   `json:"outerPermitUsers,omitempty"`                       // 使用|拼接的字符串
	OuterPermitDepts      []int    `json:"-" validate:"omitempty,lte=200"`                   // 指定本部门成员可查看的通讯录部门ID列表。,总数不能超过200.当outer_dept为true时，此参数生效。
	StrOuterPermitDepts   string   `json:"outerPermitDepts,omitempty"`                       // 使用|拼接的字符串
	OuterDept             bool     `json:"outerDept,omitempty"`                              // 是否限制本部门成员查看通讯录,开启后本部门成员只能看到限定范围内的通讯录
	Hide                  bool     `json:"deptHiding,omitempty"`                             // 是否隐藏本部门,隐藏后本部门将不会显示在公司通讯录中
	CreateDeptGroup       bool     `json:"createDeptGroup,omitempty"`                        // 是否创建一个关联此部门的企业群，默认为false即不创建。
	Order                 int      `json:"order,omitempty"`                                  // 在父部门中的排序值，order值小的排序靠前。
	SourceIdentifier      string   `json:"sourceIdentifier,omitempty"`                       // 部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射。
	AutoAddUser           bool     `json:"autoAddUser,omitempty"`                            // 当部门群已经创建后，有新人加入部门时是否会自动加入该群
	GroupContainSubDept   bool     `json:"groupContainSubDept,omitempty"`                    // 部门群是否包含子部门
	OrgDeptOwner          string   `json:"orgDeptOwner,omitempty"`                           // 企业群群主的userid
	DeptGroupChatId       string   `json:"deptGroupChatId,omitempty"`                        // 创建部门群自动生成的id。
	DeptManagerUserIds    []string `json:"-"`                                                // 部门的主管列表，取值为由主管的userid组成的字符串，不同的userid使用"|"符号进行分隔。
	StrDeptManagerUserIds string   `json:"deptManagerUseridList,omitempty"`                  // 字符串拼接
	DeptPermits           []int    `json:"-" validate:"omitempty,lte=200"`                   // 指定可以查看本部门的其他部门列表200,当hide_dept为true时，则此值生效。
	StrDeptPermits        string   `json:"deptPermits,omitempty"`                            // 使用|拼接的字符串
	UserPermits           []string `json:"-" validate:"omitempty,lte=200"`                   // 指定可以查看本部门的人员userid列表，总数不能超过200。当hide_dept为true时，则此值生效。
	StrUserPermits        string   `json:"userPermits,omitempty"`                            // 使用|拼接的字符串
	OuterDeptOnlySelf     bool     `json:"outerDeptOnlySelf,omitempty"`                      // 是否只能看到所在部门及下级部门通讯录,开启后只能看到所在部门及下级部门通讯录,当outer_dept为true时，此参数生效
	Extension             string   `json:"ext,omitempty"`                                    // 扩展字段，JSON格式。
}

// DeptUserIdsResponse:获取部门人员id返回
type DeptUserIdsResponse struct {
	Response
	UserIds []string `json:"userIds,omitempty"` // 用户列表
}

// GetDeptUserDetailResponse:获取部门人员详情反馈
type GetDeptUserDetailResponse struct {
	Response
	More     bool         `json:"hasMore"`  // 在分页查询时返回，代表是否还有下一页更多数据。
	UserList []UserDetail `json:"userlist"` // 成员列表
}

// DeptCreateResponse 创建部门返回
type DeptCreateResponse struct {
	Response
	DeptId int `json:"id"`
}

// DepartmentDetail:部门详情
type DeptDetail struct {
	Response
	Id                    int    `json:"id"`                    // 部门id
	Name                  string `json:"name"`                  // 部门名称。长度限制为1~64个字符，不允许包含字符"-"","以及","。
	ParentId              int    `json:"parentid"`              // 父部门ID，根部门ID为1。
	ParentBalanceFirst    bool   `json:"parentBalanceFirst"`    // 是否优先使用父部门的预算
	ShareBalance          bool   `json:"shareBalance"`          // 是否共享预算
	OuterPermitUsers      string `json:"outerPermitUsers"`      // 使用|拼接的字符串
	OuterPermitDepts      string `json:"outerPermitDepts"`      // 使用|拼接的字符串
	OuterDept             bool   `json:"outerDept"`             // 是否限制本部门成员查看通讯录,开启后本部门成员只能看到限定范围内的通讯录
	Hide                  bool   `json:"deptHiding"`            // 是否隐藏本部门,隐藏后本部门将不会显示在公司通讯录中
	CreateDeptGroup       bool   `json:"createDeptGroup"`       // 是否创建一个关联此部门的企业群，默认为false即不创建。
	Order                 int    `json:"order"`                 // 在父部门中的排序值，order值小的排序靠前。
	SourceIdentifier      string `json:"sourceIdentifier"`      // 部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射。
	AutoAddUser           bool   `json:"autoAddUser"`           // 当部门群已经创建后，有新人加入部门时是否会自动加入该群
	GroupContainSubDept   bool   `json:"groupContainSubDept"`   // 部门群是否包含子部门
	OrgDeptOwner          string `json:"orgDeptOwner"`          // 企业群群主的userid
	DeptGroupChatId       string `json:"deptGroupChatId"`       // 创建部门群自动生成的id。
	DeptManagerUseridList string `json:"deptManagerUseridList"` // 部门的主管列表，取值为由主管的userid组成的字符串，不同的userid使用"|"符号进行分隔。
	DeptPermits           string `json:"deptPermits"`           // 使用|拼接的字符串
	UserPermits           string `json:"userPermits"`           // 使用|拼接的字符串
	OuterDeptOnlySelf     bool   `json:"outerDeptOnlySelf"`     // 是否只能看到所在部门及下级部门通讯录,开启后只能看到所在部门及下级部门通讯录,当outer_dept为true时，此参数生效
	Extension             string `json:"ext"`                   // 扩展字段，JSON格式。
}

type GetSubDeptResponse struct {
	Response
	Depts []DeptDetail `json:"department"`
}

type GetSubDeptIdsResponse struct {
	Response
	SubDeptIds []int `json:"sub_dept_id_list"`
}

type GetParentIdsByUserIdResponse struct {
	Response
	ParentIds [][]int `json:"department"` // 指定员工的部门信息。
}

type GetParentIdsByDeptIdResponse struct {
	Response
	ParentIds []int `json:"parentIds"`
}

// CreateDetailDeptRequest
type CreateDetailDeptRequest interface {
	Request
	JoinOuterPermitUsers()
	JoinOuterPermitDepts()
	JoinDeptPermits()
	JoinUserPermits()
	JoinDeptManagerUserIds()
}

// Validate:参数验证
func (d Dept) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(d); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}

// JoinOuterPermitUsers:组装参数,去除重复
func (d *Dept) JoinOuterPermitUsers() {
	size := len(d.OuterPermitUsers)
	if size <= 0 {
		return
	}
	users := make([]string, 0, size)
	tmp := map[string]int{}
	for idx, item := range d.OuterPermitUsers {
		if _, ok := tmp[item]; !ok {
			tmp[item] = idx
			users = append(users, item)
		}
	}
	d.StrOuterPermitUsers = strings.Join(users, "|")
}

// JoinOuterPermitDepts:组装参数,去除重复
func (d *Dept) JoinOuterPermitDepts() {
	size := len(d.OuterPermitDepts)
	if size <= 0 {
		return
	}
	depts := make([]string, 0, size)
	tmp := map[int]int{}
	for idx, item := range d.OuterPermitDepts {
		if _, ok := tmp[item]; !ok {
			tmp[item] = idx
			depts = append(depts, strconv.Itoa(item))
		}
	}
	d.StrOuterPermitDepts = strings.Join(depts, "|")
}

// JoinDeptPermits:组装参数,去除重复
func (d *Dept) JoinDeptPermits() {
	size := len(d.DeptPermits)
	if size <= 0 {
		return
	}
	depts := make([]string, 0, size)
	tmp := map[int]int{}
	for idx, item := range d.DeptPermits {
		if _, ok := tmp[item]; !ok {
			tmp[item] = idx
			depts = append(depts, strconv.Itoa(item))
		}
	}
	d.StrDeptPermits = strings.Join(depts, "|")
}

// JoinUserPermits:组装参数,去除重复
func (d *Dept) JoinUserPermits() {
	size := len(d.UserPermits)
	if size <= 0 {
		return
	}
	users := make([]string, 0, size)
	tmp := map[string]int{}
	for idx, item := range d.UserPermits {
		if _, ok := tmp[item]; !ok {
			tmp[item] = idx
			users = append(users, item)
		}
	}
	d.StrUserPermits = strings.Join(users, "|")
}

// JoinDeptManagerUserIds:组装参数,去除重复
func (d *Dept) JoinDeptManagerUserIds() {
	size := len(d.DeptManagerUserIds)
	if size <= 0 {
		return
	}
	users := make([]string, 0, size)
	tmp := map[string]int{}
	for idx, item := range d.DeptManagerUserIds {
		if _, ok := tmp[item]; !ok {
			tmp[item] = idx
			users = append(users, item)
		}
	}
	d.StrDeptManagerUserIds = strings.Join(users, "|")
}
