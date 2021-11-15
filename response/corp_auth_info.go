/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package response

type CorpAuthInfo struct {
	Response
	AuthInfo        `json:"auth_info"`
	AuthCorpInfo    `json:"auth_corp_info"`
	AuthUserInfo    `json:"auth_user_info"`
	ChannelAuthInfo `json:"channel_auth_info"`
}

// ChannelAuthInfo 授权的服务窗应用信息列表
type ChannelAuthInfo struct {
	ChannelAgents []ChannelAgent `json:"channelAgent"`
}
type AuthUserInfo struct {
	UserId string `json:"userId"`
}

type AuthCorpInfo struct {
	CorpId string `json:"corpid"`

	//企业所属行业
	Industry string `json:"industry"`

	//邀请码，只有自己邀请的企业才会返回邀请码，可用该邀请码统计不同渠道的拉新，否则值为空字符串
	InviteCode string `json:"invite_code"`

	//授权方企业名称
	CorpName string `json:"corp_name"`

	//序列号
	LicenseCode string `json:"license_code"`

	//渠道码
	AuthChannel string `json:"auth_channel"`

	//渠道类型。
	//为了避免渠道码重复，可与渠道码共同确认渠道。可能为空，非空时当前只有满天星类型，值为STAR_ACTIVITY
	AuthChannelType string `json:"auth_channel_type"`

	//企业是否认证。
	Authenticated bool `json:"is_authenticated"`

	//企业认证等级：
	//
	//0：未认证
	//
	//1：高级认证
	//
	//2：中级认证
	//
	//3：初级认证
	AuthLevel int `json:"auth_level"`

	//企业邀请链接
	InviteUrl string `json:"invite_url"`

	//企业logo
	CorpLogoUrl string `json:"corp_logo_url"`
}

type AuthInfo struct {
	Agents []Agent `json:"agent"`
}

type Agent struct {
	//授权方应用名字。
	Name string `json:"agent_name"`

	//授权方应用ID。
	AgentId int `json:"agentid"`

	//授权方应用头像。
	Logo string `json:"logo_url"`

	//应用ID
	AppId int `json:"appid"`

	//对此微应用有管理权限的管理员userid。
	Admins []string `json:"admin_list"`

	UseAppRole bool `json:"use_app_role"`
}

type ChannelAgent struct {
	//授权方应用名字
	Name string `json:"agent_name"`

	//授权方应用Id
	AgentId int `json:"agentid"`

	//授权方应用头像。
	Logo string `json:"logo_url"`

	//应用ID
	AppId int `json:"appid"`
}
