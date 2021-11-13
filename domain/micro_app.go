/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package domain

//钉钉应用
type MicroApp struct {
	Name           string `json:"name"`           //应用名称
	AgentId        uint64 `json:"agentId"`        //应用id
	Icon           string `json:"appIcon"`        //应用图标
	Desc           string `json:"appDesc"`        //应用描述
	Self           bool   `json:"isSelf"`         //是否自建 false:不是
	Status         uint8  `json:"appStatus"`      //应用状态 1：启用，0：停用
	OmpLink        string `json:"ompLink"`        //应用应用的OA后台管理主页
	HomepageLink   string `json:"homepageLink"`   //应用的移动端主页
	PcHomepageLink string `json:"pcHomepageLink"` //应用的PC端主页
}

//应用列表
type MicroAppList struct {
	Response
	AppList []MicroApp `json:"appList"` //应用列表
}

//应用可见范围
type MicroAppVisibleScopes struct {
	Response
	UserVisibleScopes []string `json:"userVisibleScopes"` //应用可见的用户列表。
	DeptVisibleScopes []uint32 `json:"deptVisibleScopes"` //应用可见的部门列表。
	Hidden            bool     `json:"isHidden"`          //是否仅限管理员可见，true代表仅限管理员可见。
}
