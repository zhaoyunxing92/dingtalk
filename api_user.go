package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
	"net/url"
	"strconv"
)

type userResult struct {
	UserId string `json:"userid"`
}

//创建用户返回
type UserCreateResponse struct {
	model.Response
	userResult `json:"result,omitempty"`
}

//UserDetail:用户详情
type UserDetail struct {
	model.Response
	Id              string            `json:"userid"`          //员工唯一标识ID（不可修改），企业内必须唯一。长度为1~64个字符，如果不传，将自动生成一个userid。
	UnionId         string            `json:"unionid"`         //员工在当前开发者企业账号范围内的唯一标识，系统生成，固定值，不会改变。
	OpenId          string            `json:"OpenId"`          //openId
	Name            string            `json:"name"`            //员工姓名，长度最大64个字符。
	Email           string            `json:"email"`           //员工邮箱，长度最大64个字符。企业内必须唯一，不可重复。
	Mobile          string            `json:"mobile"`          //手机号码，企业内必须唯一，不可重复。
	Position        string            `json:"position"`        //职位信息。长度为0~64个字符
	Roles           []model.UserRoles `json:"roles"`           //用户所在角色列表。
	OrgEmail        string            `json:"orgEmail"`        //员工的企业邮箱，如果员工已经开通了企业邮箱，接口会返回，否则会报错。
	Avatar          string            `json:"avatar"`          //头像
	Senior          bool              `json:"isSenior"`        //是否开启高管模式
	Authed          bool              `json:"realAuthed"`      //是否实名认证：true：是false：否
	Boss            bool              `json:"isBoss"`          //是否为企业的老板
	Active          bool              `json:"active"`          //是否激活
	Hide            bool              `json:"isHide"`          //是否号码隐藏：管模式
	Admin           bool              `json:"isAdmin"`         //是否为企业的管理员
	JobNumber       string            `json:"jobnumber"`       //员工工号，对应显示到OA后台和客户端个人资料的工号栏目。长度为0~64个字符。码隐藏：管模式
	Remark          string            `json:"remark"`          //备注，长度最大为1024个字符。
	WorkPlace       string            `json:"workPlace"`       //办公地点。长度为0~50个字符。
	Tel             string            `json:"tel"`             //分机号。长度为0~50个字符，企业内必须唯一，不可重复。
	PositionInDepts string            `json:"positionInDepts"` //设置用户在每个部门下的职位。Key是deptId，表示部门；Value是职位，表示在这个部门下的职位。
	LeaderInDepts   string            `json:"isLeaderInDepts"` //在部门里是否主管
	OrderInDepts    string            `json:"orderInDepts"`    //在对应的部门中的排序，Map结构的json字符串。Key是部门的ID，Value是人员在这个部门的排序值。
	Department      []int             `json:"department"`      //数组类型，数组里面值为整型，成员所属部门ID列表。
	Extattr         interface{}       `json:"extattr"`         //扩展属性，可以设置多种属性
	HiredDate       int               `json:"hiredDate"`       //入职时间，Unix时间戳，单位毫秒。
	Lang            string            `json:"lang"`            //通讯录语言，默认zh_CN。如果是英文，请输入en_US。
	StateCode       string            `json:"stateCode"`       //国家地区码
	ManagerUserId   string            `json:"managerUserId"`   //主管的ID
}

//UserIdResponse:创建用户、根据unionId获取
type UserIdResponse struct {
	model.Response
	ContactType int    `json:"contactType,omitempty"` //联系人类型：0:表示企业内部员工,1:表示企业外部联系人
	UserId      string `json:"userid"`                //用户userid
}

type OrgUserCountResponse struct {
	model.Response
	Count int `json:"count"`
}

//InactiveUserResponse:未登录用户数据。
type InactiveUserResponse struct {
	model.Response
	RequestId    string             `json:"request_id"`
	InactiveUser model.InactiveUser `json:"result"`
}

type OrgAdminUserResponse struct {
	model.Response
	Admins []model.OrgAdminList `json:"adminList"`
}

type OrgAdminScopeResponse struct {
	model.Response
	RequestId string `json:"request_id"`
	Depts     []int  `json:"dept_ids"` //可管理的部门ID列表
}

//CreateUser:创建用户
//name:姓名
//mobile:手机号
//deptId:部门
func (talk *DingTalk) CreateUser(name, mobile string, deptIds []int) (req UserIdResponse, err error) {

	form := map[string]interface{}{
		"name":       name,
		"mobile":     mobile,
		"department": deptIds,
	}

	err = talk.request(http.MethodPost, global.CreateUserKey, nil, form, &req)
	return req, err
}

