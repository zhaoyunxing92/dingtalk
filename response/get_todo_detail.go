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

type GetTodoDetail struct {
	Response

	//待办ID
	Id string `json:"id"`

	//待办的标题
	Subject string `json:"subject"`

	//待办描述
	Desc string `json:"description"`

	//开始时间，Unix时间戳，单位毫秒
	StartTime int `json:"startTime"`

	//截止时间，Unix时间戳，单位毫秒
	DueTime int `json:"dueTime"`

	//完成时间，Unix时间戳，单位毫秒
	FinishTime int `json:"finishTime"`

	//完成状态
	Done bool `json:"done"`

	//执行者的unionId
	Executors []string `json:"executorIds"`

	//参与者的unionId
	Participants []string `json:"participantIds"`

	//详情页url跳转地址。
	Urls *detailUrl `json:"detailUrl"`

	//业务来源
	Source string `json:"source"`

	// 业务系统侧的唯一标识ID，即业务ID。
	SourceId string `json:"sourceId"`

	//创建时间，Unix时间戳，单位毫秒
	CreatedTime int `json:"createdTime"`

	//更新时间，Unix时间戳，单位毫秒
	ModifiedTime int `json:"modifiedTime"`

	//创建者的unionId
	CreatorId string `json:"creatorId"`

	//更新者的unionId
	ModifierId string `json:"modifierId"`

	//租户ID
	TenantId string `json:"tenantId"`

	//租户类型。
	//user：个人
	//org：企业
	//group：群
	TenantType string `json:"tenantType"`

	//接入应用标识
	BizTag string `json:"bizTag"`

	//请求ID
	RequestId string `json:"requestId"`

	//待办卡片类型ID
	CardTypeId string `json:"cardTypeId"`

	//生成的待办是否仅展示在执行者的待办列表中
	OnlyShowExecutor bool `json:"isOnlyShowExecutor"`

	//优先级，取值：
	//10：较低
	//20：普通
	//30：紧急
	//40：非常紧急
	Priority int `json:"priority"`
}
