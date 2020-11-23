package domain

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

//请求
type workNotifyRes struct {
	AgentId    uint64   `json:"agent_id" validate:"required"`         //发送消息时使用的微应用的ID。
	UserIds    []string `json:"userIds,omitempty" validate:"lte=100"` //接收者的企业内部用户的userId列表，最大用户列表长度100。
	StrUserIds string   `json:"userid_list" validate:"required"`
	StrDeptIds string   `json:"dept_ids" validate:"required"`
	DeptIds    []uint64 `json:"deptIds,omitempty" validate:"lte=20"` //接收者的部门id列表，最大列表长度20。接收者是部门Id下包括子部门下的所有用户。
	All        bool     `json:"to_all_user"`                         //是否发送给企业全部用户。当设置为false时必须指定userid_list或dept_id_list其中一个参数的值。
}

type TextWorkNotifyRes struct {
	workNotifyRes
	Msg TextMessage `json:"msg" validate:"required"`
}

//请求参数验证
func (req TextWorkNotifyRes) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(req); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}

//移除重复部门
func (req TextWorkNotifyRes) RemoveRepeatDept() []uint64 {

	return nil
}

//移除重复用户,字符串拼接
func (req TextWorkNotifyRes) RemoveRepeatUser() []string {
	var size = len(req.UserIds)
	if size == 0 {
		return []string{}
	}
	users := make([]string, 0, size)
	temp := map[string]int{}

	for i, id := range req.UserIds {
		if _, ok := temp[id]; !ok {
			temp[id] = i
			users = append(users, id)
		}
	}
	//用户id字符串拼接
	req.StrUserIds = strings.Join(users, ",")
	return users
}

//工作通知返回结果
type WorkNotifyResp struct {
	Response
	RequestId string `json:"request_id"` //请求的id
	TaskId    string `json:"task_id"`    //创建的异步发送任务id
}
