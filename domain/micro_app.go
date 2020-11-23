package domain

//钉钉应用
type MicroApp struct {
	Name           string `json:"name"`           //应用名称
	AgentId        uint64 `json:"agentId"`        //应用id
	Icon           string `json:"appIcon"`        //应用图标
	Desc           string `json:"appDesc"`        //应用描述
	Self           bool   `json:"isSelf"`         //是否自建 false:不是
	Status         uint8  `json:"appStatus"`      //应用状态 1：启用，0：停用
	OmpLink        string `json:"ompLink"`        //应用应用的OA后台管理主页
	HomepageLink   string `json:"homepageLink"`   //应用的移动端主页
	PcHomepageLink string `json:"pcHomepageLink"` //应用的PC端主页
}

//应用列表
type MicroAppList struct {
	Response
	AppList []MicroApp `json:"appList"` //应用列表
}

//应用可见范围
type MicroAppVisibleScopes struct {
	Response
	UserVisibleScopes []string `json:"userVisibleScopes"` //应用可见的用户列表。
	DeptVisibleScopes []uint32 `json:"deptVisibleScopes"` //应用可见的部门列表。
	Hidden            bool     `json:"isHidden"`          //是否仅限管理员可见，true代表仅限管理员可见。
}
