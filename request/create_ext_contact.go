package request

import json "encoding/json"

//CreateExtContact 添加外部联系人
type CreateExtContact struct {
	Contact *contact `json:"contact" validate:"required"`
}

type contact struct {
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

type extContactBuilder struct {
	c *contact
}

func NewExtContact() *extContactBuilder {
	return &extContactBuilder{c: &contact{}}
}

func (ec *extContactBuilder) SetTitle(title string) *extContactBuilder {
	ec.c.Title = title
	return ec
}

func (ec *extContactBuilder) SetLabels(id int, ids ...int) *extContactBuilder {
	labels := ec.c.Labels
	labels = append(labels, id)
	ec.c.Labels = append(labels, ids...)
	return ec
}

func (ec *extContactBuilder) SetShareDept(id int, ids ...int) *extContactBuilder {
	ds := ec.c.ShareDept
	ds = append(ds, id)
	ec.c.ShareDept = append(ds, ids...)
	return ec
}

func (ec *extContactBuilder) SetAddress(address string) *extContactBuilder {
	ec.c.Address = address
	return ec
}

func (ec *extContactBuilder) SetRemark(remark string) *extContactBuilder {
	ec.c.Remark = remark
	return ec
}

//SetFollowerUser 负责人的userId
func (ec *extContactBuilder) SetFollowerUser(userId string) *extContactBuilder {
	ec.c.FollowerUser = userId
	return ec
}

//SetStateCode 手机号国家码
func (ec *extContactBuilder) SetStateCode(stateCode string) *extContactBuilder {
	ec.c.StateCode = stateCode
	return ec
}

//SetCompanyName 外部联系人的企业名称
func (ec *extContactBuilder) SetCompanyName(name string) *extContactBuilder {
	ec.c.CompanyName = name
	return ec
}

//SetShareUser 共享给的员工userid列表
func (ec *extContactBuilder) SetShareUser(id string, ids ...string) *extContactBuilder {
	us := ec.c.ShareUser
	us = append(us, id)
	ec.c.ShareUser = append(us, ids...)
	return ec
}

//SetMobile 外部联系人的手机号
func (ec *extContactBuilder) SetMobile(mobile string) *extContactBuilder {
	ec.c.Mobile = mobile
	return ec
}

func (ec *extContactBuilder) Build() *CreateExtContact {
	ec.c.ShareUser = removeStringDuplicates(ec.c.ShareUser)
	ec.c.ShareDept = removeIntDuplicates(ec.c.ShareDept)
	ec.c.Labels = removeIntDuplicates(ec.c.Labels)
	return &CreateExtContact{ec.c}
}
