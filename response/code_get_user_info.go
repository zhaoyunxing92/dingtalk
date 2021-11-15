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

type CodeGetUserInfo struct {
	Response
	codeGetUserInfo `json:"result"`
}

type codeGetUserInfo struct {
	//用户名字
	Name string `json:"name"`

	//用户unionId
	UnionId string `json:"unionid"`

	//用户的userid
	UserId string `json:"userid"`

	//用户关联的unionId
	AssociatedUnionId string `json:"associated_unionid"`

	//级别。
	//
	//1：主管理员
	//
	//2：子管理员
	//
	//100：老板
	//
	//0：其他（如普通员工）
	Level int `json:"sys_level"`

	//是否是管理员。
	//
	//true：是
	//
	//false：不是
	Admin bool `json:"sys"`

	//设备ID
	DeviceId string `json:"device_id"`
}
