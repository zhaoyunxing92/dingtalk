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
	"encoding/json"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/language"
	"strings"
)

type UpdateUser struct {
	//员工的userid
	UserId string `json:"userid" validate:"required"`

	//员工名称，长度最大80个字符
	Name string `json:"name,omitempty"`

	//手机号码，企业内必须唯一，不可重复,如果是国际号码，请使用+xx-xxxxxx的格式
	Mobile string `json:"mobile,omitempty"`

	//是否号码隐藏： true：隐藏 隐藏手机号后，手机号在个人资料页隐藏，但仍可对其发DING、发起钉钉免费商务电话。 false：不隐藏
	HideMobile *bool `json:"hide_mobile,omitempty"`

	//分机号，长度最大50个字符。企业内必须唯一，不可重复
	Telephone string `json:"telephone,omitempty" validate:"omitempty,min=1,max=50"`

	//员工工号，长度最大50个字符
	JobNumber string `json:"job_number,omitempty" validate:"omitempty,min=1,max=50"`

	//直属主管的userId
	ManagerUserId string `json:"manager_userid,omitempty"`

	//职位，长度最大200个字符
	Title string `json:"title,omitempty" validate:"omitempty,max=200"`

	//员工邮箱，长度最大50个字符。企业内必须唯一，不可重复
	Email string `json:"email,omitempty" validate:"omitempty,max=50"`

	//员工的企业邮箱。员工的企业邮箱已开通，才能增加此字段
	OrgEmail string `json:"org_email,omitempty" validate:"omitempty,max=100"`

	//员工的企业邮箱类型，仅对支持的组织生效（profession: 标准版，base: 基础版）
	OrgEmailType string `json:"org_email_type,omitempty" validate:"omitempty,oneof=profession base"`

	//办公地点，长度最大100个字符
	WorkPlace string `json:"work_place,omitempty" validate:"omitempty,max=100"`

	//Remark 备注，长度最大2000个字符
	Remark string `json:"remark,omitempty" validate:"omitempty,max=2000"`

	//所属部门ID列表
	DeptIdList string `json:"dept_id_list,omitempty"`

	deptIds []int

	//员工在对应的部门中的排序
	DeptOrders []DeptOrder `json:"dept_order_list,omitempty"`

	//员工在对应的部门中的职位
	DeptTitles []deptTitle `json:"dept_title_list,omitempty"`

	//扩展属性，长度最大2000个字符
	Extension string `json:"extension,omitempty" validate:"omitempty,max=2000"`

	//是否开启高管模式.开启后，手机号码对所有员工隐藏。普通员工无法对其发DING、发起钉钉免费商务电话。高管之间不受影响
	SeniorMode *bool `json:"senior_mode,omitempty"`

	//入职时间，UNIX时间戳，单位毫秒
	HiredDate int `json:"hired_date,omitempty"`

	//通讯录语言，默认zh_CN。如果是英文，请传入en_US。
	Language string `json:"language,omitempty" validate:"omitempty,oneof=zh_CN en_US"`

	//强制更新的字段，支持清空指定的字段，多个字段之间使用逗号分隔。目前支持字段: manager_userid
	ForceUpdateFields string `json:"force_update_fields,omitempty"`

	//强制更新的字段，支持清空指定的字段，多个字段之间使用逗号分隔。目前支持字段: manager_userid
	forceUpdateFields []string

	//LoginId 专属帐号登录名。
	LoginId string `json:"loginId,omitempty"`
}

func (u *UpdateUser) String() string {
	str, _ := json.Marshal(u)
	return string(str)
}

type updateUserBuilder struct {
	user *UpdateUser
}

func NewUpdateUser(userId string) *updateUserBuilder {
	return &updateUserBuilder{user: &UpdateUser{UserId: userId}}
}

// SetName 员工名称，长度最大80个字符。
func (ub *updateUserBuilder) SetName(name string) *updateUserBuilder {
	ub.user.Name = name
	return ub
}

// SetMobile 手机号码，企业内必须唯一，不可重复,如果是国际号码，请使用+xx-xxxxxx的格式。
func (ub *updateUserBuilder) SetMobile(mobile string) *updateUserBuilder {
	ub.user.Mobile = mobile
	return ub
}

// SetHideMobile 是否号码隐藏： 隐藏手机号后，手机号在个人资料页隐藏，但仍可对其发DING、发起钉钉免费商务电话。
func (ub *updateUserBuilder) SetHideMobile(hide bool) *updateUserBuilder {
	ub.user.HideMobile = &hide
	return ub
}

// SetTelephone 分机号，长度最大50个字符。企业内必须唯一，不可重复
func (ub *updateUserBuilder) SetTelephone(tel string) *updateUserBuilder {
	ub.user.Telephone = tel
	return ub
}

