![dingtalk](./image/dingtalk.png)

> [创建和配置应用](https://ding-doc.dingtalk.com/document#/org-dev-guide/applicaiton-manangment) 先有应用我们才能一起玩

[dingtalk](https://github.com/zhaoyunxing92/dingtalk/v2)
是基于新版的[钉钉开发平台api文档](https://ding-doc.dingtalk.com/document#/org-dev-guide)
开发，也就是说在你使用之前是需要在[钉钉开发者后台](https://open-dev.dingtalk.com/#/index) 创建一个小程序或应用

## 为什么有它

首先我要感谢 [godingtalk](https://github.com/hugozhu/godingtalk) 库,我是参考他的代码写的。但是由于钉钉历史原因，很多api都更新了, 新的企业是无法获取`corpid`
和`corpsecret`的现在也无法获取了,其次就是我有很多钉钉预警的业务需要做,于是就写了这个

## 参考文档

[dingtalk-callback-Crypto](https://github.com/open-dingtalk/dingtalk-callback-Crypto)

## 开发进度

* 获取凭证 (**2**/5)
    - [x] [获取企业内部应用的access_token](https://developers.dingtalk.com/document/app/obtain-orgapp-token)

    - [ ] [服务商获取第三方应用授权企业的access_token](https://developers.dingtalk.com/document/app/obtains-the-enterprise-authorized-credential)

    - [x] [获取第三方企业应用的suite_access_token](https://developers.dingtalk.com/document/app/obtain-application-suite-ticket)

    - [ ] [获取jsapi_ticket](https://developers.dingtalk.com/document/app/obtain-jsapi_ticket)
 
    - [ ] [获取微应用后台免登的SsoToken](https://developers.dingtalk.com/document/app/obtain-the-ssotoken-for-micro-application-background-logon-free)

* 通讯录管理

    - 用户管理(**14**/14)
        - [x] [创建用户](https://developers.dingtalk.com/document/app/user-information-creation)

        - [x] [删除用户](https://developers.dingtalk.com/document/app/delete-a-user)

        - [x] [更新用户详情](https://developers.dingtalk.com/document/app/user-information-update)

        - [x] [根据userid获取用户详情](https://developers.dingtalk.com/document/app/query-user-details)

        - [x] [获取部门用户基础信息](https://developers.dingtalk.com/document/app/queries-the-simple-information-of-a-department-user)

        - [x] [获取部门用户userid列表](https://developers.dingtalk.com/document/app/query-the-list-of-department-userids)

        - [x] [获取部门用户详情](https://developers.dingtalk.com/document/app/queries-the-complete-information-of-a-department-user)

        - [x] [获取员工人数](https://developers.dingtalk.com/document/app/obtain-the-number-of-employees-v2)

        - [x] [获取未登录钉钉的员工列表](https://developers.dingtalk.com/document/app/queries-the-inactive-users-or-active-users-under-an-enterprise)

        - [x] [根据手机号获取userid](https://developers.dingtalk.com/document/app/query-users-by-phone-number)

        - [x] [根据unionid获取用户userid](https://developers.dingtalk.com/document/app/query-a-user-by-the-union-id)

        - [x] [获取管理员列表](https://developers.dingtalk.com/document/app/query-the-administrator-list)

        - [x] [获取管理员通讯录权限范围](https://developers.dingtalk.com/document/app/query-permissions-of-the-administrator-address-book)

        - [x] [获取管理员的应用管理权限](https://developers.dingtalk.com/document/app/obtains-the-administrator-s-microapplication-management-permission)

    - 部门管理(**8**/8)
        - [x] [创建部门](https://developers.dingtalk.com/document/app/create-a-department-v2)

        - [x] [更新部门](https://developers.dingtalk.com/document/app/update-a-department-v2)

        - [x] [删除部门](https://developers.dingtalk.com/document/app/delete-a-department-v2)

        - [x] [获取部门详情](https://developers.dingtalk.com/document/app/query-department-details0-v2)

        - [x] [获取部门列表](https://developers.dingtalk.com/document/app/obtain-the-department-list-v2)

        - [x] [获取子部门Id列表](https://developers.dingtalk.com/document/app/obtain-a-sub-department-id-list-v2)

        - [x] [查询指定用户的所有上级父部门路径](https://developers.dingtalk.com/document/app/query-the-list-of-all-parent-departments-of-a-department)

        - [x] [查询部门的所有上级父部门路径](https://developers.dingtalk.com/document/app/queries-the-list-of-all-parent-departments-of-a-user)

    - 角色管理(**11**/11)

        - [x] [获取角色组](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-role-group)

        - [x] [获取角色详情](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-role)

        - [x] [创建角色组](https://ding-doc.dingtalk.com/document#/org-dev-guide/add-role-group)

        - [x] [创建角色](https://ding-doc.dingtalk.com/document#/org-dev-guide/add-role)

        - [x] [更新角色](https://ding-doc.dingtalk.com/document#/org-dev-guide/update-role)

        - [x] [删除角色](https://ding-doc.dingtalk.com/document#/org-dev-guide/delete-role)

        - [x] [批量增加员工角色](https://ding-doc.dingtalk.com/document#/org-dev-guide/add-role-for-users)

        - [x] [批量删除员工角色](https://ding-doc.dingtalk.com/document#/org-dev-guide/remove-roles-for-users)

        - [x] [设定角色成员管理范围](https://ding-doc.dingtalk.com/document#/org-dev-guide/set-the-management-scope-of-role-members)
          未测试通过

* 群会话管理(**5**/5)
    - [x] [创建群会话](https://ding-doc.dingtalk.com/document#/org-dev-guide/create-chat)

    - [x] [获取群会话](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-chat-detail)

    - [x] [修改群会话](https://ding-doc.dingtalk.com/document#/org-dev-guide/update-chat-config)

    - [x] [设置禁止群成员私聊](https://ding-doc.dingtalk.com/document#/org-dev-guide/set-private-chat)

    - [x] [设置群管理员](https://ding-doc.dingtalk.com/document#/org-dev-guide/set-chat-admin)

* 消息通知(4/6)

    - 工作通知
        - [x] [发送工作通知](https://ding-doc.dingtalk.com/document#/org-dev-guide/send-work-notifications)

        - [x] [获取工作通知消息的发送进度](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-work-notification-progess)

        - [x] [获取工作通知消息的发送结果](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-work-notification-result)

        - [x] [撤回工作通知消息](https://ding-doc.dingtalk.com/document#/org-dev-guide/withdrawal-work-notification)

    - 企业群消息
        - [x] [发送消息到企业群](https://ding-doc.dingtalk.com/document#/org-dev-guide/send-chat-messages)

        - [x] [查询群消息已读人员列表](https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-the-list-of-people-who-have-read-a-group-1)

        - [x] [发送普通消息](https://ding-doc.dingtalk.com/document#/org-dev-guide/send-normal-messages)

* 钉钉运动

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
    - [x] [上传媒体文件](https://ding-doc.dingtalk.com/document#/org-dev-guide/upload-media-files-1)

    - [ ] [新增文件到用户自定义空间](https://ding-doc.dingtalk.com/document#/org-dev-guide/add-file-to-custom-space-2)

    - [ ] [发送钉盘文件给指定用户](https://ding-doc.dingtalk.com/document#/org-dev-guide/sends-a-dingtalk-disk-file-to-a-specified-user)

    - [ ] [获取企业下的自定义空间](https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-user-space-under-the-enterprise-2)

    - [ ] [获取应用自定义空间使用详情](https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-the-usage-details-of-a-custom-application-space)

    - [ ] [授权用户访问企业的自定义空间](https://ding-doc.dingtalk.com/document#/org-dev-guide/authorize-users-to-access-the-enterprise-s-custom-space)

    - [ ] [单步文件上传](https://ding-doc.dingtalk.com/document#/org-dev-guide/single-step-file-upload)

    - [ ] [开启分块上传事务](https://ding-doc.dingtalk.com/document#/org-dev-guide/enable-multipart-upload-transaction)

    - [ ] [提交文件上传事务](https://ding-doc.dingtalk.com/document#/org-dev-guide/submit-a-file-upload-transaction)

    - [ ] [上传文件块](https://ding-doc.dingtalk.com/document#/org-dev-guide/upload-file-blocks)
  
  
  