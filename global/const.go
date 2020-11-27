package global

const (
	Host                      = "https://oapi.dingtalk.com"           //钉钉ap域名
	GetTokenKey               = "/gettoken"                           //获取access_token
	MicroAppListKey           = "/microapp/list"                      //获取应用列表
	MicroAppVisibleScopesKey  = "/microapp/visible_scopes"            //获取应用可见范围
	corpConversation          = "/topapi/message/corpconversation"    //工作通知
	SendCorpConversationKey   = corpConversation + "/asyncsend_v2"    //发送工作通知
	GetSendProgressKey        = corpConversation + "/getsendprogress" //获取工作通知消息的发送进度
	GetSendResultKey          = corpConversation + "/getsendresult"   //获取工作通知消息的发送结果
	RecallCorpConversationKey = corpConversation + "/recall"          //撤回工作通知
	CreateUserKey             = "/topapi/v2/user/create"              //创建用户
	departmentV2              = "/topapi/v2/department"               //部门管理2.0
	DepartmentCreateKey       = departmentV2 + "/create"              //创建部门
	DepartmentUpdateKey       = departmentV2 + "/update"              //更新部门
	DepartmentDeleteKey       = departmentV2 + "/delete"              //删除部门
	DepartmentEduGetKey       = "/topapi/edu/dept/get"                //获取部门详情
	MediaUploadKey            = "/media/upload"                       //上传媒体文件
)
