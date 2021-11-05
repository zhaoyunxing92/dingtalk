package response

type DeptList struct {
	Response

	Depts []deptList `json:"result"`
}

type deptList struct {
	Id int `json:"dept_id"`

	Name string `json:"name"`

	ParentId int `json:"parent_id"`

	//是否同步创建一个关联此部门的企业群：
	//
	//true：创建
	//
	//false：不创建
	CreateDeptGroup bool `json:"create_dept_group"`

	//当部门群已经创建后，是否有新人加入部门会自动加入该群：
	//
	//true：自动加入群
	//
	//false：不会自动加入群
	AutoAddUser bool `json:"auto_add_user"`
}
