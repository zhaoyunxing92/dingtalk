package request

import (
	"strings"

	"github.com/zhaoyunxing92/dingtalk/v2/constant/employee"
)

type GetHrmEmployee struct {

	// 在职员工子状态筛选，可以查询多个状态。不同状态之间使用英文逗号分隔。
	//
	//2：试用期
	//
	//3：正式
	//
	//5：待离职
	//
	//-1：无状态
	Status string `json:"status_list"`

	// 分页游标，从0开始。根据返回结果里的next_cursor是否为空来判断是否还有下一页，且再次调用时offset设置成next_cursor的值。
	Offset int `json:"offset"`

	// 分页大小，最大50。
	Size int `json:"size"`
}

func NewGetHrmEmployee(offset, size int, status []employee.Status) *GetHrmEmployee {
	return &GetHrmEmployee{
		Status: strings.Join(removeStatusDuplicates(status), ","),
		Offset: offset, Size: size,
	}
}
