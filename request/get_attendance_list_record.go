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

type GetAttendanceListRecord struct {
	// 企业内的员工ID列表，最大值50
	UserIds []string `json:"userIds"`

	// 查询考勤打卡记录的起始工作日。格式为：yyyy-MM-dd hh:mm:ss
	// 例如，参数传"2021-12-01 10:00:00"，员工在09:00的打卡信息获取不到
	CheckDateFrom string `json:"checkDateFrom"`

	// 查询考勤打卡记录的结束工作日。格式为：yyyy-MM-dd hh:mm:ss
	// 例如，参数传"2021-12-01 18:00:00"，员工在19:00的打卡信息获取不到
	CheckDateTo string `json:"checkDateTo"`

	// 是否为海外企业使用
	IsI18n bool `json:"isI18n"`
}
