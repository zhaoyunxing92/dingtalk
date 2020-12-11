package global

const (
	Host                     = "https://oapi.dingtalk.com"        //钉钉ap域名
	GetTokenKey              = "/gettoken"                        //获取access_token
	MicroAppListKey          = "/microapp/list"                   //获取应用列表
	MicroAppVisibleScopesKey = "/microapp/visible_scopes"         //获取应用可见范围
	corpConversation         = "/topapi/message/corpconversation" //工作通知

	user                  = "/user"                      //用户模块
	CreateUserKey         = user + "/create"             //创建用户
	DeleteUserKey         = user + "/delete"             //删除用户
	UpdateUserKey         = user + "/update"             //更新用户详情
	GetUserKey            = user + "/get"                //获取用户详情
	GetDeptUserIdKey      = user + "/getDeptMember"      //获取部门用户userid列表
	GetUserIdByUnionIdKey = user + "/getUseridByUnionid" //根据unionid获取userid
	GetUserIdByMobileKey  = user + "/get_by_mobile"      //根据手机号获取userid
	GetOrgUserCountKey    = user + "/get_org_user_count" //获取企业员工人数
	GetDeptUserDetailKey  = user + "/listbypage"         //获取部门用户详情
	GetOrgAdminUserKey    = user + "/get_admin"          //获取管理员列表

	dept                    = "/department"                       //部门模块
	CreateDeptKey           = dept + "/create"                    //创建部门
	DeleteDeptKey           = dept + "/delete"                    //删除部门
	UpdateDeptKey           = dept + "/update"                    //更新部门
	GetDeptDetailKey        = dept + "/get"                       //获取部门详情
	GetSubDeptListKey       = dept + "/list"                      //获取子部门列表
	GetSubDeptIdsKey        = dept + "/list_ids"                  //获取子部门ID列表
	GetParentDeptsByUserKey = dept + "/list_parent_depts"         //查询指定用户的所有上级父部门路径
	GetParentDeptsByDeptKey = dept + "/list_parent_depts_by_dept" //查询部门的所有上级父部门路径

	role                             = "/topapi/role"               //角色模块
	GetRoleListKey                   = role + "/list"               //获取角色列表
	GetRoleUserListKey               = role + "/simplelist"         //获取指定角色的员工列表
	GetRoleGroupKey                  = role + "/getrolegroup"       //获取角色组
	GetRoleDetailKey                 = role + "/getrole"            //获取角色详情
	CreateRoleGroupKey               = "/role/add_role_group"       //创建角色组
	CreateRoleKey                    = "/role/add_role"             //创建角色
	UpdateRoleKey                    = "/role/update_role"          //更新角色
	DeleteRoleKey                    = role + "/deleterole"         //删除角色
	RoleBatchAddUserKey              = role + "/addrolesforemps"    //批量增加员工角色
	RoleBatchRemoveUserKey           = role + "/removerolesforemps" //批量删除员工角色
	RoleUpdateUserManageScopeKey = role + "/scope/update"       //设定角色成员管理范围

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

	GetOrgInactiveUserKey     = "/topapi/inactive/user/get"           //获取未登录钉钉的员工列表
	GetOrgAdminScopeKey       = "/topapi/user/get_admin_scope"        //获取管理员通讯录权限范围
	SendCorpConversationKey   = corpConversation + "/asyncsend_v2"    //发送工作通知
	GetSendProgressKey        = corpConversation + "/getsendprogress" //获取工作通知消息的发送进度
	GetSendResultKey          = corpConversation + "/getsendresult"   //获取工作通知消息的发送结果
	RecallCorpConversationKey = corpConversation + "/recall"          //撤回工作通知
	MediaUploadKey            = "/media/upload"                       //上传媒体文件
)
