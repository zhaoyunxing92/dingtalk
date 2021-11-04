package response

type OrgAdminUser struct {
	Response

	AdminUser []orgAdminUser `json:"result"`
}

type orgAdminUser struct {
	UserId string `json:"userid"`

	//管理员角色：
	//
	//1：主管理员
	//
	//2：子管理员
	Level int `json:"sys_level"`
}
