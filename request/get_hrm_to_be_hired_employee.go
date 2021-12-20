package request

type GetHrmToBeHiredEmployee struct {
	// 分页游标，从0开始。根据返回结果里的next_cursor是否为空来判断是否还有下一页，且再次调用时offset设置成next_cursor的值。
	Offset int `json:"offset"`

	// 分页大小，最大50。
	Size int `json:"size" validate:"max=50"`
}

func NewGetHrmToBeHiredEmployee(offset, size int) *GetHrmToBeHiredEmployee {
	return &GetHrmToBeHiredEmployee{offset, size}
}
