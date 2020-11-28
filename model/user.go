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

type User struct {
	Id              string      `json:"userid,omitempty" validate:"omitempty,max=64,min=1"`    //员工唯一标识ID（不可修改），企业内必须唯一。长度为1~64个字符，如果不传，将自动生成一个userid。
	UnionId         string      `json:"unionid,omitempty"`                                     //员工在当前开发者企业账号范围内的唯一标识，系统生成，固定值，不会改变。
	OpenId          string      `json:"OpenId,omitempty"`                                      //openId
	Name            string      `json:"name,omitempty" validate:"omitempty,max=64"`            //员工姓名，长度最大64个字符。
	Email           string      `json:"email,omitempty" validate:"omitempty,max=64,email"`     //员工邮箱，长度最大64个字符。企业内必须唯一，不可重复。
	Mobile          string      `json:"mobile,omitempty"`                                      //手机号码，企业内必须唯一，不可重复。
	Position        string      `json:"position,omitempty" validate:"omitempty,max=64"`        //职位信息。长度为0~64个字符
	Roles           []UserRoles `json:"roles,omitempty"`                                       //用户所在角色列表。
	OrgEmail        string      `json:"orgEmail,omitempty" validate:"omitempty,email"`         //员工的企业邮箱，如果员工已经开通了企业邮箱，接口会返回，否则会报错。
	Avatar          string      `json:"avatar,omitempty"`                                      //头像
	Senior          bool        `json:"isSenior,omitempty"`                                    //是否开启高管模式
	Authed          bool        `json:"realAuthed,omitempty"`                                  //是否实名认证：true：是false：否
	Boss            bool        `json:"isBoss,omitempty"`                                      //是否为企业的老板
	Active          bool        `json:"active,omitempty"`                                      //是否激活
	Hide            bool        `json:"isHide,omitempty"`                                      //是否号码隐藏：管模式
	Admin           bool        `json:"isAdmin,omitempty"`                                     //是否为企业的管理员
	JobNumber       string      `json:"jobnumber,omitempty" validate:"omitempty,max=64"`       //员工工号，对应显示到OA后台和客户端个人资料的工号栏目。长度为0~64个字符。码隐藏：管模式
	Remark          string      `json:"remark,omitempty" validate:"omitempty,max=1000"`        //备注，长度最大为1024个字符。
	WorkPlace       string      `json:"workPlace,omitempty" validate:"omitempty,max=50"`       //办公地点。长度为0~50个字符。
	Tel             string      `json:"tel,omitempty" validate:"omitempty,max=50"`             //分机号。长度为0~50个字符，企业内必须唯一，不可重复。
	PositionInDepts string      `json:"positionInDepts,omitempty"`                             //设置用户在每个部门下的职位。Key是deptId，表示部门；Value是职位，表示在这个部门下的职位。
	LeaderInDepts   string      `json:"isLeaderInDepts,omitempty"`                             //在部门里是否主管
	OrderInDepts    string      `json:"orderInDepts,omitempty"`                                //在对应的部门中的排序，Map结构的json字符串。Key是部门的ID，Value是人员在这个部门的排序值。
	Department      []int       `json:"department,omitempty"`                                  //数组类型，数组里面值为整型，成员所属部门ID列表。
	Extattr         string      `json:"extattr,omitempty"`                                     //扩展属性，可以设置多种属性
	HiredDate       int         `json:"hiredDate,omitempty"`                                   //入职时间，Unix时间戳，单位毫秒。
	Lang            string      `json:"lang,omitempty" validate:"omitempty,oneof=zh_CN en_US"` //通讯录语言，默认zh_CN。如果是英文，请输入en_US。
	StateCode       string      `json:"stateCode,omitempty"`                                   //国家地区码
	ManagerUserId   string      `json:"managerUserId,omitempty"`                               //主管的ID
}

//UserRoles:用户角色
type UserRoles struct {
	Id    int    `json:"id,omitempty"`        //角色id
	Name  string `json:"name,omitempty"`      //角色名称
	Group string `json:"groupName,omitempty"` //角色组
}

type userv2 struct {
	Id         string `json:"userid,omitempty" validate:"omitempty,max=64,min=1"` // 员工唯一标识ID（不可修改），企业内必须唯一。长度为1~64个字符，如果不传，将自动生成一个userid。
	Name       string `json:"name,omitempty" validate:"omitempty,max=64"`         //员工姓名，长度最大64个字符。
	Mobile     string `json:"mobile,omitempty"`                                   //手机号码，企业内必须唯一，不可重复。
	HideMobile bool   `json:"hide_mobile,omitempty"`                              //是否隐藏手机号
	Telephone  string `json:"telephone,omitempty" validate:"omitempty,max=50"`    //分机号，长度最大50个字符。企业内必须唯一，不可重复。
	JobNumber  string `json:"job_number,omitempty" validate:"omitempty,max=64"`   //员工工号，长度最大为64个字符。
	Title      string `json:"title,omitempty" validate:"omitempty,max=300"`       //职位，长度最大为300个字符
	Email      string `json:"email,omitempty" validate:"omitempty,max=64,email"`  //员工邮箱，长度最大64个字符。企业内必须唯一，不可重复。
	Remark     string `json:"remark,omitempty" validate:"omitempty,max=1024"`     //备注，长度最大为1024个字符。
	DeptIdList string `json:"dept_id_list,omitempty"`                             //所属部门ID列表。
	Extension  string `json:"extension,omitempty"`                                //扩展属性，可以设置多种属性。手机上最多只能显示10个扩展属性，
	SeniorMode bool   `json:"senior_mode,omitempty"`                              //是否开启高管模式，开启后，手机号码对所有员工隐藏。普通员工无法对其发DING、发起钉钉免费商务电话。高管之间不受影响。
	HiredDate  int    `json:"hired_date,omitempty"`                               //入职时间，Unix时间戳，单位毫秒。
}

//InactiveUser:未登录用户数据。
type InactiveUser struct {
	More    bool     `json:"has_more"` //是否还有更多
	UserIds []string `json:"list"`     //用户userid
}

//OrgAdminList:管理员列表
type OrgAdminList struct {
	Level  int    `json:"sys_level"` //1：表示主管理员,2：表示子管理员
	UserId string `json:"userid"`    //用户userid
}

func NewUser() *userv2 {
	return &userv2{}
}

//请求参数验证
func (u userv2) Validate(valid *validator.Validate, trans translator.Translator) error {
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

//请求参数验证
func (u User) Validate(valid *validator.Validate, trans translator.Translator) error {
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
