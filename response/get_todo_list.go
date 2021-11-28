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

type GetTodoList struct {
	Response

	//请求ID
	RequestId string `json:"requestId"`

	TotalCount int `json:"totalCount"`

	//下一次请求的token
	Token string `json:"nextToken"`

	//待办卡片列表
	Todos []todo `json:"todoCards"`
}

/**
{
    "todoCardView": {
        "actionType": "url",
        "cardType": "unStandard",
        "circleELType": "icon",
        "contentType": "name",
        "icon": "https://img.alicdn.com/imgextra/i3/O1CN01CNOso71Ie1s5RqcRF_!!6000000000917-2-tps-78-78.png",
        "todoCardContentList": [

        ],
        "todoCardTitle": "使用api接口创建代办"
    }
}
*/
type todo struct {
	//待办ID
	TaskId string `json:"taskId"`

	//待办的标题
	Subject string `json:"subject"`

	Category string `json:"category"`

	//截止时间，Unix时间戳，单位毫秒
	DueTime int `json:"dueTime"`

	//完成状态
	Done bool `json:"isDone"`

	//详情页url跳转地址。
	Urls *detailUrl `json:"detailUrl"`

	// 业务系统侧的唯一标识ID，即业务ID。
	SourceId string `json:"sourceId"`

	//创建时间，Unix时间戳，单位毫秒
	CreatedTime int `json:"createdTime"`

	//更新时间，Unix时间戳，单位毫秒
	ModifiedTime int `json:"modifiedTime"`

	//创建者的unionId
	CreatorId string `json:"creatorId"`

	//接入应用标识
	BizTag string `json:"bizTag"`

	//优先级，取值：
	//10：较低
	//20：普通
	//30：紧急
	//40：非常紧急
	Priority int `json:"priority"`

	Status string `json:"todoStatus"`

	orgInfo `json:"orgInfo"`

	originalSource `json:"originalSource"`
}

type orgInfo struct {
	CorpId string `json:"corpId"`

	Name string `json:"name"`
}

type originalSource struct {
	Title string `json:"sourceTitle"`
}
