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

import json "encoding/json"

//CreateExtContact 添加外部联系人
type CreateExtContact struct {
	Contact *createExtContact `json:"contact" validate:"required"`
}

type createExtContact struct {
	//职位
	Title string `json:"title,omitempty"`

	//标签列表
	Labels []int `json:"label_ids" validate:"required"`

	//共享给的部门ID
	ShareDept []int `json:"share_dept_ids,omitempty"`

	//地址
	Address string `json:"address,omitempty"`

	//备注
	Remark string `json:"remark,omitempty"`

	//负责人的userId
	FollowerUser string `json:"follower_user_id,omitempty" validate:"required"`

	//外部联系人的姓名
	Name string `json:"name,omitempty" validate:"required"`

	//手机号国家码
	StateCode string `json:"state_code,omitempty" validate:"required"`

	//外部联系人的企业名称
	CompanyName string `json:"company_name,omitempty"`

	//共享给的员工userid列表
	ShareUser []string `json:"share_user_ids,omitempty"`

	//外部联系人的手机号
	Mobile string `json:"mobile,omitempty" validate:"required"`
}

func (e *CreateExtContact) String() string {
	str, _ := json.Marshal(e)
	return string(str)
}

type createExtContactBuilder struct {
	c *createExtContact
}

func NewCreateExtContact(name, mobile, stateCode, follower string, labels ...int) *createExtContactBuilder {

	return &createExtContactBuilder{c: &createExtContact{Mobile: mobile, Name: name, StateCode: stateCode,
		FollowerUser: follower, Labels: labels}}
}

func (ec *createExtContactBuilder) SetTitle(title string) *createExtContactBuilder {
	ec.c.Title = title
	return ec
}

func (ec *createExtContactBuilder) SetShareDept(id int, ids ...int) *createExtContactBuilder {
	ds := ec.c.ShareDept
	ds = append(ds, id)
	ec.c.ShareDept = append(ds, ids...)
	return ec
}

func (ec *createExtContactBuilder) SetAddress(address string) *createExtContactBuilder {
	ec.c.Address = address
	return ec
}

func (ec *createExtContactBuilder) SetRemark(remark string) *createExtContactBuilder {
	ec.c.Remark = remark
	return ec
}

//SetCompanyName 外部联系人的企业名称
func (ec *createExtContactBuilder) SetCompanyName(name string) *createExtContactBuilder {
	ec.c.CompanyName = name
	return ec
}

//SetShareUser 共享给的员工userid列表
func (ec *createExtContactBuilder) SetShareUser(id string, ids ...string) *createExtContactBuilder {
	us := ec.c.ShareUser
	us = append(us, id)
	ec.c.ShareUser = append(us, ids...)
	return ec
}

func (ec *createExtContactBuilder) Build() *CreateExtContact {
	ec.c.ShareUser = removeStringDuplicates(ec.c.ShareUser)
	ec.c.ShareDept = removeIntDuplicates(ec.c.ShareDept)
	ec.c.Labels = removeIntDuplicates(ec.c.Labels)
	return &CreateExtContact{ec.c}
}
