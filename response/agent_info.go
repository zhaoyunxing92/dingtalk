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

type AgentInfo struct {
	Response
	//企业应用名称
	Name string `json:"name"`

	//企业应用ID。
	AgentId int `json:"agentid"`

	//授权方应用头像。
	Logo string `json:"logo_url"`

	//企业应用描述
	Description string `json:"description"`

	//授权方企业应用是否被禁用：
	//
	//0：禁用
	//
	//1：正常
	//
	//2：待激活
	Close int `json:"close"`
}