// SetJobNumber 员工工号，长度最大为50个字符
func (ub *updateUserBuilder) SetJobNumber(job string) *updateUserBuilder {
	ub.user.JobNumber = job
	return ub
}

// SetManagerUserId 直属主管的userId。
func (ub *updateUserBuilder) SetManagerUserId(managerUserId string) *updateUserBuilder {
	ub.user.ManagerUserId = managerUserId
	return ub
}

// SetTitle 职位，长度最大为200个字符。
func (ub *updateUserBuilder) SetTitle(title string) *updateUserBuilder {
	ub.user.Title = title
	return ub
}

// SetEmail 员工邮箱，长度最大50个字符。企业内必须唯一，不可重复。
func (ub *updateUserBuilder) SetEmail(email string) *updateUserBuilder {
	ub.user.Email = email
	return ub
}

// SetOrgEmail 员工的企业邮箱，长度最大100个字符。员工的企业邮箱已开通，才能增加此字段。
func (ub *updateUserBuilder) SetOrgEmail(orgEmail string) *updateUserBuilder {
	ub.user.OrgEmail = orgEmail
	return ub
}

// SetOrgEmailType 员工的企业邮箱类型，仅对支持的组织生效（profession: 标准版，base: 基础版）
func (ub *updateUserBuilder) SetOrgEmailType(orgEmailType string) *updateUserBuilder {
	ub.user.OrgEmailType = orgEmailType
	return ub
}

//SetRemark 备注，长度最大2000个字符。
func (ub *updateUserBuilder) SetRemark(remark string) *updateUserBuilder {
	ub.user.Remark = remark
	return ub
}

//SetWorkPlace 办公地点，长度最大100个字符。
func (ub *updateUserBuilder) SetWorkPlace(workPlace string) *updateUserBuilder {
	ub.user.WorkPlace = workPlace
	return ub
}

//SetDeptOrder 设置部门排序
func (ub *updateUserBuilder) SetDeptOrder(dept, order int) *updateUserBuilder {
	ub.user.DeptOrders = append(ub.user.DeptOrders, newDeptOrder(dept, order))
	return ub
}

//SetDeptTitle 员工在对应的部门中的职位
func (ub *updateUserBuilder) SetDeptTitle(dept int, title string) *updateUserBuilder {
	ub.user.DeptTitles = append(ub.user.DeptTitles, newDeptTitle(dept, title))
	return ub
}

//SetExtension 扩展属性，可以设置多种属性，最大长度2000个字符。
func (ub *updateUserBuilder) SetExtension(ext string) *updateUserBuilder {
	ub.user.Extension = ext
	return ub
}

//SetSeniorMode 是否开启高管模式：
//true：开启。 开启后，手机号码对所有员工隐藏。普通员工无法对其发DING、发起钉钉免费商务电话。高管之间不受影响。
//false：不开启。
func (ub *updateUserBuilder) SetSeniorMode(senior bool) *updateUserBuilder {
	ub.user.SeniorMode = &senior
	return ub
}

//SetHiredDate 入职时间，Unix时间戳，单位毫秒
func (ub *updateUserBuilder) SetHiredDate(hireDate int) *updateUserBuilder {
	ub.user.HiredDate = hireDate
	return ub
}

//SetLanguage 通讯录语言，默认zh_CN。如果是英文，请传入en_US。
func (ub *updateUserBuilder) SetLanguage(language language.Language) *updateUserBuilder {
	ub.user.Language = string(language)
	return ub
}

//SetForceUpdateFields 通讯录语言，默认zh_CN。如果是英文，请传入en_US。
func (ub *updateUserBuilder) SetForceUpdateFields(field string, fields ...string) *updateUserBuilder {
	updateFields := ub.user.forceUpdateFields
	updateFields = append(updateFields, strings.Split(field, ",")...)
	ub.user.forceUpdateFields = append(updateFields, fields...)
	return ub
}

//SetLoginId 专属帐号登录名。
func (ub *updateUserBuilder) SetLoginId(loginId string) *updateUserBuilder {
	ub.user.LoginId = loginId
	return ub
}

//SetDept 所属部门id，可通过获取部门列表接口获取
func (ub *updateUserBuilder) SetDept(id int, dept ...int) *updateUserBuilder {
	ub.user.deptIds = append(dept, id)
	return ub
}

func (ub *updateUserBuilder) Build() *UpdateUser {
	ub.user.deptIds = removeIntDuplicates(ub.user.deptIds)
	ub.user.DeptIdList = strings.Join(removeIntDuplicatesToString(ub.user.deptIds), ",")
	ub.user.ForceUpdateFields = strings.Join(removeStringDuplicates(ub.user.forceUpdateFields), ",")
	return ub.user
}
