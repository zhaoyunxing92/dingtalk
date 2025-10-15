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

package request

type GetAttendanceList struct {
	// 查询考勤打卡记录的起始工作日。格式为yyyy-MM-dd HH:mm:ss，HH:mm:ss可以使用00:00:00，将返回此日期从0点到24点的结果。
	// 例如，参数传"2021-12-01 10:00"，获取的是12月1日一整天的考勤结果。
	WorkDateFrom string `json:"workDateFrom"`

	// 查询考勤打卡记录的结束工作日
	WorkDateTo string `json:"workDateTo"`

	// 员工在企业内的userId列表，最大值50个
	UserIdList []string `json:"userIdList"`

	// 表示获取考勤数据的起始点
	Offset int `json:"offset"`

	// 表示获取考勤数据的条数，最大值50
	Limit int `json:"limit"`

	// 是否为海外企业使用
	IsI18n bool `json:"isI18n"`
}
