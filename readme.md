[![dingtalk](./image/dingtalk.png)](https://developers.dingtalk.com/document/app)

[![License](http://img.shields.io/:license-apache-brightgreen.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)
[![ci](https://github.com/zhaoyunxing92/dingtalk/actions/workflows/ci.yaml/badge.svg)](https://github.com/zhaoyunxing92/dingtalk/actions)
[![codecov](https://codecov.io/gh/zhaoyunxing92/dingtalk/branch/develop/graph/badge.svg)](https://codecov.io/gh/zhaoyunxing92/dingtalk)
![GitHub commit activity (branch)](https://img.shields.io/github/commit-activity/y/zhaoyunxing92/dingtalk/main?logo=CodeIgniter)
[![issues](https://img.shields.io/github/issues/zhaoyunxing92/dingtalk?color=ff69b4)](https://github.com/zhaoyunxing92/dingtalk/issues)
[![GitHub milestone](https://img.shields.io/github/milestones/progress/zhaoyunxing92/dingtalk/3)](https://github.com/zhaoyunxing92/dingtalk/milestone/2?closed=1)
[![tag](https://img.shields.io/github/v/tag/zhaoyunxing92/dingtalk?color=green&logo=AzurePipelines)](https://github.com/zhaoyunxing92/dingtalk/releases/tag/v2.0.3)

> [创建和配置应用](https://ding-doc.dingtalk.com/document#/org-dev-guide/applicaiton-manangment) 先有应用我们才能一起玩

[dingtalk](https://github.com/zhaoyunxing92/dingtalk) 是基于新版的 [钉钉开发平台api文档](https://developers.dingtalk.com/document/app)
开发，也就是说在你使用之前是需要在[钉钉开发者后台](https://open-dev.dingtalk.com/#/index) 创建一个小程序或应用

## 为什么有它

首先我要感谢 [godingtalk](https://github.com/hugozhu/godingtalk) 库,我是参考他的代码写的。但是由于钉钉历史原因，很多api都更新了,其次就是我有很多钉钉预警的业务需要做,于是就写了这个

## 怎么使用

```shell
go get github.com/zhaoyunxing92/dingtalk/v2
```

```go
client := NewClient(1244553273, "appkey", "AppSecret")
// 获取token
token, _ := client.GetAccessToken()

fmt.Println(token)
```

### 日志级别设置

```go
client:= NewClient(1244553273, "appkey", "AppSecret", WithLevel(zap.DebugLevel))
```

### `isv`和企业内部开发切换

在`NewClient`的时候如果调用了`WithTicket`和`WithCorpId`则判定为是`isv`

## 怎么参与

正如你所见钉钉的文档太多了，欢迎大家来提交pr，[怎么提交pr](https://github.com/zhaoyunxing92/dingtalk/wiki)

## 参考文档

[dingtalk-callback-Crypto](https://github.com/open-dingtalk/dingtalk-callback-Crypto)

## 开发进度

* 获取凭证 (**2**/5)
    - [x] [获取企业内部应用的access_token](https://developers.dingtalk.com/document/app/obtain-orgapp-token) `client.GetAccessToken`
    - [x] [服务商获取第三方应用授权企业的access_token](https://developers.dingtalk.com/document/app/obtains-the-enterprise-authorized-credential) `client.GetCorpAccessToken`
    - [x] [获取第三方企业应用的suite_access_token](https://developers.dingtalk.com/document/app/obtain-application-suite-ticket) `client.GetSuiteAccessToken`
    - [x] [获取jsapi_ticket](https://developers.dingtalk.com/document/app/obtain-jsapi_ticket) `client.GetJsApiTicket`
    - [x] [获取微应用后台免登的SsoToken](https://developers.dingtalk.com/document/app/obtain-the-ssotoken-for-micro-application-background-logon-free) `client.GetSSOToken`

* 通讯录管理
    - 用户管理(**14**/14)
        - [x] [创建用户](https://developers.dingtalk.com/document/app/user-information-creation) `client.CreateUser`
        - [x] [删除用户](https://developers.dingtalk.com/document/app/delete-a-user) `client.DeleteUser`
        - [x] [更新用户详情](https://developers.dingtalk.com/document/app/user-information-update) `client.UpdateUser`
        - [x] [根据userid获取用户详情](https://developers.dingtalk.com/document/app/query-user-details) `client.GetUserDetail`
        - [x] [获取部门用户基础信息](https://developers.dingtalk.com/document/app/queries-the-simple-information-of-a-department-user) `client.GetDeptUserIds`
        - [x] [获取部门用户userid列表](https://developers.dingtalk.com/document/app/query-the-list-of-department-userids) `client.GetDeptUserIds`
        - [x] [获取部门用户详情](https://developers.dingtalk.com/document/app/queries-the-complete-information-of-a-department-user) `client.GetDeptDetailUserInfo`
        - [x] [获取员工人数](https://developers.dingtalk.com/document/app/obtain-the-number-of-employees-v2)  `client.GetUserCount`
        - [x] [获取未登录钉钉的员工列表](https://developers.dingtalk.com/document/app/queries-the-inactive-users-or-active-users-under-an-enterprise)  `client.GetInactiveUser`
        - [x] [根据手机号获取userid](https://developers.dingtalk.com/document/app/query-users-by-phone-number) `client.GetUserIdByMobile`
        - [x] [根据unionid获取用户userid](https://developers.dingtalk.com/document/app/query-a-user-by-the-union-id) `client.GetUserIdByUnionId`
        - [x] [获取管理员列表](https://developers.dingtalk.com/document/app/query-the-administrator-list) `client.GetOrgAdminUser`
        - [x] [获取管理员通讯录权限范围](https://developers.dingtalk.com/document/app/query-permissions-of-the-administrator-address-book) `client.GetOrgAdminScope`
        - [x] [获取管理员的应用管理权限](https://developers.dingtalk.com/document/app/obtains-the-administrator-s-microapplication-management-permission) `client.GetUserCanAccessApplet`

    - 部门管理(**8**/8)
        - [x] [创建部门](https://developers.dingtalk.com/document/app/create-a-department-v2) `client.CreateDept`
        - [x] [更新部门](https://developers.dingtalk.com/document/app/update-a-department-v2) `client.UpdateDept`
        - [x] [删除部门](https://developers.dingtalk.com/document/app/delete-a-department-v2) `client.DeleteDept`
        - [x] [获取部门详情](https://developers.dingtalk.com/document/app/query-department-details0-v2) `client.GetDeptDetail`
        - [x] [获取部门列表](https://developers.dingtalk.com/document/app/obtain-the-department-list-v2) `client.GetDeptList`
        - [x] [获取子部门Id列表](https://developers.dingtalk.com/document/app/obtain-a-sub-department-id-list-v2)  `client.GetSubDeptList`
        - [x] [获取指定部门的所有父部门列表](https://developers.dingtalk.com/document/app/query-the-list-of-all-parent-departments-of-a-department) `client.GetParentIdsByDeptId`
        - [x] [获取指定用户的所有父部门列表](https://developers.dingtalk.com/document/app/queries-the-list-of-all-parent-departments-of-a-user)  `client.GetParentIdsByUserId`

    - 角色管理(**11**/11)
        - [x] [创建角色](https://developers.dingtalk.com/document/app/add-role) `client.CreateRole`
        - [x] [创建角色组](https://developers.dingtalk.com/document/app/add-a-role-group) `client.CreateRoleGroup`
        - [x] [更新角色](https://developers.dingtalk.com/document/app/update-role) `client.UpdateRole`
        - [x] [批量增加员工角色](https://developers.dingtalk.com/document/app/add-role-information-to-employees-in-batches) `client.BatchAddUserRole`
        - [x] [删除角色](https://developers.dingtalk.com/document/app/delete-role-information) `client.DeleteRole`
        - [x] [批量删除员工角色](https://developers.dingtalk.com/document/app/delete-the-color-information-of-employee-corners-in-batches) `client.BatchRemoveUserRole`
        - [x] [批量删除员工角色](https://developers.dingtalk.com/document/app/delete-the-color-information-of-employee-corners-in-batches) `client.BatchRemoveUserRole`
        - [x] [设定角色成员管理范围](https://developers.dingtalk.com/document/app/update-role-member-management-department-scope) `client.SetUserRoleManageScope`
          官方接口不通
        - [x] [获取角色组列表](https://developers.dingtalk.com/document/app/obtains-the-role-group-information) `client.GetGroupRoles`
        - [x] [获取角色列表](https://developers.dingtalk.com/document/app/obtains-a-list-of-enterprise-roles) `client.GetRoleList`
        - [x] [获取角色详情](https://developers.dingtalk.com/document/app/queries-role-details) `client.GetRoleDetail`
        - [x] [获取指定角色的员工列表](https://developers.dingtalk.com/document/app/obtain-the-list-of-employees-of-a-role) `client.GetRoleUserList`

    - 外部联系人(**6**/6)
        - [x] [添加外部联系人](https://developers.dingtalk.com/document/app/add-enterprise-external-contacts) `client.CreateExtContact`
        - [x] [删除外部联系人](https://developers.dingtalk.com/document/app/delete-external-contact) `client.DeleteExtContact`
        - [x] [更新外部联系人](https://developers.dingtalk.com/document/app/update-enterprise-external-contacts) `client.UpdateExtContact`
        - [x] [获取外部联系人列表](https://developers.dingtalk.com/document/app/obtain-the-external-contact-list) `client.GetExtContact`
        - [x] [获取外部联系人标签列表](https://developers.dingtalk.com/document/app/obtains-a-list-of-external-contact-tags) `client.GetExtContactLabel`
        - [x] [获取外部联系人详情](https://developers.dingtalk.com/document/app/obtains-the-external-contact-details-of-an-enterprise) `client.GetExtContactDetail`

* 群会话管理(**7**/7)
    - [x] [创建群会话](https://developers.dingtalk.com/document/app/create-group-session) `client.CreateChat`
    - [x] [获取群会话信息](https://developers.dingtalk.com/document/app/obtain-a-group-session) `client.GetChatInfo`
    - [x] [修改群会话](https://developers.dingtalk.com/document/app/modify-a-group-session)  `client.UpdateChat`
    - [x] [设置禁止群成员私聊](https://developers.dingtalk.com/document/app/set-private-chat) `client.ChatFriendSwitch`
    - [x] [获取入群二维码链接](https://developers.dingtalk.com/document/app/obtain-a-qr-code-link) `client.GetChatQRCode`
    - [x] [设置群管理员](https://developers.dingtalk.com/document/app/set-chat-admin)  `client.ChatSetSubAdmin`
    - [x] [设置群成员昵称](https://developers.dingtalk.com/document/app/set-a-group-nickname)  `client.ChatSetUserNick`

* 消息通知(4/6)
    - 工作通知
        - [x] [使用模板发送工作通知消息](https://developers.dingtalk.com/document/app/work-notification-templating-send-notification-interface) `client.SendTemplateMessage`
        - [x] [发送工作通知](https://developers.dingtalk.com/document/app/asynchronous-sending-of-enterprise-session-messages) `client.SendCorpConvMessage`
        - [x] [更新工作通知状态栏](https://developers.dingtalk.com/document/app/update-work-notification-status-bar) `client.UpdateCorpConvMessageStatus`
        - [x] [获取工作通知消息的发送进度](https://developers.dingtalk.com/document/app/obtain-the-sending-progress-of-asynchronous-sending-of-enterprise-session) `client.GetCorpConvMsgProgress`
        - [x] [获取工作通知消息的发送结果](https://developers.dingtalk.com/document/app/gets-the-result-of-sending-messages-asynchronously-to-the-enterprise) `client.GetMessageSendResult`
        - [x] [撤回工作通知消息](https://developers.dingtalk.com/document/app/notification-of-work-withdrawal) `client.RecallCorpConvMessage`

    - 企业群消息
        - [x] [发送消息到企业群](https://developers.dingtalk.com/document/app/send-group-messages) `client.SendChatMessage`
        - [x] [查询群消息已读人员列表](https://developers.dingtalk.com/document/app/queries-the-list-of-people-who-have-read-a-group) `client.GetChatMsgReadUser`

    - 普通消息
        - [x] [发送普通消息](https://developers.dingtalk.com/document/app/send-normal-messages) `client.SendMessage`

* 身份验证
    - [x] [通过免登码获取用户信息](https://developers.dingtalk.com/document/app/obtain-the-userid-of-a-user-by-using-the-log-free) `client.GetUserInfoByCode`
    - [x] [获取应用管理员的身份信息](https://developers.dingtalk.com/document/app/exchange-code-for-the-identity-information-of-a-microapplication-administrator) `client.GetSSOUserInfo`
    - [x] [根据sns临时授权码获取用户信息](https://developers.dingtalk.com/document/app/obtain-the-user-information-based-on-the-sns-temporary-authorization)

* 应用授权(**6**/6)
    - [x] [激活应用](https://developers.dingtalk.com/document/app/activate-suite) `client.ActivateSuite`
    - [x] [获取授权应用的基本信息](https://developers.dingtalk.com/document/app/obtains-application-information-of-an-enterprise)  `client.GetAgentInfo`
    - [x] [获取企业授权信息](https://developers.dingtalk.com/document/app/obtains-the-basic-information-of-an-enterprise) `client.GetAuthInfo`
    - [x] [获取授权企业的永久授权码](https://developers.dingtalk.com/document/app/obtain-a-permanent-authorization-code) `client.GetCorpPermanentCode`
    - [x] [获取应用未激活的企业列表](https://developers.dingtalk.com/document/app/obtains-a-list-of-enterprises-whose-applications-are-not-activated) `client.GetUnactiveCorpIds`
    - [x] [重新授权未激活应用的企业](https://developers.dingtalk.com/document/app/re-authorize-enterprises-whose-applications-are-not-activated) `client.ReauthCorp`

* 待办任务
    - [x] [新增钉钉待办任务](https://developers.dingtalk.com/document/app/add-dingtalk-to-do-task) `client.CreateTodo`
    - [x] [获取钉钉待办任务详情](https://developers.dingtalk.com/document/app/obtain-dingtalk-pending-tasks-details) `client.GetTodoDetail`
    - [x] [删除钉钉待办任务](https://developers.dingtalk.com/document/app/delete-dingtalk-to-do-tasks) `client.DeleteTodo`
    - [x] [更新钉钉待办任务](https://developers.dingtalk.com/document/app/update-dingtalk-to-do-tasks) `client.UpdateTodo`
    - [x] [更新钉钉待办执行者状态](https://developers.dingtalk.com/document/app/update-dingtalk-to-do-status) `client.UpdateTodoDone`
    - [x] [根据sourceId获取钉钉待办任务详情](https://developers.dingtalk.com/document/app/you-can-call-this-operation-to-query-the-details-about) `client.GetTodoListBySourceId`
    - [x] [查询企业下用户待办列表](https://developers.dingtalk.com/document/app/query-the-to-do-list-of-enterprise-users) `client.GetTodoList`

* 钉盘
    - 空间管理
        - [x] [新建空间](https://developers.dingtalk.com/document/app/new-space) `client.CreateDriveSpaces`
        - [x] [删除空间](https://developers.dingtalk.com/document/app/delete-a-space) `client.DeleteDriveSpaces`
        - [x] [获取空间列表](https://developers.dingtalk.com/document/app/queries-a-space-list) `client.GetDriveSpaces`
        - [x] [获取空间信息](https://developers.dingtalk.com/document/app/obtain-space-information) `client.GetDriveSpacesInfo`
    - 文件管理
        - [x] [查询文件（夹）列表](https://developers.dingtalk.com/document/app/obtain-the-file-list) `client.GetDriveSpacesFiles`
        - [x] [查询文件（夹）信息](https://developers.dingtalk.com/document/app/obtain-file-information) `client.GetDriveSpacesFileInfo`
        - [x] [添加文件（夹）](https://developers.dingtalk.com/document/app/add-file) `client.CreateDriveSpacesFiles`
        - [x] [删除文件（夹）](https://developers.dingtalk.com/document/app/delete-objects) `client.DeleteDriveSpacesFiles`
        - [x] [移动文件（夹）](https://developers.dingtalk.com/document/app/move-a-file) `client.MoveDriveSpacesFiles`
        - [x] [修改文件（夹）名](https://developers.dingtalk.com/document/app/rename-a-file) `client.RenameDriveSpacesFiles`
    - 权限管理
        - [x] [获取权限列表](https://developers.dingtalk.com/document/app/obtain-a-permission-list) `client.GetDriveSpacesFilesPermissions`
        - [x] [添加权限](https://developers.dingtalk.com/document/app/add-permissions) `client.AddDriveSpacesFilesPermissions`
        - [ ] [删除权限](https://open.dingtalk.com/document/orgapp-server/delete-permissions)
        - [ ] [修改权限](https://open.dingtalk.com/document/orgapp-server/modify-permissions)
        - [ ] [修改权限](https://open.dingtalk.com/document/orgapp-server/modify-permissions)
        - [ ] [添加自定义空间权限](https://open.dingtalk.com/document/orgapp-server/add-custom-workspace-permissions)
    - 回收站
        - [ ] [查询回收站文件（夹）列表](https://open.dingtalk.com/document/orgapp-server/obtain-the-recycle-bin-folder-list)
        - [ ] [还原回收站文件（夹）](https://open.dingtalk.com/document/orgapp-server/restore-recycle-bin-files-folder)
        - [ ] [删除回收站文件（夹）](https://open.dingtalk.com/document/orgapp-server/delete-recycle-bin-files-folders)
        - [ ] [清空回收站](https://open.dingtalk.com/document/orgapp-server/empty-recycle-bin-files-folders)

* 智能人事
    - [x] [获取在职员工列表](https://open.dingtalk.com/document/orgapp-server/intelligent-personnel-query-the-list-of-on-the-job-employees-of-the) `client.GetHrmEmployee`
    - [x] [获取待入职员工列表](https://open.dingtalk.com/document/orgapp-server/intelligent-personnel-query-the-list-of-employees-to-be-hired) `client.GetHrmToBeHiredEmployee`
    - [x] [获取离职员工列表](https://open.dingtalk.com/document/orgapp-server/intelligent-personnel-query-company-turnover-list) `client.GetHrmResignEmployeeIds`
    - [x] [获取员工离职信息](https://open.dingtalk.com/document/orgapp-server/intelligent-personnel-query-company-turnover-list) `client.GetHrmResignEmployee`
    - [x] [添加企业待入职员工](https://open.dingtalk.com/document/orgapp-server/add-employees-to-be-hired-through-intelligent-personnel) `client.HrmCreateEmployee`
    - 花名册
        - [x] [获取花名册字段组详情](https://open.dingtalk.com/document/orgapp-server/get-roster-field-group-details) `client.GetHrmField`
        - [x] [获取员工花名册字段信息](https://open.dingtalk.com/document/orgapp-server/intelligent-personnel-obtain-employee-roster-information) `client.GetHrmEmployeeField`
        - [x] [更新员工花名册信息](https://open.dingtalk.com/document/orgapp-server/intelligent-personnel-update-employee-file-information) `client.UpdateHrmEmployeeField`
        - [x] [获取花名册元数据](https://open.dingtalk.com/document/orgapp-server/intelligent-personnel-roster-metadata-query) `client.GetHrmMeta`
* 考勤
   - 考勤组管理
        - [x] [批量获取考勤组详情](https://open.dingtalk.com/document/orgapp-server/queries-attendance-group-list-details) `client.GetAttendanceGroups`
        - [x] [获取用户考勤组](https://open.dingtalk.com/document/orgapp-server/queries-a-user-attendance-group) `client.GetAttendanceUserGroup`
        - [x] [批量获取考勤组摘要](https://open.dingtalk.com/document/orgapp-server/batch-query-of-simple-information-of-the-attendance-group) `client.GetAttendanceGroupMinimalism`
        - [x] [获取考勤组详情](https://open.dingtalk.com/document/orgapp-server/query-a-single-attendance-group) `client.GetAttendanceGroupDetail`
        - [x] [搜索考勤组摘要](https://open.dingtalk.com/document/orgapp-server/attendance-group-search) `client.SearchAttendanceGroup`

* 应用管理 (**3**/4)
    - [x] [获取应用列表](https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-applications)
    - [x] [获取员工可见的应用列表](https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-the-list-of-applications-visible-to-an-employee)
    - [ ] [设置应用的可见范围](https://ding-doc.dingtalk.com/document#/org-dev-guide/set-the-visible-range-of-the-application)
    - [x] [获取应用可见范围](https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-the-application-visible-range)

* 群机器人
    - [x] [自定义机器人接入](https://ding-doc.dingtalk.com/document#/org-dev-guide/custom-robot) 基本的消息类型都已经支持并完成测试

* AI
    - [x] [文本翻译](https://ding-doc.dingtalk.com/document#/org-dev-guide/text-translation)
    - [x] [OCR文字识别](https://ding-doc.dingtalk.com/document#/org-dev-guide/ocr)
    - [x] [ASR 一句话语音识别](https://ding-doc.dingtalk.com/document#/org-dev-guide/asr-short-sentence-recognition)

* 文件存储 (**1**/10)
    - [x] [上传媒体文件](https://developers.dingtalk.com/document/app/upload-media-files) `client.MediaUpload`
    - [ ] [新增文件到用户自定义空间](https://ding-doc.dingtalk.com/document#/org-dev-guide/add-file-to-custom-space-2)
    - [ ] [发送钉盘文件给指定用户](https://ding-doc.dingtalk.com/document#/org-dev-guide/sends-a-dingtalk-disk-file-to-a-specified-user)
    - [ ] [获取企业下的自定义空间](https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-user-space-under-the-enterprise-2)
    - [ ] [获取应用自定义空间使用详情](https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-the-usage-details-of-a-custom-application-space)
    - [ ] [授权用户访问企业的自定义空间](https://ding-doc.dingtalk.com/document#/org-dev-guide/authorize-users-to-access-the-enterprise-s-custom-space)
    - [ ] [单步文件上传](https://ding-doc.dingtalk.com/document#/org-dev-guide/single-step-file-upload)
    - [ ] [开启分块上传事务](https://ding-doc.dingtalk.com/document#/org-dev-guide/enable-multipart-upload-transaction)
    - [ ] [提交文件上传事务](https://ding-doc.dingtalk.com/document#/org-dev-guide/submit-a-file-upload-transaction)
    - [ ] [上传文件块](https://ding-doc.dingtalk.com/document#/org-dev-guide/upload-file-blocks)
