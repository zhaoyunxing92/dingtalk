package global

const (
	Host                      = "https://oapi.dingtalk.com"           //钉钉ap域名
	GetTokenKey               = "/gettoken"                           //获取access_token
	MicroAppListKey           = "/microapp/list"                      //获取应用列表
	MicroAppVisibleScopesKey  = "/microapp/visible_scopes"            //获取应用可见范围
	corpConversation          = "/topapi/message/corpconversation"    //工作通知
	user                      = "/user"                               //用户模块
	CreateUserKey             = user + "/create"                      //创建用户
	DeleteUserKey             = user + "/delete"                      //删除用户
	UpdateUserKey             = user + "/update"                      //更新用户详情
	GetUserKey                = user + "/get"                         //获取用户详情
	GetDeptUserIdKey          = user + "/getDeptMember"               //获取部门用户userid列表
	GetUserIdByUnionIdKey     = user + "/getUseridByUnionid"          //根据unionid获取userid
	GetUserByMobileKey        = user + "/get_by_mobile"               //根据手机号获取userid
	GetOrgUserCountKey        = user + "/get_org_user_count"          //获取企业员工人数
	GetDeptUserDetailKey      = user + "/listbypage"                  //获取部门用户详情
	GetUserAdminKey           = user + "/get_admin"                   //获取管理员列表
	SendCorpConversationKey   = corpConversation + "/asyncsend_v2"    //发送工作通知
	GetSendProgressKey        = corpConversation + "/getsendprogress" //获取工作通知消息的发送进度
	GetSendResultKey          = corpConversation + "/getsendresult"   //获取工作通知消息的发送结果
	RecallCorpConversationKey = corpConversation + "/recall"          //撤回工作通知
	CreateUserV2Key           = "/topapi/v2/user/create"              //创建用户
	departmentV2              = "/topapi/v2/department"               //部门管理2.0
	DepartmentCreateV2Key     = departmentV2 + "/create"              //创建部门
	DepartmentUpdateV2Key     = departmentV2 + "/update"              //更新部门
	DepartmentDeleteV2Key     = departmentV2 + "/delete"              //删除部门
	DepartmentGetV2Key        = "/topapi/edu/dept/get"                //获取部门详情
	DepartmentGetKey          = "/department/get"                     //获取部门详情
	MediaUploadKey            = "/media/upload"                       //上传媒体文件
)
