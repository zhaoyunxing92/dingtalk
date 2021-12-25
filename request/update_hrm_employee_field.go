package request

type UpdateHrmEmpField struct {

	// 应用的AgentID
	AgentId int `json:"agentid"`

	// 员工信息
	Param UpdateHrmEmpFieldParam `json:"param"`
}

type UpdateHrmEmpFieldParam struct {

	// 被更新字段信息的员工userid
	UserId string `json:"userid,omitempty"`

	// 花名册分组
	EmpFieldGroup []UpdateHrmEmpFieldGroup `json:"groups"`
}

// UpdateHrmEmpFieldGroup 花名册分组
type UpdateHrmEmpFieldGroup struct {
	// 分组标识。
	GroupId string `json:"group_id,omitempty"`

	// 分组下明细，非明细分组仅一条明细
	Fields []EmpField `json:"sections"`
}

type EmpField struct {
	//明细下标。
	//
	//当传入该值时，表示当前传入的section为编辑员工花名册现有的第old_index条明细，此时系统会只编辑该条明细中传入的字段。
	//
	//当不传入该值时，表示当前传入的section为新增明细，此时系统会保存该条明细传入的字段，未传字段会清空。
	OldIndex *int `json:"old_index,omitempty"`

	Values []EmpFieldValue `json:"section"`
}

type EmpFieldValue struct {
	//更新的字段code。
	//
	//说明
	//sys00分组内字段code值为汉字，如姓名，更新时直接传入"姓名"作为该参数值
	//sys00分组内的字段，部门、直属主管和主部门暂不支持接口更新
	//添加的自定义字段暂不支持接口更新
	FieldCode string `json:"field_code"`

	// 更新的字段值
	Value string `json:"value"`
}
