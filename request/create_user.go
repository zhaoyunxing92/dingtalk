package request

import (
	"sort"
	"strconv"
	"strings"
)

type CreateUser struct {

	//UserId 长度为1~64个字符，如果不传，将自动生成一个userid。
	UserId string `json:"userid,omitempty" validate:"omitempty,min=1,max=64"`

	//Name 员工名称，长度最大80个字符。
	Name string `json:"name" validate:"required,min=1,max=80"`

	//Mobile 手机号码，企业内必须唯一，不可重复,如果是国际号码，请使用+xx-xxxxxx的格式。
	Mobile string `json:"mobile" validate:"required"`

	//Hide 是否号码隐藏： 隐藏手机号后，手机号在个人资料页隐藏，但仍可对其发DING、发起钉钉免费商务电话。
	Hide bool `json:"hide_mobile,omitempty"`

	//Telephone 分机号，长度最大50个字符。企业内必须唯一，不可重复
	Telephone string `json:"telephone,omitempty" validate:"omitempty,min=1,max=50"`

	//JobNumber 员工工号，长度最大为50个字符
	JobNumber string `json:"job_number,omitempty" validate:"omitempty,min=1,max=50"`

	//ManagerUserId 直属主管的userId。
	ManagerUserId string `json:"manager_userid,omitempty"`

	//Title 职位，长度最大为200个字符。
	Title string `json:"title,omitempty" validate:"omitempty,max=200"`

	//Email 员工邮箱，长度最大50个字符。企业内必须唯一，不可重复。
	Email string `json:"email,omitempty" validate:"omitempty,max=50"`

	//OrgEmail 员工的企业邮箱，长度最大100个字符。员工的企业邮箱已开通，才能增加此字段。
	OrgEmail string `json:"org_email,omitempty" validate:"omitempty,max=100"`

	//OrgEmailType 员工的企业邮箱类型，仅对支持的组织生效（profession: 标准版，base: 基础版）
	OrgEmailType string `json:"org_email_type,omitempty" validate:"omitempty,oneof=profession base"`

	WorkPlace string `json:"work_place,omitempty" validate:"omitempty,max=100"`

	Remark string `json:"remark,omitempty" validate:"omitempty,max=2000"`

	DeptIdList string `json:"dept_id_list" validate:"required"`

	DeptIds []int `json:"_"`

	DeptOrders []deptOrder `json:"dept_order_list,omitempty"`

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

//deptOrder 员工在对应的部门中的排序
type deptOrder struct {
	DeptId int `json:"dept_id"`
	Order  int `json:"order"`
}

//deptTitle 员工在对应的部门中的职位
type deptTitle struct {
	DeptId int    `json:"dept_id"`
	Title  string `json:"title"`
}

func newDeptOrder(dept, order int) deptOrder {
	return deptOrder{dept, order}
}

func newDeptTitle(dept int, title string) deptTitle {
	return deptTitle{dept, title}
}

type createUserBuilder struct {
	user *CreateUser
}

func NewCreateUser(name, mobile string, dept int, deptIds ...int) *createUserBuilder {
	ub := &createUserBuilder{user: &CreateUser{Name: name, Mobile: mobile}}
	ub.user.DeptIds = append(ub.user.DeptIds, dept)
	ub.user.DeptIds = append(ub.user.DeptIds, deptIds...)
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
	ub.user.DeptIdList = strings.Join(ub.getDeptIds(), ",")
	return ub.user
}

func (ub *createUserBuilder) getDeptIds() (ids []string) {
	deptIds := ub.user.DeptIds
	sort.Ints(deptIds)
	for idx := range deptIds {
		if idx > 1 && deptIds[idx] == deptIds[idx-1] {
			continue
		}
		ids = append(ids, strconv.Itoa(deptIds[idx]))
	}
	return ids
}

//SetDept 所属部门id，可通过获取部门列表接口获取
func (ub *createUserBuilder) setDept(dept int) *createUserBuilder {
	ids := ub.user.DeptIds
	if !ub.hasDept(dept) {
		ids = append(ids, dept)
	}
	ub.user.DeptIds = ids
	return ub
}

func (ub createUserBuilder) hasDept(dept int) bool {
	for _, id := range ub.user.DeptIds {
		if id == dept {
			return true
		}
	}
	return false
}
