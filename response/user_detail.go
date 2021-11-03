package response

import "github.com/zhaoyunxing92/dingtalk/v2/request"

type UserDetail struct {
	Response
	userGetResponse `json:"result"`
}

type userGetResponse struct {
	request.CreateUser `json:"result"`

	//员工在当前开发者企业账号范围内的唯一标识
	UnionId string `json:"unionid"`

	//头像
	Avatar string `json:"avatar"`

	//国际电话区号
	StateCode string `json:"state_code"`

	//所属部门ID列表
	DeptIds []int `json:"dept_id_list"`

	//员工在对应的部门中的排序
	DeptOrders []request.DeptOrder `json:"dept_order_list"`

	//是否激活了钉钉
	Active bool `json:"active"`

	//是否完成了实名认证
	RealAuthed bool `json:"real_authed"`

	//是否为企业的高管
	Senior bool `json:"senior"`

	//是否为企业的老板
	Boss bool `json:"boss"`

	//员工在对应的部门中是否领导
	LeaderInDept []LeaderInDept `json:"leader_in_dept"`
}

type LeaderInDept struct {
	DeptId int `json:"dept_id"`

	Leader bool `json:"leader"`
}
