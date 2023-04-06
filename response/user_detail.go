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

import (
	"github.com/zhaoyunxing92/dingtalk/v2/request"
)

type UserDetail struct {
	Response
	UserInfoDetail `json:"result"`
}

type UserInfoDetail struct {
	UserId string `json:"userid"`

	// 员工在当前开发者企业账号范围内的唯一标识
	UnionId string `json:"unionid"`

	// 员工名称
	Name string `json:"name"`

	// 头像
	Avatar string `json:"avatar"`

	// 国际电话区号
	StateCode string `json:"state_code"`

	// 员工的直属主管
	ManagerUserId string `json:"manager_userid"`

	// 手机号码
	Mobile string `json:"mobile"`

	// 是否号码隐藏
	HideMobile bool `json:"hide_mobile"`

	// 分机号
	Telephone string `json:"telephone"`

	// 员工工号
	JobNumber string `json:"job_number"`

	// 职位
	Title string `json:"title"`

	// 员工邮箱 (需要开通对应权限才会返回)
	Email string `json:"email,omitempty"`

	// 办公地点
	WorkPlace string `json:"work_place"`

	// 备注
	Remark string `json:"remark"`

	// 专属帐号登录名
	LoginId string `json:"loginId"`

	// 是否专属帐号
	ExclusiveAccount bool `json:"exclusive_account"`

	//专属帐号类型：
	//
	//sso：企业自建专属帐号
	//
	//dingtalk：钉钉自建专属帐号
	ExclusiveAccountType string `json:"exclusive_account_type"`

	// 所属部门ID列表
	DeptIds []int `json:"dept_id_list"`

	// 所属部门ID列表
	DeptOrders []request.DeptOrder `json:"dept_order_list"`

	// 员工在对应的部门中的排序
	Extension string `json:"extension,omitempty" validate:"omitempty"`

	// 入职时间
	HiredDate int `json:"hired_date"`

	// 是否激活了钉钉
	Active bool `json:"active"`

	// 是否完成了实名认证
	RealAuthed bool `json:"real_authed"`

	//员工的企业邮箱
	//
	//如果员工的企业邮箱没有开通，返回信息中不包含该数据
	OrgEmail string `json:"org_email,omitempty" validate:"omitempty,max=100"`

	// 员工的企业邮箱类型
	OrgEmailType string `json:"org_email_type,omitempty" validate:"omitempty,oneof=profession base"`

	// 是否为企业的高管
	Senior bool `json:"senior"`

	//是否为企业的管理员：
	//
	//true：是
	//
	//false：不是
	Admin bool `json:"admin"`

	// 是否为企业的老板
	Boss bool `json:"boss"`

	// 员工在对应的部门中是否领导
	LeaderInDept []LeaderInDept `json:"leader_in_dept"`

	// 角色列表
	UserRoles []struct {
		// 角色id
		Id int `json:"id"`

		// 角色名称
		Name string `json:"name"`

		// 角色组名称
		GroupName string `json:"group_name"`
	} `json:"role_list"`

	// 当用户来自于关联组织时的关联信息
	UnionOrg `json:"union_emp_ext"`
}

type LeaderInDept struct {
	DeptId int `json:"dept_id"`

	Leader bool `json:"leader"`
}

type UnionOrg struct {
	// 员工的userid
	UserId string `json:"userid"`

	// 当前用户所属的组织的企业corpId
	CorpId string `json:"corp_id"`
}

// AssociatedOrg 关联映射关系
type AssociatedOrg struct {
	UnionOrgList []UnionOrg `json:"union_emp_map_list"`
}
