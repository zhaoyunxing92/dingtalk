package model

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

/*
{
	"extension":"{"爱好":"旅游","年龄":"24"}",
	"mobile":"18513027676",
	"telephone":"010-86123456-2345",
	"remark":"备注备注",
	"hide_mobile":"false",
	"title":"技术总监",
	"hired_date":"1597573616828",
	"userid":"zhangsan",
	"work_place":"未来park",
	"dept_title_list":{
		"dept_id":"2",
		"title":"资深产品经理"
	},
	"dept_order_list":{
		"dept_id":"2",
		"order":"1"
	},
	"senior_mode":"false",
	"org_email":"test@xxx.com",
	"name":"张三",
	"dept_id_list":"\"2,3,4\"",
	"job_number":"4",
	"email":"test@xxx.com"
}
*/
type user struct {
	Id         string `json:"userid" validate:"omitempty,max=64,min=1"` // 员工唯一标识ID（不可修改），企业内必须唯一。长度为1~64个字符，如果不传，将自动生成一个userid。
	Name       string `json:"name" validate:"required,max=64"`          //员工姓名，长度最大64个字符。
	Mobile     string `json:"mobile" validate:"required"`               //手机号码，企业内必须唯一，不可重复。
	Hide       bool   `json:"hide_mobile"`                              //是否隐藏手机号
	Telephone  string `json:"telephone" validate:"omitempty,max=50"`    //分机号，长度最大50个字符。企业内必须唯一，不可重复。
	JobNumber  string `json:"job_number" validate:"omitempty,max=64"`   //员工工号，长度最大为64个字符。
	Title      string `json:"title" validate:"omitempty,max=300"`       //职位，长度最大为300个字符
	Email      string `json:"email" validate:"omitempty,max=64,email"`  //员工邮箱，长度最大64个字符。企业内必须唯一，不可重复。
	Remark     string `json:"remark" validate:"omitempty,max=1024"`     //备注，长度最大为1024个字符。
	DeptIdList string `json:"dept_id_list"`                             //所属部门ID列表。
	Extension  string `json:"extension"`                                //扩展属性，可以设置多种属性。手机上最多只能显示10个扩展属性，
	SeniorMode bool   `json:"senior_mode"`                              //是否开启高管模式，开启后，手机号码对所有员工隐藏。普通员工无法对其发DING、发起钉钉免费商务电话。高管之间不受影响。
	HiredDate  int    `json:"hired_date"`                               //入职时间，Unix时间戳，单位毫秒。
}

func NewUser(name, mobile string) *user {
	return &user{Name: name, Mobile: mobile}
}

//请求参数验证
func (u user) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(u); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
