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

type MessageSendResult struct {
	Response
	sendResult `json:"send_result"`
}

type sendResult struct {
	//无效的userid。
	InvalidUserIds []string `json:"invalid_user_id_list"`

	//因发送消息过于频繁或超量而被流控过滤后实际未发送的userid。
	//
	//未被限流的接收者仍会被成功发送。
	//
	//限流规则包括：
	//
	//给同一用户发相同内容消息一天仅允许一次
	//
	//同一个应用给同一个用户发送消息：
	//
	//如果是ISV接入方式，给同一用户发消息一天不得超过100次
	//
	//如果是企业接入方式，此上限为500
	ForbiddenUserIds []string `json:"forbidden_user_id_list"`

	//发送失败的userid。
	FailedUserIds []string `json:"failed_user_id_list"`

	//已读消息的userid
	ReadUserIds []string `json:"read_user_id_list"`

	//未读消息的userid
	UnReadUserIds []string `json:"unread_user_id_list"`

	//无效的部门ID
	InvalidDeptIds []string `json:"invalid_dept_id_list"`

	//推送被禁止的具体原因。
	Forbiddings []forbidden `json:"forbidden_list"`
}

type forbidden struct {

	//流控code：
	//
	//143105表示企业自建应用每日推送给用户的消息超过上限
	//
	//143106表示企业自建应用推送给用户的消息重复
	Code string `json:"code"`

	//流控阀值
	Count int `json:"count"`

	//被流控员工的userid
	UserId string `json:"userid"`
}
