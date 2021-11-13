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

package response

type GetExtContact struct {
	Response

	Contacts []extContact `json:"results"`
}

type extContact struct {
	UserId string `json:"userid"`

	//邮箱
	Email string `json:"email"`

	//职位
	Title string `json:"title"`

	//标签列表
	Labels []int `json:"label_ids"`

	//共享给的部门ID
	ShareDept []int `json:"share_dept_ids"`

	//地址
	Address string `json:"address"`

	//备注
	Remark string `json:"remark"`

	//负责人的userId
	FollowerUser string `json:"follower_user_id"`

	//外部联系人的姓名
	Name string `json:"name"`

	//手机号国家码
	StateCode string `json:"state_code"`

	//外部联系人的企业名称
	CompanyName string `json:"company_name"`

	//共享给的员工userid列表
	ShareUser []string `json:"share_user_ids"`

	//外部联系人的手机号
	Mobile string `json:"mobile"`
}
