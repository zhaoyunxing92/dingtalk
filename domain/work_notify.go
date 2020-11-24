package domain

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

type WorkNotify struct {
	AgentId    int         `json:"agent_id"` //发送消息时使用的微应用的ID。
	UserIds    []string    `json:"-" validate:"lte=100"`
	StrUserIds string      `json:"userid_list,omitempty"`  //接收者的企业内部用户的userId列表，最大用户列表长度100。
	StrDeptIds string      `json:"dept_id_list,omitempty"` //接收者的部门id列表，最大列表长度20。接收者是部门Id下包括子部门下的所有用户。
	DeptIds    []int       `json:"-" validate:"lte=20"`
	All        bool        `json:"to_all_user"` //是否发送给企业全部用户。当设置为false时必须指定userid_list或dept_id_list其中一个参数的值。
	Msg        interface{} `json:"msg" validate:"required"`
}

func NewOAWorkNotify(oa OA) *WorkNotify {
	return &WorkNotify{Msg: newOaMessage(oa)}
}

func NewTextWorkNotify(text string) *WorkNotify {
	return &WorkNotify{Msg: newTextMessages(text)}
}

//组装部门，移除重复部门
func (notify *WorkNotify) AssembleDept() {
	var size = len(notify.DeptIds)
	if size == 0 {
		return
	}
	depts := make([]string, 0, size)
	temp := map[int]int{}
	for idx, id := range notify.DeptIds {
		if _, ok := temp[id]; !ok {
			temp[id] = idx
			depts = append(depts, strconv.Itoa(id))
		}
	}
	//部门id字符串拼接
	notify.StrDeptIds = strings.Join(depts, ",")
}

//组装用户,移除重复数据
func (notify *WorkNotify) AssembleUser() {
	var size = len(notify.UserIds)
	if size == 0 {
		return
	}
	users := make([]string, 0, size)
	temp := map[string]int{}

	for i, id := range notify.UserIds {
		if _, ok := temp[id]; !ok {
			temp[id] = i
			users = append(users, id)
		}
	}
	//用户id字符串拼接
	notify.StrUserIds = strings.Join(users, ",")
}

//设置agentId
func (notify *WorkNotify) SetAgentId(agentId int) {
	notify.AgentId = agentId
}

//请求参数验证
func (notify *WorkNotify) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(notify); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	if notify.All == false && len(notify.UserIds) == 0 && len(notify.DeptIds) == 0 {
		return errors.New("当字段All设置为false时必须指定UserIds或DeptIds其中一个参数的值")
	}
	return nil
}

//工作通知返回结果
type WorkNotifyResp struct {
	Response
	RequestId string `json:"request_id"` //请求的id
	TaskId    int    `json:"task_id"`    //创建的异步发送任务id
}

//工作通知请求
type WorkNotifyRes interface {
	//拼接部门
	AssembleDept()
	//拼接用户
	AssembleUser()
	//设置agentId
	SetAgentId(id int)
	//参数验证
	Validate(valid *validator.Validate, trans translator.Translator) error
}
