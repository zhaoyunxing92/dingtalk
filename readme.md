![dingtalk](./image/dingtalk.png)

> [创建和配置应用](https://ding-doc.dingtalk.com/document#/org-dev-guide/applicaiton-manangment) 先有应用我们才能一起玩

[dingtalk](https://github.com/zhaoyunxing92/dingtalk) 是基于新版的[钉钉开发平台api文档](https://ding-doc.dingtalk.com/document#/org-dev-guide) 开发，也就是说在你使用之前是需要在[钉钉开发者后台](https://open-dev.dingtalk.com/#/index) 创建一个小程序或应用

## 为什么有它

首先我要感谢 [godingtalk](https://github.com/hugozhu/godingtalk) 库,我是参考他的代码写的。但是由于钉钉历史原因，很多api都更新了,
新的企业是无法获取`corpid`和`corpsecret`的现在也无法获取了,其次就是我有很多钉钉预警的业务需要做,于是就写了这个

## 怎么使用

基本上每个接口我都会在`tests`目录里面有测试用例可以去查看

## 开发进度

* 获取凭证 (**1**/3)
  - [x] [获取access_token](https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-access_token)

  - [ ] [获取jsapi_ticket](https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-the-jsapi-credential-jsapi_ticket)
  
  - [ ] [获取微应用后台免登的SsoToken](https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-the-ssotoken-for-micro-application-background-logon-free)

* 通讯录管理

  - 用户管理(**10**/10)
      - [x] [创建用户](https://ding-doc.dingtalk.com/document#/org-dev-guide/create-user)

      - [x] [删除用户](https://ding-doc.dingtalk.com/document#/org-dev-guide/delete-user)
      
      - [x] [更新用户详情](https://ding-doc.dingtalk.com/document#/org-dev-guide/update-user-info)

      - [x] [获取用户详情](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-user-detail)
      
      - [x] [获取部门用户userid列表](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-deptment-users)

      - [x] [根据unionid获取userid](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-Userid-By-Unionid)
      
      - [x] [根据手机号获取userid](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-userid-By-Mobile)

      - [x] [获取企业员工人数](https://ding-doc.dingtalk.com/document#/org-dev-guide/Get-user-count)
      
      - [x] [获取部门用户详情](https://ding-doc.dingtalk.com/document#/org-dev-guide/list-dept-user-details)
     
      - [x] [获取未登录钉钉的员工列表](https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-the-list-of-employees-who-have-not-logged-on)
      
      - [x] [获取管理员列表](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-Admin-list)
      
      - [x] [获取管理员通讯录权限范围](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-permission-of-admin)

  - 部门管理(**8**/8)
      - [x] [创建部门](https://ding-doc.dingtalk.com/document#/org-dev-guide/create-a-department)
      
      - [x] [删除部门](https://ding-doc.dingtalk.com/document#/org-dev-guide/delete-a-department)
      
      - [x] [更新部门](https://ding-doc.dingtalk.com/document#/org-dev-guide/update-a-department-v1)
      
      - [x] [获取部门详情](https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-department-details-v1)
      
      - [x] [获取部门列表](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-department-list)
      
      - [x] [获取子部门Id列表](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-sub-departmentId-list)
      
      - [x] [查询指定用户的所有上级父部门路径](https://ding-doc.dingtalk.com/document#/org-dev-guide/GET-parent-depts)
      
      - [x] [查询部门的所有上级父部门路径](https://ding-doc.dingtalk.com/document#/org-dev-guide/GET-parent-deptsby-dept)
      
  - 角色管理(**3**/11)
     - [x] [获取角色列表](https://ding-doc.dingtalk.com/document#/org-dev-guide/list-roles)
     
     - [x] [获取指定角色的员工列表](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-role-simplelist)
     
     - [x] [获取角色组](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-role-group)
     
     - [x] [获取角色详情](https://ding-doc.dingtalk.com/document#/org-dev-guide/get-role)
     
     - [x] [创建角色组](https://ding-doc.dingtalk.com/document#/org-dev-guide/add-role-group)
     
     - [x] [创建角色](https://ding-doc.dingtalk.com/document#/org-dev-guide/add-role)
     
     - [x] [更新角色](https://ding-doc.dingtalk.com/document#/org-dev-guide/update-role)
     
     - [x] [删除角色](https://ding-doc.dingtalk.com/document#/org-dev-guide/delete-role)
     
     - [x] [批量增加员工角色](https://ding-doc.dingtalk.com/document#/org-dev-guide/add-role-for-users)
     
     - [x] [批量删除员工角色](https://ding-doc.dingtalk.com/document#/org-dev-guide/remove-roles-for-users)
     
     - [x] [设定角色成员管理范围](https://ding-doc.dingtalk.com/document#/org-dev-guide/set-the-management-scope-of-role-members) 官方接口不通
     
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
  
  
  