package request

type GetAttendanceGroup struct {
	// 支持分页查询，与size参数同时设置时才生效，此参数代表偏移量，偏移量从0开始，下次调用传上次调用时的size与offset之和
	Offset int `json:"offset"`

	// 分页大小，最大10。
	Size int `json:"size" validate:"max=10"`
}

func NewGetAttendanceGroup(offset, size int) *GetAttendanceGroup {
	return &GetAttendanceGroup{offset, size}
}
