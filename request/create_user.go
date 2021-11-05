package request

import (
	"encoding/json"
	"strings"
)

type CreateUser struct {
	UserId string `json:"userid,omitempty" validate:"omitempty,min=1,max=64"`

	Name string `json:"name" validate:"required,min=1,max=80"`

	Mobile string `json:"mobile" validate:"required"`

	HideMobile bool `json:"hide_mobile,omitempty"`

	Telephone string `json:"telephone,omitempty" validate:"omitempty,min=1,max=50"`

	JobNumber string `json:"job_number,omitempty" validate:"omitempty,min=1,max=50"`

	ManagerUserId string `json:"manager_userid,omitempty"`

	Title string `json:"title,omitempty" validate:"omitempty,max=200"`

	Email string `json:"email,omitempty" validate:"omitempty,max=50"`

	OrgEmail string `json:"org_email,omitempty" validate:"omitempty,max=100"`

	OrgEmailType string `json:"org_email_type,omitempty" validate:"omitempty,oneof=profession base"`

	WorkPlace string `json:"work_place,omitempty" validate:"omitempty,max=100"`

	Remark string `json:"remark,omitempty" validate:"omitempty,max=2000"`

	DeptIdList string `json:"dept_id_list" validate:"required"`

	deptIds []int

	DeptOrders []DeptOrder `json:"dept_order_list,omitempty"`

	DeptTitles []deptTitle `json:"dept_title_list,omitempty"`

	Extension string `json:"extension,omitempty" validate:"omitempty,max=2000"`

	SeniorMode bool `json:"senior_mode,omitempty"`

	HiredDate int `json:"hired_date,omitempty"`

	LoginEmail string `json:"login_email,omitempty"`

	ExclusiveAccount bool `json:"exclusive_account,omitempty"`

	ExclusiveAccountType string `json:"exclusive_account_type,omitempty" validate:"omitempty,oneof=sso dingtalk"`

	LoginId string `json:"loginId,omitempty"`

	Password string `json:"init_password,omitempty"`
}

func (cu *CreateUser) String() string {
	str, _ := json.Marshal(cu)
	return string(str)
}

//DeptOrder 员工在对应的部门中的排序
type DeptOrder struct {
	DeptId int `json:"dept_id"`

	Order int `json:"order"`
}

//deptTitle 员工在对应的部门中的职位
type deptTitle struct {
	DeptId int `json:"dept_id"`

	Title string `json:"title"`
}

func newDeptOrder(dept, order int) DeptOrder {
	return DeptOrder{dept, order}
}

func newDeptTitle(dept int, title string) deptTitle {
	return deptTitle{dept, title}
}

type createUserBuilder struct {
	user *CreateUser
}

func NewCreateUser(name, mobile string, dept int, deptIds ...int) *createUserBuilder {
	ub := &createUserBuilder{user: &CreateUser{Name: name, Mobile: mobile}}
	ub.user.deptIds = append(ub.user.deptIds, dept)
	ub.user.deptIds = append(ub.user.deptIds, deptIds...)
	return ub
}

// SetName 员工名称，长度最大80个字符。
func (ub *createUserBuilder) SetName(name string) *createUserBuilder {
	ub.user.Name = name
	return ub
}

// SetMobile 手机号码，企业内必须唯一，不可重复,如果是国际号码，请使用+xx-xxxxxx的格式。
func (ub *createUserBuilder) SetMobile(mobile string) *createUserBuilder {
	ub.user.Mobile = mobile
	return ub
}

// SetHideMobile 是否号码隐藏： 隐藏手机号后，手机号在个人资料页隐藏，但仍可对其发DING、发起钉钉免费商务电话。
func (ub *createUserBuilder) SetHideMobile(hide bool) *createUserBuilder {
	ub.user.HideMobile = hide
	return ub
}

// SetTelephone 分机号，长度最大50个字符。企业内必须唯一，不可重复
func (ub *createUserBuilder) SetTelephone(tel string) *createUserBuilder {
	ub.user.Telephone = tel
	return ub
}

// SetJobNumber 员工工号，长度最大为50个字符
func (ub *createUserBuilder) SetJobNumber(job string) *createUserBuilder {
	ub.user.JobNumber = job
	return ub
}

