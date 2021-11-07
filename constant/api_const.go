package constant

const (
	Host                     = "https://oapi.dingtalk.com" //钉钉ap域名
	GetTokenKey              = "/gettoken"                 //获取access_token
	MicroAppListKey          = "/microapp/list"            //获取应用列表
	MicroAppVisibleScopesKey = "/microapp/visible_scopes"  //获取应用可见范围

	// SuiteAccessToken 获取第三方企业应用的
	SuiteAccessToken = "/service/get_suite_token"
	//CorpAccessToken 服务商获取第三方应用授权企业的access_token
	CorpAccessToken = "/service/get_corp_token"
	//GetAuthInfo 获取企业授权信息
	GetAuthInfo = "/service/get_auth_info"

	CreateUserKey             = "/topapi/v2/user/create"       //创建用户
	DeleteUserKey             = "/topapi/v2/user/delete"       //删除用户
	UpdateUserKey             = "/topapi/v2/user/update"       //更新用户详情
	GetUserDetailKey          = "/topapi/v2/user/get"          //获取用户详情
	GetDeptUserIdKey          = "/topapi/user/listid"          //获取部门用户userid列表
	GetUserIdByUnionIdKey     = "/topapi/user/getbyunionid"    //根据unionid获取userid
	GetUserIdByMobileKey      = "/topapi/v2/user/getbymobile"  //根据手机号获取userid
	GetDeptSimpleUserKey      = "/topapi/user/listsimple"      //获取部门用户基础信息
	GetDeptDetailUserKey      = "/topapi/v2/user/list"         //获取部门用户详情
	GetOrgAdminUserKey        = "/topapi/user/listadmin"       //获取管理员列表
	GetUserCanAccessAppletKey = "/user/can_access_microapp"    //获取管理员的应用管理权限
	GetUserCountKey           = "/topapi/user/count"           //获取员工人数
	GetInactiveUserKey        = "/topapi/inactive/user/v2/get" //获取未登录钉钉的员工列表
	GetOrgAdminScopeKey       = "/topapi/user/get_admin_scope" //获取管理员通讯录权限范围

	CreateDeptKey           = "/topapi/v2/department/create"           //创建部门
	DeleteDeptKey           = "/topapi/v2/department/delete"           //删除部门
	UpdateDeptKey           = "/topapi/v2/department/update"           //更新部门
	GetDeptDetailKey        = "/topapi/v2/department/get"              //获取部门详情
	GetDeptListKey          = "/topapi/v2/department/listsub"          //获取部门详情
	GetSubDeptListKey       = "/topapi/v2/department/listsubid"        //获取子部门列表
	GetParentDeptsByUserKey = "/topapi/v2/department/listparentbyuser" //查询指定用户的所有上级父部门路径
	GetParentDeptsByDeptKey = "/topapi/v2/department/listparentbydept" //查询部门的所有上级父部门路径

	CreateExtContactKey    = "/topapi/extcontact/create"          //添加外部联系人
	DeleteExtContactKey    = "/topapi/extcontact/delete"          //删除外部联系人
	UpdateExtContactKey    = "/topapi/extcontact/update"          //更新外部联系人
	GetExtContactKey       = "/topapi/extcontact/list"            //获取外部联系人列表
	GetExtContactLabelKey  = "/topapi/extcontact/listlabelgroups" //获取外部联系人标签列表
	GetExtContactDetailKey = "/topapi/extcontact/get"             //获取外部联系人详情

	GetIndustryDeptDetailKey = "/topapi/industry/department/get"  //获取外部联系人详情
	GetIndustryDeptKey       = "/topapi/industry/department/list" //获取部门列表

	role                         = "/topapi/role"                 //角色模块
	GetRoleListKey               = "/topapi/role/list"            //获取角色列表
	GetRoleUserListKey           = "/topapi/role/simplelist"      //获取指定角色的员工列表
	GetRoleGroupKey              = "/topapi/role/getrolegroup"    //获取角色组
	GetRoleDetailKey             = "/topapi/role/getrole"         //获取角色详情
	CreateRoleGroupKey           = "/role/add_role_group"         //创建角色组
	CreateRoleKey                = "/role/add_role"               //创建角色
	UpdateRoleKey                = "/role/update_role"            //更新角色
	DeleteRoleKey                = "/topapi/role/deleterole"      //删除角色
	RoleBatchAddUserKey          = "/topapi/role/addrolesforemps" //批量增加员工角色
	RoleBatchRemoveUserKey       = role + "/removerolesforemps"   //批量删除员工角色
	RoleUpdateUserManageScopeKey = "/topapi/role/scope/update"    //设定角色成员管理范围

	chat                = "/chat"                                   //群模块
	CreateChatKey       = chat + "/create"                          //创建群
	GetChatInfoKey      = chat + "/get"                             //获取群会话
	UpdateChatKey       = chat + "/update"                          //修改群会话
	SendMsgToChatKey    = chat + "/send"                            //发送消息到企业群
	GetChatReadUserKey  = chat + "/getReadList"                     //查询群消息已读人员列表
	ChatFriendSwitchKey = "/topapi/chat/member/friendswitch/update" //设置禁止群成员私聊
	ChatSubAdminKey     = "/topapi/chat/subadmin/update"            //设置群管理员

	SendToConversationKey = "/message/send_to_conversation" //发送普通消息

	SendRobotMsgKey = "/robot/send" //发送机器人消息

	TranslateKey      = "/topapi/ai/mt/translate"          //文本翻译
	OrcRecognizeKey   = "/topapi/ocr/structured/recognize" //OCR文字识别
	VoiceTranslateKey = "/topapi/asr/voice/translate"      //ASR 一句话语音识别

	corpConversation          = "/topapi/message/corpconversation"    //工作通知
	SendCorpConversationKey   = corpConversation + "/asyncsend_v2"    //发送工作通知
	GetSendProgressKey        = corpConversation + "/getsendprogress" //获取工作通知消息的发送进度
	GetSendResultKey          = corpConversation + "/getsendresult"   //获取工作通知消息的发送结果
	RecallCorpConversationKey = corpConversation + "/recall"          //撤回工作通知

	MediaUploadKey = "/media/upload" //上传媒体文件

	callback                   = "/call_back"                              //事件回调模板
	RegisterCallBackKey        = callback + "/register_call_back"          //注册事件回调接口
	GetCallBackKey             = callback + "/get_call_back"               //查询事件回调接口
	UpdateCallBackKey          = callback + "/update_call_back"            //更新事件回调接口
	GetCallBackFailedDetailKey = callback + "/get_call_back_failed_result" //获取回调失败的结果
	DeleteCallBackKey          = callback + "/delete_call_back"            //删除事件回调接口
)
