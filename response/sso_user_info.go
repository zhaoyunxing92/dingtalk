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

package response

type SSOUserInfo struct {
	Response

	//是否是管理员
	Admin bool `json:"is_sys"`

	//企业信息
	corpInfo `json:"corp_info"`

	//用户信息
	userInfo `json:"user_info"`
}

type corpInfo struct {
	//授权方企业CorpId
	CorpId string `json:"corpid"`

	//授权方企业名称。
	CorpName string `json:"corp_name"`
}

type userInfo struct {

	//用户名字
	Name string `json:"name"`

	//用户的userid
	UserId string `json:"userid"`

	//email地址
	Email string `json:"email"`

	//头像地址
	Avatar string `json:"avatar"`
}
