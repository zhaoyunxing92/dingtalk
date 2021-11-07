package response

type GetExtContactDetail struct {
	Response

	getExtContactDetail `json:"result"`
}

type getExtContactDetail struct {
	//职位
	Title string `json:"title"`

	UserId string `json:"userid"`

	//共享给的部门ID
	ShareDept []int `json:"share_dept_ids"`

	//备注
	Remark string `json:"remark"`

	//标签列表
	Labels []int `json:"label_ids"`

	//邮箱
	Email string `json:"email"`

	//地址
	Address string `json:"address"`

	//负责人的userId
	FollowerUser string `json:"follower_user_id"`

	//外部联系人的姓名
	Name string `json:"name"`

	//手机号国家码
	StateCode string `json:"state_code"`

	//外部联系人的企业名称
	CompanyName string `json:"company_name"`

	//共享给的员工userid列表
	ShareUser []string `json:"share_user_ids"`

	//外部联系人的手机号
	Mobile string `json:"mobile"`
}