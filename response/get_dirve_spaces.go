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

type GetDriveSpaces struct {
	Response

	Spaces []struct {
		// 空间ID
		SpaceId string `json:"spaceId"`

		// 空间名称
		Name string `json:"spaceName"`

		// 空间类型
		Type string `json:"spaceType"`

		// 空间总额度
		Quota int `json:"quota"`

		// 空间已使用额度
		UsedQuota int `json:"usedQuota"`

		//授权模式，取值：
		//
		//acl：acl授权
		//custom：自定义授权
		PermissionMode string `json:"permissionMode"`

		// 创建时间
		CreateTime string `json:"createTime"`

		// 修改时间
		ModifyTime string `json:"modifyTime"`
	} `json:"spaces"`

	Token string `json:"nextToken"`
}
