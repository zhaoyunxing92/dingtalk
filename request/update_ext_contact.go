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

import "encoding/json"

//UpdateExtContact 更新外部联系人
type UpdateExtContact struct {
	Contact *updateExtContact `json:"contact" validate:"required"`
}

type updateExtContact struct {
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

	//该外部联系人的userid
	UserId string `json:"user_id,omitempty" validate:"required"`

	//外部联系人的企业名称
	CompanyName string `json:"company_name,omitempty"`

	//共享给的员工userid列表
	ShareUser []string `json:"share_user_ids,omitempty"`
}

func (e *UpdateExtContact) String() string {
	str, _ := json.Marshal(e)
	return string(str)
}

type updateExtContactBuilder struct {
	c *updateExtContact
}

func NewUpdateExtContact(userId, name, follower string, labels ...int) *updateExtContactBuilder {
	return &updateExtContactBuilder{c: &updateExtContact{UserId: userId, Name: name, FollowerUser: follower,
		Labels: labels}}
}

func (ec *updateExtContactBuilder) SetTitle(title string) *updateExtContactBuilder {
	ec.c.Title = title
	return ec
}

func (ec *updateExtContactBuilder) SetShareDept(id int, ids ...int) *updateExtContactBuilder {
	ec.c.ShareDept = append(ids, id)
	return ec
}

func (ec *updateExtContactBuilder) SetAddress(address string) *updateExtContactBuilder {
	ec.c.Address = address
	return ec
}

func (ec *updateExtContactBuilder) SetRemark(remark string) *updateExtContactBuilder {
	ec.c.Remark = remark
	return ec
}

//SetCompanyName 外部联系人的企业名称
func (ec *updateExtContactBuilder) SetCompanyName(name string) *updateExtContactBuilder {
	ec.c.CompanyName = name
	return ec
}

//SetShareUser 共享给的员工userid列表
func (ec *updateExtContactBuilder) SetShareUser(id string, ids ...string) *updateExtContactBuilder {
	ec.c.ShareUser = append(ids, id)
	return ec
}

func (ec *updateExtContactBuilder) Build() *UpdateExtContact {
	ec.c.ShareUser = removeStringDuplicates(ec.c.ShareUser)
	ec.c.ShareDept = removeIntDuplicates(ec.c.ShareDept)
	ec.c.Labels = removeIntDuplicates(ec.c.Labels)
	return &UpdateExtContact{ec.c}
}
