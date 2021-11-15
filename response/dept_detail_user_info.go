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

type DeptDetailUserInfo struct {
	Response
	deptDetailUserInfoPage `json:"result"`
}

type deptDetailUserInfoPage struct {
	//是否还有更多的数据
	HasMore bool `json:"has_more"`

	//下一次分页的游标，如果has_more为false，表示没有更多的分页数据。
	NextCursor int `json:"next_cursor"`

	DeptDetailUsers []deptDetailUserInfo `json:"list"`
}

//deptDetailUserInfo 部门用户详情
type deptDetailUserInfo struct {
	UserId string `json:"userid"`

	//员工在当前开发者企业账号范围内的唯一标识
	UnionId string `json:"unionid"`

	//员工名称
	Name string `json:"name"`

	//头像
	Avatar string `json:"avatar"`

	//国际电话区号
	StateCode string `json:"state_code"`

	//员工的直属主管
	ManagerUserId string `json:"manager_userid"`

	//手机号码
	Mobile string `json:"mobile"`

	//是否号码隐藏
	HideMobile bool `json:"hide_mobile"`

	//分机号
	Telephone string `json:"telephone"`

	//员工工号
	JobNumber string `json:"job_number"`

	//职位
	Title string `json:"title"`

	//办公地点
	WorkPlace string `json:"work_place"`

	//备注
	Remark string `json:"remark"`

	//专属帐号登录名
	LoginId string `json:"loginId"`

	//所属部门ID列表
	DeptIds []int `json:"dept_id_list"`

	//员工在部门中的排序
	DeptOrder int `json:"dept_order"`

	//员工在对应的部门中的排序
	Extension string `json:"extension"`

	//入职时间
	HiredDate int `json:"hired_date"`

	//是否激活了钉钉
	Active bool `json:"active"`

	//是否为企业的管理员：
	//
	//true：是
	//
	//false：不是
	Admin bool `json:"admin"`

	//是否为企业的老板
	Boss bool `json:"boss"`

	//是否专属帐号
	ExclusiveAccount bool `json:"exclusive_account"`

	//是否是部门的主管
	Leader bool `json:"leader"`

	//专属帐号类型：
	//
	//sso：企业自建专属帐号
	//
	//dingtalk：钉钉自建专属帐号
	ExclusiveAccountType string `json:"exclusive_account_type"`
}
