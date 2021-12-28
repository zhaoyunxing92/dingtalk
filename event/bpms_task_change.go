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

package event

// BpmTaskChange 审批任务开始，结束，转交
//{
//    "processInstanceId": "60406628-b642-4740-81e7-68cc67ebaa45",
//    "corpId": "dingc7c5220402493357f2c783f7214b6d69",
//    "EventType": "bpms_task_change",
//    "businessId": "202112281713000090163",
//    "title": "赵云提交的审批任务",
//    "type": "start",
//    "createTime": 1640682790000,
//    "processCode": "PROC-5CB63684-096E-40DD-B326-CF60691C507A",
//    "bizCategoryId": "",
//    "staffId": "manager164",
//    "taskId": 72062030325
//}
type BpmTaskChange struct {
	Event

	ProcessInstanceId string `json:"processInstanceId"`

	FinishTime int `json:"finishTime"`

	BusinessId string `json:"businessId"`

	Title string `json:"title"`

	// 备注
	Remark string `json:"remark"`

	Type string `json:"type"`

	Result string `json:"result"`

	CreateTime int `json:"createTime"`

	ProcessCode string `json:"processCode"`

	BizCategoryId string `json:"bizCategoryId"`

	StaffId string `json:"staffId"`

	TaskId int `json:"taskId"`
}