//DeleteUser:删除用户
//userId:员工唯一标识userid
func (talk *DingTalk) DeleteUser(userId string) (req model.Response, err error) {

	params := url.Values{}
	params.Set("userid", userId)

	err = talk.request(http.MethodGet, global.DeleteUserKey, params, nil, &req)
	return req, err
}

//UpdateUser:更新用户
//res:员工信息
//TODO:PositionInDepts:只用写一个，因为写多了也是无用的
func (talk *DingTalk) UpdateUser(res model.Request) (req model.Response, err error) {

	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return req, err
	}

	err = talk.request(http.MethodPost, global.UpdateUserKey, nil, res, &req)
	return req, err
}

//GetUserDetail:获取用户
//userId:用户id
//lang:语言，默认是zh_CN
//TODO:PositionInDepts:只用写一个，因为写多了也是无用的
func (talk *DingTalk) GetUserDetail(userId, lang string) (req UserDetail, err error) {

	if lang != "en_US" {
		lang = "zh_CN"
	}

	params := url.Values{}
	params.Set("userid", userId)
	params.Set("lang", lang)

	err = talk.request(http.MethodGet, global.GetUserKey, params, nil, &req)
	return req, err
}

//GetUserIdByUnionId:根据unionid获取userid
//unionId:员工在当前开发者企业账号范围内的唯一标识
func (talk *DingTalk) GetUserIdByUnionId(unionId string) (req UserIdResponse, err error) {

	params := url.Values{}
	params.Set("unionid", unionId)

	err = talk.request(http.MethodGet, global.GetUserIdByUnionIdKey, params, nil, &req)
	return req, err
}

//GetUserIdByMobile:根据手机号获取userid
//unionId:员工在当前开发者企业账号范围内的唯一标识
func (talk *DingTalk) GetUserIdByMobile(mobile string) (req UserIdResponse, err error) {

	params := url.Values{}
	params.Set("mobile", mobile)

	err = talk.request(http.MethodGet, global.GetUserIdByMobileKey, params, nil, &req)
	return req, err
}

//GetOrgUserCount:获取企业员工人数
//active:0:包含未激活钉钉的人员数量,1:不包含未激活钉钉的人员数量
func (talk *DingTalk) GetOrgUserCount(active int) (req OrgUserCountResponse, err error) {

	if active < 0 || active > 2 {
		active = 0
	}

	params := url.Values{}
	params.Set("onlyActive", strconv.Itoa(active))

	err = talk.request(http.MethodGet, global.GetOrgUserCountKey, params, nil, &req)
	return req, err
}

//GetOrgInactiveUser:获取未登录钉钉的员工列表
//active:0:包含未激活钉钉的人员数量,1:不包含未激活钉钉的人员数量
func (talk *DingTalk) GetOrgInactiveUser(date string, offset, size int) (req InactiveUserResponse, err error) {
	if size < 0 || size > 100 {
		size = 100
	}

	form := make(map[string]interface{}, 3)
	form["query_date"] = date
	form["offset"] = offset
	form["size"] = size

	err = talk.request(http.MethodPost, global.GetOrgInactiveUserKey, nil, form, &req)
	return req, err
}

//GetOrgAdminUser:获取未登录钉钉的员工列表
func (talk *DingTalk) GetOrgAdminUser() (req OrgAdminUserResponse, err error) {
	err = talk.request(http.MethodGet, global.GetOrgAdminUserKey, nil, nil, &req)
	return req, err
}

//GetOrgAdminScope:获取管理员通讯录权限范围
func (talk *DingTalk) GetOrgAdminScope(userId string) (req OrgAdminScopeResponse, err error) {
	form := make(map[string]string, 1)
	form["userid"] = userId

	err = talk.request(http.MethodPost, global.GetOrgAdminScopeKey, nil, form, &req)
	return req, err
}

//CreateUserV2:创建用户
//name:姓名
//mobile:手机号
//deptId:部门
func (talk *DingTalk) CreateUserV2(name, mobile string, deptId int) (req UserCreateResponse, err error) {

	form := map[string]interface{}{
		"name":         name,
		"mobile":       mobile,
		"dept_id_list": deptId,
	}

	err = talk.request(http.MethodGet, global.CreateUserKey, nil, form, &req)
	return req, err
}

//创建详细的用户
func (talk *DingTalk) CreateDetailUserV2(res model.Request) (req UserCreateResponse, err error) {

	if err = res.Validate(talk.validate, talk.trans); err != nil {
		return req, err
	}

	err = talk.request(http.MethodGet, global.CreateUserKey, nil, res, &req)
	return req, err
}
