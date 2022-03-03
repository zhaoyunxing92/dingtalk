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

// BpmInstanceChange 审批实例开始，结束
//{
//    "processInstanceId": "60406628-b642-4740-81e7-68cc67ebaa45",
//    "finishTime": 1640683876000,
//    "corpId": "dingc7c5220402493357f2c783f7214b6d69",
//    "EventType": "bpms_instance_change",
//    "businessId": "202112281713000090163",
//    "title": "赵云提交的审批任务",
//    "type": "finish",
//    "url": "https://aflow.dingtalk.com/dingtalk/mobile/hom",
//    "result": "agree",
//    "createTime": 1640682789000,
//    "processCode": "PROC-5CB63684-096E-40DD-B326-CF60691C507A",
//    "bizCategoryId": "",
//    "staffId": "manager164"
//}
type BpmInstanceChange struct {
	Event

	ProcessInstanceId string `json:"processInstanceId"`

	FinishTime int `json:"finishTime"`

	BusinessId string `json:"businessId"`

	Title string `json:"title"`

	Type string `json:"type"`

	Url string `json:"url"`

	Result string `json:"result"`

	CreateTime int `json:"createTime"`

	ProcessCode string `json:"processCode"`

	BizCategoryId string `json:"bizCategoryId"`

	StaffId string `json:"staffId"`
}