// SetManagerUserId 直属主管的userId。
func (ub *createUserBuilder) SetManagerUserId(managerUserId string) *createUserBuilder {
	ub.user.ManagerUserId = managerUserId
	return ub
}

// SetTitle 职位，长度最大为200个字符。
func (ub *createUserBuilder) SetTitle(title string) *createUserBuilder {
	ub.user.Title = title
	return ub
}

// SetEmail 员工邮箱，长度最大50个字符。企业内必须唯一，不可重复。
func (ub *createUserBuilder) SetEmail(email string) *createUserBuilder {
	ub.user.Email = email
	return ub
}

// SetOrgEmail 员工的企业邮箱，长度最大100个字符。员工的企业邮箱已开通，才能增加此字段。
func (ub *createUserBuilder) SetOrgEmail(orgEmail string) *createUserBuilder {
	ub.user.OrgEmail = orgEmail
	return ub
}

// SetOrgEmailType 员工的企业邮箱类型，仅对支持的组织生效（profession: 标准版，base: 基础版）
func (ub *createUserBuilder) SetOrgEmailType(orgEmailType string) *createUserBuilder {
	ub.user.OrgEmailType = orgEmailType
	return ub
}

// SetRemark 备注，长度最大2000个字符。
func (ub *createUserBuilder) SetRemark(remark string) *createUserBuilder {
	ub.user.Remark = remark
	return ub
}

//SetWorkPlace 办公地点，长度最大100个字符。
func (ub *createUserBuilder) SetWorkPlace(workPlace string) *createUserBuilder {
	ub.user.WorkPlace = workPlace
	return ub
}

// SetDeptOrder 设置部门排序
func (ub *createUserBuilder) SetDeptOrder(dept, order int) *createUserBuilder {
	ub.user.DeptOrders = append(ub.user.DeptOrders, newDeptOrder(dept, order))
	return ub
}

// SetDeptTitle 员工在对应的部门中的职位
func (ub *createUserBuilder) SetDeptTitle(dept int, title string) *createUserBuilder {
	ub.user.DeptTitles = append(ub.user.DeptTitles, newDeptTitle(dept, title))
	return ub
}

//SetExtension 扩展属性，可以设置多种属性，最大长度2000个字符。
func (ub *createUserBuilder) SetExtension(ext string) *createUserBuilder {
	ub.user.Extension = ext
	return ub
}

//SetSeniorMode 是否开启高管模式：
//true：开启。 开启后，手机号码对所有员工隐藏。普通员工无法对其发DING、发起钉钉免费商务电话。高管之间不受影响。
//false：不开启。
func (ub *createUserBuilder) SetSeniorMode(senior bool) *createUserBuilder {
	ub.user.SeniorMode = senior
	return ub
}

//SetHiredDate 入职时间，Unix时间戳，单位毫秒
func (ub *createUserBuilder) SetHiredDate(hireDate int) *createUserBuilder {
	ub.user.HiredDate = hireDate
	return ub
}

//SetLoginEmail 登录邮箱
func (ub *createUserBuilder) SetLoginEmail(loginEmail string) *createUserBuilder {
	ub.user.LoginEmail = loginEmail
	return ub
}

//SetExclusiveAccount 是否专属帐号。 为true时，不能指定loginEmail或mobile）
func (ub *createUserBuilder) SetExclusiveAccount(exclusive bool) *createUserBuilder {
	ub.user.ExclusiveAccount = exclusive
	return ub
}

func (ub *createUserBuilder) SetSetExclusiveAccountType(exclusiveType string) *createUserBuilder {
	ub.user.ExclusiveAccountType = exclusiveType
	return ub
}

func (ub *createUserBuilder) SetLoginId(loginId string) *createUserBuilder {
	ub.user.LoginId = loginId
	return ub
}

func (ub *createUserBuilder) SetPassword(pwd string) *createUserBuilder {
	ub.user.Password = pwd
	return ub
}

func (ub *createUserBuilder) Build() *CreateUser {
	ub.user.DeptIdList = strings.Join(removeIntDuplicates(ub.user.deptIds), ",")
	return ub.user
}
