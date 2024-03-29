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

package domain

// MicroApp 钉钉应用
type MicroApp struct {
	// 应用名称
	Name string `json:"name"`

	// 应用id
	AgentId int `json:"agentId"`

	// 应用图标
	Icon string `json:"appIcon"`

	// 应用描述
	Desc string `json:"appDesc"`

	// 是否自建 false:不是
	Self bool `json:"isSelf"`

	// 应用状态 1：启用，0：停用
	Status int `json:"appStatus"`

	// 应用应用的OA后台管理主页
	OmpLink string `json:"ompLink"`

	// 应用的移动端主页
	HomepageLink string `json:"homepageLink"`

	// 应用的PC端主页
	PcHomepageLink string `json:"pcHomepageLink"`
}
