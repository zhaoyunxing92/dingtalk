package request

type RoleList struct {

	//支持分页查询，与size参数同时设置时才生效，此参数代表偏移量，偏移量从0开始
	Offset int `json:"offset"`

	//支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，默认值20，最大值200
	Size int `json:"size" validate:"max=200"`
}

func NewRoleList(offset, size int) *RoleList {

	return &RoleList{offset, size}
}
