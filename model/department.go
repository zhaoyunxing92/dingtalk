package model

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

//Department:部门实体
type Department struct {
	Id                     int      `json:"id,omitempty"`                                              //部门id
	Name                   string   `json:"name,omitempty" validate:"omitempty,max=64,min=1"`          //部门名称。长度限制为1~64个字符，不允许包含字符"-"","以及","。
	ParentId               int      `json:"parent_id,omitempty"`                                       //父部门ID，根部门ID为1。
	SourceIdentifier       string   `json:"source_identifier,omitempty"`                               //部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射。
	Hide                   bool     `json:"hide_dept,omitempty"`                                       //是否隐藏本部门,隐藏后本部门将不会显示在公司通讯录中
	DeptPermits            []int    `json:"dept_permits,omitempty" validate:"omitempty,lte=200"`       //指定可以查看本部门的其他部门列表200,当hide_dept为true时，则此值生效。
	UserPermits            []string `json:"user_permits,omitempty" validate:"omitempty,lte=200"`       //指定可以查看本部门的人员userid列表，总数不能超过200。当hide_dept为true时，则此值生效。
	OuterDept              bool     `json:"outer_dept,omitempty"`                                      //是否限制本部门成员查看通讯录,开启后本部门成员只能看到限定范围内的通讯录
	OuterDeptOnlySelf      bool     `json:"outer_dept_only_self,omitempty"`                            //是否只能看到所在部门及下级部门通讯录,开启后只能看到所在部门及下级部门通讯录,当outer_dept为true时，此参数生效
	OuterPermitUsers       []string `json:"outer_permit_users,omitempty"`                              //指定本部门成员可查看的通讯录用户userid列表，总数不能超过200。当outer_dept为true时，此参数生效。
	OuterPermitDepts       []int    `json:"outer_permit_depts,omitempty"`                              //指定本部门成员可查看的通讯录部门ID列表。当outer_dept为true时，此参数生效。
	CreateDeptGroup        bool     `json:"create_dept_group,omitempty"`                               //是否创建一个关联此部门的企业群，默认为false即不创建。
	Order                  int      `json:"order,omitempty"`                                           //在父部门中的排序值，order值小的排序靠前。
	Extension              string   `json:"extension,omitempty"`                                       //扩展字段，JSON格式。
	Language               string   `json:"language,omitempty" validate:"omitempty,oneof=zh_CN en_US"` //通讯录语言。
	AutoAddUser            bool     `json:"auto_add_user,omitempty"`                                   //当部门群已经创建后，有新人加入部门时是否会自动加入该群
	GroupContainSubDept    bool     `json:"group_contain_sub_dept,omitempty"`                          //部门群是否包含子部门
	GroupContainHiddenDept bool     `json:"group_contain_hidden_dept,omitempty"`                       //部门群是否包含隐藏部门
	DeptManagerUseridList  []string `json:"dept_manager_userid_list,omitempty"`                        //部门的主管userid列表
	OrgDeptOwner           string   `json:"org_dept_owner,omitempty"`                                  //企业群群主的userid
}

//DepartmentDetail:部门详情
type DepartmentDetail struct {
	Name  string `json:"name,omitempty"`  //部门名称
	Nick  string `json:"nick,omitempty"`  //部门别名。
	Chain []int  `json:"chain,omitempty"` //节点链。从顶层节点到当前节点的中间节点，其内容不包含当前节点。

}

//Validate:参数验证
func (d Department) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(d); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
